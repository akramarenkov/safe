package inspect

// Options of inspecting. A reference and inspected functions must be specified.
type Opts[TypeFrom, TypeTo UpTo32Bits, TypeRef SixtyFourBits] struct {
	// Number of nested loops, i.e. the number of generated arguments for reference and
	// inspected functions. A value greater than three is not recommended due to low
	// performance
	LoopsQuantity uint

	// Inspected function
	Inspected func(args ...TypeFrom) (TypeTo, error)
	// Function that returns a reference value
	Reference func(args ...TypeRef) (TypeRef, error)

	// Minimum and maximum value for specified TypeTo type
	min TypeRef
	max TypeRef
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
	Result[TypeFrom, TypeTo, TypeRef],
	error,
) {
	if err := opts.IsValid(); err != nil {
		return Result[TypeFrom, TypeTo, TypeRef]{}, err
	}

	opts.min, opts.max = PickUpRange[TypeTo, TypeRef]()

	return opts.main(), nil
}

func (opts *Opts[TypeFrom, TypeTo, TypeRef]) main() Result[TypeFrom, TypeTo, TypeRef] {
	result := Result[TypeFrom, TypeTo, TypeRef]{}

	// Protection against changes from the inspected and reference functions
	argsFrom := make([]TypeFrom, opts.LoopsQuantity)
	argsRef := make([]TypeRef, opts.LoopsQuantity)

	do := func(args ...TypeFrom) bool {
		copy(argsFrom, args)

		for id := range len(args) {
			argsRef[id] = TypeRef(args[id])
		}

		reference, fault := opts.Reference(argsRef...)

		actual, err := opts.Inspected(argsFrom...)

		if fault != nil {
			if err == nil {
				result.Actual = actual
				result.Conclusion = ErrErrorExpected

				result.Args = append([]TypeFrom(nil), args...)

				return true
			}

			result.ReferenceFaults++

			return false
		}

		if reference > opts.max || reference < opts.min {
			if err == nil {
				result.Actual = actual
				result.Conclusion = ErrErrorExpected
				result.Reference = reference

				result.Args = append([]TypeFrom(nil), args...)

				return true
			}

			result.Overflows++

			return false
		}

		if err != nil {
			result.Conclusion = ErrUnexpectedError
			result.Err = err
			result.Reference = reference

			result.Args = append([]TypeFrom(nil), args...)

			return true
		}

		if TypeRef(actual) != reference {
			result.Actual = actual
			result.Conclusion = ErrNotEqual
			result.Reference = reference

			result.Args = append([]TypeFrom(nil), args...)

			return true
		}

		result.NoOverflows++

		return false
	}

	_ = loop[TypeRef](opts.LoopsQuantity, do)

	return result
}

func loop[TypeRef SixtyFourBits, TypeFrom UpTo32Bits](
	level uint,
	do func(args ...TypeFrom) bool,
	args ...TypeFrom,
) bool {
	if level == 0 {
		return do(args...)
	}

	level--

	args = append(args, 0)

	begin, end := PickUpRange[TypeFrom, TypeRef]()

	for number := begin; number <= end; number++ {
		args[len(args)-1] = TypeFrom(number)

		if level == 0 {
			if stop := do(args...); stop {
				return true
			}

			continue
		}

		if stop := loop[TypeRef](level, do, args...); stop {
			return true
		}
	}

	return false
}
