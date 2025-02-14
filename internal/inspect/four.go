package inspect

import (
	"runtime"
	"sync"

	"github.com/akramarenkov/safe/internal/inspect/types"

	"github.com/akramarenkov/intspec"
)

// Inspected function with four arguments.
type Inspected4[Type types.I8] func(first, second, third, fourth Type) (Type, error)

// Function with four arguments that returns a reference value.
type Reference4 func(first, second, third, fourth int64) (int64, error)

type inspector4[Type types.I8] struct {
	// Inspected function with four arguments
	inspected Inspected4[Type]
	// Function with four arguments that returns a reference value
	reference Reference4

	// Minimum and maximum value for specified type
	minimum int64
	maximum int64
}

// Performs inspection with four arguments.
//
// A inspected and reference functions must be specified.
func Do4[Type types.I8](
	inspected Inspected4[Type],
	reference Reference4,
) (types.Result[Type, Type, int64], error) {
	if inspected == nil {
		return types.Result[Type, Type, int64]{}, ErrInspectedNotSpecified
	}

	if reference == nil {
		return types.Result[Type, Type, int64]{}, ErrReferenceNotSpecified
	}

	minimum, maximum := intspec.Range[Type]()

	insp := &inspector4[Type]{
		inspected: inspected,
		reference: reference,

		minimum: int64(minimum),
		maximum: int64(maximum),
	}

	return insp.do(), nil
}

func (insp *inspector4[Type]) do() types.Result[Type, Type, int64] {
	parallelization := runtime.NumCPU()

	// Buffer size is chosen for simplicity: so that all goroutines can
	// definitely write the result and not block on writing even without reading
	// these results
	results := make(chan types.Result[Type, Type, int64], parallelization)
	defer close(results)

	wg := &sync.WaitGroup{}
	defer wg.Wait()

	// Buffer size is chosen for simplicity: so that the goroutine can
	// definitely write all the first numbers and not block on writing even without
	// reading these first numbers
	//
	// insp.maximum and insp.minimum accept values ​​in the range int8|uint8, and
	// themselves have type int64, so overflow is impossible
	firsts := make(chan int64, insp.maximum-insp.minimum)

	for range parallelization {
		wg.Add(1)

		go func() {
			defer wg.Done()

			results <- insp.loop(firsts)
		}()
	}

	for first := insp.minimum; first <= insp.maximum; first++ {
		firsts <- first
	}

	// In case of a single error, there is a wait for all loops to complete. Firstly,
	// such an error is rare. Secondly, the performance in case of using a separate
	// interrupter drops by two or more times. Thirdly, the number of iterations and
	// the execution time in this case will be the same as in the absence of an error
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
func (insp *inspector4[Type]) loop(firsts chan int64) types.Result[Type, Type, int64] {
	result := types.Result[Type, Type, int64]{}

	for first := range firsts {
		for second := insp.minimum; second <= insp.maximum; second++ {
			for third := insp.minimum; third <= insp.maximum; third++ {
				for fourth := insp.minimum; fourth <= insp.maximum; fourth++ {
					reference, fault := insp.reference(first, second, third, fourth)

					actual, err := insp.inspected(
						Type(first),
						Type(second),
						Type(third),
						Type(fourth),
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
							)

							return result
						}

						result.ReferenceFaults++

						continue
					}

					if reference > insp.maximum || reference < insp.minimum {
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
