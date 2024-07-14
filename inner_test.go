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

func TestGetMinusOne(t *testing.T) {
	value8, successfully := getMinusOne[int8]()
	require.True(t, successfully)
	require.Equal(t, int8(-1), value8)

	_, successfully = getMinusOne[uint8]()
	require.False(t, successfully)

	value16, successfully := getMinusOne[int16]()
	require.True(t, successfully)
	require.Equal(t, int16(-1), value16)

	_, successfully = getMinusOne[uint16]()
	require.False(t, successfully)

	value32, successfully := getMinusOne[int32]()
	require.True(t, successfully)
	require.Equal(t, int32(-1), value32)

	_, successfully = getMinusOne[uint32]()
	require.False(t, successfully)

	value, successfully := getMinusOne[int]()
	require.True(t, successfully)
	require.Equal(t, int(-1), value)

	_, successfully = getMinusOne[uint]()
	require.False(t, successfully)

	value64, successfully := getMinusOne[int64]()
	require.True(t, successfully)
	require.Equal(t, int64(-1), value64)

	_, successfully = getMinusOne[uint64]()
	require.False(t, successfully)
}

func TestGetMinusOneUnsure(t *testing.T) {
	require.Equal(t, int8(-1), getMinusOneUnsure[int8]())
	require.Equal(t, uint8(math.MaxUint8), getMinusOneUnsure[uint8]())

	require.Equal(t, int16(-1), getMinusOneUnsure[int16]())
	require.Equal(t, uint16(math.MaxUint16), getMinusOneUnsure[uint16]())

	require.Equal(t, int32(-1), getMinusOneUnsure[int32]())
	require.Equal(t, uint32(math.MaxUint32), getMinusOneUnsure[uint32]())

	require.Equal(t, int(-1), getMinusOneUnsure[int]())
	require.Equal(t, uint(math.MaxUint), getMinusOneUnsure[uint]())

	require.Equal(t, int64(-1), getMinusOneUnsure[int64]())
	require.Equal(t, uint64(math.MaxUint64), getMinusOneUnsure[uint64]())
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

func BenchmarkIsMinReference(b *testing.B) {
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

func BenchmarkIsMaxReference(b *testing.B) {
	conclusion := false

	// b.N, conclusion and require is used to prevent compiler optimizations
	for range b.N {
		if b.N == math.MaxInt {
			conclusion = true
		}
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

func BenchmarkIsMinusOneReference(b *testing.B) {
	conclusion := false

	// b.N, conclusion and require is used to prevent compiler optimizations
	for range b.N {
		if b.N == -1 {
			conclusion = true
		}
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

func BenchmarkGetMinusOne(b *testing.B) {
	value := 0

	// value and require is used to prevent compiler optimizations
	for range b.N {
		value, _ = getMinusOne[int]()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, value)
}

func BenchmarkGetMinusOneUnsure(b *testing.B) {
	value := 0

	// value and require is used to prevent compiler optimizations
	for range b.N {
		value = getMinusOneUnsure[int]()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, value)
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
