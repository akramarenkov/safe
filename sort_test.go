package safe

import (
	"slices"
	"testing"

	"github.com/akramarenkov/safe/internal/iterator"
	"github.com/stretchr/testify/require"
)

func TestSortMulM(t *testing.T) {
	testSortMulM(
		t,
		[]int{7, 6, 5, 4, 0, 3, 2, 1, -1, -2, -3, 0, -4, -5, -6, -7},
		[]int{-1, -2, -3, -4, -5, -6, -7, 0, 0, 1, 2, 3, 4, 5, 6, 7},
	)

	testSortMulM(
		t,
		[]int{7, 6, 5, 4, 0, 3, 2, 1, -1, -2, -3, 0, -4, -5, -6, -7, -8},
		[]int{-1, -2, -3, -4, -5, -6, -7, -8, 0, 0, 1, 2, 3, 4, 5, 6, 7},
	)

	testSortMulM(
		t,
		[]int{8, 7, 6, 5, 4, 0, 3, 2, 1, -1, -2, -3, 0, -4, -5, -6, -7, -8},
		[]int{-1, -2, -3, -4, -5, -6, -7, -8, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8},
	)

	testSortMulM(
		t,
		[]int{-1, -2, -3, -4, -5, -6, -7, -8, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8},
		[]int{-1, -2, -3, -4, -5, -6, -7, -8, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8},
	)
}

func testSortMulM(t *testing.T, original, expected []int) {
	sorted := slices.Clone(original)

	sortMulM(sorted)
	require.Equal(t, expected, sorted)
}

func BenchmarkSortMulM_1(b *testing.B) {
	benchmarkSortMulM(b, 1)
}

func BenchmarkSortMulMStd_1(b *testing.B) {
	benchmarkSortMulMStd(b, 1)
}

func BenchmarkSortMulM_4(b *testing.B) {
	benchmarkSortMulM(b, 4)
}

func BenchmarkSortMulMStd_4(b *testing.B) {
	benchmarkSortMulMStd(b, 4)
}

func BenchmarkSortMulM_12(b *testing.B) {
	benchmarkSortMulM(b, 12)
}

func BenchmarkSortMulMStd_12(b *testing.B) {
	benchmarkSortMulMStd(b, 12)
}

func BenchmarkSortMulM_16(b *testing.B) {
	benchmarkSortMulM(b, 16)
}

func BenchmarkSortMulMStd_16(b *testing.B) {
	benchmarkSortMulMStd(b, 16)
}

func BenchmarkSortMulM_17(b *testing.B) {
	benchmarkSortMulM(b, 17)
}

func BenchmarkSortMulMStd_17(b *testing.B) {
	benchmarkSortMulMStd(b, 17)
}

func BenchmarkSortMulM_20(b *testing.B) {
	benchmarkSortMulM(b, 20)
}

func BenchmarkSortMulMStd_20(b *testing.B) {
	benchmarkSortMulMStd(b, 20)
}

func BenchmarkSortMulM_24(b *testing.B) {
	benchmarkSortMulM(b, 24)
}

func BenchmarkSortMulMStd_24(b *testing.B) {
	benchmarkSortMulMStd(b, 24)
}

func BenchmarkSortMulM_32(b *testing.B) {
	benchmarkSortMulM(b, 32)
}

func BenchmarkSortMulMStd_32(b *testing.B) {
	benchmarkSortMulMStd(b, 32)
}

func benchmarkSortMulM(b *testing.B, quantity int) {
	original, expected := benchmarkPrepareMulM(b, quantity)

	sorted := make([]int, len(original))

	b.ResetTimer()

	for range b.N {
		copy(sorted, original)

		sortMulM(sorted)
	}

	require.Equal(b, expected, sorted)
}

func benchmarkSortMulMStd(b *testing.B, quantity int) {
	original, expected := benchmarkPrepareMulM(b, quantity)

	sorted := make([]int, len(original))

	b.ResetTimer()

	for range b.N {
		copy(sorted, original)

		slices.SortFunc(sorted, compareMulM)
	}

	require.Equal(b, expected, sorted)
}

func benchmarkPrepareMulM(b *testing.B, quantity int) ([]int, []int) {
	positive := quantity/2 + quantity%2
	negative := -quantity / 2

	original := make([]int, 0, quantity)
	expected := make([]int, 0, quantity)

	for number := range iterator.Iter(positive, 1) {
		original = append(original, number)
	}

	for number := range iterator.Iter(negative, -1) {
		original = append(original, number)
	}

	for number := range iterator.Iter(-1, negative) {
		expected = append(expected, number)
	}

	for number := range iterator.Iter(1, positive) {
		expected = append(expected, number)
	}

	require.NotEqual(b, expected, original)

	return original, expected
}
