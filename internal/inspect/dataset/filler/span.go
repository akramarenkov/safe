package filler

import (
	"github.com/akramarenkov/safe/internal/inspect/types"
)

// Returns values from begin to end inclusive.
func Span[Type types.UpToUSI32](begin Type, end Type) []Type {
	span := make([]Type, 0, int64(end)-int64(begin)+int64(1))

	for value := begin; value < end; value++ {
		span = append(span, value)
	}

	// prevents infinite loop when end is equal maximum value of type
	return append(span, end)
}
