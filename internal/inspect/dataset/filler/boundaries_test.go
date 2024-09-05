package filler

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoundaries(t *testing.T) {
	expectedInt := []int8{
		math.MinInt8,
		math.MinInt8 + 1,
		math.MinInt8 + 2,
		-2,
		-1,
		0,
		1,
		2,
		math.MaxInt8 - 2,
		math.MaxInt8 - 1,
		math.MaxInt8,
	}

	expectedUint := []uint8{
		0,
		1,
		2,
		math.MaxUint8 - 2,
		math.MaxUint8 - 1,
		math.MaxUint8,
	}

	require.Equal(t, expectedInt, Boundaries[int8]())
	require.Equal(t, expectedUint, Boundaries[uint8]())
}
