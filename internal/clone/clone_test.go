package clone

import (
	"slices"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
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

func Benchmark1(b *testing.B) {
	benchmark(b, 1)
}

func BenchmarkAppend1(b *testing.B) {
	benchmarkAppend(b, 1)
}

func BenchmarkSlicesClone1(b *testing.B) {
	benchmarkSlicesClone(b, 1)
}

func Benchmark4(b *testing.B) {
	benchmark(b, 4)
}

func BenchmarkAppend4(b *testing.B) {
	benchmarkAppend(b, 4)
}

func BenchmarkSlicesClone4(b *testing.B) {
	benchmarkSlicesClone(b, 4)
}

func Benchmark8(b *testing.B) {
	benchmark(b, 8)
}

func BenchmarkAppend8(b *testing.B) {
	benchmarkAppend(b, 8)
}

func BenchmarkSlicesClone8(b *testing.B) {
	benchmarkSlicesClone(b, 8)
}

func Benchmark32(b *testing.B) {
	benchmark(b, 32)
}

func BenchmarkAppend32(b *testing.B) {
	benchmarkAppend(b, 32)
}

func BenchmarkSlicesClone32(b *testing.B) {
	benchmarkSlicesClone(b, 32)
}

func Benchmark64(b *testing.B) {
	benchmark(b, 64)
}

func BenchmarkAppend64(b *testing.B) {
	benchmarkAppend(b, 64)
}

func BenchmarkSlicesClone64(b *testing.B) {
	benchmarkSlicesClone(b, 64)
}

func Benchmark128(b *testing.B) {
	benchmark(b, 128)
}

func BenchmarkAppend128(b *testing.B) {
	benchmarkAppend(b, 128)
}

func BenchmarkSlicesClone128(b *testing.B) {
	benchmarkSlicesClone(b, 128)
}

func benchmark(b *testing.B, quantity int) {
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
