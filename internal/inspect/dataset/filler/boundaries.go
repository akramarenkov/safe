package filler

import (
	"math"

	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/is"
)

// Returns values ​​equal to and close to the minimum and maximum values ​​for the used type.
func Boundaries[Type types.USI8]() []Type {
	const (
		one = 1
		two = 2
	)

	if is.Signed[Type]() {
		zero := 0
		min := math.MinInt8
		max := math.MaxInt8

		boundaries := []Type{
			Type(min),
			Type(min + one),
			Type(min + two),
			Type(zero - two),
			Type(zero - one),
			Type(zero),
			Type(zero + one),
			Type(zero + two),
			Type(max - two),
			Type(max - one),
			Type(max),
		}

		return boundaries
	}

	min := 0
	max := math.MaxUint8

	boundaries := []Type{
		Type(min),
		Type(min + one),
		Type(min + two),
		Type(max - two),
		Type(max - one),
		Type(max),
	}

	return boundaries
}
