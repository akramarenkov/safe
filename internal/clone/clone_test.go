package clone

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5}
	original := []int{0, 1, 2, 3, 4, 5}

	copied := Slice(original)

	require.Equal(t, expected, copied)
	require.Equal(t, expected, original)
	require.NotSame(t, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark(b *testing.B) {
	slice := make([]bool, 1<<6)

	b.ResetTimer()

	for range b.N {
		slice = Slice(slice)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, slice)
}

func BenchmarkCloneSliceAppend(b *testing.B) {
	slice := make([]bool, 1<<6)

	b.ResetTimer()

	for range b.N {
		slice = append([]bool(nil), slice...)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, slice)
}
