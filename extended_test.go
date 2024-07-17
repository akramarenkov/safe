package safe

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddT(t *testing.T) {
	testAddTInt(t)
	testAddTUint(t, AddT)
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

func testAddTUint(t *testing.T, add func(uint8, uint8, uint8) (uint8, error)) {
	faults := 0
	successful := 0

	for first := 0; first <= math.MaxUint8; first++ {
		for second := 0; second <= math.MaxUint8; second++ {
			for third := 0; third <= math.MaxUint8; third++ {
				sum, err := add(uint8(first), uint8(second), uint8(third))

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

	add := func(first uint8, second uint8, third uint8) (uint8, error) {
		return AddUM(first, second, third)
	}

	testAddTUint(t, add)
}

func TestSubT(t *testing.T) {
	testSubTInt(t)
	testSubTUint(t, SubT)
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

func testSubTUint(t *testing.T, sub func(uint8, uint8, uint8) (uint8, error)) {
	faults := 0
	successful := 0

	for minuend := 0; minuend <= math.MaxUint8; minuend++ {
		for subtrahend := 0; subtrahend <= math.MaxUint8; subtrahend++ {
			for secondSub := 0; secondSub <= math.MaxUint8; secondSub++ {
				diff, err := sub(uint8(minuend), uint8(subtrahend), uint8(secondSub))

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

	sub := func(minuend uint8, subtrahend uint8, secondSubtrahend uint8) (uint8, error) {
		return SubUM(minuend, subtrahend, secondSubtrahend)
	}

	testSubTUint(t, sub)
}

func TestMulT(t *testing.T) {
	testMulTInt(t)
	testMulTUint(t, MulT)
}

func testMulTInt(t *testing.T) {
	faults := 0
	successful := 0

	for first := math.MinInt8; first <= math.MaxInt8; first++ {
		for second := math.MinInt8; second <= math.MaxInt8; second++ {
			for third := math.MinInt8; third <= math.MaxInt8; third++ {
				product, err := MulT(int8(first), int8(second), int8(third))

				reference := first * second * third

				if reference > math.MaxInt8 || reference < math.MinInt8 {
					require.Error(
						t,
						err,
						"first: %v, second: %v, third: %v, product: %v, reference: %v",
						first,
						second,
						third,
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
					"first: %v, second: %v, third: %v, product: %v, reference: %v",
					first,
					second,
					third,
					product,
					reference,
				)

				require.Equal(
					t,
					reference,
					int(product),
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

func testMulTUint(t *testing.T, mul func(uint8, uint8, uint8) (uint8, error)) {
	faults := 0
	successful := 0

	for first := 0; first <= math.MaxUint8; first++ {
		for second := 0; second <= math.MaxUint8; second++ {
			for third := 0; third <= math.MaxUint8; third++ {
				product, err := mul(uint8(first), uint8(second), uint8(third))

				reference := first * second * third

				if reference > math.MaxUint8 {
					require.Error(
						t,
						err,
						"first: %v, second: %v, third: %v, product: %v, reference: %v",
						first,
						second,
						third,
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
					"first: %v, second: %v, third: %v, product: %v, reference: %v",
					first,
					second,
					third,
					product,
					reference,
				)

				require.Equal(
					t,
					reference,
					int(product),
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

func TestMulUM(t *testing.T) {
	product, err := MulUM(uint(math.MaxUint))
	require.NoError(t, err)
	require.Equal(t, uint(math.MaxUint), product)

	mul := func(first uint8, second uint8, third uint8) (uint8, error) {
		return MulUM(first, second, third)
	}

	testMulTUint(t, mul)
}

func TestDivM(t *testing.T) {
	quotient, err := DivM(uint(math.MaxUint))
	require.NoError(t, err)
	require.Equal(t, uint(math.MaxUint), quotient)

	testDivMInt(t)
	testDivMUint(t)
}

func testDivMInt(t *testing.T) {
	zeros := 0
	faults := 0
	successful := 0

	for dividend := math.MinInt8; dividend <= math.MaxInt8; dividend++ {
		for divisor := math.MinInt8; divisor <= math.MaxInt8; divisor++ {
			for secondDiv := math.MinInt8; secondDiv <= math.MaxInt8; secondDiv++ {
				quotient, err := DivM(int8(dividend), int8(divisor), int8(secondDiv))

				if divisor == 0 || secondDiv == 0 {
					require.Error(
						t,
						err,
						"dividend: %v, divisor: %v, second divisor: %v",
						dividend,
						divisor,
						secondDiv,
					)

					zeros++

					continue
				}

				reference := dividend / divisor / secondDiv

				if reference > math.MaxInt8 || reference < math.MinInt8 {
					require.Error(
						t,
						err,
						"dividend: %v, divisor: %v, second divisor: %v, "+
							"quotient: %v, reference: %v",
						dividend,
						divisor,
						secondDiv,
						quotient,
						reference,
					)

					faults++

					continue
				}

				successful++

				require.NoError(
					t,
					err,
					"dividend: %v, divisor: %v, second divisor: %v, "+
						"quotient: %v, reference: %v",
					dividend,
					divisor,
					secondDiv,
					quotient,
					reference,
				)

				require.Equal(
					t,
					reference,
					int(quotient),
					"dividend: %v, divisor: %v, second divisor: %v",
					dividend,
					divisor,
					secondDiv,
				)
			}
		}
	}

	require.NotZero(t, zeros)
	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func testDivMUint(t *testing.T) {
	zeros := 0
	faults := 0
	successful := 0

	for dividend := 0; dividend <= math.MaxUint8; dividend++ {
		for divisor := 0; divisor <= math.MaxUint8; divisor++ {
			for secondDiv := 0; secondDiv <= math.MaxUint8; secondDiv++ {
				quotient, err := DivM(uint8(dividend), uint8(divisor), uint8(secondDiv))

				if divisor == 0 || secondDiv == 0 {
					require.Error(
						t,
						err,
						"dividend: %v, divisor: %v, second divisor: %v",
						dividend,
						divisor,
						secondDiv,
					)

					zeros++

					continue
				}

				reference := dividend / divisor / secondDiv

				if reference > math.MaxUint8 {
					require.Error(
						t,
						err,
						"dividend: %v, divisor: %v, second divisor: %v, "+
							"quotient: %v, reference: %v",
						dividend,
						divisor,
						secondDiv,
						quotient,
						reference,
					)

					faults++

					continue
				}

				successful++

				require.NoError(
					t,
					err,
					"dividend: %v, divisor: %v, second divisor: %v, "+
						"quotient: %v, reference: %v",
					dividend,
					divisor,
					secondDiv,
					quotient,
					reference,
				)

				require.Equal(
					t,
					reference,
					int(quotient),
					"dividend: %v, divisor: %v, second divisor: %v",
					dividend,
					divisor,
					secondDiv,
				)
			}
		}
	}

	require.NotZero(t, zeros)
	require.Zero(t, faults)
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
	faults := 0
	successful := 0

	// Is used int32 because in its value range float64 does not lose the precision of
	// the integer part
	maxInt32 := float64(math.MaxInt32)
	minInt32 := float64(math.MinInt32)

	for base := int32(math.MinInt8); base <= math.MaxInt8; base++ {
		for power := int32(math.MinInt8); power <= math.MaxInt8; power++ {
			reference := math.Pow(float64(base), float64(power))
			require.False(t, math.IsNaN(reference))

			// It is separately checked that in areas close to the maximum/minimum int32
			// values, the overflow conditions are satisfied correctly
			if reference >= maxInt32-2 && reference <= maxInt32+2 {
				require.InDelta(t, maxInt32+1, reference, 0)

				// reference > maxInt32 == true
				require.Greater(t, reference, maxInt32)
			}

			if reference >= minInt32-2 && reference <= minInt32+2 {
				require.InDelta(t, minInt32, reference, 0)

				// reference < minInt32 == false
				require.GreaterOrEqual(t, reference, minInt32)
			}

			product, err := Pow(base, power)

			// Converting reference to any integer type can and will overflow,
			// so the comparison is done in float64
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

func BenchmarkMulT(b *testing.B) {
	product := 0

	// b.N, product and require is used to prevent compiler optimizations
	for range b.N {
		product, _ = MulT(b.N, 3, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMulUM(b *testing.B) {
	product := uint(0)

	// b.N, product and require is used to prevent compiler optimizations
	for range b.N {
		product, _ = MulUM(uint(b.N), 3, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkDivM(b *testing.B) {
	quotient := 0

	// b.N, quotient and require is used to prevent compiler optimizations
	for range b.N {
		quotient, _ = DivM(b.N, 3, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, quotient)
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
