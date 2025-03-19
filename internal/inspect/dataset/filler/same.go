package filler

import "github.com/akramarenkov/safe/internal/inspect/confines"

// Filler that always returns the same item.
type Same[Type confines.UpToI32] struct {
	maximum int
	value   Type
}

// Creates filler that always returns the same item.
//
// Arguments are filled with 'value' value, after 'max' number of calls method Fill
// will stop filling arguments and will return true.
func NewSame[Type confines.UpToI32](value Type, maximum int) *Same[Type] {
	sm := &Same[Type]{
		maximum: maximum,
		value:   value,
	}

	return sm
}

// Returns the same item always.
func (sm *Same[Type]) Fill(args []Type, args64 []int64) (bool, error) {
	if sm.maximum <= 0 {
		return true, nil
	}

	sm.maximum--

	for id := range args64 {
		args[id] = sm.value
		args64[id] = int64(sm.value)
	}

	return false, nil
}
