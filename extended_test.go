package safe

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddT(t *testing.T) {
	testAddTInt(t)
	testAddTUint(t)
}

func testAddTInt(t *testing.T) {
	faults := 0
	successful := 0

	for first := math.MinInt8; first <= math.MaxInt8; first++ {
		for second := math.MinInt8; second <= math.MaxInt8; second++ {
			for third := math.MinInt8; third <= math.MaxInt8; third++ {
				sum, err := AddT(int8(first), int8(second), int8(third))

				reference := first + second + third

				if reference > math.MaxInt8 || reference < math.MinInt8 {
					require.Error(
						t,
						err,
						"first: %v, second: %v, third: %v, sum: %v, reference: %v",
						first,
						second,
						third,
						sum,
						reference,
					)

					faults++

					continue
				}

				successful++

				require.NoError(
					t,
					err,
					"first: %v, second: %v, third: %v, sum: %v, reference: %v",
					first,
					second,
					third,
					sum,
					reference,
				)

				require.Equal(
					t,
					reference,
					int(sum),
					"first: %v, second: %v, third: %v",
					first,
					second,
					third,
				)
			}
		}
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func testAddTUint(t *testing.T) {
	faults := 0
	successful := 0

	for first := 0; first <= math.MaxUint8; first++ {
		for second := 0; second <= math.MaxUint8; second++ {
			for third := 0; third <= math.MaxUint8; third++ {
				sum, err := AddT(uint8(first), uint8(second), uint8(third))

				reference := first + second + third

				if reference > math.MaxUint8 {
					require.Error(
						t,
						err,
						"first: %v, second: %v, third: %v, sum: %v, reference: %v",
						first,
						second,
						third,
						sum,
						reference,
					)

					faults++

					continue
				}

				successful++

				require.NoError(
					t,
					err,
					"first: %v, second: %v, third: %v, sum: %v, reference: %v",
					first,
					second,
					third,
					sum,
					reference,
				)

				require.Equal(
					t,
					reference,
					int(sum),
					"first: %v, second: %v, third: %v",
					first,
					second,
					third,
				)
			}
		}
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func TestPow10(t *testing.T) {
	testPow10Manually(t)
	testPow10Diff(t)
}

func testPow10Manually(t *testing.T) {
	product, err := Pow10[uint64](-3)
	require.NoError(t, err)
	require.Equal(t, uint64(0), product)

	product, err = Pow10[uint64](-2)
	require.NoError(t, err)
	require.Equal(t, uint64(0), product)

	product, err = Pow10[uint64](-1)
	require.NoError(t, err)
	require.Equal(t, uint64(0), product)

	product, err = Pow10[uint64](0)
	require.NoError(t, err)
	require.Equal(t, uint64(1), product)

	product, err = Pow10[uint64](1)
	require.NoError(t, err)
	require.Equal(t, uint64(10), product)

	product, err = Pow10[uint64](2)
	require.NoError(t, err)
	require.Equal(t, uint64(100), product)

	product, err = Pow10[uint64](3)
	require.NoError(t, err)
	require.Equal(t, uint64(1000), product)

	product, err = Pow10[uint64](19)
	require.NoError(t, err)
	require.Equal(t, uint64(1e19), product)

	product, err = Pow10[uint64](20)
	require.Error(t, err)
	require.Equal(t, uint64(0), product)

	product, err = Pow10[uint64](21)
	require.Error(t, err)
	require.Equal(t, uint64(0), product)
}

func testPow10Diff(t *testing.T) {
	for power := 1; power <= 19; power++ {
		previous, err := Pow10[uint64](power - 1)
		require.NoError(t, err, "power: %v", power)

		current, err := Pow10[uint64](power)
		require.NoError(t, err, "power: %v", power)

		require.Equal(t, uint64(10), current/previous, "power: %v", power)
		require.Equal(t, uint64(0), current%previous, "power: %v", power)
	}
}

func TestPow(t *testing.T) {
	product, err := Pow[int32](2, 30)
	require.NoError(t, err)
	require.Equal(t, int32(1<<30), product)

	product, err = Pow[int32](2, 31)
	require.Error(t, err)
	require.Equal(t, int32(0), product)

	testPow(t, 31)
}

func testPow(t *testing.T, maxPower int8) {
	faults := 0
	successful := 0

	// int32 is used because in its value range float64 does not lose the precision of
	// the integer part
	maxInt32, loss := IToF[float64](math.MaxInt32)
	require.False(t, loss)

	minInt32, loss := IToF[float64](math.MinInt32)
	require.False(t, loss)

	// int8 range is used because when using base and power in its values,
	// the result of exponentiation in floating point numbers is less than infinity,
	// except in the case of raising 0 to a negative power
	for base := int32(math.MinInt8); base <= math.MaxInt8; base++ {
		for power := int32(math.MinInt8); power <= int32(maxPower); power++ {
			bf, loss := IToF[float64](base)
			require.False(t, loss)

			pf, loss := IToF[float64](power)
			require.False(t, loss)

			reference := math.Pow(bf, pf)
			require.False(t, math.IsNaN(reference))

			if base != 0 || power >= 0 {
				require.False(
					t,
					math.IsInf(reference, 0),
					"base: %v, power: %v",
					base,
					power,
				)
			}

			product, err := Pow(base, power)

			if reference > maxInt32 || reference < minInt32 {
				require.Error(
					t,
					err,
					"base: %v, power: %v, product: %v, reference: %f",
					base,
					power,
					product,
					reference,
				)

				faults++

				continue
			}

			successful++

			require.NoError(
				t,
				err,
				"base: %v, power: %v, product: %v, reference: %f",
				base,
				power,
				product,
				reference,
			)

			require.Equal(
				t,
				int32(reference),
				product,
				"base: %v, power: %v, reference: %f",
				base,
				power,
				reference,
			)

			require.InDelta(
				t,
				reference,
				float64(product),
				0.5,
				"base: %v, power: %v, product: %v",
				base,
				power,
				product,
			)
		}
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func BenchmarkAddT(b *testing.B) {
	sum := 0

	// b.N, sum and require is used to prevent compiler optimizations
	for range b.N {
		sum, _ = AddT(b.N, 3, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkPow10Reference(b *testing.B) {
	product := float64(0)

	// product and require is used to prevent compiler optimizations
	for range b.N {
		product = math.Pow10(19)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkPow10(b *testing.B) {
	product := uint64(0)

	// product and require is used to prevent compiler optimizations
	for range b.N {
		product, _ = Pow10[uint64](19)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkPowReference(b *testing.B) {
	product := float64(0)

	// product and require is used to prevent compiler optimizations
	for range b.N {
		product = math.Pow(14, 14)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkPow(b *testing.B) {
	product := uint64(0)

	// product and require is used to prevent compiler optimizations
	for range b.N {
		product, _ = Pow(uint64(14), 14)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}
