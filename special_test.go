package safe

import (
	"math"
	"testing"

	"github.com/akramarenkov/safe/internal/iterator"
	"github.com/stretchr/testify/require"
)

func TestDist(t *testing.T) {
	testDistMan(t)
	testDistSig(t)
	testDistUns(t)
}

func testDistMan(t *testing.T) {
	require.Equal(t, uint64(3), Dist[int8](2, 5))
	require.Equal(t, uint64(3), Dist[int8](5, 2))
	require.Equal(t, uint64(7), Dist[int8](-2, 5))
	require.Equal(t, uint64(7), Dist[int8](5, -2))
	require.Equal(t, uint64(3), Dist[int8](-2, -5))
	require.Equal(t, uint64(3), Dist[int8](-5, -2))
	require.Equal(t, uint64(255), Dist[int8](-128, 127))
	require.Equal(t, uint64(255), Dist[int8](127, -128))

	require.Equal(t, uint64(3), Dist[uint8](2, 5))
	require.Equal(t, uint64(3), Dist[uint8](5, 2))
	require.Equal(t, uint64(255), Dist[uint8](0, 255))
	require.Equal(t, uint64(255), Dist[uint8](255, 0))

	require.Equal(t, uint64(math.MaxUint64), Dist[int64](math.MinInt64, math.MaxInt64))
	require.Equal(t, uint64(math.MaxUint64), Dist[int64](math.MaxInt64, math.MinInt64))

	require.Equal(t, uint64(math.MaxUint64), Dist[uint64](0, math.MaxUint64))
	require.Equal(t, uint64(math.MaxUint64), Dist[uint64](math.MaxUint64, 0))
}

func testDistSig(t *testing.T) {
	for first := range iterator.Iter[int8](math.MinInt8, math.MaxInt8) {
		for second := range iterator.Iter[int8](math.MinInt8, math.MaxInt8) {
			reference := int(second) - int(first)

			if first > second {
				reference = int(first) - int(second)
			}

			require.GreaterOrEqual(t, reference, 0)

			//nolint:gosec // Is the result of subtracting the smaller value
			// from the larger value for int8 values
			referenceU := uint64(reference)

			actual := Dist(first, second)
			require.Equal(t, referenceU, actual, "first: %v, second: %v", first, second)
		}
	}
}

func testDistUns(t *testing.T) {
	for first := range iterator.Iter[uint8](0, math.MaxUint8) {
		for second := range iterator.Iter[uint8](0, math.MaxUint8) {
			reference := uint64(second) - uint64(first)

			if first > second {
				reference = uint64(first) - uint64(second)
			}

			actual := Dist(first, second)
			require.Equal(t, reference, actual, "first: %v, second: %v", first, second)
		}
	}
}

func BenchmarkDistReference(b *testing.B) {
	dist := 0

	for range b.N {
		dist = b.N - 1
	}

	require.NotNil(b, dist)
}

func BenchmarkDist(b *testing.B) {
	dist := uint64(0)

	for range b.N {
		dist = Dist(b.N, 1)
	}

	require.NotNil(b, dist)
}
