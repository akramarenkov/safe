package safe

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompareMulM(t *testing.T) {
	testCompareMulM(
		t,
		[]int{4, 4, 2, 3, -1, 3, 0, -2, 1, -4, 0, -3},
		[]int{-1, -2, -3, -4, 0, 0, 1, 2, 3, 3, 4, 4},
	)

	testCompareMulM(
		t,
		[]int{-1, -2, -3, -4, 0, 0, 1, 2, 3, 3, 4, 4},
		[]int{-1, -2, -3, -4, 0, 0, 1, 2, 3, 3, 4, 4},
	)

	testCompareMulM(
		t,
		[]int{2, 4, 3, 0, 1, 4, 2, 3, 0, 1},
		[]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4},
	)

	testCompareMulM(
		t,
		[]int{-2, -4, -3, 0, -1, -4, -2, -3, 0, -1},
		[]int{-1, -1, -2, -2, -3, -3, -4, -4, 0, 0},
	)
}

func testCompareMulM(t *testing.T, original, expected []int) {
	sorted := slices.Clone(original)

	slices.SortFunc(sorted, compareMulM)
	require.Equal(t, expected, sorted)
}
