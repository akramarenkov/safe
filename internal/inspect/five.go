package inspect

import (
	"runtime"
	"sync"

	"github.com/akramarenkov/safe/internal/inspect/types"
)

// Inspected function with five arguments.
type Inspected5[Type types.USI8] func(first, second, third, fourth, fifth Type) (Type, error)

// Function with five arguments that returns a reference value.
type Reference5 func(first, second, third, fourth, fifth int64) (int64, error)

// Options of inspecting function with five arguments. A reference and inspected
// functions must be specified.
type Opts5[Type types.USI8] struct {
	// Inspected function with five arguments
	Inspected Inspected5[Type]
	// Function with five arguments that returns a reference value
	Reference Reference5

	// Minimum and maximum value for specified type
	minimum int64
	maximum int64
}

// Validates options. A reference and inspected functions must be specified.
func (opts Opts5[Type]) IsValid() error {
	if opts.Reference == nil {
		return ErrReferenceNotSpecified
	}

	if opts.Inspected == nil {
		return ErrInspectedNotSpecified
	}

	return nil
}

// Performs inspection.
func (opts Opts5[Type]) Do() (types.Result[Type, Type, int64], error) {
	if err := opts.IsValid(); err != nil {
		return types.Result[Type, Type, int64]{}, err
	}

	opts.minimum, opts.maximum = ConvSpan[Type, int64]()

	return opts.main(), nil
}

func (opts *Opts5[Type]) main() types.Result[Type, Type, int64] {
	parallelization := runtime.NumCPU()

	// buffer size is chosen for simplicity: so that all goroutines can
	// definitely write the result and not block on writing even without reading
	// these results
	results := make(chan types.Result[Type, Type, int64], parallelization)
	defer close(results)

	wg := &sync.WaitGroup{}
	defer wg.Wait()

	// buffer size is chosen for simplicity: so that the goroutine can
	// definitely write all the first numbers and not block on writing even without
	// reading these first numbers
	//
	// opts.maximum and opts.minimum accept values ​​in the range int8|uint8, and
	// themselves have type int64, so overflow is impossible
	firsts := make(chan int64, opts.maximum-opts.minimum)

	for range parallelization {
		wg.Add(1)

		go func() {
			defer wg.Done()

			results <- opts.loop(firsts)
		}()
	}

	for first := opts.minimum; first <= opts.maximum; first++ {
		firsts <- first
	}

	// In case of a single error, there is a wait for all loops to complete. Firstly,
	// such an error is rare. Secondly, the performance in case of using a separate
	// interrupter drops by two or more times. Thirdly, this code is launched manually
	// and after receiving the first error, the test can be interrupted manually
	close(firsts)

	received := 0
	result := types.Result[Type, Type, int64]{}

	for interim := range results {
		received++

		if interim.Conclusion != nil {
			return interim
		}

		result.NoOverflows += interim.NoOverflows
		result.Overflows += interim.Overflows
		result.ReferenceFaults += interim.ReferenceFaults

		if received == parallelization {
			break // for coverage
		}
	}

	return result
}

//nolint:gocognit // When the complexity decreases, the performance drops by half.
func (opts *Opts5[Type]) loop(firsts chan int64) types.Result[Type, Type, int64] {
	result := types.Result[Type, Type, int64]{}

	for first := range firsts {
		for second := opts.minimum; second <= opts.maximum; second++ {
			for third := opts.minimum; third <= opts.maximum; third++ {
				for fourth := opts.minimum; fourth <= opts.maximum; fourth++ {
					for fifth := opts.minimum; fifth <= opts.maximum; fifth++ {
						reference, fault := opts.Reference(
							first,
							second,
							third,
							fourth,
							fifth,
						)

						actual, err := opts.Inspected(
							Type(first),
							Type(second),
							Type(third),
							Type(fourth),
							Type(fifth),
						)

						if fault != nil {
							if err == nil {
								result.Actual = actual
								result.Conclusion = ErrErrorExpected

								result.Args = append(
									[]Type(nil),
									Type(first),
									Type(second),
									Type(third),
									Type(fourth),
									Type(fifth),
								)

								return result
							}

							result.ReferenceFaults++

							continue
						}

						if reference > opts.maximum || reference < opts.minimum {
							if err == nil {
								result.Actual = actual
								result.Conclusion = ErrErrorExpected
								result.Reference = reference

								result.Args = append(
									[]Type(nil),
									Type(first),
									Type(second),
									Type(third),
									Type(fourth),
									Type(fifth),
								)

								return result
							}

							result.Overflows++

							continue
						}

						if err != nil {
							result.Conclusion = ErrUnexpectedError
							result.Err = err
							result.Reference = reference

							result.Args = append(
								[]Type(nil),
								Type(first),
								Type(second),
								Type(third),
								Type(fourth),
								Type(fifth),
							)

							return result
						}

						if int64(actual) != reference {
							result.Actual = actual
							result.Conclusion = ErrNotEqual
							result.Reference = reference

							result.Args = append(
								[]Type(nil),
								Type(first),
								Type(second),
								Type(third),
								Type(fourth),
								Type(fifth),
							)

							return result
						}

						result.NoOverflows++
					}
				}
			}
		}
	}

	return result
}
