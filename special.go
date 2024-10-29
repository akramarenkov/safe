package safe

import "golang.org/x/exp/constraints"

// Used to safely (using a method that avoids integer overflow) calculate the distance
// (difference in absolute value) between two numbers.
func Dist[Type constraints.Integer](first, second Type) uint64 {
	firstU64 := u64(first)
	secondU64 := u64(second)

	// first < 0 && second > 0 || first > 0 && second < 0
	if first^second < 0 {
		return secondU64 + firstU64
	}

	if firstU64 > secondU64 {
		return firstU64 - secondU64
	}

	return secondU64 - firstU64
}
