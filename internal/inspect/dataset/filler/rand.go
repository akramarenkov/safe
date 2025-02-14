package filler

import (
	"crypto/rand"
	"math/big"

	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/is"

	"github.com/akramarenkov/intspec"
)

// Fills arguments with random values.
//
// Data in this filler never ends, so it must be specified last.
type Rand[Type types.UpToI32] struct {
	maximum     int64
	maximumRand *big.Int
}

// Creates filler that fill arguments with random values.
//
// Data in this filler never ends, so it must be specified last.
func NewRand[Type types.UpToI32]() *Rand[Type] {
	_, maximum := intspec.Range[Type]()

	rnd := &Rand[Type]{
		maximum:     int64(maximum),
		maximumRand: big.NewInt(1 << intspec.BitSize[Type]()),
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
