package filler

import (
	"github.com/akramarenkov/safe/internal/inspect/types"
)

// Returns values from begin to end inclusive.
func Span[Type types.USI8](begin Type, end Type) []Type {
	span := make([]Type, 0, int(end)-int(begin)+1)

	for value := begin; value < end; value++ {
		span = append(span, value)
	}

	// prevents infinite loop when end is equal max value of type
	return append(span, end)
}
