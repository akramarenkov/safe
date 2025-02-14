package filler

import (
	"errors"

	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/iterator"
)

var (
	ErrBeginGreaterEnd = errors.New("begin is greater than end")
)

// Returns values from begin to end inclusive.
func Span[Type types.UpToI32](begin, end Type) []Type {
	if begin > end {
		panic(ErrBeginGreaterEnd)
	}

	span := make([]Type, 0, int64(end)-int64(begin)+int64(1))

	for value := range iterator.Iter(begin, end) {
		span = append(span, value)
	}

	return span
}
