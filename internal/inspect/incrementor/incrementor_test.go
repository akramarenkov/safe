package incrementor

import (
	"math"
	"testing"

	"github.com/akramarenkov/safe/internal/iterator"

	"github.com/stretchr/testify/require"
)

func TestIncrementor(t *testing.T) {
	testIncrementorSig(t, math.MinInt8, math.MaxInt8)
	testIncrementorUns(t, 0, math.MaxUint8)
	testIncrementorSig(t, 0, 0)
	testIncrementorUns(t, 0, 0)
	testIncrementorSig(t, 0, 1)
	testIncrementorUns(t, 0, 1)
	testIncrementorSig(t, -1, 0)
	testIncrementorSig(t, -1, 1)
}

func testIncrementorSig(t *testing.T, begin, end int8) {
	incrementor, err := New(3, begin, end)
	require.NoError(t, err)

	for first := range iterator.Iter(begin, end) {
		for second := range iterator.Iter(begin, end) {
			for third := range iterator.Iter(begin, end) {
				err := incrementor.Test(first, second, third)
				if err != nil {
					require.NoError(t, err)
				}
			}
		}
	}
}

func testIncrementorUns(t *testing.T, begin, end uint8) {
	incrementor, err := New(3, begin, end)
	require.NoError(t, err)

	for first := range iterator.Iter(begin, end) {
		for second := range iterator.Iter(begin, end) {
			for third := range iterator.Iter(begin, end) {
				err := incrementor.Test(first, second, third)
				if err != nil {
					require.NoError(t, err)
				}
			}
		}
	}
}

func TestIncrementorError(t *testing.T) {
	_, err := New[int8](3, 1, -1)
	require.Error(t, err)

	incrementor, err := New[int8](3, -1, 1)
	require.NoError(t, err)

	err = incrementor.Test(-1, -1)
	require.Error(t, err)

	err = incrementor.Test(-1, -1, -1, -1)
	require.Error(t, err)

	err = incrementor.Test(0, 0, 0)
	require.Error(t, err)
}
