package safe

import (
	"slices"
	"testing"

	"github.com/akramarenkov/safe/internal/clone"
	"github.com/akramarenkov/safe/internal/iterator"

	"github.com/stretchr/testify/require"
)

func TestSortAddM(t *testing.T) {
	testSortAddM(
		t,
		[]int{5, 4, 0, 3, 2, 1, -1, -2, -3, -4, -5},
		[]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
	)

	testSortAddM(
		t,
		[]int{6, 5, 0, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6},
		[]int{-6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6},
	)
}

func testSortAddM(t *testing.T, unsorted, expected []int) {
	sorted := clone.Slice(unsorted)

	sortAddM(sorted)

	require.NotEqual(t, unsorted, sorted)
	require.Equal(t, expected, sorted)
}

func TestSortMulM(t *testing.T) {
	testSortMulM(
		t,
		[]int{7, 6, 5, 4, 0, 3, 2, 1, -1, -2, -3, 0, -4, -5, -6, -7},
		[]int{-1, -2, -3, -4, -5, -6, -7, 0, 0, 1, 2, 3, 4, 5, 6, 7},
	)

	testSortMulM(
		t,
		[]int{8, 7, 6, 5, 4, 0, 3, 2, 1, -1, -2, -3, 0, -4, -5, -6, -7, -8},
		[]int{-1, -2, -3, -4, -5, -6, -7, -8, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8},
	)
}

func testSortMulM(t *testing.T, unsorted, expected []int) {
	sorted := clone.Slice(unsorted)
	sortedStd := clone.Slice(unsorted)

	sortMulM(sorted)
	slices.SortFunc(sortedStd, compareMulM)

	require.NotEqual(t, unsorted, sorted)
	require.NotEqual(t, unsorted, sortedStd)
	require.Equal(t, expected, sorted)
	require.Equal(t, expected, sortedStd)
}

func BenchmarkSortAddMSmall_1(b *testing.B) {
	benchmarkSortAddMSmall(b, 1)
}

func BenchmarkSortAddM_1(b *testing.B) {
	benchmarkSortAddM(b, 1)
}

func BenchmarkSortAddMStd_1(b *testing.B) {
	benchmarkSortAddMStd(b, 1)
}

func BenchmarkSortAddMSmall_4(b *testing.B) {
	benchmarkSortAddMSmall(b, 4)
}

func BenchmarkSortAddM_4(b *testing.B) {
	benchmarkSortAddM(b, 4)
}

func BenchmarkSortAddMStd_4(b *testing.B) {
	benchmarkSortAddMStd(b, 4)
}

func BenchmarkSortAddMSmall_12(b *testing.B) {
	benchmarkSortAddMSmall(b, 12)
}

func BenchmarkSortAddM_12(b *testing.B) {
	benchmarkSortAddM(b, 12)
}

func BenchmarkSortAddMStd_12(b *testing.B) {
	benchmarkSortAddMStd(b, 12)
}

func BenchmarkSortAddMSmall_13(b *testing.B) {
	benchmarkSortAddMSmall(b, 13)
}

func BenchmarkSortAddM_13(b *testing.B) {
	benchmarkSortAddM(b, 13)
}

func BenchmarkSortAddMStd_13(b *testing.B) {
	benchmarkSortAddMStd(b, 13)
}

func BenchmarkSortAddMSmall_32(b *testing.B) {
	benchmarkSortAddMSmall(b, 32)
}

func BenchmarkSortAddM_32(b *testing.B) {
	benchmarkSortAddM(b, 32)
}

func BenchmarkSortAddMStd_32(b *testing.B) {
	benchmarkSortAddMStd(b, 32)
}

func benchmarkSortAddMSmall(b *testing.B, quantity int) {
	unsorted, expected := prepareAddM(quantity)

	sorted := make([]int, len(unsorted))

	b.ResetTimer()

	for range b.N {
		copy(sorted, unsorted)

		sortAddMSmall(sorted)
	}

	require.Equal(b, expected, sorted)
}

func benchmarkSortAddM(b *testing.B, quantity int) {
	unsorted, expected := prepareAddM(quantity)

	sorted := make([]int, len(unsorted))

	b.ResetTimer()

	for range b.N {
		copy(sorted, unsorted)

		sortAddM(sorted)
	}

	require.Equal(b, expected, sorted)
}

func benchmarkSortAddMStd(b *testing.B, quantity int) {
	unsorted, expected := prepareAddM(quantity)

	sorted := make([]int, len(unsorted))

	b.ResetTimer()

	for range b.N {
		copy(sorted, unsorted)

		slices.Sort(sorted)
	}

	require.Equal(b, expected, sorted)
}

func BenchmarkSortMulMSmall_1(b *testing.B) {
	benchmarkSortMulMSmall(b, 1)
}

func BenchmarkSortMulM_1(b *testing.B) {
	benchmarkSortMulM(b, 1)
}

func BenchmarkSortMulMStd_1(b *testing.B) {
	benchmarkSortMulMStd(b, 1)
}

func BenchmarkSortMulMSmall_4(b *testing.B) {
	benchmarkSortMulMSmall(b, 4)
}

func BenchmarkSortMulM_4(b *testing.B) {
	benchmarkSortMulM(b, 4)
}

func BenchmarkSortMulMStd_4(b *testing.B) {
	benchmarkSortMulMStd(b, 4)
}

func BenchmarkSortMulMSmall_12(b *testing.B) {
	benchmarkSortMulMSmall(b, 12)
}

func BenchmarkSortMulM_12(b *testing.B) {
	benchmarkSortMulM(b, 12)
}

func BenchmarkSortMulMStd_12(b *testing.B) {
	benchmarkSortMulMStd(b, 12)
}

func BenchmarkSortMulMSmall_13(b *testing.B) {
	benchmarkSortMulMSmall(b, 13)
}

func BenchmarkSortMulM_13(b *testing.B) {
	benchmarkSortMulM(b, 13)
}

func BenchmarkSortMulMStd_13(b *testing.B) {
	benchmarkSortMulMStd(b, 13)
}

func BenchmarkSortMulMSmall_16(b *testing.B) {
	benchmarkSortMulMSmall(b, 16)
}

func BenchmarkSortMulM_16(b *testing.B) {
	benchmarkSortMulM(b, 16)
}

func BenchmarkSortMulMStd_16(b *testing.B) {
	benchmarkSortMulMStd(b, 16)
}

func BenchmarkSortMulMSmall_17(b *testing.B) {
	benchmarkSortMulMSmall(b, 17)
}

func BenchmarkSortMulM_17(b *testing.B) {
	benchmarkSortMulM(b, 17)
}

func BenchmarkSortMulMStd_17(b *testing.B) {
	benchmarkSortMulMStd(b, 17)
}

func BenchmarkSortMulMSmall_32(b *testing.B) {
	benchmarkSortMulMSmall(b, 32)
}

func BenchmarkSortMulM_32(b *testing.B) {
	benchmarkSortMulM(b, 32)
}

func BenchmarkSortMulMStd_32(b *testing.B) {
	benchmarkSortMulMStd(b, 32)
}

func benchmarkSortMulMSmall(b *testing.B, quantity int) {
	unsorted, expected := prepareMulM(quantity)

	sorted := make([]int, len(unsorted))

	b.ResetTimer()

	for range b.N {
		copy(sorted, unsorted)

		sortMulMSmall(sorted)
	}

	require.Equal(b, expected, sorted)
}

func benchmarkSortMulM(b *testing.B, quantity int) {
	unsorted, expected := prepareMulM(quantity)

	sorted := make([]int, len(unsorted))

	b.ResetTimer()

	for range b.N {
		copy(sorted, unsorted)

		sortMulM(sorted)
	}

	require.Equal(b, expected, sorted)
}

func benchmarkSortMulMStd(b *testing.B, quantity int) {
	unsorted, expected := prepareMulM(quantity)

	sorted := make([]int, len(unsorted))

	b.ResetTimer()

	for range b.N {
		copy(sorted, unsorted)

		slices.SortFunc(sorted, compareMulM)
	}

	require.Equal(b, expected, sorted)
}

func prepareAddM(quantity int) ([]int, []int) {
	unsorted := make([]int, 0, quantity)
	expected := make([]int, 0, quantity)

	for number := range iterator.Iter(quantity, 1) {
		unsorted = append(unsorted, number)
	}

	for number := range iterator.Iter(1, quantity) {
		expected = append(expected, number)
	}

	return unsorted, expected
}

func prepareMulM(quantity int) ([]int, []int) {
	positive := quantity/2 + quantity%2
	negative := -quantity / 2

	unsorted := make([]int, 0, quantity)
	expected := make([]int, 0, quantity)

	for number := range iterator.Iter(positive, 1) {
		unsorted = append(unsorted, number)
	}

	for number := range iterator.Iter(negative, -1) {
		unsorted = append(unsorted, number)
	}

	for number := range iterator.Iter(-1, negative) {
		expected = append(expected, number)
	}

	for number := range iterator.Iter(1, positive) {
		expected = append(expected, number)
	}

	return unsorted, expected
}
