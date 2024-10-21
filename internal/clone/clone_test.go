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

func Benchmark64(b *testing.B) {
	benchmark(b, 1<<6)
}

func BenchmarkAppend64(b *testing.B) {
	benchmarkAppend(b, 1<<6)
}

func BenchmarkSlicesClone64(b *testing.B) {
	benchmarkSlicesClone(b, 1<<6)
}

func Benchmark256(b *testing.B) {
	benchmark(b, 1<<8)
}

func BenchmarkAppend256(b *testing.B) {
	benchmarkAppend(b, 1<<8)
}

func BenchmarkSlicesClone256(b *testing.B) {
	benchmarkSlicesClone(b, 1<<8)
}

func Benchmark512(b *testing.B) {
	benchmark(b, 1<<9)
}

func BenchmarkAppend512(b *testing.B) {
	benchmarkAppend(b, 1<<9)
}

func BenchmarkSlicesClone512(b *testing.B) {
	benchmarkSlicesClone(b, 1<<9)
}

func Benchmark1024(b *testing.B) {
	benchmark(b, 1<<10)
}

func BenchmarkAppend1024(b *testing.B) {
	benchmarkAppend(b, 1<<10)
}

func BenchmarkSlicesClone1024(b *testing.B) {
	benchmarkSlicesClone(b, 1<<10)
}

func benchmark(b *testing.B, size int) {
	slice := make([]bool, size)

	for range b.N {
		slice = Slice(slice)
	}

	require.NotNil(b, slice)
}

func benchmarkAppend(b *testing.B, size int) {
	slice := make([]bool, size)

	for range b.N {
		slice = append([]bool(nil), slice...)
	}

	require.NotNil(b, slice)
}

func benchmarkSlicesClone(b *testing.B, size int) {
	slice := make([]bool, size)

	for range b.N {
		slice = slices.Clone(slice)
	}

	require.NotNil(b, slice)
}
