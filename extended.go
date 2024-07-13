package safe

import (
	"golang.org/x/exp/constraints"
)

var pow10table = [...]uint64{ //nolint:gochecknoglobals
	1e00, 1e01, 1e02, 1e03, 1e04, 1e05, 1e06, 1e07, 1e08, 1e09,
	1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18, 1e19,
}

// Adds three integers and determines whether an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func AddT[Type constraints.Integer](first Type, second Type, third Type) (Type, error) {
	interim, err := Add(first, second)
	if err == nil {
		sum, err := Add(interim, third)
		if err == nil {
			return sum, nil
		}
	}

	interim, err = Add(first, third)
	if err == nil {
		sum, err := Add(interim, second)
		if err == nil {
			return sum, nil
		}
	}

	interim, err = Add(second, third)
	if err == nil {
		return Add(first, interim)
	}

	return 0, ErrOverflow
}

// Adds up multiple unsigned integers and determines whether an overflow has occurred or
// not.
//
// Slower than a function with two arguments.
//
// In case of overflow, an error is returned.
func AddUM[Type constraints.Unsigned](first Type, others ...Type) (Type, error) {
	sum := first

	for _, next := range others {
		interim, err := AddU(sum, next)
		if err != nil {
			return 0, err
		}

		sum = interim
	}

	return sum, nil
}

// Subtracts three integers (subtrahend from minuend) and determines whether an overflow
// has occurred or not.
//
// In case of overflow, an error is returned.
func SubT[Type constraints.Integer](
	minuend Type,
	subtrahend Type,
	secondSubtrahend Type,
) (Type, error) {
	interim, err := Sub(minuend, subtrahend)
	if err == nil {
		diff, err := Sub(interim, secondSubtrahend)
		if err == nil {
			return diff, nil
		}
	}

	interim, err = Sub(minuend, secondSubtrahend)
	if err == nil {
		diff, err := Sub(interim, subtrahend)
		if err == nil {
			return diff, nil
		}
	}

	interim, err = Add(subtrahend, secondSubtrahend)
	if err == nil {
		return Sub(minuend, interim)
	}

	return 0, ErrOverflow
}

// Raises 10 to a power and determines whether an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func Pow10[Type constraints.Integer, TypeP constraints.Integer](
	power TypeP,
) (Type, error) {
	if power < 0 {
		return 0, nil
	}

	// Value of pow10table length fits into any integer type
	if power >= TypeP(len(pow10table)) {
		return 0, ErrOverflow
	}

	return IToI[Type](pow10table[power])
}

// Raises base to a power and determines whether an overflow has occurred or not.
//
// Straightforward and slow implementation. Be careful.
//
// In case of overflow, an error is returned.
func Pow[Type constraints.Integer, TypeP constraints.Integer](
	base Type,
	power TypeP,
) (Type, error) {
	if power == 0 {
		return 1, nil
	}

	if base == 1 {
		return 1, nil
	}

	if base == 0 {
		if power < 0 {
			return 0, ErrDivisionByZero
		}

		return 0, nil
	}

	if power < 0 {
		if isMinusOne(base) {
			if isEven(power) {
				return 1, nil
			}

			return base, nil
		}

		return 0, nil
	}

	powered := base

	for step := TypeP(1); step < power; step++ {
		// overflow must be checked at each multiplication step
		product, err := Mul(powered, base)
		if err != nil {
			return 0, err
		}

		powered = product
	}

	return powered, nil
}
