package filler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSame(t *testing.T) {
	const argsQuantity = 5

	filler := NewSame[int8](3, 1<<8)

	args := make([]int8, argsQuantity)
	args64 := make([]int64, argsQuantity)

	expected := [argsQuantity]int8{3, 3, 3, 3, 3}
	expected64 := [argsQuantity]int64{3, 3, 3, 3, 3}

	for range 1 << 8 {
		completed, err := filler.Fill(args, args64)
		require.False(t, completed)
		require.NoError(t, err)
		require.Equal(t, expected[:], args)
		require.Equal(t, expected64[:], args64)
	}

	for range 1 << 8 {
		completed, err := filler.Fill(args, args64)
		require.True(t, completed)
		require.NoError(t, err)
	}
}
