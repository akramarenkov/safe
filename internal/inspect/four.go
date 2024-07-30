package inspect

import (
	"runtime"
	"sync"
)

// Function with four arguments that returns a reference value.
type Reference4 func(first, second, third, fourth int64) (int64, error)

// Options of inspecting function with four arguments. A reference and inspected
// functions must be specified.
type Opts4[Type EightBits] struct {
	// Inspected function with four arguments
	Inspected func(first, second, third, fourth Type) (Type, error)
	// Function with four arguments that returns a reference value
	Reference4

	// Minimum and maximum value for specified type
	min int64
	max int64
}

// Validates options. A reference and inspected functions must be specified.
func (opts *Opts4[Type]) IsValid() error {
	if opts.Reference4 == nil {
		return ErrReferenceNotSpecified
	}

	if opts.Inspected == nil {
		return ErrInspectedNotSpecified
	}

	return nil
}

// Performs inspection.
func (opts *Opts4[Type]) Do() (Result[Type], error) {
	if err := opts.IsValid(); err != nil {
		return Result[Type]{}, err
	}

	opts.min, opts.max = pickUpRange[Type]()

	return opts.main(), nil
}

func (opts *Opts4[Type]) main() Result[Type] {
	parallelization := runtime.NumCPU()

	// so the buffer size is chosen for simplicity: so that all goroutines can
	// definitely write the result and not block on writing even without reading
	// these results
	results := make(chan Result[Type], parallelization)
	defer close(results)

	wg := &sync.WaitGroup{}
	defer wg.Wait()

	// so the buffer size is chosen for simplicity: so that the goroutine can
	// definitely write all the first numbers and not block on writing even without
	// reading these first numbers
	//
	// opts.max and opts.min accept values ​​in the range int8|uint8, and themselves have
	// type int64, so overflow is impossible
	firsts := make(chan int64, opts.max-opts.min)

	for range parallelization {
		wg.Add(1)

		go func() {
			defer wg.Done()

			results <- opts.loop(firsts)
		}()
	}

	for first := opts.min; first <= opts.max; first++ {
		firsts <- first
	}

	// In case of a single error, there is a wait for all loops to complete. Firstly,
	// such an error is rare. Secondly, the performance in case of using a separate
	// interrupter drops by two or more times. Thirdly, the number of iterations and
	// the execution time in this case will be the same as in the absence of an error
	close(firsts)

	received := 0
	result := Result[Type]{}

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
func (opts *Opts4[Type]) loop(firsts chan int64) Result[Type] {
	result := Result[Type]{}

	for first := range firsts {
		for second := opts.min; second <= opts.max; second++ {
			for third := opts.min; third <= opts.max; third++ {
				for fourth := opts.min; fourth <= opts.max; fourth++ {
					reference, fault := opts.Reference4(first, second, third, fourth)

					actual, err := opts.Inspected(
						Type(first),
						Type(second),
						Type(third),
						Type(fourth),
					)

					if fault != nil {
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
							)

							return result
						}

						result.ReferenceFaults++

						continue
					}

					if reference > opts.max || reference < opts.min {
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
						)

						return result
					}

					result.NoOverflows++
				}
			}
		}
	}

	return result
}
