package safe

import (
	"slices"

	"golang.org/x/exp/constraints"
)

var pow10table = [...]uint64{ //nolint:gochecknoglobals
	1e00, 1e01, 1e02, 1e03, 1e04, 1e05, 1e06, 1e07, 1e08, 1e09,
	1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18, 1e19,
}

// Adds up multiple integers and determines whether an overflow has occurred or not.
//
// Slower than the Add, AddT functions.
//
// In case of overflow or missing arguments, an error is returned.
func AddM[Type constraints.Integer](addends ...Type) (Type, error) {
	//nolint:mnd
	switch len(addends) {
	case 0:
		return 0, ErrMissinArguments
	case 1:
		return addends[0], nil
	case 2:
		return Add(addends[0], addends[1])
	}

	for len(addends) != 3 {
		sortAddM(addends)

		interim, err := Add(addends[0], addends[len(addends)-1])
		if err != nil {
			return 0, err
		}

		addends[0] = interim
		addends = addends[:len(addends)-1]
	}

	return AddT(addends[0], addends[1], addends[2])
}

func sortAddM[Type constraints.Integer](addends []Type) {
	for first := 1; first < len(addends); first++ {
		for second := first; second > 0 && addends[second] < addends[second-1]; second-- {
			addends[second], addends[second-1] = addends[second-1], addends[second]
		}
	}
}

// Adds three integers and determines whether an overflow has occurred or not.
//
// Faster than the AddM function.
//
// In case of overflow, an error is returned.
func AddT[Type constraints.Integer](first, second, third Type) (Type, error) {
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
// Slower than the AddU function, faster than the AddM function.
//
// In case of overflow or missing arguments, an error is returned.
func AddUM[Type constraints.Unsigned](addends ...Type) (Type, error) {
	if len(addends) == 0 {
		return 0, ErrMissinArguments
	}

	sum := addends[0]

	for _, addend := range addends[1:] {
		interim, err := AddU(sum, addend)
		if err != nil {
			return 0, err
		}

		sum = interim
	}

	return sum, nil
}

// Subtracts three integers (subtrahend, deductible from minuend) and determines whether
// an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func SubT[Type constraints.Integer](minuend, subtrahend, deductible Type) (Type, error) {
	interim, err := Sub(minuend, subtrahend)
	if err == nil {
		diff, err := Sub(interim, deductible)
		if err == nil {
			return diff, nil
		}
	}

	interim, err = Sub(minuend, deductible)
	if err == nil {
		diff, err := Sub(interim, subtrahend)
		if err == nil {
			return diff, nil
		}
	}

	interim, err = Add(subtrahend, deductible)
	if err == nil {
		return Sub(minuend, interim)
	}

	return 0, ErrOverflow
}

// Subtracts multiple unsigned integers (subtrahends from minuend) and determines
// whether an overflow has occurred or not.
//
// Slower than the SubU function.
//
// In case of overflow, an error is returned.
func SubUM[Type constraints.Unsigned](minuend Type, subtrahends ...Type) (Type, error) {
	diff := minuend

	for _, subtrahend := range subtrahends {
		interim, err := SubU(diff, subtrahend)
		if err != nil {
			return 0, err
		}

		diff = interim
	}

	return diff, nil
}

// Multiplies multiple integers and determines whether an overflow has occurred or not.
//
// Slower than the Mul function.
//
// In case of overflow or missing arguments, an error is returned.
func MulM[Type constraints.Integer](factors ...Type) (Type, error) {
	if len(factors) == 0 {
		return 0, ErrMissinArguments
	}

	slices.SortFunc(factors, cmpMulM)

	for _, factor := range factors {
		if factor == 0 {
			return 0, nil
		}
	}

	product := factors[0]

	for _, factor := range factors[1:] {
		interim, err := Mul(product, factor)
		if err != nil {
			return 0, err
		}

		product = interim
	}

	return product, nil
}

func cmpMulM[Type constraints.Integer](first, second Type) int {
	if first < 0 && second < 0 {
		switch {
		case first > second:
			return -1
		case first < second:
			return 1
		}

		return 0
	}

	switch {
	case first < second:
		return -1
	case first > second:
		return 1
	}

	return 0
}

// Multiplies three integers and determines whether an overflow has occurred or not.
//
// Faster than the MulM function.
//
// In case of overflow, an error is returned.
func MulT[Type constraints.Integer](first, second, third Type) (Type, error) {
	interim, err := Mul(first, second)
	if err == nil {
		product, err := Mul(interim, third)
		if err == nil {
			return product, nil
		}
	}

	interim, err = Mul(first, third)
	if err == nil {
		product, err := Mul(interim, second)
		if err == nil {
			return product, nil
		}
	}

	interim, err = Mul(second, third)
	if err == nil {
		return Mul(first, interim)
	}

	return 0, ErrOverflow
}

// Multiplies multiple unsigned integers and determines whether an overflow has
// occurred or not.
//
// Faster than the MulM function.
//
// In case of overflow or missing arguments, an error is returned.
func MulUM[Type constraints.Unsigned](factors ...Type) (Type, error) {
	if len(factors) == 0 {
		return 0, ErrMissinArguments
	}

	product := factors[0]
	factors = factors[1:]

	for _, factor := range factors {
		if factor == 0 {
			return 0, nil
		}
	}

	for _, factor := range factors {
		interim, err := Mul(product, factor)
		if err != nil {
			return 0, err
		}

		product = interim
	}

	return product, nil
}

// Divides multiple integers (dividend to divisors) and determines whether an overflow
// has occurred or not.
//
// The divisors is also checked for equality to zero.
//
// Slower than the Div function.
//
// In case of overflow or divisors equal to zero, an error is returned.
func DivM[Type constraints.Integer](dividend Type, divisors ...Type) (Type, error) {
	quotient := dividend

	minusOnes := 0

	for _, divisor := range divisors {
		if divisor == 0 {
			return 0, ErrDivisionByZero
		}

		if isMinusOne(divisor) {
			minusOnes++
			continue
		}

		quotient /= divisor
	}

	// Paired minus ones cancel each other out
	if isEven(minusOnes) {
		return quotient, nil
	}

	if isMin(quotient) {
		return 0, ErrOverflow
	}

	return -quotient, nil
}

// Raises 10 to a power and determines whether an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func Pow10[Type, TypePower constraints.Integer](power TypePower) (Type, error) {
	if power < 0 {
		return 0, nil
	}

	// Value of pow10table length fits into any integer type
	if power >= TypePower(len(pow10table)) {
		return 0, ErrOverflow
	}

	return IToI[Type](pow10table[power])
}

// Raises base to a power and determines whether an overflow has occurred or not.
//
// Straightforward and slow implementation. Be careful.
//
// In case of overflow, an error is returned.
func Pow[Type, TypePower constraints.Integer](base Type, power TypePower) (Type, error) {
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

	for step := TypePower(1); step < power; step++ {
		// overflow must be checked at each multiplication step
		product, err := Mul(powered, base)
		if err != nil {
			return 0, err
		}

		powered = product
	}

	return powered, nil
}
