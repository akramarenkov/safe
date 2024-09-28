package iterator

import (
	"golang.org/x/exp/constraints"
)

func toUint64[Type constraints.Integer](number Type) uint64 {
	if number < 0 {
		return uint64(-(number + 1)) + 1
	}

	return uint64(number)
}
