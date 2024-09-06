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

	args := make([]int64, 3)

	require.Panics(t, func() { _, _ = filler.Fill(args) })
}

func testSet(t *testing.T, filler *Set[int8], set []int8) {
	args := make([]int64, 3)

	for fid, first := range set {
		for sid, second := range set {
			for tid, third := range set {
				isLastIteration := func() bool {
					return fid == len(set)-1 && sid == fid && tid == fid
				}

				completed, err := filler.Fill(args)
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

				require.Equal(
					t,
					[]int64{int64(third), int64(second), int64(first)},
					args,
				)
			}
		}
	}

	for range set {
		completed, err := filler.Fill(args)
		require.NoError(t, err)
		require.True(t, completed)
	}
}
