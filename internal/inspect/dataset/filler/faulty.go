package filler

import (
	"errors"

	"github.com/akramarenkov/safe/internal/inspect/confines"
)

var (
	ErrFaulty = errors.New("faulty")
)

// Filler that always returns an error.
type Faulty[Type confines.UpToI32] struct{}

// Creates filler that always returns an error.
func NewFaulty[Type confines.UpToI32]() *Faulty[Type] {
	return &Faulty[Type]{}
}

// Returns an error always.
func (*Faulty[Type]) Fill([]Type, []int64) (bool, error) {
	return false, ErrFaulty
}
