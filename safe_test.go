package safe

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumInt8(t *testing.T) {
	sum, err := SumInt[int8](0, 0)
	require.NoError(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](126, 0)
	require.NoError(t, err)
	require.Equal(t, int8(126), sum)

	sum, err = SumInt[int8](0, 126)
	require.NoError(t, err)
	require.Equal(t, int8(126), sum)

	sum, err = SumInt[int8](126, 1)
	require.NoError(t, err)
	require.Equal(t, int8(127), sum)

	sum, err = SumInt[int8](1, 126)
	require.NoError(t, err)
	require.Equal(t, int8(127), sum)

	sum, err = SumInt[int8](-127, 0)
	require.NoError(t, err)
	require.Equal(t, int8(-127), sum)

	sum, err = SumInt[int8](0, -127)
	require.NoError(t, err)
	require.Equal(t, int8(-127), sum)

	sum, err = SumInt[int8](-127, -1)
	require.NoError(t, err)
	require.Equal(t, int8(-128), sum)

	sum, err = SumInt[int8](-1, -127)
	require.NoError(t, err)
	require.Equal(t, int8(-128), sum)

	sum, err = SumInt[int8](-128, 127)
	require.NoError(t, err)
	require.Equal(t, int8(-1), sum)

	sum, err = SumInt[int8](127, -128)
	require.NoError(t, err)
	require.Equal(t, int8(-1), sum)
}

func TestSumInt8Overflow(t *testing.T) {
	sum, err := SumInt[int8](127, 1)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](1, 127)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](127, 2)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](2, 127)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](127, 3)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](3, 127)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](127, 125)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](125, 127)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](127, 126)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](126, 127)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](127, 127)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](-128, -1)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](-1, -128)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](-128, -2)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](-2, -128)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](-128, -3)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](-3, -128)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](-128, -126)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](-126, -128)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](-128, -127)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](-127, -128)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)

	sum, err = SumInt[int8](-128, -128)
	require.Error(t, err)
	require.Equal(t, int8(0), sum)
}

func TestSumUint8(t *testing.T) {
	sum, err := SumInt[uint8](0, 0)
	require.NoError(t, err)
	require.Equal(t, uint8(0), sum)

	sum, err = SumInt[uint8](254, 0)
	require.NoError(t, err)
	require.Equal(t, uint8(254), sum)

	sum, err = SumInt[uint8](0, 254)
	require.NoError(t, err)
	require.Equal(t, uint8(254), sum)

	sum, err = SumInt[uint8](254, 1)
	require.NoError(t, err)
	require.Equal(t, uint8(255), sum)

	sum, err = SumInt[uint8](1, 254)
	require.NoError(t, err)
	require.Equal(t, uint8(255), sum)
}

func TestSumUint8Overflow(t *testing.T) {
	sum, err := SumInt[uint8](255, 1)
	require.Error(t, err)
	require.Equal(t, uint8(0), sum)

	sum, err = SumInt[uint8](1, 255)
	require.Error(t, err)
	require.Equal(t, uint8(0), sum)

	sum, err = SumInt[uint8](255, 2)
	require.Error(t, err)
	require.Equal(t, uint8(0), sum)

	sum, err = SumInt[uint8](2, 255)
	require.Error(t, err)
	require.Equal(t, uint8(0), sum)

	sum, err = SumInt[uint8](255, 3)
	require.Error(t, err)
	require.Equal(t, uint8(0), sum)

	sum, err = SumInt[uint8](3, 255)
	require.Error(t, err)
	require.Equal(t, uint8(0), sum)

	sum, err = SumInt[uint8](255, 253)
	require.Error(t, err)
	require.Equal(t, uint8(0), sum)

	sum, err = SumInt[uint8](253, 255)
	require.Error(t, err)
	require.Equal(t, uint8(0), sum)

	sum, err = SumInt[uint8](255, 254)
	require.Error(t, err)
	require.Equal(t, uint8(0), sum)

	sum, err = SumInt[uint8](254, 255)
	require.Error(t, err)
	require.Equal(t, uint8(0), sum)

	sum, err = SumInt[uint8](255, 255)
	require.Error(t, err)
	require.Equal(t, uint8(0), sum)

	sum, err = SumInt[uint8](128, 128)
	require.Error(t, err)
	require.Equal(t, uint8(0), sum)

	sum, err = SumInt[uint8](127, 129)
	require.Error(t, err)
	require.Equal(t, uint8(0), sum)

	sum, err = SumInt[uint8](129, 127)
	require.Error(t, err)
	require.Equal(t, uint8(0), sum)
}

func TestIsMaxNegative(t *testing.T) {
	require.Equal(t, true, IsMaxNegative[int8](-128))
	require.Equal(t, false, IsMaxNegative[int8](-127))
	require.Equal(t, false, IsMaxNegative[int8](-126))
	require.Equal(t, false, IsMaxNegative[int8](-125))
	require.Equal(t, false, IsMaxNegative[int8](-3))
	require.Equal(t, false, IsMaxNegative[int8](-2))
	require.Equal(t, false, IsMaxNegative[int8](-1))
	require.Equal(t, false, IsMaxNegative[int8](0))
	require.Equal(t, false, IsMaxNegative[int8](1))
	require.Equal(t, false, IsMaxNegative[int8](2))
	require.Equal(t, false, IsMaxNegative[int8](3))
	require.Equal(t, false, IsMaxNegative[int8](124))
	require.Equal(t, false, IsMaxNegative[int8](125))
	require.Equal(t, false, IsMaxNegative[int8](126))
	require.Equal(t, false, IsMaxNegative[int8](127))

	require.Equal(t, false, IsMaxNegative[uint8](0))
	require.Equal(t, false, IsMaxNegative[uint8](1))
	require.Equal(t, false, IsMaxNegative[uint8](2))
	require.Equal(t, false, IsMaxNegative[uint8](3))
	require.Equal(t, false, IsMaxNegative[uint8](255))
}

func TestIsMaxPositive(t *testing.T) {
	require.Equal(t, false, IsMaxPositive[int8](-128))
	require.Equal(t, false, IsMaxPositive[int8](-127))
	require.Equal(t, false, IsMaxPositive[int8](-126))
	require.Equal(t, false, IsMaxPositive[int8](-125))
	require.Equal(t, false, IsMaxPositive[int8](-3))
	require.Equal(t, false, IsMaxPositive[int8](-2))
	require.Equal(t, false, IsMaxPositive[int8](-1))
	require.Equal(t, false, IsMaxPositive[int8](0))
	require.Equal(t, false, IsMaxPositive[int8](1))
	require.Equal(t, false, IsMaxPositive[int8](2))
	require.Equal(t, false, IsMaxPositive[int8](3))
	require.Equal(t, false, IsMaxPositive[int8](124))
	require.Equal(t, false, IsMaxPositive[int8](125))
	require.Equal(t, false, IsMaxPositive[int8](126))
	require.Equal(t, true, IsMaxPositive[int8](127))

	require.Equal(t, false, IsMaxPositive[uint8](0))
	require.Equal(t, false, IsMaxPositive[uint8](1))
	require.Equal(t, false, IsMaxPositive[uint8](2))
	require.Equal(t, false, IsMaxPositive[uint8](3))
	require.Equal(t, true, IsMaxPositive[uint8](255))
}

func TestProductInt8(t *testing.T) {
	product, err := ProductInt[int8](0, 0)
	require.NoError(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](2, 0)
	require.NoError(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](0, 2)
	require.NoError(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](3, 0)
	require.NoError(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](0, 3)
	require.NoError(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](-2, 0)
	require.NoError(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](0, -2)
	require.NoError(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](-3, 0)
	require.NoError(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](0, -3)
	require.NoError(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](2, 3)
	require.NoError(t, err)
	require.Equal(t, int8(6), product)

	product, err = ProductInt[int8](3, 2)
	require.NoError(t, err)
	require.Equal(t, int8(6), product)

	product, err = ProductInt[int8](-2, 3)
	require.NoError(t, err)
	require.Equal(t, int8(-6), product)

	product, err = ProductInt[int8](3, -2)
	require.NoError(t, err)
	require.Equal(t, int8(-6), product)

	product, err = ProductInt[int8](2, -3)
	require.NoError(t, err)
	require.Equal(t, int8(-6), product)

	product, err = ProductInt[int8](-3, 2)
	require.NoError(t, err)
	require.Equal(t, int8(-6), product)

	product, err = ProductInt[int8](-2, -3)
	require.NoError(t, err)
	require.Equal(t, int8(6), product)

	product, err = ProductInt[int8](-3, -2)
	require.NoError(t, err)
	require.Equal(t, int8(6), product)

	product, err = ProductInt[int8](1, -127)
	require.NoError(t, err)
	require.Equal(t, int8(-127), product)

	product, err = ProductInt[int8](-127, 1)
	require.NoError(t, err)
	require.Equal(t, int8(-127), product)

	product, err = ProductInt[int8](127, -1)
	require.NoError(t, err)
	require.Equal(t, int8(-127), product)

	product, err = ProductInt[int8](-1, 127)
	require.NoError(t, err)
	require.Equal(t, int8(-127), product)

	product, err = ProductInt[int8](-127, -1)
	require.NoError(t, err)
	require.Equal(t, int8(127), product)

	product, err = ProductInt[int8](-1, -127)
	require.NoError(t, err)
	require.Equal(t, int8(127), product)

	product, err = ProductInt[int8](-128, 1)
	require.NoError(t, err)
	require.Equal(t, int8(-128), product)

	product, err = ProductInt[int8](1, -128)
	require.NoError(t, err)
	require.Equal(t, int8(-128), product)
}

func TestProductInt8Overflow(t *testing.T) {
	product, err := ProductInt[int8](127, 2)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](2, 127)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](64, 2)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](2, 64)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](127, 127)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](-127, 2)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](2, -127)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](-2, 127)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](127, -2)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](-127, -2)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](-2, -127)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](127, -127)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](-127, 127)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](-127, -127)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](-128, -1)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](-1, -128)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](-128, 2)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](2, -128)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](-128, -2)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](-2, -128)
	require.Error(t, err)
	require.Equal(t, int8(0), product)

	product, err = ProductInt[int8](-128, -128)
	require.Error(t, err)
	require.Equal(t, int8(0), product)
}

func TestProductUint8(t *testing.T) {
	product, err := ProductInt[uint8](0, 0)
	require.NoError(t, err)
	require.Equal(t, uint8(0), product)

	product, err = ProductInt[uint8](2, 0)
	require.NoError(t, err)
	require.Equal(t, uint8(0), product)

	product, err = ProductInt[uint8](0, 2)
	require.NoError(t, err)
	require.Equal(t, uint8(0), product)

	product, err = ProductInt[uint8](3, 0)
	require.NoError(t, err)
	require.Equal(t, uint8(0), product)

	product, err = ProductInt[uint8](0, 3)
	require.NoError(t, err)
	require.Equal(t, uint8(0), product)

	product, err = ProductInt[uint8](3, 2)
	require.NoError(t, err)
	require.Equal(t, uint8(6), product)

	product, err = ProductInt[uint8](2, 3)
	require.NoError(t, err)
	require.Equal(t, uint8(6), product)

	product, err = ProductInt[uint8](85, 3)
	require.NoError(t, err)
	require.Equal(t, uint8(255), product)

	product, err = ProductInt[uint8](3, 85)
	require.NoError(t, err)
	require.Equal(t, uint8(255), product)

	product, err = ProductInt[uint8](255, 1)
	require.NoError(t, err)
	require.Equal(t, uint8(255), product)

	product, err = ProductInt[uint8](1, 255)
	require.NoError(t, err)
	require.Equal(t, uint8(255), product)
}

func TestProductUint8Overflow(t *testing.T) {
	product, err := ProductInt[uint8](255, 2)
	require.Error(t, err)
	require.Equal(t, uint8(0), product)

	product, err = ProductInt[uint8](2, 255)
	require.Error(t, err)
	require.Equal(t, uint8(0), product)

	product, err = ProductInt[uint8](128, 2)
	require.Error(t, err)
	require.Equal(t, uint8(0), product)

	product, err = ProductInt[uint8](2, 128)
	require.Error(t, err)
	require.Equal(t, uint8(0), product)

	product, err = ProductInt[uint8](86, 3)
	require.Error(t, err)
	require.Equal(t, uint8(0), product)

	product, err = ProductInt[uint8](3, 86)
	require.Error(t, err)
	require.Equal(t, uint8(0), product)

	product, err = ProductInt[uint8](64, 4)
	require.Error(t, err)
	require.Equal(t, uint8(0), product)

	product, err = ProductInt[uint8](4, 64)
	require.Error(t, err)
	require.Equal(t, uint8(0), product)

	product, err = ProductInt[uint8](255, 254)
	require.Error(t, err)
	require.Equal(t, uint8(0), product)

	product, err = ProductInt[uint8](254, 255)
	require.Error(t, err)
	require.Equal(t, uint8(0), product)

	product, err = ProductInt[uint8](255, 255)
	require.Error(t, err)
	require.Equal(t, uint8(0), product)
}

func TestFloatToInt(t *testing.T) {
	converted, err := FloatToInt[float64, int8](126.1)
	require.NoError(t, err)
	require.Equal(t, int8(126), converted)

	converted, err = FloatToInt[float64, int8](126.6)
	require.NoError(t, err)
	require.Equal(t, int8(126), converted)

	converted, err = FloatToInt[float64, int8](127.0)
	require.NoError(t, err)
	require.Equal(t, int8(127), converted)

	converted, err = FloatToInt[float64, int8](3.0)
	require.NoError(t, err)
	require.Equal(t, int8(3), converted)

	converted, err = FloatToInt[float64, int8](2.0)
	require.NoError(t, err)
	require.Equal(t, int8(2), converted)

	converted, err = FloatToInt[float64, int8](1.0)
	require.NoError(t, err)
	require.Equal(t, int8(1), converted)

	converted, err = FloatToInt[float64, int8](0.6)
	require.NoError(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](0.1)
	require.NoError(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](0.0)
	require.NoError(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-0.1)
	require.NoError(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-0.6)
	require.NoError(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-1.0)
	require.NoError(t, err)
	require.Equal(t, int8(-1), converted)

	converted, err = FloatToInt[float64, int8](-2.0)
	require.NoError(t, err)
	require.Equal(t, int8(-2), converted)

	converted, err = FloatToInt[float64, int8](-3.0)
	require.NoError(t, err)
	require.Equal(t, int8(-3), converted)

	converted, err = FloatToInt[float64, int8](-127.1)
	require.NoError(t, err)
	require.Equal(t, int8(-127), converted)

	converted, err = FloatToInt[float64, int8](-127.6)
	require.NoError(t, err)
	require.Equal(t, int8(-127), converted)

	converted, err = FloatToInt[float64, int8](-128.0)
	require.NoError(t, err)
	require.Equal(t, int8(-128), converted)
}

func TestFloatToIntOverflow(t *testing.T) {
	converted, err := FloatToInt[float64, int8](127.1)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](127.6)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](128)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](129)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](130)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](127 * 2)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](127*2 - 1)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](127*2 - 2)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](127*2 - 3)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](127*2 + 1)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](127*2 + 2)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](127*2 + 3)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](127 * 5 / 3)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-128.1)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-128.6)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-129)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-130)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-131)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-128 * 2)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-128*2 - 1)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-128*2 - 2)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-128*2 - 3)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-128*2 + 1)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-128*2 + 2)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-128*2 + 3)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = FloatToInt[float64, int8](-128 * 5 / 3)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)
}

func TestFloatToUint(t *testing.T) {
	converted, err := FloatToInt[float64, uint8](0.0)
	require.NoError(t, err)
	require.Equal(t, uint8(0), converted)

	converted, err = FloatToInt[float64, uint8](0.1)
	require.NoError(t, err)
	require.Equal(t, uint8(0), converted)

	converted, err = FloatToInt[float64, uint8](0.6)
	require.NoError(t, err)
	require.Equal(t, uint8(0), converted)

	converted, err = FloatToInt[float64, uint8](1.0)
	require.NoError(t, err)
	require.Equal(t, uint8(1), converted)

	converted, err = FloatToInt[float64, uint8](2.0)
	require.NoError(t, err)
	require.Equal(t, uint8(2), converted)

	converted, err = FloatToInt[float64, uint8](3.0)
	require.NoError(t, err)
	require.Equal(t, uint8(3), converted)

	converted, err = FloatToInt[float64, uint8](254.1)
	require.NoError(t, err)
	require.Equal(t, uint8(254), converted)

	converted, err = FloatToInt[float64, uint8](254.6)
	require.NoError(t, err)
	require.Equal(t, uint8(254), converted)

	converted, err = FloatToInt[float64, uint8](255.0)
	require.NoError(t, err)
	require.Equal(t, uint8(255), converted)
}

func TestFloatToUintOverflow(t *testing.T) {
	converted, err := FloatToInt[float64, uint8](255.1)
	require.Error(t, err)
	require.Equal(t, uint8(0), converted)

	converted, err = FloatToInt[float64, uint8](255.6)
	require.Error(t, err)
	require.Equal(t, uint8(0), converted)

	converted, err = FloatToInt[float64, uint8](256)
	require.Error(t, err)
	require.Equal(t, uint8(0), converted)

	converted, err = FloatToInt[float64, uint8](257)
	require.Error(t, err)
	require.Equal(t, uint8(0), converted)

	converted, err = FloatToInt[float64, uint8](258)
	require.Error(t, err)
	require.Equal(t, uint8(0), converted)

	converted, err = FloatToInt[float64, uint8](255 * 2)
	require.Error(t, err)
	require.Equal(t, uint8(0), converted)

	converted, err = FloatToInt[float64, uint8](255*2 - 1)
	require.Error(t, err)
	require.Equal(t, uint8(0), converted)

	converted, err = FloatToInt[float64, uint8](255*2 - 2)
	require.Error(t, err)
	require.Equal(t, uint8(0), converted)

	converted, err = FloatToInt[float64, uint8](255*2 - 3)
	require.Error(t, err)
	require.Equal(t, uint8(0), converted)

	converted, err = FloatToInt[float64, uint8](255*2 + 1)
	require.Error(t, err)
	require.Equal(t, uint8(0), converted)

	converted, err = FloatToInt[float64, uint8](255*2 + 2)
	require.Error(t, err)
	require.Equal(t, uint8(0), converted)

	converted, err = FloatToInt[float64, uint8](255*2 + 3)
	require.Error(t, err)
	require.Equal(t, uint8(0), converted)

	converted, err = FloatToInt[float64, uint8](255 * 5 / 3)
	require.Error(t, err)
	require.Equal(t, uint8(0), converted)
}

func TestInvert(t *testing.T) {
	inverted, err := Invert[int8](127)
	require.NoError(t, err)
	require.Equal(t, int8(-127), inverted)

	inverted, err = Invert[int8](126)
	require.NoError(t, err)
	require.Equal(t, int8(-126), inverted)

	inverted, err = Invert[int8](2)
	require.NoError(t, err)
	require.Equal(t, int8(-2), inverted)

	inverted, err = Invert[int8](1)
	require.NoError(t, err)
	require.Equal(t, int8(-1), inverted)

	inverted, err = Invert[int8](0)
	require.NoError(t, err)
	require.Equal(t, int8(0), inverted)

	inverted, err = Invert[int8](-1)
	require.NoError(t, err)
	require.Equal(t, int8(1), inverted)

	inverted, err = Invert[int8](-2)
	require.NoError(t, err)
	require.Equal(t, int8(2), inverted)

	inverted, err = Invert[int8](-126)
	require.NoError(t, err)
	require.Equal(t, int8(126), inverted)

	inverted, err = Invert[int8](-127)
	require.NoError(t, err)
	require.Equal(t, int8(127), inverted)
}

func TestInvertOverflow(t *testing.T) {
	inverted, err := Invert[int8](-128)
	require.Error(t, err)
	require.Equal(t, int8(0), inverted)
}

func TestPowUnsigned(t *testing.T) {
	power, err := PowUnsigned[uint8](0, 4)
	require.NoError(t, err)
	require.Equal(t, uint8(0), power)

	power, err = PowUnsigned[uint8](4, 0)
	require.NoError(t, err)
	require.Equal(t, uint8(1), power)

	power, err = PowUnsigned[uint8](1, 4)
	require.NoError(t, err)
	require.Equal(t, uint8(1), power)

	power, err = PowUnsigned[uint8](4, 1)
	require.NoError(t, err)
	require.Equal(t, uint8(4), power)

	power, err = PowUnsigned[uint8](2, 3)
	require.NoError(t, err)
	require.Equal(t, uint8(8), power)

	power, err = PowUnsigned[uint8](2, 4)
	require.NoError(t, err)
	require.Equal(t, uint8(16), power)

	power, err = PowUnsigned[uint8](2, 5)
	require.NoError(t, err)
	require.Equal(t, uint8(32), power)

	power, err = PowUnsigned[uint8](3, 2)
	require.NoError(t, err)
	require.Equal(t, uint8(9), power)

	power, err = PowUnsigned[uint8](3, 3)
	require.NoError(t, err)
	require.Equal(t, uint8(27), power)

	power, err = PowUnsigned[uint8](3, 4)
	require.NoError(t, err)
	require.Equal(t, uint8(81), power)

	power, err = PowUnsigned[uint8](10, 2)
	require.NoError(t, err)
	require.Equal(t, uint8(100), power)

	power, err = PowUnsigned[uint8](3, 5)
	require.NoError(t, err)
	require.Equal(t, uint8(243), power)
}

func TestPowUnsignedOverflow(t *testing.T) {
	power, err := PowUnsigned[uint8](2, 8)
	require.Error(t, err)
	require.Equal(t, uint8(0), power)

	power, err = PowUnsigned[uint8](10, 3)
	require.Error(t, err)
	require.Equal(t, uint8(0), power)
}

func TestUnsignedToSigned(t *testing.T) {
	converted, err := UnsignedToSigned[uint8, int8](0)
	require.NoError(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = UnsignedToSigned[uint8, int8](1)
	require.NoError(t, err)
	require.Equal(t, int8(1), converted)

	converted, err = UnsignedToSigned[uint8, int8](2)
	require.NoError(t, err)
	require.Equal(t, int8(2), converted)

	converted, err = UnsignedToSigned[uint8, int8](3)
	require.NoError(t, err)
	require.Equal(t, int8(3), converted)

	converted, err = UnsignedToSigned[uint8, int8](125)
	require.NoError(t, err)
	require.Equal(t, int8(125), converted)

	converted, err = UnsignedToSigned[uint8, int8](126)
	require.NoError(t, err)
	require.Equal(t, int8(126), converted)

	converted, err = UnsignedToSigned[uint8, int8](127)
	require.NoError(t, err)
	require.Equal(t, int8(127), converted)

	converted16, err := UnsignedToSigned[uint8, int16](128)
	require.NoError(t, err)
	require.Equal(t, int16(128), converted16)

	converted16, err = UnsignedToSigned[uint8, int16](129)
	require.NoError(t, err)
	require.Equal(t, int16(129), converted16)

	converted16, err = UnsignedToSigned[uint8, int16](130)
	require.NoError(t, err)
	require.Equal(t, int16(130), converted16)

	converted16, err = UnsignedToSigned[uint8, int16](254)
	require.NoError(t, err)
	require.Equal(t, int16(254), converted16)

	converted16, err = UnsignedToSigned[uint8, int16](255)
	require.NoError(t, err)
	require.Equal(t, int16(255), converted16)
}

func TestUintToIntOverflow(t *testing.T) {
	converted, err := UnsignedToSigned[uint8, int8](128)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = UnsignedToSigned[uint8, int8](129)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = UnsignedToSigned[uint8, int8](130)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = UnsignedToSigned[uint8, int8](253)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = UnsignedToSigned[uint8, int8](254)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = UnsignedToSigned[uint8, int8](255)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = UnsignedToSigned[uint16, int8](256)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = UnsignedToSigned[uint16, int8](257)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = UnsignedToSigned[uint16, int8](258)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)

	converted, err = UnsignedToSigned[uint16, int8](259)
	require.Error(t, err)
	require.Equal(t, int8(0), converted)
}
