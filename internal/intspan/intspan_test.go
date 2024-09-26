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
	minimum64, maximum64 := Get[int64]()
	require.Equal(t, int64(math.MinInt64), minimum64)
	require.Equal(t, int64(math.MaxInt64), maximum64)

	minimum32, maximum32 := Get[int32]()
	require.Equal(t, int32(math.MinInt32), minimum32)
	require.Equal(t, int32(math.MaxInt32), maximum32)

	minimum16, maximum16 := Get[int16]()
	require.Equal(t, int16(math.MinInt16), minimum16)
	require.Equal(t, int16(math.MaxInt16), maximum16)

	minimum8, maximum8 := Get[int8]()
	require.Equal(t, int8(math.MinInt8), minimum8)
	require.Equal(t, int8(math.MaxInt8), maximum8)

	minimum, maximum := Get[int]()
	require.Equal(t, math.MinInt, minimum)
	require.Equal(t, math.MaxInt, maximum)
}

func testGetUint(t *testing.T) {
	minimum64, maximum64 := Get[uint64]()
	require.Equal(t, uint64(0), minimum64)
	require.Equal(t, uint64(math.MaxUint64), maximum64)

	minimum32, maximum32 := Get[uint32]()
	require.Equal(t, uint32(0), minimum32)
	require.Equal(t, uint32(math.MaxUint32), maximum32)

	minimum16, maximum16 := Get[uint16]()
	require.Equal(t, uint16(0), minimum16)
	require.Equal(t, uint16(math.MaxUint16), maximum16)

	minimum8, maximum8 := Get[uint8]()
	require.Equal(t, uint8(0), minimum8)
	require.Equal(t, uint8(math.MaxUint8), maximum8)

	minimum, maximum := Get[uint]()
	require.Equal(t, uint(0), minimum)
	require.Equal(t, uint(math.MaxUint), maximum)
}

func TestBitSize(t *testing.T) {
	require.Equal(t, bitSize64, BitSize[int64]())
	require.Equal(t, bitSize32, BitSize[int32]())
	require.Equal(t, bitSize16, BitSize[int16]())
	require.Equal(t, bitSize8, BitSize[int8]())
	require.Equal(t, bitSizeInt, BitSize[int]())

	require.Equal(t, bitSize64, BitSize[uint64]())
	require.Equal(t, bitSize32, BitSize[uint32]())
	require.Equal(t, bitSize16, BitSize[uint16]())
	require.Equal(t, bitSize8, BitSize[uint8]())
	require.Equal(t, bitSizeInt, BitSize[uint]())
}

func BenchmarkGet(b *testing.B) {
	var (
		minimum int8
		maximum int8

		minimumU uint8
		maximumU uint8
	)

	for range b.N {
		minimum, maximum = Get[int8]()
		minimumU, maximumU = Get[uint8]()
	}

	require.NotNil(b, minimum)
	require.NotNil(b, maximum)
	require.NotNil(b, minimumU)
	require.NotNil(b, maximumU)
}

func BenchmarkBitSize(b *testing.B) {
	var (
		size  int
		sizeU int
	)

	for range b.N {
		size = BitSize[int8]()
		sizeU = BitSize[uint8]()
	}

	require.NotNil(b, size)
	require.NotNil(b, sizeU)
}
