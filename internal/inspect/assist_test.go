package inspect

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPickUpRange(t *testing.T) {
	min, max := PickUpRange[int8, int64]()
	require.Equal(t, int64(math.MinInt8), min)
	require.Equal(t, int64(math.MaxInt8), max)

	min, max = PickUpRange[uint8, int64]()
	require.Equal(t, int64(0), min)
	require.Equal(t, int64(math.MaxUint8), max)

	min, max = PickUpRange[int16, int64]()
	require.Equal(t, int64(math.MinInt16), min)
	require.Equal(t, int64(math.MaxInt16), max)

	min, max = PickUpRange[uint16, int64]()
	require.Equal(t, int64(0), min)
	require.Equal(t, int64(math.MaxUint16), max)

	min, max = PickUpRange[int32, int64]()
	require.Equal(t, int64(math.MinInt32), min)
	require.Equal(t, int64(math.MaxInt32), max)

	min, max = PickUpRange[uint32, int64]()
	require.Equal(t, int64(0), min)
	require.Equal(t, int64(math.MaxUint32), max)
}

func TestPickUpRangeFloat(t *testing.T) {
	min, max := PickUpRange[int8, float64]()
	require.InDelta(t, float64(math.MinInt8), min, 0)
	require.InDelta(t, float64(math.MaxInt8), max, 0)

	min, max = PickUpRange[uint8, float64]()
	require.InDelta(t, float64(0), min, 0)
	require.InDelta(t, float64(math.MaxUint8), max, 0)

	min, max = PickUpRange[int16, float64]()
	require.InDelta(t, float64(math.MinInt16), min, 0)
	require.InDelta(t, float64(math.MaxInt16), max, 0)

	min, max = PickUpRange[uint16, float64]()
	require.InDelta(t, float64(0), min, 0)
	require.InDelta(t, float64(math.MaxUint16), max, 0)

	min, max = PickUpRange[int32, float64]()
	require.InDelta(t, float64(math.MinInt32), min, 0)
	require.InDelta(t, float64(math.MaxInt32), max, 0)

	min, max = PickUpRange[uint32, float64]()
	require.InDelta(t, float64(0), min, 0)
	require.InDelta(t, float64(math.MaxUint32), max, 0)
}
