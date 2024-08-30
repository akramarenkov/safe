package filler

import (
	"math"

	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/is"
)

// Fills arguments with values ​​equal to and close to the minimum and maximum values
// ​​for the used type.
type Boundary[Type types.USI8] struct {
	boundaries []Type
	completed  bool
	indices    []int
}

// Creates filler that fill arguments with values ​​equal to and close to the minimum and
// maximum values ​for the used type.
func NewBoundary[Type types.USI8]() *Boundary[Type] {
	bnd := &Boundary[Type]{
		boundaries: getBoundaries[Type](),
	}

	return bnd
}

// Fills arguments with values ​​equal to and close to the minimum and maximum values
// ​​for the used type.
func (bnd *Boundary[Type]) Fill(args []Type, args64 []int64) (bool, error) {
	if bnd.completed {
		return true, nil
	}

	bnd.extendIndices(args)

	for id := range args {
		args[id] = bnd.boundaries[bnd.indices[id]]
		args64[id] = int64(args[id])
	}

	if bnd.isIncreasedToMax() {
		bnd.completed = true
		return true, nil
	}

	bnd.increase()

	return false, nil
}

func (bnd *Boundary[Type]) extendIndices(args []Type) {
	// extension corresponds to the addition of top-level loops
	if len(args) > len(bnd.indices) {
		bnd.indices = append(bnd.indices, make([]int, len(args)-len(bnd.indices))...)
	}
}

func (bnd *Boundary[Type]) increase() {
	// smaller the id, the more nested the loop it corresponds to and vice versa
	for id := range bnd.indices {
		bnd.indices[id]++

		if bnd.indices[id] == len(bnd.boundaries) {
			bnd.indices[id] = 0
			continue
		}

		return
	}
}

func (bnd *Boundary[Type]) isIncreasedToMax() bool {
	for id := range bnd.indices {
		if bnd.indices[id] != len(bnd.boundaries)-1 {
			return false
		}
	}

	return true
}

func (bnd *Boundary[Type]) Reset() {
	clear(bnd.indices)
	bnd.completed = false
}

func getBoundaries[Type types.USI8]() []Type {
	const (
		one = 1
		two = 2
	)

	if is.Signed[Type]() {
		zero := 0
		min := math.MinInt8
		max := math.MaxInt8

		boundaries := []Type{
			Type(min),
			Type(min + one),
			Type(min + two),
			Type(zero - two),
			Type(zero - one),
			Type(zero),
			Type(zero + one),
			Type(zero + two),
			Type(max - two),
			Type(max - one),
			Type(max),
		}

		return boundaries
	}

	min := 0
	max := math.MaxUint8

	boundaries := []Type{
		Type(min),
		Type(min + one),
		Type(min + two),
		Type(max - two),
		Type(max - one),
		Type(max),
	}

	return boundaries
}
