package inspect

func Conversion[TypeFrom, TypeTo UpToSixteenBits](
	inspected func(number TypeFrom) (TypeTo, error),
) (Result[TypeFrom, TypeTo], error) {
	if inspected == nil {
		return Result[TypeFrom, TypeTo]{}, ErrInspectedNotSpecified
	}

	return conversion(inspected), nil
}

func conversion[TypeFrom, TypeTo UpToSixteenBits](
	inspected func(number TypeFrom) (TypeTo, error),
) Result[TypeFrom, TypeTo] {
	result := Result[TypeFrom, TypeTo]{}

	minFrom, maxFrom := PickUpRange[TypeFrom]()
	minTo, maxTo := PickUpRange[TypeTo]()

	for reference := minFrom; reference <= maxFrom; reference++ {
		arg := TypeFrom(reference)

		actual, err := inspected(arg)

		if reference > maxTo || reference < minTo {
			if err == nil {
				result.Actual = actual
				result.Args = []TypeFrom{arg}
				result.Conclusion = ErrErrorExpected
				result.Reference = reference

				return result
			}

			result.Overflows++

			continue
		}

		if err != nil {
			result.Args = []TypeFrom{arg}
			result.Conclusion = ErrUnexpectedError
			result.Err = err
			result.Reference = reference

			return result
		}

		if int64(actual) != reference {
			result.Actual = actual
			result.Args = []TypeFrom{arg}
			result.Conclusion = ErrNotEqual
			result.Reference = reference

			return result
		}

		result.NoOverflows++
	}

	return result
}
