package filler

import (
	"crypto/rand"
	"math"
	"math/big"

	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/is"
)

// Fills arguments with random values.
type Rand[Type types.USI8] struct {
	maxRand *big.Int
}

// Creates filler that fill arguments with random values.
func NewRand[Type types.USI8]() *Rand[Type] {
	rnd := &Rand[Type]{
		maxRand: big.NewInt(math.MaxUint8 + 1),
	}

	return rnd
}

// Fills arguments with random values.
func (rnd *Rand[Type]) Fill(args []Type, args64 []int64) (bool, error) {
	for id := range args {
		value, err := rand.Int(rand.Reader, rnd.maxRand)
		if err != nil {
			return false, err
		}

		if is.Signed[Type]() {
			args64[id] = int64(math.MaxInt8) - value.Int64()
			args[id] = Type(args64[id])

			continue
		}

		args64[id] = int64(math.MaxUint8) - value.Int64()
		args[id] = Type(args64[id])
	}

	return false, nil
}
