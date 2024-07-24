package inspect

import (
	"math"

	"golang.org/x/exp/constraints"
)

func pickUpRange[Type EightBits]() (int64, int64) {
	if isSigned[Type]() {
		return math.MinInt8, math.MaxInt8
	}

	return 0, math.MaxUint8
}

func isSigned[Type constraints.Integer]() bool {
	number := Type(0)

	number--

	return number < 0
}
