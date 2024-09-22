package inspect

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvSpan(t *testing.T) {
	minimum, maximum := ConvSpan[int8, int64]()
	require.Equal(t, int64(math.MinInt8), minimum)
	require.Equal(t, int64(math.MaxInt8), maximum)

	minimum, maximum = ConvSpan[uint8, int64]()
	require.Equal(t, int64(0), minimum)
	require.Equal(t, int64(math.MaxUint8), maximum)

	minimum, maximum = ConvSpan[int16, int64]()
	require.Equal(t, int64(math.MinInt16), minimum)
	require.Equal(t, int64(math.MaxInt16), maximum)

	minimum, maximum = ConvSpan[uint16, int64]()
	require.Equal(t, int64(0), minimum)
	require.Equal(t, int64(math.MaxUint16), maximum)

	minimum, maximum = ConvSpan[int32, int64]()
	require.Equal(t, int64(math.MinInt32), minimum)
	require.Equal(t, int64(math.MaxInt32), maximum)

	minimum, maximum = ConvSpan[uint32, int64]()
	require.Equal(t, int64(0), minimum)
	require.Equal(t, int64(math.MaxUint32), maximum)
}

func TestConvSpanFloat(t *testing.T) {
	minimum, maximum := ConvSpan[int8, float64]()
	require.InDelta(t, float64(math.MinInt8), minimum, 0)
	require.InDelta(t, float64(math.MaxInt8), maximum, 0)

	minimum, maximum = ConvSpan[uint8, float64]()
	require.InDelta(t, float64(0), minimum, 0)
	require.InDelta(t, float64(math.MaxUint8), maximum, 0)

	minimum, maximum = ConvSpan[int16, float64]()
	require.InDelta(t, float64(math.MinInt16), minimum, 0)
	require.InDelta(t, float64(math.MaxInt16), maximum, 0)

	minimum, maximum = ConvSpan[uint16, float64]()
	require.InDelta(t, float64(0), minimum, 0)
	require.InDelta(t, float64(math.MaxUint16), maximum, 0)

	minimum, maximum = ConvSpan[int32, float64]()
	require.InDelta(t, float64(math.MinInt32), minimum, 0)
	require.InDelta(t, float64(math.MaxInt32), maximum, 0)

	minimum, maximum = ConvSpan[uint32, float64]()
	require.InDelta(t, float64(0), minimum, 0)
	require.InDelta(t, float64(math.MaxUint32), maximum, 0)
}
