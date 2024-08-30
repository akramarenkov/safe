package inspect

import (
	"github.com/akramarenkov/safe/internal/inspect/types"
)

// Options of inspecting. A reference and inspected functions must be specified.
type Opts[TypeFrom, TypeTo types.UpToUSI32, TypeRef types.SIF64] struct {
	// Number of nested loops, i.e. the number of generated arguments for reference and
	// inspected functions. A value greater than three is not recommended due to low
	// performance
	LoopsQuantity uint

	// Inspected function
	Inspected func(args ...TypeFrom) (TypeTo, error)
	// Function that returns a reference value
	Reference func(args ...TypeRef) (TypeRef, error)
	// Optional function that customize arg values span
	Span func() (TypeRef, TypeRef)

	// Minimum and maximum value for specified TypeTo type
	min TypeRef
	max TypeRef

	// Buffers used to decrease allocations
	argsFrom []TypeFrom
	argsRef  []TypeRef

	// Result of inspecting
	result types.Result[TypeFrom, TypeTo, TypeRef]
}

// Validates options. A reference and inspected functions must be specified.
func (opts Opts[TypeFrom, TypeTo, TypeRef]) IsValid() error {
	if opts.Reference == nil {
		return ErrReferenceNotSpecified
	}

	if opts.Inspected == nil {
		return ErrInspectedNotSpecified
	}

	return nil
}

// Performs inspection.
func (opts Opts[TypeFrom, TypeTo, TypeRef]) Do() (
	types.Result[TypeFrom, TypeTo, TypeRef],
	error,
) {
	if err := opts.IsValid(); err != nil {
		return types.Result[TypeFrom, TypeTo, TypeRef]{}, err
	}

	opts.min, opts.max = PickUpSpan[TypeTo, TypeRef]()

	opts.argsFrom = make([]TypeFrom, opts.LoopsQuantity)
	opts.argsRef = make([]TypeRef, opts.LoopsQuantity)

	if err := opts.main(); err != nil {
		return types.Result[TypeFrom, TypeTo, TypeRef]{}, err
	}

	return opts.result, nil
}

func (opts *Opts[TypeFrom, TypeTo, TypeRef]) main() error {
	if _, err := loop(opts.LoopsQuantity, opts.Span, opts.do); err != nil {
		return err
	}

	return nil
}

func (opts *Opts[TypeFrom, TypeTo, TypeRef]) do(args ...TypeFrom) bool {
	// Protection against changes from the inspected and reference functions
	copy(opts.argsFrom, args)

	for id := range len(args) {
		opts.argsRef[id] = TypeRef(args[id])
	}

	reference, fault := opts.Reference(opts.argsRef...)

	actual, err := opts.Inspected(opts.argsFrom...)

	if fault != nil {
		if err == nil {
			opts.result.Actual = actual
			opts.result.Conclusion = ErrErrorExpected

			opts.result.Args = append([]TypeFrom(nil), args...)

			return true
		}

		opts.result.ReferenceFaults++

		return false
	}

	if reference > opts.max || reference < opts.min {
		if err == nil {
			opts.result.Actual = actual
			opts.result.Conclusion = ErrErrorExpected
			opts.result.Reference = reference

			opts.result.Args = append([]TypeFrom(nil), args...)

			return true
		}

		opts.result.Overflows++

		return false
	}

	if err != nil {
		opts.result.Conclusion = ErrUnexpectedError
		opts.result.Err = err
		opts.result.Reference = reference

		opts.result.Args = append([]TypeFrom(nil), args...)

		return true
	}

	// A universal, in this case, condition for comparing for inequality
	// (actual != reference), both for integers and for floating point numbers
	if TypeRef(actual)-reference >= 1 || TypeRef(actual)-reference <= -1 {
		opts.result.Actual = actual
		opts.result.Conclusion = ErrNotEqual
		opts.result.Reference = reference

		opts.result.Args = append([]TypeFrom(nil), args...)

		return true
	}

	opts.result.NoOverflows++

	return false
}

func loop[TypeRef types.SIF64, TypeFrom types.UpToUSI32](
	level uint,
	span func() (TypeRef, TypeRef),
	do func(args ...TypeFrom) bool,
	args ...TypeFrom,
) (bool, error) {
	if level == 0 {
		return do(args...), nil
	}

	level--

	args = append(args, 0)

	begin, end, err := getSpan[TypeFrom](span)
	if err != nil {
		return false, err
	}

	for number := begin; number <= end; number++ {
		args[len(args)-1] = TypeFrom(number)

		if level == 0 {
			if stop := do(args...); stop {
				return true, nil
			}

			continue
		}

		// The only error that can occur will be detected on the first (topmost) call
		if stop, _ := loop(level, span, do, args...); stop {
			return true, nil
		}
	}

	return false, nil
}

func getSpan[TypeFrom types.UpToUSI32, TypeRef types.SIF64](
	span func() (TypeRef, TypeRef),
) (TypeRef, TypeRef, error) {
	begin, end := PickUpSpan[TypeFrom, TypeRef]()

	if span == nil {
		return begin, end, nil
	}

	beginCustom, endCustom := span()

	if beginCustom < begin || endCustom > end {
		return 0, 0, ErrInvalidCustomSpan
	}

	return beginCustom, endCustom, nil
}
