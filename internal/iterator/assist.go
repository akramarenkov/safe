package iterator

import (
	"github.com/akramarenkov/safe/internal/intspec"
	"golang.org/x/exp/constraints"
)

func toUint64[Type constraints.Integer](number Type) uint64 {
	if number < 0 {
		return uint64(-(number + 1)) + 1
	}

	return uint64(number)
}

func toInt(number uint64) int {
	if number > uint64(intspec.MaxInt) {
		return intspec.MaxInt
	}

	return int(number)
}
