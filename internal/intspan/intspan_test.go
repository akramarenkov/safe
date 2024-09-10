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
	min8, max8 := Get[int8]()
	require.Equal(t, int8(math.MinInt8), min8)
	require.Equal(t, int8(math.MaxInt8), max8)

	min16, max16 := Get[int16]()
	require.Equal(t, int16(math.MinInt16), min16)
	require.Equal(t, int16(math.MaxInt16), max16)

	min32, max32 := Get[int32]()
	require.Equal(t, int32(math.MinInt32), min32)
	require.Equal(t, int32(math.MaxInt32), max32)

	min, max := Get[int]()
	require.Equal(t, math.MinInt, min)
	require.Equal(t, math.MaxInt, max)

	min64, max64 := Get[int64]()
	require.Equal(t, int64(math.MinInt64), min64)
	require.Equal(t, int64(math.MaxInt64), max64)
}

func testGetUint(t *testing.T) {
	min8, max8 := Get[uint8]()
	require.Equal(t, uint8(0), min8)
	require.Equal(t, uint8(math.MaxUint8), max8)

	min16, max16 := Get[uint16]()
	require.Equal(t, uint16(0), min16)
	require.Equal(t, uint16(math.MaxUint16), max16)

	min32, max32 := Get[uint32]()
	require.Equal(t, uint32(0), min32)
	require.Equal(t, uint32(math.MaxUint32), max32)

	min, max := Get[uint]()
	require.Equal(t, uint(0), min)
	require.Equal(t, uint(math.MaxUint), max)

	min64, max64 := Get[uint64]()
	require.Equal(t, uint64(0), min64)
	require.Equal(t, uint64(math.MaxUint64), max64)
}

func BenchmarkGet(b *testing.B) {
	var (
		min uint8
		max uint8
	)

	for range b.N {
		min, max = Get[uint8]()
	}

	require.NotNil(b, min)
	require.NotNil(b, max)
}
