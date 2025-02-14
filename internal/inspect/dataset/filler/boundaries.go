package filler

import (
	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/is"

	"github.com/akramarenkov/intspec"
)

// Returns values ​​equal to and close to the minimum and maximum values ​​for the used type.
func Boundaries[Type types.UpToI32]() []Type {
	minimum, maximum := intspec.Range[Type]()

	zero := Type(0)

	if is.Signed[Type]() {
		boundaries := []Type{
			minimum,
			minimum + 1,
			minimum + 2,
			zero - 2,
			zero - 1,
			zero,
			zero + 1,
			zero + 2,
			maximum - 2,
			maximum - 1,
			maximum,
		}

		return boundaries
	}

	boundaries := []Type{
		minimum,
		minimum + 1,
		minimum + 2,
		maximum - 2,
		maximum - 1,
		maximum,
	}

	return boundaries
}
