package inspect

// Function that returns a reference value. Number of arguments is equal to the
// number of nested loops specified in the options.
type Reference func(args ...int64) (int64, error)

// Options of inspecting. A reference and inspected functions must be specified.
type Opts[Type EightBits] struct {
	// Number of nested loops, i.e. the number of generated arguments for reference and
	// inspected functions. A value greater than three is not recommended due to low
	// performance
	LoopsQuantity uint

	// Inspected function
	Inspected func(args ...Type) (Type, error)
	// Function that returns a reference value
	Reference

	// Minimum and maximum value for specified type
	min int64
	max int64
}

// Validates options. A reference and inspected functions must be specified.
func (opts *Opts[Type]) IsValid() error {
	if opts.Reference == nil {
		return ErrReferenceNotSpecified
	}

	if opts.Inspected == nil {
		return ErrInpectedNotSpecified
	}

	return nil
}

// Performs inspection.
func (opts *Opts[Type]) Do() (Result[Type], error) {
	if err := opts.IsValid(); err != nil {
		return Result[Type]{}, err
	}

	opts.min, opts.max = pickUpRange[Type]()

	return opts.main(), nil
}

func (opts *Opts[Type]) main() Result[Type] {
	result := Result[Type]{}

	// Protection against changes from the inspected and reference functions
	args8 := make([]Type, opts.LoopsQuantity)
	args64 := make([]int64, opts.LoopsQuantity)

	do := func(args ...Type) bool {
		copy(args8, args)

		for id := range len(args) {
			args64[id] = int64(args[id])
		}

		reference, fault := opts.Reference(args64...)

		actual, err := opts.Inspected(args8...)

		if fault != nil {
			if err == nil {
				result.Actual = actual
				result.Conclusion = ErrErrorExpected
				result.Reference = reference

				result.Args = append([]Type(nil), args...)

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

				result.Args = append([]Type(nil), args...)

				return true
			}

			result.Overflows++

			return false
		}

		if err != nil {
			result.Conclusion = ErrUnexpectedError
			result.Err = err
			result.Reference = reference

			result.Args = append([]Type(nil), args...)

			return true
		}

		if int64(actual) != reference {
			result.Actual = actual
			result.Conclusion = ErrNotEqual
			result.Reference = reference

			result.Args = append([]Type(nil), args...)

			return true
		}

		result.NoOverflows++

		return false
	}

	_ = loop(opts.LoopsQuantity, do)

	return result
}

func loop[Type EightBits](level uint, do func(args ...Type) bool, args ...Type) bool {
	if level == 0 {
		return do(args...)
	}

	level--

	args = append(args, 0)

	begin, end := pickUpRange[Type]()

	for number := begin; number <= end; number++ {
		args[len(args)-1] = Type(number)

		if level == 0 {
			if stop := do(args...); stop {
				return true
			}

			continue
		}

		if stop := loop(level, do, args...); stop {
			return true
		}
	}

	return false
}
