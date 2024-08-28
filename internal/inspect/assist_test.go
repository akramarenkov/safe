package inspect

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPickUpRange(t *testing.T) {
	min, max := PickUpRange[int8]()
	require.Equal(t, int64(math.MinInt8), min)
	require.Equal(t, int64(math.MaxInt8), max)

	min, max = PickUpRange[uint8]()
	require.Equal(t, int64(0), min)
	require.Equal(t, int64(math.MaxUint8), max)

	min, max = PickUpRange[int16]()
	require.Equal(t, int64(math.MinInt16), min)
	require.Equal(t, int64(math.MaxInt16), max)

	min, max = PickUpRange[uint16]()
	require.Equal(t, int64(0), min)
	require.Equal(t, int64(math.MaxUint16), max)
}
