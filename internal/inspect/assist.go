package inspect

import (
	"math"

	"github.com/akramarenkov/safe/internal/is"
)

// Picks up maximum and minimum values for specified type.
func PickUpRange[Type EightBits]() (int64, int64) {
	if is.Signed[Type]() {
		return math.MinInt8, math.MaxInt8
	}

	return 0, math.MaxUint8
}
