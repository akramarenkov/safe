package filler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetBoundaries(t *testing.T) {
	filler := NewSet[int8]()

	testSet(t, filler, Boundaries[int8]())

	filler.Reset()

	testSet(t, filler, Boundaries[int8]())
}

func TestSetSpan(t *testing.T) {
	span := func() []int8 {
		return Span[int8](-10, 10)
	}

	filler := NewSet(span)

	testSet(t, filler, span())

	filler.Reset()

	testSet(t, filler, span())
}

func TestSetPanic(t *testing.T) {
	filler := NewSet[int8](nil)

	const argsQuantity = 1

	args := make([]int8, argsQuantity)
	args64 := make([]int64, argsQuantity)

	require.Panics(t, func() { _, _ = filler.Fill(args, args64) })
}

func testSet(t *testing.T, filler *Set[int8], set []int8) {
	const argsQuantity = 3

	args := make([]int8, argsQuantity)
	args64 := make([]int64, argsQuantity)

	for fid, first := range set {
		for sid, second := range set {
			for tid, third := range set {
				isLastIteration := func() bool {
					return fid == len(set)-1 && sid == fid && tid == fid
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

	for range set {
		completed, err := filler.Fill(args, args64)
		require.NoError(t, err)
		require.True(t, completed)
	}
}
