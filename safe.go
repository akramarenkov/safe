package safe

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var (
	ErrValueOverflow = errors.New("value overflow")
)

func SumInt[Type constraints.Integer](first Type, second Type) (Type, error) {
	sum := first + second

	switch {
	case first > 0 && second > 0:
		if sum < first {
			return 0, ErrValueOverflow
		}
	case first < 0 && second < 0:
		if sum > first {
			return 0, ErrValueOverflow
		}
	}

	return sum, nil
}

func IsMaxNegative[Type constraints.Integer](number Type) bool {
	if number >= 0 {
		return false
	}

	number--

	return number >= 0
}

func IsMaxPositive[Type constraints.Integer](number Type) bool {
	if number <= 0 {
		return false
	}

	number++

	return number <= 0
}

func ProductInt[Type constraints.Integer](first Type, second Type) (Type, error) {
	if second == 0 {
		return 0, nil
	}

	if IsMaxNegative(first) && second < 0 {
		return 0, ErrValueOverflow
	}

	product := first * second

	if product/second != first {
		return 0, ErrValueOverflow
	}

	return product, nil
}

func FloatToInt[
	Float constraints.Float,
	Integer constraints.Integer,
](float Float) (Integer, error) {
	converted := Integer(float)
	reverted := Float(converted)

	if reverted > float && IsMaxNegative(converted) {
		return 0, ErrValueOverflow
	}

	if reverted < float && IsMaxPositive(converted) {
		return 0, ErrValueOverflow
	}

	if reverted > float+1 {
		return 0, ErrValueOverflow
	}

	if reverted < float-1 {
		return 0, ErrValueOverflow
	}

	return converted, nil
}

func Invert[Type constraints.Signed](number Type) (Type, error) {
	if IsMaxNegative(number) {
		return 0, ErrValueOverflow
	}

	return -number, nil
}

func PowUnsigned[Type constraints.Unsigned](base Type, exponent Type) (Type, error) {
	if base == 0 {
		return 0, nil
	}

	if base == 1 {
		return 1, nil
	}

	if exponent == 0 {
		return 1, nil
	}

	powered := base

	for stage := Type(1); stage < exponent; stage++ {
		product, err := ProductInt(powered, base)
		if err != nil {
			return 0, err
		}

		powered = product
	}

	return powered, nil
}

func UnsignedToSigned[
	Unsigned constraints.Unsigned,
	Signed constraints.Signed,
](number Unsigned) (Signed, error) {
	converted := Signed(number)

	if converted < 0 {
		return 0, ErrValueOverflow
	}

	reverted := Unsigned(converted)

	if reverted != number {
		return 0, ErrValueOverflow
	}

	return converted, nil
}
