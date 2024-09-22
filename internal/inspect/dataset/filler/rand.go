package filler

import (
	"crypto/rand"
	"math/big"

	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/intspan"
	"github.com/akramarenkov/safe/internal/is"
)

// Fills arguments with random values.
//
// Data in this filler never ends, so it must be specified last.
type Rand[Type types.UpToUSI32] struct {
	bitSize     int
	maximum     int64
	maximumRand *big.Int
}

// Creates filler that fill arguments with random values.
//
// Data in this filler never ends, so it must be specified last.
func NewRand[Type types.UpToUSI32]() *Rand[Type] {
	_, maximum, bitSize := intspan.Get[Type]()

	rnd := &Rand[Type]{
		bitSize:     bitSize,
		maximum:     int64(maximum),
		maximumRand: big.NewInt(1 << bitSize),
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
			conv := rnd.maximum - value.Int64()
			return Type(conv), conv
		}
	}

	for id := range args64 {
		value, err := rand.Int(rand.Reader, rnd.maximumRand)
		if err != nil {
			return false, err
		}

		args[id], args64[id] = conv(value)
	}

	return false, nil
}
