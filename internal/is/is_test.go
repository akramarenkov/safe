package is

import (
	"math"
	"testing"

	"github.com/akramarenkov/safe/internal/iterator"
	"github.com/stretchr/testify/require"
)

func TestMin(t *testing.T) {
	testMinSig(t)
	testMinUns(t)
}

func testMinSig(t *testing.T) {
	for number := range iterator.Iter[int8](math.MinInt8, math.MaxInt8) {
		if number == math.MinInt8 {
			require.True(t, Min(number))
			continue
		}

		require.False(t, Min(number))
	}
}

func testMinUns(t *testing.T) {
	for number := range iterator.Iter[uint8](0, math.MaxUint8) {
		if number == 0 {
			require.True(t, Min(number))
			continue
		}

		require.False(t, Min(number))
	}
}

func TestMax(t *testing.T) {
	testMaxSig(t)
	testMaxUns(t)
}

func testMaxSig(t *testing.T) {
	for number := range iterator.Iter[int8](math.MinInt8, math.MaxInt8) {
		if number == math.MaxInt8 {
			require.True(t, Max(number))
			continue
		}

		require.False(t, Max(number))
	}
}

func testMaxUns(t *testing.T) {
	for number := range iterator.Iter[uint8](0, math.MaxUint8) {
		if number == math.MaxUint8 {
			require.True(t, Max(number))
			continue
		}

		require.False(t, Max(number))
	}
}

func TestMinusOne(t *testing.T) {
	testMinusOneSig(t)
	testMinusOneUns(t)
}

func testMinusOneSig(t *testing.T) {
	for number := range iterator.Iter[int8](math.MinInt8, math.MaxInt8) {
		if number == -1 {
			require.True(t, MinusOne(number))
			continue
		}

		require.False(t, MinusOne(number))
	}
}

func testMinusOneUns(t *testing.T) {
	for number := range iterator.Iter[uint8](0, math.MaxUint8) {
		require.False(t, MinusOne(number))
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

func TestSigned(t *testing.T) {
	require.True(t, Signed[int]())
	require.False(t, Signed[uint]())
}

func BenchmarkReference(b *testing.B) {
	conclusion := false

	for range b.N {
		if b.N == math.MinInt {
			conclusion = true
		}
	}

	require.NotNil(b, conclusion)
}

func BenchmarkMin(b *testing.B) {
	conclusion := false

	for range b.N {
		conclusion = Min(b.N)
	}

	require.NotNil(b, conclusion)
}

func BenchmarkMax(b *testing.B) {
	conclusion := false

	for range b.N {
		conclusion = Max(b.N)
	}

	require.NotNil(b, conclusion)
}

func BenchmarkMinusOne(b *testing.B) {
	conclusion := false

	for range b.N {
		conclusion = MinusOne(b.N)
	}

	require.NotNil(b, conclusion)
}

func BenchmarkEven(b *testing.B) {
	conclusion := false

	for range b.N {
		conclusion = Even(b.N)
	}

	require.NotNil(b, conclusion)
}

func BenchmarkSigned(b *testing.B) {
	conclusion := false

	for range b.N {
		conclusion = Signed[int]()
	}

	require.NotNil(b, conclusion)
}
