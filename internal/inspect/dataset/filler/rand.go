package filler

import (
	"crypto/rand"
	"math"
	"math/big"

	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/is"
)

// Fills arguments with random values.
//
// Data in this filler never ends, so it must be specified last.
type Rand[Type types.USI8] struct {
	maxRand *big.Int
}

// Creates filler that fill arguments with random values.
//
// Data in this filler never ends, so it must be specified last.
func NewRand[Type types.USI8]() *Rand[Type] {
	rnd := &Rand[Type]{
		maxRand: big.NewInt(math.MaxUint8 + 1),
	}

	return rnd
}

// Fills arguments with random values.
//
// Data in this filler never ends, so it must be specified last.
func (rnd *Rand[Type]) Fill(args []Type, args64 []int64) (bool, error) {
	conv := func(value *big.Int) (Type, int64) {
		conv := value.Int64()
		return Type(conv), conv
	}

	if is.Signed[Type]() {
		conv = func(value *big.Int) (Type, int64) {
			conv := int64(math.MaxInt8) - value.Int64()
			return Type(conv), conv
		}
	}

	for id := range args64 {
		value, err := rand.Int(rand.Reader, rnd.maxRand)
		if err != nil {
			return false, err
		}

		args[id], args64[id] = conv(value)
	}

	return false, nil
}
