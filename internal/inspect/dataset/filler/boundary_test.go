package filler

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoundary(t *testing.T) {
	filler := NewBoundary[int8]()

	testBoundary(t, filler)

	filler.Reset()

	testBoundary(t, filler)
}

func testBoundary(t *testing.T, filler *Boundary[int8]) {
	const argsQuantity = 3

	args := make([]int8, argsQuantity)
	args64 := make([]int64, argsQuantity)

	boundaries := getBoundaries[int8]()

	for fid, first := range boundaries {
		for sid, second := range boundaries {
			for tid, third := range boundaries {
				isLastIteration := func() bool {
					return fid == len(boundaries)-1 && sid == fid && tid == fid
				}

				completed, err := filler.Fill(args, args64)
				require.NoError(
					t,
					err,
					"first: %v, second: %v, third: %v",
					first,
					second,
					third,
				)

				if isLastIteration() {
					require.True(
						t,
						completed,
						"first: %v, second: %v, third: %v",
						first,
						second,
						third,
					)
				} else {
					require.False(
						t,
						completed,
						"first: %v, second: %v, third: %v",
						first,
						second,
						third,
					)
				}

				require.Equal(t, []int8{third, second, first}, args)
				require.Equal(
					t,
					[]int64{int64(third), int64(second), int64(first)},
					args64,
				)
			}
		}
	}

	for range boundaries {
		completed, err := filler.Fill(args, args64)
		require.NoError(t, err)
		require.True(t, completed)
	}
}

func TestGetBoundaries(t *testing.T) {
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

	require.Equal(t, expectedInt, getBoundaries[int8]())
	require.Equal(t, expectedUint, getBoundaries[uint8]())
}
