package intspan

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	testGetInt(t)
	testGetUint(t)
}

func testGetInt(t *testing.T) {
	min8, max8, size := Get[int8]()
	require.Equal(t, int8(math.MinInt8), min8)
	require.Equal(t, int8(math.MaxInt8), max8)
	require.Equal(t, 8, size)

	min16, max16, size := Get[int16]()
	require.Equal(t, int16(math.MinInt16), min16)
	require.Equal(t, int16(math.MaxInt16), max16)
	require.Equal(t, 16, size)

	min32, max32, size := Get[int32]()
	require.Equal(t, int32(math.MinInt32), min32)
	require.Equal(t, int32(math.MaxInt32), max32)
	require.Equal(t, 32, size)

	min, max, size := Get[int]()
	require.Equal(t, math.MinInt, min)
	require.Equal(t, math.MaxInt, max)
	require.Equal(t, BitSizeInt, size)

	min64, max64, size := Get[int64]()
	require.Equal(t, int64(math.MinInt64), min64)
	require.Equal(t, int64(math.MaxInt64), max64)
	require.Equal(t, 64, size)
}

func testGetUint(t *testing.T) {
	min8, max8, size := Get[uint8]()
	require.Equal(t, uint8(0), min8)
	require.Equal(t, uint8(math.MaxUint8), max8)
	require.Equal(t, 8, size)

	min16, max16, size := Get[uint16]()
	require.Equal(t, uint16(0), min16)
	require.Equal(t, uint16(math.MaxUint16), max16)
	require.Equal(t, 16, size)

	min32, max32, size := Get[uint32]()
	require.Equal(t, uint32(0), min32)
	require.Equal(t, uint32(math.MaxUint32), max32)
	require.Equal(t, 32, size)

	min, max, size := Get[uint]()
	require.Equal(t, uint(0), min)
	require.Equal(t, uint(math.MaxUint), max)
	require.Equal(t, BitSizeInt, size)

	min64, max64, size := Get[uint64]()
	require.Equal(t, uint64(0), min64)
	require.Equal(t, uint64(math.MaxUint64), max64)
	require.Equal(t, 64, size)
}

func BenchmarkGet(b *testing.B) {
	var (
		min  uint8
		max  uint8
		size int
	)

	for range b.N {
		min, max, size = Get[uint8]()
	}

	require.NotNil(b, min)
	require.NotNil(b, max)
	require.NotNil(b, size)
}
