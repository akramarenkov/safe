package filler

import (
	"github.com/akramarenkov/safe/internal/inspect/types"
)

// Filler that always returns the same item.
type Same[Type types.USI8] struct {
	max   int
	value Type
}

// Creates filler that always returns the same item.
//
// Arguments are filled with 'value' value, after 'max' number of calls method Fill
// will stop filling arguments and will return true.
func NewSame[Type types.USI8](value Type, max int) *Same[Type] {
	sm := &Same[Type]{
		max:   max,
		value: value,
	}

	return sm
}

// Returns the same item always.
func (sm *Same[Type]) Fill(args []int64) (bool, error) {
	if sm.max <= 0 {
		return true, nil
	}

	sm.max--

	for id := range args {
		args[id] = int64(sm.value)
	}

	return false, nil
}
