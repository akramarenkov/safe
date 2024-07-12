package safe

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var (
	ErrDivisionByZero = errors.New("division by zero")
	ErrOverflow       = errors.New("overflow")
)

var pow10table = [...]uint64{ //nolint:gochecknoglobals
	1e00, 1e01, 1e02, 1e03, 1e04, 1e05, 1e06, 1e07, 1e08, 1e09,
	1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18, 1e19,
}

// Adds two integers and determines whether an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func Add[Type constraints.Integer](first Type, second Type) (Type, error) {
	sum := first + second

	// When adding or subtracting, only one times overflow is possible

	// When overflowing, the sum turns out to be less than both terms, so you can
	// compare it with any of them. For an illustration of overflow when adding
	// unsigned numbers, see add.svg. For signed numbers it turns out a little more
	// complicated as indicated in the conditions below

	// When adding positive and negative number, overflow is impossible
	switch {
	case first > 0 && second > 0:
		if sum < first {
			return 0, ErrOverflow
		}
	case first < 0 && second < 0:
		if sum > first {
			return 0, ErrOverflow
		}
	}

	return sum, nil
}

// Adds two unsigned integers and determines whether an overflow has occurred or not.
//
// Slightly faster than the Add function.
//
// In case of overflow, an error is returned.
func AddU[Type constraints.Unsigned](first Type, second Type) (Type, error) {
	sum := first + second

	// When adding or subtracting, only one times overflow is possible

	if sum < first {
		return 0, ErrOverflow
	}

	return sum, nil
}

// Subtracts two integers (subtrahend from minuend) and determines whether an overflow
// has occurred or not.
//
// In case of overflow, an error is returned.
func Sub[Type constraints.Integer](minuend Type, subtrahend Type) (Type, error) {
	diff := minuend - subtrahend

	// When adding or subtracting, only one times overflow is possible

	switch {
	case subtrahend > 0:
		if diff > minuend {
			return 0, ErrOverflow
		}
	case subtrahend < 0:
		if diff < minuend {
			return 0, ErrOverflow
		}
	}

	return diff, nil
}

// Subtracts two unsigned integers (subtrahend from minuend) and determines whether an
// overflow has occurred or not.
//
// Slightly faster than the Sub function.
//
// In case of overflow, an error is returned.
func SubU[Type constraints.Unsigned](minuend Type, subtrahend Type) (Type, error) {
	diff := minuend - subtrahend

	// When adding or subtracting, only one times overflow is possible

	if diff > minuend {
		return 0, ErrOverflow
	}

	return diff, nil
}

// Multiplies two integers and determines whether an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func Mul[Type constraints.Integer](first Type, second Type) (Type, error) {
	if second == 0 {
		return 0, nil
	}

	// When multiplying, many times overflows are possible

	// When multiplying the minimum negative value by -1, an overflow occurs resulting
	// in the same minimum negative value. The division operation with which overflow
	// is checked in this case is also performed with exactly the same overflow.
	// Therefore, this case is checked separately. Since the constraints in the type
	// definition are equal to constraints.Integer i.e. Signed | Unsigned, then a
	// simple check for equality of the second -1 fails, therefore second is checked
	// for a negative value
	if isMin(first) && second < 0 {
		return 0, ErrOverflow
	}

	product := first * second

	// It would be possible to represent the multiplication as an addition in a loop
	// and check the sum at each iteration, but this is slower
	if product/second != first {
		return 0, ErrOverflow
	}

	return product, nil
}

// Divides two integers and determines whether an overflow has occurred or not.
//
// The divisor is also checked for equality to zero.
//
// In case of overflow or the equality of the divisor to zero, an error is returned.
func Div[Type constraints.Integer](dividend Type, divisor Type) (Type, error) {
	if divisor == 0 {
		return 0, ErrDivisionByZero
	}

	quotient := dividend / divisor

	// The only time division overflow occurs is when the dividend is equal to the
	// minimum negative value and the divisor is -1. In this case, the dividend simply
	// changes sign and the quotient becomes equal to the maximum positive value +1 i.e.
	// due to overflow - minimum negative value. Since the constraints in the type
	// definition are equal to constraints.Integer i.e. Signed | Unsigned, then a simple
	// check for equality of the divisor -1 fails, therefore divisor is checked for a
	// negative value
	if isMin(quotient) && divisor < 0 {
		return 0, ErrOverflow
	}

	return quotient, nil
}

// Changes the sign of a signed integer and determines whether an overflow has
// occurred or not.
//
// In case of overflow, an error is returned.
func Negate[Type constraints.Signed](number Type) (Type, error) {
	if isMin(number) {
		return 0, ErrOverflow
	}

	return -number, nil
}

// Converts an integer of one type to an integer of another type and determines whether
// an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func IToI[TypeS constraints.Integer, TypeD constraints.Integer](
	number TypeS,
) (TypeD, error) {
	converted := TypeD(number)

	switch {
	case converted < 0:
		if number > 0 {
			return 0, ErrOverflow
		}
	case converted > 0:
		if number < 0 {
			return 0, ErrOverflow
		}
	}

	// When converting from a variable with a larger bit width to a variable with a
	// smaller bit width, multiple overflows are possible
	reverted := TypeS(converted)

	if reverted != number {
		return 0, ErrOverflow
	}

	return converted, nil
}

// Converts an unsigned integer to a signed integer and determines whether an overflow
// has occurred or not.
//
// Slightly faster than the IToI function.
//
// In case of overflow, an error is returned.
func UToS[Uns constraints.Unsigned, Sgn constraints.Signed](number Uns) (Sgn, error) {
	converted := Sgn(number)

	if converted < 0 {
		return 0, ErrOverflow
	}

	// When converting from a variable with a larger bit width to a variable with a
	// smaller bit width, multiple overflows are possible
	reverted := Uns(converted)

	if reverted != number {
		return 0, ErrOverflow
	}

	return converted, nil
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

	return IToI[uint64, Type](pow10table[power])
}

// Converts an integer to a floating point number and determines whether loss of
// precision has occurred or not.
//
// Loss of precision can lead to overflow when converting back to an integer number.
//
// Returns true if precision is lost.
func IToF[Int constraints.Integer, Flt constraints.Float](number Int) (Flt, bool) {
	converted := Flt(number)
	reverted := Int(converted)

	return converted, number != reverted
}

// Converts a floating point number to an integer and determines whether an overflow
// has occurred or not.
//
// In case of overflow, an error is returned.
func FToI[Flt constraints.Float, Int constraints.Integer](number Flt) (Int, error) {
	converted := Int(number)
	reverted := Flt(converted)

	switch {
	case number-reverted >= 1:
		return 0, ErrOverflow
	case number-reverted <= -1:
		return 0, ErrOverflow
	}

	return converted, nil
}
