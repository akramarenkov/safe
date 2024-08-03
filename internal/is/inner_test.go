package is

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMin(t *testing.T) {
	TestMinInt(t)
	TestMinUint(t)
}

func TestMinInt(t *testing.T) {
	for number := math.MinInt8; number <= math.MaxInt8; number++ {
		if number == math.MinInt8 {
			require.True(t, Min(int8(number)))
			continue
		}

		require.False(t, Min(int8(number)))
	}
}

func TestMinUint(t *testing.T) {
	for number := 0; number <= math.MaxUint8; number++ {
		if number == 0 {
			require.True(t, Min(uint8(number)))
			continue
		}

		require.False(t, Min(uint8(number)))
	}
}

func TestMax(t *testing.T) {
	testMaxInt(t)
	testMaxUint(t)
}

func testMaxInt(t *testing.T) {
	for number := math.MinInt8; number <= math.MaxInt8; number++ {
		if number == math.MaxInt8 {
			require.True(t, Max(int8(number)))
			continue
		}

		require.False(t, Max(int8(number)))
	}
}

func testMaxUint(t *testing.T) {
	for number := 0; number <= math.MaxUint8; number++ {
		if number == math.MaxUint8 {
			require.True(t, Max(uint8(number)))
			continue
		}

		require.False(t, Max(uint8(number)))
	}
}

func TestMinusOne(t *testing.T) {
	testMinusOneInt(t)
	testMinusOneUint(t)
}

func testMinusOneInt(t *testing.T) {
	for number := math.MinInt8; number <= math.MaxInt8; number++ {
		if number == -1 {
			require.True(t, MinusOne(int8(number)))
			continue
		}

		require.False(t, MinusOne(int8(number)))
	}
}

func testMinusOneUint(t *testing.T) {
	for number := 0; number <= math.MaxUint8; number++ {
		require.False(t, MinusOne(uint8(number)))
	}
}

func TestEven(t *testing.T) {
	require.True(t, Even(-8))
	require.False(t, Even(-7))
	require.True(t, Even(-6))
	require.False(t, Even(-5))
	require.True(t, Even(-4))
	require.False(t, Even(-3))
	require.True(t, Even(-2))
	require.False(t, Even(-1))
	require.True(t, Even(0))
	require.False(t, Even(1))
	require.True(t, Even(2))
	require.False(t, Even(3))
	require.True(t, Even(4))
	require.False(t, Even(5))
	require.True(t, Even(6))
	require.False(t, Even(7))
	require.True(t, Even(8))
}

func BenchmarkReference(b *testing.B) {
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

func BenchmarkMin(b *testing.B) {
	conclusion := false

	// b.N, conclusion and require is used to prevent compiler optimizations
	for range b.N {
		conclusion = Min(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, conclusion)
}

func BenchmarkMax(b *testing.B) {
	conclusion := false

	// b.N, conclusion and require is used to prevent compiler optimizations
	for range b.N {
		conclusion = Max(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, conclusion)
}

func BenchmarkMinusOne(b *testing.B) {
	conclusion := false

	// b.N, conclusion and require is used to prevent compiler optimizations
	for range b.N {
		conclusion = MinusOne(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, conclusion)
}

func BenchmarkEven(b *testing.B) {
	conclusion := false

	// b.N, conclusion and require is used to prevent compiler optimizations
	for range b.N {
		conclusion = Even(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, conclusion)
}
