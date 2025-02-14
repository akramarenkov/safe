package research

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindLastLosslessly(t *testing.T) {
	number, finded, err := FindLastLosslessly[float32, uint64](1, 0)
	require.NoError(t, err)
	require.True(t, finded)
	require.Equal(t, uint64(10000000000), number)

	number, finded, err = FindLastLosslessly[float64, uint64](1, 0)
	require.NoError(t, err)
	require.True(t, finded)
	require.Equal(t, uint64(10000000000000000000), number)

	number, finded, err = FindLastLosslessly[float32, uint64](2, 0)
	require.NoError(t, err)
	require.True(t, finded)
	require.Equal(t, uint64(1<<24), number)

	number, finded, err = FindLastLosslessly[float64, uint64](2, 0)
	require.NoError(t, err)
	require.True(t, finded)
	require.Equal(t, uint64(1<<53), number)

	number, finded, err = FindLastLosslessly[float32, uint64](1000, 0)
	require.NoError(t, err)
	require.True(t, finded)
	require.Equal(t, uint64(1<<24), number)

	number, finded, err = FindLastLosslessly[float64, uint64](1000, 0)
	require.NoError(t, err)
	require.True(t, finded)
	require.Equal(t, uint64(1<<53), number)
}

func TestFindLastLosslesslyNegative(t *testing.T) {
	number, finded, err := FindLastLosslessly[float32, int64](1, -1)
	require.NoError(t, err)
	require.True(t, finded)
	require.Equal(t, int64(-10000000000), number)

	number, finded, err = FindLastLosslessly[float64, int64](1, -1)
	require.NoError(t, err)
	require.True(t, finded)
	require.Equal(t, int64(-1000000000000000000), number)

	number, finded, err = FindLastLosslessly[float32, int64](2, -1)
	require.NoError(t, err)
	require.True(t, finded)
	require.Equal(t, int64(-1<<24), number)

	number, finded, err = FindLastLosslessly[float64, int64](2, -1)
	require.NoError(t, err)
	require.True(t, finded)
	require.Equal(t, int64(-1<<53), number)

	number, finded, err = FindLastLosslessly[float32, int64](1000, -1)
	require.NoError(t, err)
	require.True(t, finded)
	require.Equal(t, int64(-1<<24), number)

	number, finded, err = FindLastLosslessly[float64, int64](1000, -1)
	require.NoError(t, err)
	require.True(t, finded)
	require.Equal(t, int64(-1<<53), number)
}

func TestFindLastLosslesslyUnsuccessful(t *testing.T) {
	_, _, err := FindLastLosslessly[float64, int8](-100, -1)
	require.Error(t, err)

	_, _, err = FindLastLosslessly[float64, uint8](100, 0)
	require.Error(t, err)

	_, _, err = FindLastLosslessly[float64, int8](100, -1)
	require.Error(t, err)

	_, finded, err := FindLastLosslessly[float64, int8](10, -1)
	require.NoError(t, err)
	require.False(t, finded)

	_, finded, err = FindLastLosslessly[float64, uint8](10, 1)
	require.NoError(t, err)
	require.False(t, finded)
}

func TestIsSequenceLosslessly(t *testing.T) {
	conclusion, err := IsSequenceLosslessly[float32](1<<24, 1<<25)
	require.NoError(t, err)
	require.False(t, conclusion)

	conclusion, err = IsSequenceLosslessly[float64, int64](1<<53, 1<<54)
	require.NoError(t, err)
	require.False(t, conclusion)

	conclusion, err = IsSequenceLosslessly[float32](-1<<24, 1<<24)
	require.NoError(t, err)
	require.True(t, conclusion)

	conclusion, err = IsSequenceLosslessly[float64, uint32](0, math.MaxUint32)
	require.NoError(t, err)
	require.True(t, conclusion)

	conclusion, err = IsSequenceLosslessly[float64, int64](-1<<53, -1<<53+1<<24)
	require.NoError(t, err)
	require.True(t, conclusion)

	conclusion, err = IsSequenceLosslessly[float64, int64](1<<53-1<<24, 1<<53)
	require.NoError(t, err)
	require.True(t, conclusion)
}

func TestIsSequenceLosslesslyError(t *testing.T) {
	conclusion, err := IsSequenceLosslessly[float32](-1<<24, -1<<25)
	require.Error(t, err)
	require.False(t, conclusion)
}

func TestPassThroughFloat(t *testing.T) {
	require.Equal(t, -3, passThroughFloat[float32](-3))
	require.Equal(t, -3, passThroughFloat[float64](-3))
	require.Equal(t, -2, passThroughFloat[float32](-2))
	require.Equal(t, -2, passThroughFloat[float64](-2))
	require.Equal(t, -1, passThroughFloat[float32](-1))
	require.Equal(t, -1, passThroughFloat[float64](-1))
	require.Equal(t, 0, passThroughFloat[float32](0))
	require.Equal(t, 0, passThroughFloat[float64](0))
	require.Equal(t, 1, passThroughFloat[float32](1))
	require.Equal(t, 1, passThroughFloat[float64](1))
	require.Equal(t, 2, passThroughFloat[float32](2))
	require.Equal(t, 2, passThroughFloat[float64](2))
	require.Equal(t, 3, passThroughFloat[float32](3))
	require.Equal(t, 3, passThroughFloat[float64](3))

	require.Equal(
		t,
		math.MinInt32,
		passThroughFloat[float32](math.MinInt32),
	)

	// There is a loss of precision
	require.NotEqual(
		t,
		math.MaxInt32,
		passThroughFloat[float32](math.MaxInt32),
	)
	require.NotEqual(
		t,
		uint32(math.MaxUint32),
		passThroughFloat[float32](uint32(math.MaxUint32)),
	)

	require.Equal(
		t,
		math.MinInt32,
		passThroughFloat[float64](math.MinInt32),
	)
	require.Equal(
		t,
		math.MaxInt32,
		passThroughFloat[float64](math.MaxInt32),
	)
	require.Equal(
		t,
		uint32(math.MaxUint32),
		passThroughFloat[float64](uint32(math.MaxUint32)),
	)

	require.Equal(
		t,
		int64(math.MinInt64),
		passThroughFloat[float64](int64(math.MinInt64)),
	)

	// There is a loss of precision
	require.NotEqual(
		t,
		int64(math.MaxInt64),
		passThroughFloat[float64](int64(math.MaxInt64)),
	)
	require.NotEqual(
		t,
		uint64(math.MaxUint64),
		passThroughFloat[float64](uint64(math.MaxUint64)),
	)
}
