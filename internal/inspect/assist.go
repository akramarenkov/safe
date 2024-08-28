package inspect

import (
	"math"

	"github.com/akramarenkov/safe/internal/is"
)

// Picks up maximum and minimum values for specified type.
func PickUpRange[Type UpToSixteenBits]() (int64, int64) {
	if is.Signed[Type]() {
		return pickUpRangeSigned[Type]()
	}

	return pickUpRangeUnsigned[Type]()
}

func pickUpRangeSigned[Type UpToSixteenBits]() (int64, int64) {
	reference := int64(math.MaxInt16)

	if int64(Type(reference)) == math.MaxInt16 {
		return math.MinInt16, math.MaxInt16
	}

	return math.MinInt8, math.MaxInt8
}

func pickUpRangeUnsigned[Type UpToSixteenBits]() (int64, int64) {
	reference := int64(math.MaxUint16)

	if int64(Type(reference)) == math.MaxUint16 {
		return 0, math.MaxUint16
	}

	return 0, math.MaxUint8
}
