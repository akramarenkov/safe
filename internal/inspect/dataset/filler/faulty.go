package filler

import (
	"errors"

	"github.com/akramarenkov/safe/internal/inspect"
)

var (
	ErrFaulty = errors.New("faulty")
)

// Filler that always returns an error.
type Faulty[Type inspect.EightBits] struct{}

// Creates filler that always returns an error.
func NewFaulty[Type inspect.EightBits]() *Faulty[Type] {
	return &Faulty[Type]{}
}

// Returns an error always.
func (flt *Faulty[Type]) Fill([]Type, []int64) (bool, error) {
	return false, ErrFaulty
}
