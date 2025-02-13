package safe

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompareMulM(t *testing.T) {
	testCompareMulM(
		t,
		[]int{7, 6, 5, 4, 0, 3, 2, 1, -1, -2, -3, 0, -4, -5, -6, -7},
		[]int{-1, -2, -3, -4, -5, -6, -7, 0, 0, 1, 2, 3, 4, 5, 6, 7},
	)

	testCompareMulM(
		t,
		[]int{8, 7, 6, 5, 4, 0, 3, 2, 1, -1, -2, -3, 0, -4, -5, -6, -7, -8},
		[]int{-1, -2, -3, -4, -5, -6, -7, -8, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8},
	)
}

func testCompareMulM(t *testing.T, unsorted, expected []int) {
	sorted := slices.Clone(unsorted)

	slices.SortFunc(sorted, compareMulM)
	require.NotEqual(t, unsorted, sorted)
	require.Equal(t, expected, sorted)
}
