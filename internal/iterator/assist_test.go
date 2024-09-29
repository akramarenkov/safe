package iterator

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestU64(t *testing.T) {
	require.Equal(t, uint64(math.MaxInt8), u64[int8](math.MaxInt8))
	require.Equal(t, uint64(-(math.MinInt8 + 1)), u64[int8](math.MinInt8+1))
	require.Equal(t, uint64(-math.MinInt8), u64[int8](math.MinInt8))

	require.Equal(t, uint64(math.MaxInt64), u64[int64](math.MaxInt64))
	require.Equal(t, uint64(-(math.MinInt64 + 1)), u64[int64](math.MinInt64+1))
	require.Equal(t, uint64(-math.MinInt64), u64[int64](math.MinInt64))
}
