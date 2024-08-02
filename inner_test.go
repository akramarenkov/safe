package safe

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsMin(t *testing.T) {
	testIsMinInt(t)
	testIsMinUint(t)
}

func testIsMinInt(t *testing.T) {
	for number := math.MinInt8; number <= math.MaxInt8; number++ {
		if number == math.MinInt8 {
			require.True(t, isMin(int8(number)))
			continue
		}

		require.False(t, isMin(int8(number)))
	}
}

func testIsMinUint(t *testing.T) {
	for number := 0; number <= math.MaxUint8; number++ {
		if number == 0 {
			require.True(t, isMin(uint8(number)))
			continue
		}

		require.False(t, isMin(uint8(number)))
	}
}

func TestIsMax(t *testing.T) {
	testIsMaxInt(t)
	testIsMaxUint(t)
}

func testIsMaxInt(t *testing.T) {
	for number := math.MinInt8; number <= math.MaxInt8; number++ {
		if number == math.MaxInt8 {
			require.True(t, isMax(int8(number)))
			continue
		}

		require.False(t, isMax(int8(number)))
	}
}

func testIsMaxUint(t *testing.T) {
	for number := 0; number <= math.MaxUint8; number++ {
		if number == math.MaxUint8 {
			require.True(t, isMax(uint8(number)))
			continue
		}

		require.False(t, isMax(uint8(number)))
	}
}

func TestIsMinusOne(t *testing.T) {
	testIsMinusOneInt(t)
	testIsMinusOneUint(t)
}

func testIsMinusOneInt(t *testing.T) {
	for number := math.MinInt8; number <= math.MaxInt8; number++ {
		if number == -1 {
			require.True(t, isMinusOne(int8(number)))
			continue
		}

		require.False(t, isMinusOne(int8(number)))
	}
}

func testIsMinusOneUint(t *testing.T) {
	for number := 0; number <= math.MaxUint8; number++ {
		require.False(t, isMinusOne(uint8(number)))
	}
}

func TestIsEven(t *testing.T) {
	require.True(t, isEven(-8))
	require.False(t, isEven(-7))
	require.True(t, isEven(-6))
	require.False(t, isEven(-5))
	require.True(t, isEven(-4))
	require.False(t, isEven(-3))
	require.True(t, isEven(-2))
	require.False(t, isEven(-1))
	require.True(t, isEven(0))
	require.False(t, isEven(1))
	require.True(t, isEven(2))
	require.False(t, isEven(3))
	require.True(t, isEven(4))
	require.False(t, isEven(5))
	require.True(t, isEven(6))
	require.False(t, isEven(7))
	require.True(t, isEven(8))
}

func TestCloneSlice(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5}
	original := []int{0, 1, 2, 3, 4, 5}

	copied := cloneSlice(original)

	require.Equal(t, expected, copied)
	require.Equal(t, expected, original)
	require.NotSame(t, original, copied)
}

func BenchmarkIsReference(b *testing.B) {
	conclusion := false

	// b.N, conclusion and require is used to prevent compiler optimizations
	for range b.N {
		if b.N == math.MinInt {
			conclusion = true
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, conclusion)
}

func BenchmarkIsMin(b *testing.B) {
	conclusion := false

	// b.N, conclusion and require is used to prevent compiler optimizations
	for range b.N {
		conclusion = isMin(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, conclusion)
}

func BenchmarkIsMax(b *testing.B) {
	conclusion := false

	// b.N, conclusion and require is used to prevent compiler optimizations
	for range b.N {
		conclusion = isMax(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, conclusion)
}

func BenchmarkIsMinusOne(b *testing.B) {
	conclusion := false

	// b.N, conclusion and require is used to prevent compiler optimizations
	for range b.N {
		conclusion = isMinusOne(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, conclusion)
}

func BenchmarkIsEven(b *testing.B) {
	conclusion := false

	// b.N, conclusion and require is used to prevent compiler optimizations
	for range b.N {
		conclusion = isEven(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, conclusion)
}

func BenchmarkCloneSlice(b *testing.B) {
	slice := make([]bool, 1<<6)

	b.ResetTimer()

	for range b.N {
		slice = cloneSlice(slice)
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
