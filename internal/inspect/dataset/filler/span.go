package filler

import (
	"github.com/akramarenkov/safe/internal/inspect/types"
)

// Returns values from begin to end inclusive.
//
// Allocating slices larger than 2^16 for tests is expensive.
func Span[Type types.UpToUSI16](begin Type, end Type) []Type {
	// The target type is obviously smaller than int type, so there can be no integer
	// overflow here
	span := make([]Type, 0, int(end)-int(begin)+1)

	for value := begin; value < end; value++ {
		span = append(span, value)
	}

	// prevents infinite loop when end is equal max value of type
	return append(span, end)
}
