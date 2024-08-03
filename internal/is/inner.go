package is

import (
	"golang.org/x/exp/constraints"
)

// Determines whether a number is the minimum number for a given integer type.
func Min[Type constraints.Integer](number Type) bool {
	if number > 0 {
		return false
	}

	number--

	return number > 0
}

// Determines whether a number is the maximum number for a given integer type.
func Max[Type constraints.Integer](number Type) bool {
	if number <= 0 {
		return false
	}

	number++

	return number <= 0
}

// Determines whether a number is equal to minus one.
func MinusOne[Type constraints.Integer](number Type) bool {
	if number >= 0 {
		return false
	}

	number++

	return number == 0
}

// Determines whether a number is a multiple of two.
func Even[Type constraints.Integer](number Type) bool {
	return number%2 == 0
}
