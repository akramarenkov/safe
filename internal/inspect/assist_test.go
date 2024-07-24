package inspect

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPickUpRange(t *testing.T) {
	min, max := pickUpRange[int8]()
	require.Equal(t, int64(math.MinInt8), min)
	require.Equal(t, int64(math.MaxInt8), max)

	min, max = pickUpRange[uint8]()
	require.Equal(t, int64(0), min)
	require.Equal(t, int64(math.MaxUint8), max)
}

func TestIsSigned(t *testing.T) {
	require.True(t, isSigned[int8]())
	require.False(t, isSigned[uint8]())
}
