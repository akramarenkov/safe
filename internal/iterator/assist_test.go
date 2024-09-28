package iterator

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToUint64(t *testing.T) {
	require.Equal(t, uint64(math.MaxInt8), toUint64[int8](math.MaxInt8))
	require.Equal(t, uint64(-(math.MinInt8 + 1)), toUint64[int8](math.MinInt8+1))
	require.Equal(t, uint64(-math.MinInt8), toUint64[int8](math.MinInt8))

	require.Equal(t, uint64(math.MaxInt64), toUint64[int64](math.MaxInt64))
	require.Equal(t, uint64(-(math.MinInt64 + 1)), toUint64[int64](math.MinInt64+1))
	require.Equal(t, uint64(-math.MinInt64), toUint64[int64](math.MinInt64))
}
