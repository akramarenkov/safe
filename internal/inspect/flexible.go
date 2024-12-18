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
	types.Inspected[TypeFrom, TypeTo]
	// Function that returns a reference value
	types.Reference[TypeRef]
	// Optional function that customize arg values span
	Span func() (TypeFrom, TypeFrom)

	// Minimum and maximum value for specified TypeTo type
	minimum TypeRef
	maximum TypeRef

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

	opts.minimum, opts.maximum = ConvSpan[TypeTo, TypeRef]()

	opts.argsFrom = make([]TypeFrom, opts.LoopsQuantity)
	opts.argsRef = make([]TypeRef, opts.LoopsQuantity)

	opts.main()

	return opts.result, nil
}

func (opts *Opts[TypeFrom, TypeTo, TypeRef]) main() {
	_ = loop[TypeRef](opts.LoopsQuantity, opts.Span, opts.do)
}

func (opts *Opts[TypeFrom, TypeTo, TypeRef]) do(args ...TypeFrom) bool {
	// Protection against changes from the inspected and reference functions
	copy(opts.argsFrom, args)

	for id := range args {
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

	if reference > opts.maximum || reference < opts.minimum {
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

	// An universal, in this case, condition for comparing for inequality
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
	span func() (TypeFrom, TypeFrom),
	do func(args ...TypeFrom) bool,
	args ...TypeFrom,
) bool {
	if level == 0 {
		return do(args...)
	}

	level--

	args = append(args, 0)

	begin, end := getSpan[TypeFrom, TypeRef](span)

	for number := begin; number <= end; number++ {
		args[len(args)-1] = TypeFrom(number)

		if level == 0 {
			if stop := do(args...); stop {
				return true
			}

			continue
		}

		if stop := loop[TypeRef](level, span, do, args...); stop {
			return true
		}
	}

	return false
}

func getSpan[TypeFrom types.UpToUSI32, TypeRef types.SIF64](
	span func() (TypeFrom, TypeFrom),
) (TypeRef, TypeRef) {
	if span == nil {
		return ConvSpan[TypeFrom, TypeRef]()
	}

	begin, end := span()

	return TypeRef(begin), TypeRef(end)
}
