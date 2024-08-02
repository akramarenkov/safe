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

// Determines whether a number is a multiple of two.
func isEven[Type constraints.Integer](number Type) bool {
	return number%2 == 0
}

// Creates a shallow copy of a slice. Slightly faster than append on small number of
// elements.
func cloneSlice[Type any](slice []Type) []Type {
	copied := make([]Type, len(slice))

	copy(copied, slice)

	return copied
}
