package filler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRand(t *testing.T) {
	testRandInt(t)
	testRandUint(t)
}

func testRandInt(t *testing.T) {
	const argsQuantity = 3

	args := make([]int8, argsQuantity)
	args64 := make([]int64, argsQuantity)
	previous := make([]int8, argsQuantity)

	filler := NewRand[int8]()

	for range 1 << 16 {
		copy(previous, args)

		completed, err := filler.Fill(args, args64)
		require.NoError(t, err)
		require.False(t, completed)
		require.NotEqual(t, previous, args)
	}
}

func testRandUint(t *testing.T) {
	const argsQuantity = 3

	args := make([]uint8, argsQuantity)
	args64 := make([]int64, argsQuantity)
	previous := make([]uint8, argsQuantity)

	filler := NewRand[uint8]()

	for range 1 << 16 {
		copy(previous, args)

		completed, err := filler.Fill(args, args64)
		require.NoError(t, err)
		require.False(t, completed)
		require.NotEqual(t, previous, args)
	}
}
