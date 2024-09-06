package filler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSame(t *testing.T) {
	filler := NewSame[int8](3, 1<<8)

	args := make([]int64, 5)

	expected := []int64{3, 3, 3, 3, 3}

	for range 1 << 8 {
		completed, err := filler.Fill(args)
		require.False(t, completed)
		require.NoError(t, err)
		require.Equal(t, expected, args)
	}

	for range 1 << 8 {
		completed, err := filler.Fill(args)
		require.True(t, completed)
		require.NoError(t, err)
	}
}
