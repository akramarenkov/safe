package safe

import (
	"golang.org/x/exp/constraints"
)

func u64[Type constraints.Integer](number Type) uint64 {
	if number < 0 {
		return uint64(-(number + 1)) + 1
	}

	return uint64(number)
}
