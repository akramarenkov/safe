// Internal package that research overflow when adding integers.
package add

import (
	"github.com/akramarenkov/safe"
)

// Calculates the minimum and maximum value for the specified type when overflow occurs
// when adding two positive signed numbers.
func calcSpanAddPositive[Uns uint8, Type int8](first, second, overflowed Type) (Type, Type) {
	const midpointDivider = 2

	op, err := safe.Negate(overflowed)
	if err != nil {
		return overflowed, overflowed - 1
	}

	fu := Uns(first)  //nolint:gosec // There can only be positive numbers of the same bit depth here
	su := Uns(second) //nolint:gosec // There can only be positive numbers of the same bit depth here
	ou := Uns(op)     //nolint:gosec // There can only be positive numbers of the same bit depth here

	//nolint:gosec // The value cannot exceed the range of the specified type
	minimum := overflowed - Type((fu+su-ou)/midpointDivider)

	return minimum, minimum - 1
}

// Calculates the minimum and maximum value for the specified type when overflow occurs
// when adding two negative signed numbers.
func calcSpanAddNegative[Uns uint8, Type int8](first, second, overflowed Type) (Type, Type) {
	const midpointDivider = 2

	fp, err := safe.Negate(first)
	if err != nil {
		return first, first - 1
	}

	sp, err := safe.Negate(second)
	if err != nil {
		return second, second - 1
	}

	fu := Uns(fp)         //nolint:gosec // There can only be positive numbers of the same bit depth here
	su := Uns(sp)         //nolint:gosec // There can only be positive numbers of the same bit depth here
	ou := Uns(overflowed) //nolint:gosec // There can only be positive numbers of the same bit depth here

	//nolint:gosec // The value cannot exceed the range of the specified type
	minimum := first - Type((ou+su-fu)/midpointDivider)

	return minimum, minimum - 1
}
