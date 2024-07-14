package safe

import (
	"golang.org/x/exp/constraints"
)

// Determines whether a number is the minimum number for a given integer type.
func isMin[Type constraints.Integer](number Type) bool {
	if number > 0 {
		return false
	}

	number--

	return number > 0
}

// Determines whether a number is the maximum number for a given integer type.
func isMax[Type constraints.Integer](number Type) bool {
	if number <= 0 {
		return false
	}

	number++

	return number <= 0
}

// Determines whether a number is equal to minus one.
func isMinusOne[Type constraints.Integer](number Type) bool {
	if number >= 0 {
		return false
	}

	number++

	return number == 0
}

// Tries to get minus one for the specified type and determines whether it was obtained
// or not.
func getMinusOne[Type constraints.Integer]() (Type, bool) {
	value := Type(0)

	value--

	return value, isMinusOne(value)
}

// Tries to get minus one for the specified type without determining whether it
// succeeded in getting it or not.
func getMinusOneUnsure[Type constraints.Integer]() Type {
	value := Type(0)

	value--

	return value
}

// Determines whether a number is a multiple of two.
func isEven[Type constraints.Integer](number Type) bool {
	return number%2 == 0
}
