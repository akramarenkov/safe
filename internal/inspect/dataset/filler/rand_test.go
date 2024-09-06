package filler

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRand(t *testing.T) {
	testRandInt(t)
	testRandUint(t)
}

func testRandInt(t *testing.T) {
	const argsQuantity = 3

	args := make([]int64, argsQuantity)
	previous := make([]int64, argsQuantity)

	filler := NewRand[int8]()

	negative := 0

	for range 1 << 16 {
		copy(previous, args)

		completed, err := filler.Fill(args)
		require.NoError(t, err)
		require.False(t, completed)
		require.NotEqual(t, previous, args)

		for _, arg := range args {
			if arg < 0 {
				negative++
			}

			require.LessOrEqual(t, arg, int64(math.MaxInt8))
			require.GreaterOrEqual(t, arg, int64(math.MinInt8))
		}
	}

	require.NotZero(t, negative)
}

func testRandUint(t *testing.T) {
	const argsQuantity = 3

	args := make([]int64, argsQuantity)
	previous := make([]int64, argsQuantity)

	filler := NewRand[uint8]()

	for range 1 << 16 {
		copy(previous, args)

		completed, err := filler.Fill(args)
		require.NoError(t, err)
		require.False(t, completed)
		require.NotEqual(t, previous, args)

		for _, arg := range args {
			require.LessOrEqual(t, arg, int64(math.MaxUint8))
			require.GreaterOrEqual(t, arg, int64(0))
		}
	}
}

func BenchmarkRand(b *testing.B) {
	args := make([]int64, 3)

	filler := NewRand[uint8]()

	for range b.N {
		if _, err := filler.Fill(args); err != nil {
			require.NoError(b, err)
		}
	}
}
