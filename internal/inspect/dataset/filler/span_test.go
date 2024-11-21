package filler

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpan(t *testing.T) {
	require.Equal(
		t,
		[]int8{math.MaxInt8 - 2, math.MaxInt8 - 1, math.MaxInt8},
		Span[int8](math.MaxInt8-2, math.MaxInt8),
	)

	require.Equal(
		t,
		[]int8{math.MaxInt8},
		Span[int8](math.MaxInt8, math.MaxInt8),
	)

	require.Equal(
		t,
		[]uint8{math.MaxUint8 - 2, math.MaxUint8 - 1, math.MaxUint8},
		Span[uint8](math.MaxUint8-2, math.MaxUint8),
	)

	require.Equal(
		t,
		[]uint8{math.MaxUint8},
		Span[uint8](math.MaxUint8, math.MaxUint8),
	)

	require.Len(t, Span[int8](math.MinInt8, math.MaxInt8), 1<<8)
	require.Len(t, Span[uint8](0, math.MaxUint8), 1<<8)

	require.Panics(t, func() { Span[int8](math.MaxInt8, math.MaxInt8-1) })
	require.Panics(t, func() { Span[uint8](math.MaxUint8, math.MaxUint8-1) })
	require.Panics(t, func() { Span[int8](math.MaxInt8, 2) })
	require.Panics(t, func() { Span[uint8](math.MaxUint8, 2) })
}
