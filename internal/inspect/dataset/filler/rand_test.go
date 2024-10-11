package filler

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRand(t *testing.T) {
	testRandSig(t)
	testRandUns(t)
}

func testRandSig(t *testing.T) {
	const argsQuantity = 3

	args := make([]int8, argsQuantity)
	args64 := make([]int64, argsQuantity)
	previous := make([]int64, argsQuantity)

	filler := NewRand[int8]()

	negative := 0

	for range 1 << 16 {
		copy(previous, args64)

		completed, err := filler.Fill(args, args64)
		require.NoError(t, err)
		require.False(t, completed)
		require.NotEqual(t, previous, args64)

		for _, arg := range args64 {
			if arg < 0 {
				negative++
			}

			require.LessOrEqual(t, arg, int64(math.MaxInt8))
			require.GreaterOrEqual(t, arg, int64(math.MinInt8))
		}
	}

	require.NotZero(t, negative)
}

func testRandUns(t *testing.T) {
	const argsQuantity = 3

	args := make([]uint8, argsQuantity)
	args64 := make([]int64, argsQuantity)
	previous := make([]int64, argsQuantity)

	filler := NewRand[uint8]()

	for range 1 << 16 {
		copy(previous, args64)

		completed, err := filler.Fill(args, args64)
		require.NoError(t, err)
		require.False(t, completed)
		require.NotEqual(t, previous, args64)

		for _, arg := range args64 {
			require.LessOrEqual(t, arg, int64(math.MaxUint8))
			require.GreaterOrEqual(t, arg, int64(0))
		}
	}
}

func BenchmarkRand(b *testing.B) {
	const argsQuantity = 3

	args := make([]int8, argsQuantity)
	args64 := make([]int64, argsQuantity)

	filler := NewRand[int8]()

	for range b.N {
		if _, err := filler.Fill(args, args64); err != nil {
			require.NoError(b, err)
		}
	}
}
