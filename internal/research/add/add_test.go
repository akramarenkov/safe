package add

import (
	"math"
	"testing"

	"github.com/akramarenkov/safe"
	"github.com/akramarenkov/safe/internal/iterator"

	"github.com/stretchr/testify/require"
)

func TestCalcSpanAdd(t *testing.T) {
	for first := range iterator.Iter[int8](math.MinInt8, math.MaxInt8) {
		for second := range iterator.Iter[int8](math.MinInt8, math.MaxInt8) {
			if _, err := safe.Add(first, second); err == nil {
				continue
			}

			overflowed := first + second

			if first > 0 {
				minimum, maximum := calcSpanAddPositive(first, second, overflowed)
				require.Equal(t, int8(math.MinInt8), minimum)
				require.Equal(t, int8(math.MaxInt8), maximum)

				continue
			}

			minimum, maximum := calcSpanAddNegative(first, second, overflowed)
			require.Equal(t, int8(math.MinInt8), minimum)
			require.Equal(t, int8(math.MaxInt8), maximum)
		}
	}
}
