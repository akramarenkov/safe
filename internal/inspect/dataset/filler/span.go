package filler

import (
	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/iterator"
)

// Returns values from begin to end inclusive.
func Span[Type types.UpToUSI32](begin Type, end Type) []Type {
	span := make([]Type, 0, int64(end)-int64(begin)+int64(1))

	for value := range iterator.Iter(begin, end) {
		span = append(span, value)
	}

	return span
}
