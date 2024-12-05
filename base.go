package safe

import (
	"github.com/akramarenkov/safe/is"
	"golang.org/x/exp/constraints"
)

// Adds two integers and detects whether an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func Add[Type constraints.Integer](first, second Type) (Type, error) {
	sum := first + second

	// When adding or subtracting two integers, only one times overflow is possible
	// When adding signed numbers of different signs, overflow is impossible

	// When overflowing, the sum turns out to be less than both terms, so you can
	// compare it with any of them. For an illustration of overflow when adding
	// unsigned numbers, see internal/research/add. For signed numbers it turns out a
	// little more complicated as indicated in the conditions below

	if sum < first {
		if second > 0 {
			return 0, ErrOverflow
		}

		return sum, nil
	}

	if second < 0 {
		return 0, ErrOverflow
	}

	return sum, nil
}

// Adds two unsigned integers and detects whether an overflow has occurred or not.
//
// Faster than the [Add] function about 20%.
//
// In case of overflow, an error is returned.
func AddU[Type constraints.Unsigned](first, second Type) (Type, error) {
	sum := first + second

	// When adding or subtracting two integers, only one times overflow is possible

	if sum < first {
		return 0, ErrOverflow
	}

	return sum, nil
}

// Subtracts two integers (subtrahend from minuend) and detects whether an overflow
// has occurred or not.
//
// In case of overflow, an error is returned.
func Sub[Type constraints.Integer](minuend, subtrahend Type) (Type, error) {
	diff := minuend - subtrahend

	// When adding or subtracting two integers, only one times overflow is possible
	// When subtracting signed numbers of the same sign, overflow is impossible

	if diff > minuend {
		if subtrahend > 0 {
			return 0, ErrOverflow
		}

		return diff, nil
	}

	if subtrahend < 0 {
		return 0, ErrOverflow
	}

	return diff, nil
}

// Subtracts two unsigned integers (subtrahend from minuend) and detects whether an
// overflow has occurred or not.
//
// Faster than the [Sub] function about 10%.
//
// In case of overflow, an error is returned.
func SubU[Type constraints.Unsigned](minuend, subtrahend Type) (Type, error) {
	diff := minuend - subtrahend

	// When adding or subtracting two integers, only one times overflow is possible

	if diff > minuend {
		return 0, ErrOverflow
	}

	return diff, nil
}

// Multiplies two integers and detects whether an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func Mul[Type constraints.Integer](first, second Type) (Type, error) {
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
	if is.Min(first) && second < 0 {
		return 0, ErrOverflow
	}

	product := first * second

	// It would be possible to represent the multiplication as an addition in a loop
	// and check the sum at each iteration, but this is much slower
	if product/second != first {
		return 0, ErrOverflow
	}

	return product, nil
}

// Divides two integers (dividend to divisor) and detects whether an overflow has
// occurred or not.
//
// The divisor is also checked for equality to zero.
//
// In case of overflow or divisor equal to zero, an error is returned.
func Div[Type constraints.Integer](dividend, divisor Type) (Type, error) {
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
	if is.Min(quotient) && divisor < 0 {
		return 0, ErrOverflow
	}

	return quotient, nil
}

// Changes the sign of a integer and detects whether an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func Negate[Type constraints.Integer](number Type) (Type, error) {
	if is.Min(number) {
		return 0, ErrOverflow
	}

	negated := -number

	if number|negated > 0 { // number > 0 && negated > 0
		return 0, ErrOverflow
	}

	return negated, nil
}

// Converts an integer of one type to an integer of another type and detects whether
// an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func IToI[TypeTo, TypeFrom constraints.Integer](number TypeFrom) (TypeTo, error) {
	converted := TypeTo(number)

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
	reverted := TypeFrom(converted)

	if reverted != number {
		return 0, ErrOverflow
	}

	return converted, nil
}

// Converts an integer to a floating point number and detects whether loss of
// precision has occurred or not.
//
// Loss of precision can lead to overflow when converting back to an integer number.
//
// In case of precision is lost, an error is returned.
func IToF[Flt constraints.Float, Int constraints.Integer](number Int) (Flt, error) {
	converted := Flt(number)
	reverted := Int(converted)

	if number != reverted {
		return 0, ErrPrecisionLoss
	}

	return converted, nil
}

// Converts a floating point number to an integer and detects whether an overflow
// has occurred or not.
//
// Number is also checked for equality to NaN.
//
// In case of overflow or number is equality to NaN, an error is returned.
func FToI[Int constraints.Integer, Flt constraints.Float](number Flt) (Int, error) {
	// It was not possible to find cases where, in the absence of overflow, the
	// difference between a number with a fractional part and its integer part would
	// exceed or equal 1. However, to guarantee the absence of false overflow
	// detection, a difference of 2 was chosen. In the case of real overflow,
	// the difference is always greater.
	const absenceOverflowDiff = 2

	if number != number { //nolint:gocritic
		return 0, ErrNaN
	}

	converted := Int(number)
	reverted := Flt(converted)

	switch {
	case number-reverted > absenceOverflowDiff:
		return 0, ErrOverflow
	case number-reverted < -absenceOverflowDiff:
		return 0, ErrOverflow
	}

	return converted, nil
}

// Shifts an integer left to specified shift count and detects whether an overflow
// has occurred or not.
//
// Shift count is also checked for negativity.
//
// In case of overflow or shift count is negative, an error is returned.
func Shift[Type, CountType constraints.Integer](number Type, count CountType) (Type, error) {
	if count < 0 {
		return 0, ErrNegativeShift
	}

	shifted := number << count
	reverted := shifted >> count

	if reverted != number {
		return 0, ErrOverflow
	}

	return shifted, nil
}
