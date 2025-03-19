package inspect

import (
	"github.com/akramarenkov/safe/internal/inspect/constraints"

	"github.com/akramarenkov/intspec"
)

// Options of inspecting. A inspected and reference functions must be specified.
type Opts[TypeFrom, TypeTo constraints.UpToI32, TypeRef constraints.IF64] struct {
	// Number of nested loops, i.e. the number of generated arguments for reference and
	// inspected functions. A value greater than three is not recommended due to low
	// performance
	LoopsQuantity uint

	// Optional function that customize arg values span
	Span func() (TypeFrom, TypeFrom)

	// Inspected function
	Inspected[TypeFrom, TypeTo]

	// Function that returns a reference value
	Reference[TypeRef]
}

type inspector[TypeFrom, TypeTo constraints.UpToI32, TypeRef constraints.IF64] struct {
	opts Opts[TypeFrom, TypeTo, TypeRef]

	// Minimum and maximum value for specified TypeTo type
	minimum TypeRef
	maximum TypeRef

	// Buffers used to decrease allocations
	argsFrom []TypeFrom
	argsRef  []TypeRef

	// Result of inspecting
	result Result[TypeFrom, TypeTo, TypeRef]
}

func (opts Opts[TypeFrom, TypeTo, TypeRef]) isValid() error {
	if opts.Inspected == nil {
		return ErrInspectedNotSpecified
	}

	if opts.Reference == nil {
		return ErrReferenceNotSpecified
	}

	return nil
}

// Performs inspection.
func Do[TypeFrom, TypeTo constraints.UpToI32, TypeRef constraints.IF64](
	opts Opts[TypeFrom, TypeTo, TypeRef],
) (Result[TypeFrom, TypeTo, TypeRef], error) {
	if err := opts.isValid(); err != nil {
		return Result[TypeFrom, TypeTo, TypeRef]{}, err
	}

	minimum, maximum := intspec.Range[TypeTo]()

	insp := &inspector[TypeFrom, TypeTo, TypeRef]{
		opts: opts,

		minimum: TypeRef(minimum),
		maximum: TypeRef(maximum),

		argsFrom: make([]TypeFrom, opts.LoopsQuantity),
		argsRef:  make([]TypeRef, opts.LoopsQuantity),
	}

	insp.main()

	return insp.result, nil
}

func (insp *inspector[TypeFrom, TypeTo, TypeRef]) main() {
	_ = loop[TypeRef](insp.opts.LoopsQuantity, insp.opts.Span, insp.do)
}

func (insp *inspector[TypeFrom, TypeTo, TypeRef]) do(args ...TypeFrom) bool {
	// Protection against changes from the inspected and reference functions
	copy(insp.argsFrom, args)

	for id := range args {
		insp.argsRef[id] = TypeRef(args[id])
	}

	reference, fault := insp.opts.Reference(insp.argsRef...)

	actual, err := insp.opts.Inspected(insp.argsFrom...)

	if fault != nil {
		if err == nil {
			insp.result.Actual = actual
			insp.result.Conclusion = ErrErrorExpected

			insp.result.Args = append([]TypeFrom(nil), args...)

			return true
		}

		insp.result.ReferenceFaults++

		return false
	}

	if reference > insp.maximum || reference < insp.minimum {
		if err == nil {
			insp.result.Actual = actual
			insp.result.Conclusion = ErrErrorExpected
			insp.result.Reference = reference

			insp.result.Args = append([]TypeFrom(nil), args...)

			return true
		}

		insp.result.Overflows++

		return false
	}

	if err != nil {
		insp.result.Conclusion = ErrUnexpectedError
		insp.result.Err = err
		insp.result.Reference = reference

		insp.result.Args = append([]TypeFrom(nil), args...)

		return true
	}

	// An universal, in this case, condition for comparing for inequality
	// (actual != reference), both for integers and for floating point numbers
	if TypeRef(actual)-reference >= 1 || TypeRef(actual)-reference <= -1 {
		insp.result.Actual = actual
		insp.result.Conclusion = ErrNotEqual
		insp.result.Reference = reference

		insp.result.Args = append([]TypeFrom(nil), args...)

		return true
	}

	insp.result.NoOverflows++

	return false
}

func loop[TypeRef constraints.IF64, TypeFrom constraints.UpToI32](
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

func getSpan[TypeFrom constraints.UpToI32, TypeRef constraints.IF64](
	span func() (TypeFrom, TypeFrom),
) (TypeRef, TypeRef) {
	if span == nil {
		minimum, maximum := intspec.Range[TypeFrom]()
		return TypeRef(minimum), TypeRef(maximum)
	}

	begin, end := span()

	return TypeRef(begin), TypeRef(end)
}
