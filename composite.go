package safe

import (
	"github.com/akramarenkov/safe/internal/is"

	"github.com/akramarenkov/intspec"
	"golang.org/x/exp/constraints"
)

// Calculates the value of the expression first + second - subtrahend and detects
// whether an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func AddSub[Type constraints.Integer](first, second, subtrahend Type) (Type, error) {
	if interim, err := Add(first, second); err == nil {
		return Sub(interim, subtrahend)
	}

	if interim, err := Sub(first, subtrahend); err == nil {
		if sum, err := Add(interim, second); err == nil {
			return sum, nil
		}
	}

	if interim, err := Sub(second, subtrahend); err == nil {
		if sum, err := Add(interim, first); err == nil {
			return sum, nil
		}
	}

	interim, err := Sub(subtrahend, first)
	if err != nil {
		return 0, err
	}

	return Sub(second, interim)
}

// Calculates the quotient of dividing the sum of two integers by divisor and
// detects whether an overflow has occurred or not.
//
// In case of overflow or divisor equal to zero, an error is returned.
func AddDiv[Type constraints.Integer](first, second, divisor Type) (Type, error) {
	if sum, err := Add(first, second); err == nil {
		return Div(sum, divisor)
	}

	if divisor == 0 {
		return 0, ErrDivisionByZero
	}

	// If overflow occurs during addition, the addition arguments have the same signs

	minimum, maximum := intspec.Range[Type]()

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
func AddDivRem[Type constraints.Integer](first, second, divisor Type) (Type, error) {
	if divisor == 0 {
		return 0, ErrDivisionByZero
	}

	if sum, err := Add(first, second); err == nil {
		return sum % divisor, nil
	}

	minimum, maximum := intspec.Range[Type]()

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
// detects whether an overflow has occurred or not.
//
// Not faster than the [AddDiv] function.
//
// In case of overflow or divisor equal to zero, an error is returned.
func AddDivU[Type constraints.Unsigned](first, second, divisor Type) (Type, error) {
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
	//
	// complement := second - excess
	// maximum := first + complement
	maximum := ^Type(0)

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
// detects whether an overflow has occurred or not.
//
// In case of overflow or divisor equal to zero, an error is returned.
func SubDiv[Type constraints.Integer](minuend, subtrahend, divisor Type) (Type, error) {
	if diff, err := Sub(minuend, subtrahend); err == nil {
		return Div(diff, divisor)
	}

	if divisor == 0 {
		return 0, ErrDivisionByZero
	}

	// If an overflow occurs during subtraction and the subtraction arguments are
	// signed, then they have different signs

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

	minimum, maximum := intspec.Range[Type]()

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

	// Excess is overflowed, sign inversion performed
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
// detects whether an overflow has occurred or not.
//
// In case of overflow or divisor equal to zero, an error is returned.
func SubDivRem[Type constraints.Integer](minuend, subtrahend, divisor Type) (Type, error) {
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

	minimum, maximum := intspec.Range[Type]()

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
// divisor and detects whether an overflow has occurred or not.
//
// Faster than the [SubDiv] function about 30%.
//
// In case of overflow or divisor equal to zero, an error is returned.
func SubDivU[Type constraints.Unsigned](minuend, subtrahend, divisor Type) (Type, error) {
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

// Calculates the quotient of dividing of the expression first + second - subtrahend by
// divisor and detects whether an overflow has occurred or not.
//
// In case of overflow or divisor equal to zero, an error is returned.
func AddSubDiv[Type constraints.Integer](first, second, subtrahend, divisor Type) (Type, error) {
	if interim, err := AddSub(first, second, subtrahend); err == nil {
		return Div(interim, divisor)
	}

	if interim, err := Add(first, second); err == nil {
		return SubDiv(interim, subtrahend, divisor)
	}

	if interim, err := Sub(subtrahend, second); err == nil {
		return SubDiv(first, interim, divisor)
	}

	qm, err := SubDiv(first, subtrahend, divisor)
	if err != nil {
		return 0, err
	}

	rm, _ := SubDivRem(first, subtrahend, divisor)

	qe, err := Div(second, divisor)
	if err != nil {
		return 0, err
	}

	re := second % divisor

	remainder, _ := AddDiv(rm, re, divisor)

	return Add3(qm, qe, remainder)
}

// Calculates the quotient of dividing of the expression minuend + 1 - subtrahend by
// divisor and detects whether an overflow has occurred or not.
//
// Faster than the [AddSubDiv] function about 15%.
//
// In case of overflow or divisor equal to zero, an error is returned.
func AddOneSubDiv[Type constraints.Integer](minuend, subtrahend, divisor Type) (Type, error) {
	if interim, err := Add(minuend, 1); err == nil {
		return SubDiv(interim, subtrahend, divisor)
	}

	if interim, err := Sub(subtrahend, 1); err == nil {
		return SubDiv(minuend, interim, divisor)
	}

	qm, err := SubDiv(minuend, subtrahend, divisor)
	if err != nil {
		return 0, err
	}

	rm, _ := SubDivRem(minuend, subtrahend, divisor)

	qe := 1 / divisor
	re := 1 % divisor

	remainder, _ := AddDiv(rm, re, divisor)

	return Add3(qm, qe, remainder)
}
