package filler

import (
	"github.com/akramarenkov/safe/internal/inspect/types"
)

// Fills arguments with values ​from set.
type Set[Type types.USI8] struct {
	completed bool
	indices   []int
	set       []Type
}

// Creates filler that fill arguments with values ​from set.
//
// If setter functions was not specified then [Boundaries] function will be used.
func NewSet[Type types.USI8](setters ...func() []Type) *Set[Type] {
	bnd := &Set[Type]{}

	for _, setter := range setters {
		if setter != nil {
			bnd.set = append(bnd.set, setter()...)
		}
	}

	if len(setters) == 0 {
		bnd.set = Boundaries[Type]()
	}

	return bnd
}

// Fills arguments with values ​from set.
func (bnd *Set[Type]) Fill(args []Type, args64 []int64) (bool, error) {
	if bnd.completed {
		return true, nil
	}

	bnd.extendIndices(args)

	for id := range args {
		args[id] = bnd.set[bnd.indices[id]]
		args64[id] = int64(args[id])
	}

	if bnd.isIncreasedToMax() {
		bnd.completed = true
		return true, nil
	}

	bnd.increase()

	return false, nil
}

func (bnd *Set[Type]) extendIndices(args []Type) {
	// extension corresponds to the addition of top-level loops
	if len(args) > len(bnd.indices) {
		bnd.indices = append(bnd.indices, make([]int, len(args)-len(bnd.indices))...)
	}
}

func (bnd *Set[Type]) increase() {
	// smaller the id, the more nested the loop it corresponds to and vice versa
	for id := range bnd.indices {
		bnd.indices[id]++

		if bnd.indices[id] == len(bnd.set) {
			bnd.indices[id] = 0
			continue
		}

		return
	}
}

func (bnd *Set[Type]) isIncreasedToMax() bool {
	for id := range bnd.indices {
		if bnd.indices[id] != len(bnd.set)-1 {
			return false
		}
	}

	return true
}

func (bnd *Set[Type]) Reset() {
	clear(bnd.indices)
	bnd.completed = false
}
