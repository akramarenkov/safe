package safe

import (
	"golang.org/x/exp/constraints"
)

func compareMulM[Type constraints.Integer](first, second Type) int {
	// first < 0 && second < 0
	if first&second < 0 {
		if first < second {
			return 1
		}

		return -1
	}

	switch {
	case first < second:
		return -1
	case first > second:
		return 1
	}

	return 0
}
