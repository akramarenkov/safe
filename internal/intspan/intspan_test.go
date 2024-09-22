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
	minimum8, maximum8, size := Get[int8]()
	require.Equal(t, int8(math.MinInt8), minimum8)
	require.Equal(t, int8(math.MaxInt8), maximum8)
	require.Equal(t, 8, size)

	minimum16, maximum16, size := Get[int16]()
	require.Equal(t, int16(math.MinInt16), minimum16)
	require.Equal(t, int16(math.MaxInt16), maximum16)
	require.Equal(t, 16, size)

	minimum32, maximum32, size := Get[int32]()
	require.Equal(t, int32(math.MinInt32), minimum32)
	require.Equal(t, int32(math.MaxInt32), maximum32)
	require.Equal(t, 32, size)

	minimum, maximum, size := Get[int]()
	require.Equal(t, math.MinInt, minimum)
	require.Equal(t, math.MaxInt, maximum)
	require.Equal(t, BitSizeInt, size)

	minimum64, maximum64, size := Get[int64]()
	require.Equal(t, int64(math.MinInt64), minimum64)
	require.Equal(t, int64(math.MaxInt64), maximum64)
	require.Equal(t, 64, size)
}

func testGetUint(t *testing.T) {
	minimum8, maximum8, size := Get[uint8]()
	require.Equal(t, uint8(0), minimum8)
	require.Equal(t, uint8(math.MaxUint8), maximum8)
	require.Equal(t, 8, size)

	minimum16, maximum16, size := Get[uint16]()
	require.Equal(t, uint16(0), minimum16)
	require.Equal(t, uint16(math.MaxUint16), maximum16)
	require.Equal(t, 16, size)

	minimum32, maximum32, size := Get[uint32]()
	require.Equal(t, uint32(0), minimum32)
	require.Equal(t, uint32(math.MaxUint32), maximum32)
	require.Equal(t, 32, size)

	minimum, maximum, size := Get[uint]()
	require.Equal(t, uint(0), minimum)
	require.Equal(t, uint(math.MaxUint), maximum)
	require.Equal(t, BitSizeInt, size)

	minimum64, maximum64, size := Get[uint64]()
	require.Equal(t, uint64(0), minimum64)
	require.Equal(t, uint64(math.MaxUint64), maximum64)
	require.Equal(t, 64, size)
}

func BenchmarkGet(b *testing.B) {
	var (
		minimum uint8
		maximum uint8
		size    int
	)

	for range b.N {
		minimum, maximum, size = Get[uint8]()
	}

	require.NotNil(b, minimum)
	require.NotNil(b, maximum)
	require.NotNil(b, size)
}
