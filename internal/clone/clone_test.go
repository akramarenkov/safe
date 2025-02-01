package clone

import (
	"slices"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestClone(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5}
	original := []int{0, 1, 2, 3, 4, 5}

	copied := Slice(original)

	require.Equal(t, expected, original)
	require.Equal(t, expected, copied)
	require.NotSame(t, unsafe.SliceData(original), unsafe.SliceData(copied))

	for id := range copied {
		copied[id]--
	}

	require.Equal(t, expected, original)
	require.NotEqual(t, expected, copied)
}

func BenchmarkClone1(b *testing.B) {
	benchmarkClone(b, 1)
}

func BenchmarkAppend1(b *testing.B) {
	benchmarkAppend(b, 1)
}

func BenchmarkSlicesClone1(b *testing.B) {
	benchmarkSlicesClone(b, 1)
}

func BenchmarkClone4(b *testing.B) {
	benchmarkClone(b, 4)
}

func BenchmarkAppend4(b *testing.B) {
	benchmarkAppend(b, 4)
}

func BenchmarkSlicesClone4(b *testing.B) {
	benchmarkSlicesClone(b, 4)
}

func BenchmarkClone8(b *testing.B) {
	benchmarkClone(b, 8)
}

func BenchmarkAppend8(b *testing.B) {
	benchmarkAppend(b, 8)
}

func BenchmarkSlicesClone8(b *testing.B) {
	benchmarkSlicesClone(b, 8)
}

func BenchmarkClone32(b *testing.B) {
	benchmarkClone(b, 32)
}

func BenchmarkAppend32(b *testing.B) {
	benchmarkAppend(b, 32)
}

func BenchmarkSlicesClone32(b *testing.B) {
	benchmarkSlicesClone(b, 32)
}

func BenchmarkClone64(b *testing.B) {
	benchmarkClone(b, 64)
}

func BenchmarkAppend64(b *testing.B) {
	benchmarkAppend(b, 64)
}

func BenchmarkSlicesClone64(b *testing.B) {
	benchmarkSlicesClone(b, 64)
}

func BenchmarkClone128(b *testing.B) {
	benchmarkClone(b, 128)
}

func BenchmarkAppend128(b *testing.B) {
	benchmarkAppend(b, 128)
}

func BenchmarkSlicesClone128(b *testing.B) {
	benchmarkSlicesClone(b, 128)
}

func benchmarkClone(b *testing.B, quantity int) {
	slice := make([]int, quantity)

	for range b.N {
		slice = Slice(slice)
	}

	require.NotNil(b, slice)
}

func benchmarkAppend(b *testing.B, quantity int) {
	slice := make([]int, quantity)

	for range b.N {
		slice = append([]int(nil), slice...)
	}

	require.NotNil(b, slice)
}

func benchmarkSlicesClone(b *testing.B, quantity int) {
	slice := make([]int, quantity)

	for range b.N {
		slice = slices.Clone(slice)
	}

	require.NotNil(b, slice)
}
