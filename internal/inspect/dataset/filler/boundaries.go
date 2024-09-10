package filler

import (
	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/intspan"
	"github.com/akramarenkov/safe/internal/is"
)

// Returns values ​​equal to and close to the minimum and maximum values ​​for the used type.
func Boundaries[Type types.USI8]() []Type {
	min, max := intspan.Get[Type]()

	zero := Type(0)

	if is.Signed[Type]() {
		boundaries := []Type{
			min,
			min + 1,
			min + 2,
			zero - 2,
			zero - 1,
			zero,
			zero + 1,
			zero + 2,
			max - 2,
			max - 1,
			max,
		}

		return boundaries
	}

	boundaries := []Type{
		min,
		min + 1,
		min + 2,
		max - 2,
		max - 1,
		max,
	}

	return boundaries
}
