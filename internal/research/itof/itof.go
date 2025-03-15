// Internal package that research problems of converting integer values to floating
// point values.
package research

import (
	"errors"

	"github.com/akramarenkov/safe"
	"github.com/akramarenkov/safe/internal/consts"

	"golang.org/x/exp/constraints"
)

var (
	ErrBeginGreaterEnd          = errors.New("begin is greater than end")
	ErrLosslesslyLengthNegative = errors.New("losslessly length is negative")
	ErrLosslesslyLengthTooLong  = errors.New("losslessly length is too long for specified type")
)

const (
	maxIntPower = 19
)

// Finds last number in a continuous sequence of integers that convert to floating point
// numbers without loss followed by a number that converts with loss.
//
// It is necessary to specify the required length of a continuous sequence of integers
// that must be converted to floating point numbers without loss.
//
// It is also necessary to indicate the sign (search direction on the number line) of
// the numbers for which you want to search. If the sign is negative, the search will be
// performed for negative numbers, in other cases - for positive numbers.
func FindLastLosslessly[TypeTo constraints.Float, TypeFrom constraints.Integer](
	losslesslyLength TypeFrom,
	sign TypeFrom,
) (TypeFrom, bool, error) {
	if losslesslyLength < 0 {
		return 0, false, ErrLosslesslyLengthNegative
	}

	if sign < 0 {
		return findLastLosslesslyNegative[TypeTo](losslesslyLength, sign)
	}

	return findLastLosslesslyPositive[TypeTo](losslesslyLength, sign)
}

func findLastLosslesslyPositive[TypeTo constraints.Float, TypeFrom constraints.Integer](
	losslesslyLength TypeFrom,
	sign TypeFrom,
) (TypeFrom, bool, error) {
	previous := TypeFrom(0)

	for power := maxIntPower; power >= 0; power-- {
		begun := false

		exponent, err := safe.Pow10[TypeFrom](power)
		if err != nil {
			// Too large a power for a given integer type, try a smaller power
			continue
		}

		for multiplier := TypeFrom(1); multiplier <= consts.DecimalBase; multiplier++ {
			current, err := safe.Mul(multiplier, exponent)
			if err != nil {
				// Too large a number for a given integer type, try a smaller power
				break
			}

			base, err := safe.Add(previous, current)
			if err != nil {
				// Too large a number for a given integer type, try a smaller power
				break
			}

			number, beginning, found, err := findLosslessly[TypeTo](
				base,
				losslesslyLength,
				sign,
			)
			if err != nil {
				return 0, false, err
			}

			if found {
				return number, true, nil
			}

			if beginning {
				begun = true
				continue
			}

			if begun {
				// Since at the current step we managed to multiply the exponent and
				// add/subtract the resulting number with a multiplier equal to
				// 'multiplier' without overflow, then for 'multiplier-1' this will
				// definitely be possible without overflow
				previous += (multiplier - 1) * exponent
				break
			}
		}
	}

	return 0, false, nil
}

func findLastLosslesslyNegative[TypeTo constraints.Float, TypeFrom constraints.Integer](
	losslesslyLength TypeFrom,
	sign TypeFrom,
) (TypeFrom, bool, error) {
	previous := TypeFrom(0)

	for power := maxIntPower; power >= 0; power-- {
		begun := false

		exponent, err := safe.Pow10[TypeFrom](power)
		if err != nil {
			// Too large a power for a given integer type, try a smaller power
			continue
		}

		for multiplier := TypeFrom(1); multiplier <= consts.DecimalBase; multiplier++ {
			current, err := safe.Mul(multiplier, exponent)
			if err != nil {
				// Too large a number for a given integer type, try a smaller power
				break
			}

			base, err := safe.Sub(previous, current)
			if err != nil {
				// Too large a number for a given integer type, try a smaller power
				break
			}

			number, beginning, found, err := findLosslessly[TypeTo](
				base,
				losslesslyLength,
				sign,
			)
			if err != nil {
				return 0, false, err
			}

			if found {
				return number, true, nil
			}

			if beginning {
				begun = true
				continue
			}

			if begun {
				// Since at the current step we managed to multiply the exponent and
				// add/subtract the resulting number with a multiplier equal to
				// 'multiplier' without overflow, then for 'multiplier-1' this will
				// definitely be possible without overflow
				previous -= (multiplier - 1) * exponent
				break
			}
		}
	}

	return 0, false, nil
}

// Finds last number in a continuous sequence of integers that convert to floating point
// numbers without loss followed by a number that converts with loss.
//
// Search range is limited by the base number and required length of a continuous
// sequence of integers that must be converted to floating point numbers without loss.
//
// See svg file for visualization.
func findLosslessly[TypeTo constraints.Float, TypeFrom constraints.Integer](
	base TypeFrom,
	losslesslyLength TypeFrom,
	sign TypeFrom,
) (TypeFrom, bool, bool, error) {
	if sign < 0 {
		return findLosslesslyNegative[TypeTo](base, losslesslyLength)
	}

	return findLosslesslyPositive[TypeTo](base, losslesslyLength)
}

func findLosslesslyPositive[TypeTo constraints.Float, TypeFrom constraints.Integer](
	base TypeFrom,
	losslesslyLength TypeFrom,
) (TypeFrom, bool, bool, error) {
	greatest, err := safe.Add3(base, losslesslyLength, 1)
	if err != nil {
		return 0, false, false, ErrLosslesslyLengthTooLong
	}

	for number := base; number < greatest; number++ {
		reverted := passThroughFloat[TypeTo](number)

		if reverted == number {
			if losslesslyLength == 0 {
				return 0, true, false, nil
			}

			losslesslyLength--

			continue
		}

		if losslesslyLength == 0 {
			return number - 1, true, true, nil
		}
	}

	return 0, false, false, nil
}

func findLosslesslyNegative[TypeTo constraints.Float, TypeFrom constraints.Integer](
	base TypeFrom,
	losslesslyLength TypeFrom,
) (TypeFrom, bool, bool, error) {
	leastest, err := safe.Sub3(base, losslesslyLength, 1)
	if err != nil {
		return 0, false, false, ErrLosslesslyLengthTooLong
	}

	for number := base; number > leastest; number-- {
		reverted := passThroughFloat[TypeTo](number)

		if reverted == number {
			if losslesslyLength == 0 {
				return 0, true, false, nil
			}

			losslesslyLength--

			continue
		}

		if losslesslyLength == 0 {
			return number + 1, true, true, nil
		}
	}

	return 0, false, false, nil
}

// Checks whether precision loss occurs when converting an integer to a floating point
// number within a given range of values.
func IsSequenceLosslessly[TypeTo constraints.Float, TypeFrom constraints.Integer](
	begin TypeFrom,
	end TypeFrom,
) (bool, error) {
	if begin > end {
		return false, ErrBeginGreaterEnd
	}

	for number := begin; number < end; number++ {
		if passThroughFloat[TypeTo](number) != number {
			return false, nil
		}
	}

	// Prevents infinite loop when end is equal maximum value of type
	return passThroughFloat[TypeTo](end) == end, nil
}

// Converts an integer to a floating point number and back.
func passThroughFloat[TypeTo constraints.Float, TypeFrom constraints.Integer](
	number TypeFrom,
) TypeFrom {
	return TypeFrom(TypeTo(number))
}
