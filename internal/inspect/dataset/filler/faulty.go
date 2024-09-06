package filler

import (
	"errors"
)

var (
	ErrFaulty = errors.New("faulty")
)

// Filler that always returns an error.
type Faulty struct{}

// Creates filler that always returns an error.
func NewFaulty() *Faulty {
	return &Faulty{}
}

// Returns an error always.
func (flt *Faulty) Fill([]int64) (bool, error) {
	return false, ErrFaulty
}
