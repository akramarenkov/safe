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

func TestAddUM(t *testing.T) {
	sum, err := AddUM(uint(math.MaxUint))
	require.NoError(t, err)
	require.Equal(t, uint(math.MaxUint), sum)

	testAddUM(t)
}

func testAddUM(t *testing.T) {
	faults := 0
	successful := 0

	for first := 0; first <= math.MaxUint8; first++ {
		for second := 0; second <= math.MaxUint8; second++ {
			for third := 0; third <= math.MaxUint8; third++ {
				sum, err := AddUM(uint8(first), uint8(second), uint8(third))

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

func TestSubT(t *testing.T) {
	testSubTInt(t)
	testSubTUint(t)
}

func testSubTInt(t *testing.T) {
	faults := 0
	successful := 0

	for minuend := math.MinInt8; minuend <= math.MaxInt8; minuend++ {
		for subtrahend := math.MinInt8; subtrahend <= math.MaxInt8; subtrahend++ {
			for secondSub := math.MinInt8; secondSub <= math.MaxInt8; secondSub++ {
				diff, err := SubT(int8(minuend), int8(subtrahend), int8(secondSub))

				reference := minuend - subtrahend - secondSub

				if reference > math.MaxInt8 || reference < math.MinInt8 {
					require.Error(
						t,
						err,
						"minuend: %v, subtrahend: %v, second subtrahend: %v, "+
							"diff: %v, reference: %v",
						minuend,
						subtrahend,
						secondSub,
						diff,
						reference,
					)

					faults++

					continue
				}

				successful++

				require.NoError(
					t,
					err,
					"minuend: %v, subtrahend: %v, second subtrahend: %v, "+
						"diff: %v, reference: %v",
					minuend,
					subtrahend,
					secondSub,
					diff,
					reference,
				)

				require.Equal(
					t,
					reference,
					int(diff),
					"minuend: %v, subtrahend: %v, second subtrahend: %v",
					minuend,
					subtrahend,
					secondSub,
				)
			}
		}
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func testSubTUint(t *testing.T) {
	faults := 0
	successful := 0

	for minuend := 0; minuend <= math.MaxUint8; minuend++ {
		for subtrahend := 0; subtrahend <= math.MaxUint8; subtrahend++ {
			for secondSub := 0; secondSub <= math.MaxUint8; secondSub++ {
				diff, err := SubT(uint8(minuend), uint8(subtrahend), uint8(secondSub))

				reference := minuend - subtrahend - secondSub

				if reference < 0 {
					require.Error(
						t,
						err,
						"minuend: %v, subtrahend: %v, second subtrahend: %v, "+
							"diff: %v, reference: %v",
						minuend,
						subtrahend,
						secondSub,
						diff,
						reference,
					)

					faults++

					continue
				}

				successful++

				require.NoError(
					t,
					err,
					"minuend: %v, subtrahend: %v, second subtrahend: %v, "+
						"diff: %v, reference: %v",
					minuend,
					subtrahend,
					secondSub,
					diff,
					reference,
				)

				require.Equal(
					t,
					reference,
					int(diff),
					"minuend: %v, subtrahend: %v, second subtrahend: %v",
					minuend,
					subtrahend,
					secondSub,
				)
			}
		}
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func TestSubUM(t *testing.T) {
	diff, err := SubUM(uint(math.MaxUint))
	require.NoError(t, err)
	require.Equal(t, uint(math.MaxUint), diff)

	testSubUM(t)
}

func testSubUM(t *testing.T) {
	faults := 0
	successful := 0

	for minuend := 0; minuend <= math.MaxUint8; minuend++ {
		for subtrahend := 0; subtrahend <= math.MaxUint8; subtrahend++ {
			for secondSub := 0; secondSub <= math.MaxUint8; secondSub++ {
				diff, err := SubUM(uint8(minuend), uint8(subtrahend), uint8(secondSub))

				reference := minuend - subtrahend - secondSub

				if reference < 0 {
					require.Error(
						t,
						err,
						"minuend: %v, subtrahend: %v, second subtrahend: %v, "+
							"diff: %v, reference: %v",
						minuend,
						subtrahend,
						secondSub,
						diff,
						reference,
					)

					faults++

					continue
				}

				successful++

				require.NoError(
					t,
					err,
					"minuend: %v, subtrahend: %v, second subtrahend: %v, "+
						"diff: %v, reference: %v",
					minuend,
					subtrahend,
					secondSub,
					diff,
					reference,
				)

				require.Equal(
					t,
					reference,
					int(diff),
					"minuend: %v, subtrahend: %v, second subtrahend: %v",
					minuend,
					subtrahend,
					secondSub,
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

func BenchmarkAddUM(b *testing.B) {
	sum := uint(0)

	// b.N, sum and require is used to prevent compiler optimizations
	for range b.N {
		sum, _ = AddUM(uint(b.N), 3, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkSubT(b *testing.B) {
	diff := 0

	// b.N, diff and require is used to prevent compiler optimizations
	for range b.N {
		diff, _ = SubT(b.N, 3, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubUM(b *testing.B) {
	diff := uint(0)

	// b.N, diff and require is used to prevent compiler optimizations
	for range b.N {
		diff, _ = SubUM(uint(b.N), 3, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
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
