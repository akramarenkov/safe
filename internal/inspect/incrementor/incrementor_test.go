package incrementor

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIncrementor(t *testing.T) {
	testIncrementorInt(t, math.MinInt8, math.MaxInt8)
	testIncrementorUint(t, 0, math.MaxUint8)
	testIncrementorInt(t, 0, 0)
	testIncrementorUint(t, 0, 0)
	testIncrementorInt(t, 0, 1)
	testIncrementorUint(t, 0, 1)
	testIncrementorInt(t, -1, 0)
	testIncrementorInt(t, -1, 1)
}

func testIncrementorInt(t *testing.T, begin, end int8) {
	incrementor, err := New(3, begin, end)
	require.NoError(t, err)

	for first := int(begin); first <= int(end); first++ {
		for second := int(begin); second <= int(end); second++ {
			for third := int(begin); third <= int(end); third++ {
				err := incrementor.Test(int8(first), int8(second), int8(third))
				if err != nil {
					require.NoError(t, err)
				}
			}
		}
	}
}

func testIncrementorUint(t *testing.T, begin, end uint8) {
	incrementor, err := New(3, begin, end)
	require.NoError(t, err)

	for first := int(begin); first <= int(end); first++ {
		for second := int(begin); second <= int(end); second++ {
			for third := int(begin); third <= int(end); third++ {
				err := incrementor.Test(uint8(first), uint8(second), uint8(third))
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
