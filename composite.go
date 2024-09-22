package safe

import (
	"github.com/akramarenkov/safe/internal/intspan"
	"github.com/akramarenkov/safe/internal/is"
	"golang.org/x/exp/constraints"
)

// Calculates the quotient of dividing the sum of two integers by divisor and
// determines whether an overflow has occurred or not.
//
// In case of overflow or divisor equal to zero, an error is returned.
func AddDiv[Type constraints.Integer](first Type, second Type, divisor Type) (Type, error) {
	if sum, err := Add(first, second); err == nil {
		return Div(sum, divisor)
	}

	if divisor == 0 {
		return 0, ErrDivisionByZero
	}

	// If overflow occurs during addition, the addition arguments have the same signs

	minimum, maximum, _ := intspan.Get[Type]()

	overflowed := first + second

	// The idea is to complement one of the arguments being added to the maximum
	// (if the arguments are positive) or minimum (if the arguments are negative)
	// values for a given type ​​so that the division result is never zero. Then divide
	// separately the maximum/minimum value and the excess that arose during overflow,
	// add the quotients from these divisions and add the quotient from dividing the
	// sum of the remainders from the two previous divisions

	// Since the arguments of addition have the same signs, it is sufficient to check
	// either of them
	if first > 0 {
		// If the divisor is equal to the minimum (negative) value for a given type,
		// then the approach of complementing one of the arguments to the maximum
		// (positive) value stops working because the minimum value is greater than
		// the maximum in absolute value and the division results in zero, and the
		// remainder is equal to the dividend
		if is.Min(divisor) {
			// If the divisor is equal to the minimum (negative) value for the given
			// type, then when dividing the sum of two positive arguments (in case of
			// overflow during addition), the result is always -1
			return -Type(1), nil
		}

		excess := -(minimum - overflowed) + 1

		qm := maximum / divisor
		rm := maximum % divisor

		qe := excess / divisor
		re := excess % divisor

		interim, err := Add(qm, qe)
		if err != nil {
			return 0, err
		}

		quotient := interim + (rm+re)/divisor

		return quotient, nil
	}

	excess := -(maximum - overflowed + 1)

	qm, err := Div(minimum, divisor)
	if err != nil {
		return 0, err
	}

	rm := minimum % divisor

	qe := excess / divisor
	re := excess % divisor

	interim, err := Add(qm, qe)
	if err != nil {
		return 0, err
	}

	quotient := interim + (rm+re)/divisor

	return quotient, nil
}

// Calculates the remainder of dividing the sum of two integers by divisor.
//
// In case of divisor equal to zero, an error is returned.
func AddDivRem[Type constraints.Integer](first Type, second Type, divisor Type) (Type, error) {
	if divisor == 0 {
		return 0, ErrDivisionByZero
	}

	if sum, err := Add(first, second); err == nil {
		return sum % divisor, nil
	}

	minimum, maximum, _ := intspan.Get[Type]()

	overflowed := first + second

	if first > 0 {
		if is.Min(divisor) {
			return overflowed - divisor, nil
		}

		excess := -(minimum - overflowed) + 1

		rm := maximum % divisor
		re := excess % divisor

		remainder := (rm + re) % divisor

		return remainder, nil
	}

	excess := -(maximum - overflowed + 1)

	rm := minimum % divisor
	re := excess % divisor

	remainder := (rm + re) % divisor

	return remainder, nil
}

// Calculates the quotient of dividing the sum of two unsigned integers by divisor and
// determines whether an overflow has occurred or not.
//
// Slightly faster than the [AddDiv] function.
//
// In case of overflow or divisor equal to zero, an error is returned.
func AddDivU[Type constraints.Unsigned](first Type, second Type, divisor Type) (Type, error) {
	if divisor == 0 {
		return 0, ErrDivisionByZero
	}

	if sum, err := AddU(first, second); err == nil {
		return sum / divisor, nil
	}

	overflowed := first + second
	excess := overflowed + 1

	// For unsigned types, the maximum value for the type in case of overflow can be
	// calculated
	complement := second - excess
	maximum := first + complement

	qm := maximum / divisor
	rm := maximum % divisor

	qe := excess / divisor
	re := excess % divisor

	interim, err := AddU(qm, qe)
	if err != nil {
		return 0, err
	}

	quotient := interim + (rm+re)/divisor

	return quotient, nil
}

// Calculates the quotient of dividing the difference of two integers by divisor and
// determines whether an overflow has occurred or not.
//
// In case of overflow or divisor equal to zero, an error is returned.
func SubDiv[Type constraints.Integer](minuend Type, subtrahend Type, divisor Type) (Type, error) {
	if diff, err := Sub(minuend, subtrahend); err == nil {
		return Div(diff, divisor)
	}

	if divisor == 0 {
		return 0, ErrDivisionByZero
	}

	// When subtracting unsigned arguments, overflow cannot be compensated by
	// division because the result of the subtraction and subsequent division must be
	// negative. Except for one case when the divisor is greater than the difference and
	// the division result is zero
	if !is.Signed[Type]() {
		excess := subtrahend - minuend

		if excess/divisor == 0 {
			return 0, nil
		}

		return 0, ErrOverflow
	}

	// If an overflow occurs during subtraction and the subtraction arguments are
	// signed, then they have different signs

	minimum, maximum, _ := intspan.Get[Type]()

	overflowed := minuend - subtrahend

	if subtrahend > 0 {
		excess := -(maximum - overflowed + 1)

		qm, err := Div(minimum, divisor)
		if err != nil {
			return 0, err
		}

		rm := minimum % divisor

		qe := excess / divisor
		re := excess % divisor

		interim, err := Add(qm, qe)
		if err != nil {
			return 0, err
		}

		quotient := interim + (rm+re)/divisor

		return quotient, nil
	}

	if is.Min(divisor) {
		return -Type(1), nil
	}

	// Overflow is possible in these calculations, but only in one case: if
	// minuend is equal to the maximum value for the given type, and subtrahend is
	// equal to the minimum value for the given type. In this case, the result of
	// the calculation is equal to the minimum value for the given type. This value
	// will be used in the division, and the result of the division will be inverted
	excess := -(minimum - overflowed) + 1

	qm := maximum / divisor
	rm := maximum % divisor

	qe := excess / divisor
	re := excess % divisor

	// excess is overflowed, sign inversion performed
	if excess < 0 {
		negated, err := Negate(qe)
		if err != nil {
			return 0, err
		}

		qe = negated
		re = -re
	}

	interim, err := Add(qm, qe)
	if err != nil {
		return 0, err
	}

	quotient := interim + (rm+re)/divisor

	return quotient, nil
}

// Calculates the remainder of dividing the difference of two integers by divisor and
// determines whether an overflow has occurred or not.
//
// In case of overflow or divisor equal to zero, an error is returned.
func SubDivRem[Type constraints.Integer](minuend Type, subtrahend Type, divisor Type) (Type, error) {
	if divisor == 0 {
		return 0, ErrDivisionByZero
	}

	if diff, err := Sub(minuend, subtrahend); err == nil {
		return diff % divisor, nil
	}

	if !is.Signed[Type]() {
		excess := subtrahend - minuend

		if excess%divisor == 0 {
			return 0, nil
		}

		return 0, ErrOverflow
	}

	minimum, maximum, _ := intspan.Get[Type]()

	overflowed := minuend - subtrahend

	if subtrahend > 0 {
		excess := -(maximum - overflowed + 1)

		rm := minimum % divisor
		re := excess % divisor

		remainder := (rm + re) % divisor

		return remainder, nil
	}

	if is.Min(divisor) {
		return overflowed - divisor, nil
	}

	excess := -(minimum - overflowed) + 1

	rm := maximum % divisor
	re := excess % divisor

	if excess < 0 {
		re = -re
	}

	remainder := (rm + re) % divisor

	return remainder, nil
}

// Calculates the quotient of dividing the difference of two unsigned integers by
// divisor and determines whether an overflow has occurred or not.
//
// Slightly faster than the [SubDiv] function.
//
// In case of overflow or divisor equal to zero, an error is returned.
func SubDivU[Type constraints.Unsigned](minuend Type, subtrahend Type, divisor Type) (Type, error) {
	if divisor == 0 {
		return 0, ErrDivisionByZero
	}

	if diff, err := SubU(minuend, subtrahend); err == nil {
		return diff / divisor, nil
	}

	excess := subtrahend - minuend

	if excess/divisor == 0 {
		return 0, nil
	}

	return 0, ErrOverflow
}
