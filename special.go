package safe

import "golang.org/x/exp/constraints"

// Converts absolute value of an integer of specified type to an integer of uint64 type.
func Abs[Type constraints.Integer](number Type) uint64 {
	if number < 0 {
		return uint64(-(number + 1)) + 1
	}

	return uint64(number)
}

// Used to safely (using a method that avoids integer overflow) calculate the distance
// (difference in absolute value) between two numbers.
func Dist[Type constraints.Integer](first, second Type) uint64 {
	firstU64 := Abs(first)
	secondU64 := Abs(second)

	// first < 0 && second > 0 || first > 0 && second < 0
	if first^second < 0 {
		return secondU64 + firstU64
	}

	if firstU64 > secondU64 {
		return firstU64 - secondU64
	}

	return secondU64 - firstU64
}
