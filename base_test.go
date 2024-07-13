package safe

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	testAddInt(t)
	testAddUint(t, Add)
}

func TestAddU(t *testing.T) {
	testAddUint(t, AddU)
}

func testAddInt(t *testing.T) {
	faults := 0
	successful := 0

	for first := math.MinInt8; first <= math.MaxInt8; first++ {
		for second := math.MinInt8; second <= math.MaxInt8; second++ {
			sum, err := Add(int8(first), int8(second))

			reference := first + second

			if reference > math.MaxInt8 || reference < math.MinInt8 {
				require.Error(
					t,
					err,
					"first: %v, second: %v, sum: %v, reference: %v",
					first,
					second,
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
				"first: %v, second: %v, sum: %v, reference: %v",
				first,
				second,
				sum,
				reference,
			)

			require.Equal(
				t,
				reference,
				int(sum),
				"first: %v, second: %v",
				first,
				second,
			)
		}
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func testAddUint(t *testing.T, add func(uint8, uint8) (uint8, error)) {
	faults := 0
	successful := 0

	for first := 0; first <= math.MaxUint8; first++ {
		for second := 0; second <= math.MaxUint8; second++ {
			sum, err := add(uint8(first), uint8(second))

			reference := first + second

			if reference > math.MaxUint8 {
				require.Error(
					t,
					err,
					"first: %v, second: %v, sum: %v, reference: %v",
					first,
					second,
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
				"first: %v, second: %v, sum: %v, reference: %v",
				first,
				second,
				sum,
				reference,
			)

			require.Equal(
				t,
				reference,
				int(sum),
				"first: %v, second: %v",
				first,
				second,
			)
		}
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func TestSub(t *testing.T) {
	testSubInt(t)
	testSubUint(t, Sub)
}

func TestSubU(t *testing.T) {
	testSubUint(t, SubU)
}

func testSubInt(t *testing.T) {
	faults := 0
	successful := 0

	for minuend := math.MinInt8; minuend <= math.MaxInt8; minuend++ {
		for subtrahend := math.MinInt8; subtrahend <= math.MaxInt8; subtrahend++ {
			diff, err := Sub(int8(minuend), int8(subtrahend))

			reference := minuend - subtrahend

			if reference > math.MaxInt8 || reference < math.MinInt8 {
				require.Error(
					t,
					err,
					"minuend: %v, subtrahend: %v, diff: %v, reference: %v",
					minuend,
					subtrahend,
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
				"minuend: %v, subtrahend: %v, diff: %v, reference: %v",
				minuend,
				subtrahend,
				diff,
				reference,
			)

			require.Equal(
				t,
				reference,
				int(diff),
				"minuend: %v, subtrahend: %v",
				minuend,
				subtrahend,
			)
		}
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func testSubUint(t *testing.T, sub func(uint8, uint8) (uint8, error)) {
	faults := 0
	successful := 0

	for minuend := 0; minuend <= math.MaxUint8; minuend++ {
		for subtrahend := 0; subtrahend <= math.MaxUint8; subtrahend++ {
			diff, err := sub(uint8(minuend), uint8(subtrahend))

			reference := minuend - subtrahend

			if reference < 0 {
				require.Error(
					t,
					err,
					"minuend: %v, subtrahend: %v, diff: %v, reference: %v",
					minuend,
					subtrahend,
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
				"minuend: %v, subtrahend: %v, diff: %v, reference: %v",
				minuend,
				subtrahend,
				diff,
				reference,
			)

			require.Equal(
				t,
				reference,
				int(diff),
				"minuend: %v, subtrahend: %v",
				minuend,
				subtrahend,
			)
		}
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func TestMul(t *testing.T) {
	testMulInt(t)
	testMulUint(t)
}

func testMulInt(t *testing.T) {
	faults := 0
	successful := 0

	for first := math.MinInt8; first <= math.MaxInt8; first++ {
		for second := math.MinInt8; second <= math.MaxInt8; second++ {
			product, err := Mul(int8(first), int8(second))

			reference := first * second

			if reference > math.MaxInt8 || reference < math.MinInt8 {
				require.Error(
					t,
					err,
					"first: %v, second: %v, product: %v, reference: %v",
					first,
					second,
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
				"first: %v, second: %v, product: %v, reference: %v",
				first,
				second,
				product,
				reference,
			)

			require.Equal(
				t,
				reference,
				int(product),
				"first: %v, second: %v",
				first,
				second,
			)
		}
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func testMulUint(t *testing.T) {
	faults := 0
	successful := 0

	for first := 0; first <= math.MaxUint8; first++ {
		for second := 0; second <= math.MaxUint8; second++ {
			product, err := Mul(uint8(first), uint8(second))

			reference := first * second

			if reference > math.MaxUint8 {
				require.Error(
					t,
					err,
					"first: %v, second: %v, product: %v, reference: %v",
					first,
					second,
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
				"first: %v, second: %v, product: %v, reference: %v",
				first,
				second,
				product,
				reference,
			)

			require.Equal(
				t,
				reference,
				int(product),
				"first: %v, second: %v",
				first,
				second,
			)
		}
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func TestDiv(t *testing.T) {
	testDivInt(t)
	testDivUint(t)
}

func testDivInt(t *testing.T) {
	faults := 0
	successful := 0

	for dividend := math.MinInt8; dividend <= math.MaxInt8; dividend++ {
		for divisor := math.MinInt8; divisor <= math.MaxInt8; divisor++ {
			quotient, err := Div(int8(dividend), int8(divisor))

			if divisor == 0 {
				require.Error(t, err, "dividend: %v, divisor: %v", dividend, divisor)
				continue
			}

			reference := dividend / divisor

			if reference > math.MaxInt8 || reference < math.MinInt8 {
				require.Error(
					t,
					err,
					"dividend: %v, divisor: %v, quotient: %v, reference: %v",
					dividend,
					divisor,
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
				"dividend: %v, divisor: %v, quotient: %v, reference: %v",
				dividend,
				divisor,
				quotient,
				reference,
			)

			require.Equal(
				t,
				reference,
				int(quotient),
				"dividend: %v, divisor: %v",
				dividend,
				divisor,
			)
		}
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func testDivUint(t *testing.T) {
	faults := 0
	successful := 0

	for dividend := 0; dividend <= math.MaxUint8; dividend++ {
		for divisor := 0; divisor <= math.MaxUint8; divisor++ {
			quotient, err := Div(uint8(dividend), uint8(divisor))

			if divisor == 0 {
				require.Error(t, err, "dividend: %v, divisor: %v", dividend, divisor)
				continue
			}

			reference := dividend / divisor

			if reference > math.MaxUint8 {
				require.Error(
					t,
					err,
					"dividend: %v, divisor: %v, quotient: %v, reference: %v",
					dividend,
					divisor,
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
				"dividend: %v, divisor: %v, quotient: %v, reference: %v",
				dividend,
				divisor,
				quotient,
				reference,
			)

			require.Equal(
				t,
				reference,
				int(quotient),
				"dividend: %v, divisor: %v",
				dividend,
				divisor,
			)
		}
	}

	require.Zero(t, faults)
	require.NotZero(t, successful)
}

func TestNegate(t *testing.T) {
	faults := 0
	successful := 0

	for number := math.MinInt8; number <= math.MaxInt8; number++ {
		negated, err := Negate(int8(number))

		reference := -number

		if reference > math.MaxInt8 || reference < math.MinInt8 {
			require.Error(
				t,
				err,
				"number: %v, negated: %v, reference: %v",
				number,
				negated,
				reference,
			)

			faults++

			continue
		}

		successful++

		require.NoError(
			t,
			err,
			"number: %v, negated: %v, reference: %v",
			number,
			negated,
			reference,
		)

		require.Equal(t, reference, int(negated), "number: %v", number)
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func TestIToI(t *testing.T) {
	testIToIU8ToS8(t)
	testIToIS8ToU8(t)
	testIToIS8ToU16(t)
	testIToIU16ToS8(t)
	testIToIU16ToU8(t)
	testIToIS16ToS8(t)
	testIToIS16ToU8(t)
}

func testIToIU8ToS8(t *testing.T) {
	faults := 0
	successful := 0

	for number := 0; number <= math.MaxUint8; number++ {
		converted, err := IToI[uint8, int8](uint8(number))

		reference := number

		if reference > math.MaxInt8 {
			require.Error(
				t,
				err,
				"number: %v, converted: %v, reference: %v",
				number,
				converted,
				reference,
			)

			faults++

			continue
		}

		successful++

		require.NoError(
			t,
			err,
			"number: %v, converted: %v, reference: %v",
			number,
			converted,
			reference,
		)

		require.Equal(t, reference, int(converted), "number: %v", number)
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func testIToIS8ToU8(t *testing.T) {
	faults := 0
	successful := 0

	for number := math.MinInt8; number <= math.MaxInt8; number++ {
		converted, err := IToI[int8, uint8](int8(number))

		reference := number

		if reference < 0 {
			require.Error(
				t,
				err,
				"number: %v, converted: %v, reference: %v",
				number,
				converted,
				reference,
			)

			faults++

			continue
		}

		successful++

		require.NoError(
			t,
			err,
			"number: %v, converted: %v, reference: %v",
			number,
			converted,
			reference,
		)

		require.Equal(t, reference, int(converted), "number: %v", number)
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func testIToIS8ToU16(t *testing.T) {
	faults := 0
	successful := 0

	for number := math.MinInt8; number <= math.MaxInt8; number++ {
		converted, err := IToI[int8, uint16](int8(number))

		reference := number

		if reference < 0 {
			require.Error(
				t,
				err,
				"number: %v, converted: %v, reference: %v",
				number,
				converted,
				reference,
			)

			faults++

			continue
		}

		successful++

		require.NoError(
			t,
			err,
			"number: %v, converted: %v, reference: %v",
			number,
			converted,
			reference,
		)

		require.Equal(t, reference, int(converted), "number: %v", number)
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func testIToIU16ToS8(t *testing.T) {
	faults := 0
	successful := 0

	for number := 0; number <= math.MaxUint16; number++ {
		converted, err := IToI[uint16, int8](uint16(number))

		reference := number

		if reference > math.MaxInt8 {
			require.Error(
				t,
				err,
				"number: %v, converted: %v, reference: %v",
				number,
				converted,
				reference,
			)

			faults++

			continue
		}

		successful++

		require.NoError(
			t,
			err,
			"number: %v, converted: %v, reference: %v",
			number,
			converted,
			reference,
		)

		require.Equal(t, reference, int(converted), "number: %v", number)
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func testIToIU16ToU8(t *testing.T) {
	faults := 0
	successful := 0

	for number := 0; number <= math.MaxUint16; number++ {
		converted, err := IToI[uint16, uint8](uint16(number))

		reference := number

		if reference > math.MaxUint8 {
			require.Error(
				t,
				err,
				"number: %v, converted: %v, reference: %v",
				number,
				converted,
				reference,
			)

			faults++

			continue
		}

		successful++

		require.NoError(
			t,
			err,
			"number: %v, converted: %v, reference: %v",
			number,
			converted,
			reference,
		)

		require.Equal(t, reference, int(converted), "number: %v", number)
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func testIToIS16ToS8(t *testing.T) {
	faults := 0
	successful := 0

	for number := math.MinInt16; number <= math.MaxInt16; number++ {
		converted, err := IToI[int16, int8](int16(number))

		reference := number

		if reference > math.MaxInt8 || reference < math.MinInt8 {
			require.Error(
				t,
				err,
				"number: %v, converted: %v, reference: %v",
				number,
				converted,
				reference,
			)

			faults++

			continue
		}

		successful++

		require.NoError(
			t,
			err,
			"number: %v, converted: %v, reference: %v",
			number,
			converted,
			reference,
		)

		require.Equal(t, reference, int(converted), "number: %v", number)
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func testIToIS16ToU8(t *testing.T) {
	faults := 0
	successful := 0

	for number := math.MinInt16; number <= math.MaxInt16; number++ {
		converted, err := IToI[int16, uint8](int16(number))

		reference := number

		if reference > math.MaxUint8 || reference < 0 {
			require.Error(
				t,
				err,
				"number: %v, converted: %v, reference: %v",
				number,
				converted,
				reference,
			)

			faults++

			continue
		}

		successful++

		require.NoError(
			t,
			err,
			"number: %v, converted: %v, reference: %v",
			number,
			converted,
			reference,
		)

		require.Equal(t, reference, int(converted), "number: %v", number)
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func TestUToS(t *testing.T) {
	testUToS8To8(t)
	testUToS16To8(t)
}

func testUToS8To8(t *testing.T) {
	faults := 0
	successful := 0

	for number := 0; number <= math.MaxUint8; number++ {
		converted, err := UToS[uint8, int8](uint8(number))

		reference := number

		if reference > math.MaxInt8 {
			require.Error(
				t,
				err,
				"number: %v, converted: %v, reference: %v",
				number,
				converted,
				reference,
			)

			faults++

			continue
		}

		successful++

		require.NoError(
			t,
			err,
			"number: %v, converted: %v, reference: %v",
			number,
			converted,
			reference,
		)

		require.Equal(t, reference, int(converted), "number: %v", number)
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func testUToS16To8(t *testing.T) {
	faults := 0
	successful := 0

	for number := 0; number <= math.MaxUint16; number++ {
		converted, err := UToS[uint16, int8](uint16(number))

		reference := number

		if reference > math.MaxInt8 {
			require.Error(
				t,
				err,
				"number: %v, converted: %v, reference: %v",
				number,
				converted,
				reference,
			)

			faults++

			continue
		}

		successful++

		require.NoError(
			t,
			err,
			"number: %v, converted: %v, reference: %v",
			number,
			converted,
			reference,
		)

		require.Equal(t, reference, int(converted), "number: %v", number)
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

func TestIToF(t *testing.T) {
	testIToF32(t)
	testIToF64(t)
}

func testIToF32(t *testing.T) {
	converted, loss := IToF[uint64, float32](math.MaxUint8)
	require.False(t, loss)
	require.InDelta(t, float32(math.MaxUint8), converted, 0.0)

	converted, loss = IToF[uint64, float32](math.MaxUint16)
	require.False(t, loss)
	require.InDelta(t, float32(math.MaxUint16), converted, 0.0)

	converted, loss = IToF[uint64, float32](1 << 24)
	require.False(t, loss)
	require.InDelta(t, float32(1<<24), converted, 0.0)

	_, loss = IToF[uint64, float32](1<<24 + 1)
	require.True(t, loss)

	_, loss = IToF[uint64, float32](math.MaxUint32)
	require.True(t, loss)
}

func testIToF64(t *testing.T) {
	converted, loss := IToF[uint64, float64](math.MaxUint8)
	require.False(t, loss)
	require.InDelta(t, float64(math.MaxUint8), converted, 0.0)

	converted, loss = IToF[uint64, float64](math.MaxUint16)
	require.False(t, loss)
	require.InDelta(t, float64(math.MaxUint16), converted, 0.0)

	converted, loss = IToF[uint64, float64](math.MaxUint32)
	require.False(t, loss)
	require.InDelta(t, float64(math.MaxUint32), converted, 0.0)

	converted, loss = IToF[uint64, float64](1 << 53)
	require.False(t, loss)
	require.InDelta(t, float64(1<<53), converted, 0.0)

	_, loss = IToF[uint64, float64](1<<53 + 1)
	require.True(t, loss)

	_, loss = IToF[uint64, float64](math.MaxUint64)
	require.True(t, loss)
}

func TestFToI(t *testing.T) {
	fractionals := []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}

	faultsInt := 0
	faultsUint := 0
	successfulInt := 0
	successfulUint := 0

	for reference := math.MinInt16; reference <= math.MaxInt16; reference++ {
		integer, loss := IToF[int, float64](reference)
		require.False(t, loss)

		for id, fractional := range fractionals {
			number := integer + fractional

			if integer < 0 {
				number = integer - fractional
			}

			if id == 0 {
				require.InDelta(t, integer, number, 0.0)
			} else {
				require.InDelta(t, integer, number, fractional+fractional/2)
			}

			fi, si := testFToIInt(t, number, reference)
			fu, su := testFToIUint(t, number, reference)

			faultsInt += fi
			faultsUint += fu
			successfulInt += si
			successfulUint += su
		}
	}

	// Previous loop does not generate a sequence of 0-fractionals
	for _, fractional := range fractionals {
		fi, si := testFToIInt(t, -fractional, 0)
		fu, su := testFToIUint(t, -fractional, 0)

		faultsInt += fi
		faultsUint += fu
		successfulInt += si
		successfulUint += su
	}

	require.NotZero(t, faultsInt)
	require.NotZero(t, faultsUint)
	require.NotZero(t, successfulInt)
	require.NotZero(t, successfulUint)
}

func testFToIInt(t *testing.T, number float64, reference int) (int, int) {
	converted, err := FToI[float64, int8](number)

	if reference > math.MaxInt8 || reference < math.MinInt8 {
		require.Error(
			t,
			err,
			"number: %f, converted: %v, reference: %v",
			number,
			converted,
			reference,
		)

		return 1, 0
	}

	require.NoError(
		t,
		err,
		"number: %f, converted: %v, reference: %v",
		number,
		converted,
		reference,
	)

	require.Equal(t, reference, int(converted))

	return 0, 1
}

func testFToIUint(t *testing.T, number float64, reference int) (int, int) {
	converted, err := FToI[float64, uint8](number)

	if reference > math.MaxUint8 || reference < 0 {
		require.Error(
			t,
			err,
			"number: %f, converted: %v, reference: %v",
			number,
			converted,
			reference,
		)

		return 1, 0
	}

	require.NoError(
		t,
		err,
		"number: %f, converted: %v, reference: %v",
		number,
		converted,
		reference,
	)

	require.Equal(t, reference, int(converted))

	return 0, 1
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

func testPow(t *testing.T, maxPower int32) {
	faults := 0
	successful := 0

	// int32 is used because in its value range float64 does not lose the precision of
	// the integer part
	maxInt32, loss := IToF[int32, float64](math.MaxInt32)
	require.False(t, loss)

	minInt32, loss := IToF[int32, float64](math.MinInt32)
	require.False(t, loss)

	for base := int32(math.MinInt8); base <= math.MaxInt8; base++ {
		for power := int32(math.MinInt8); power <= maxPower; power++ {
			bf, loss := IToF[int32, float64](base)
			require.False(t, loss)

			pf, loss := IToF[int32, float64](power)
			require.False(t, loss)

			reference := math.Pow(bf, pf)

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

func BenchmarkIdle(b *testing.B) {
	for range b.N { //nolint:revive
	}
}

func BenchmarkAddReference(b *testing.B) {
	sum := 0

	// b.N, sum and require is used to prevent compiler optimizations
	for range b.N {
		sum = b.N + 3
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAdd(b *testing.B) {
	sum := 0

	// b.N, sum and require is used to prevent compiler optimizations
	for range b.N {
		sum, _ = Add(b.N, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddU(b *testing.B) {
	sum := uint(0)

	// b.N, sum and require is used to prevent compiler optimizations
	for range b.N {
		sum, _ = AddU(uint(b.N), 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkSubReference(b *testing.B) {
	diff := 0

	// b.N, diff and require is used to prevent compiler optimizations
	for range b.N {
		diff = b.N - 3
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSub(b *testing.B) {
	diff := 0

	// b.N, diff and require is used to prevent compiler optimizations
	for range b.N {
		diff, _ = Sub(b.N, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubU(b *testing.B) {
	diff := uint(0)

	// b.N, diff and require is used to prevent compiler optimizations
	for range b.N {
		diff, _ = SubU(uint(b.N), 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkMulReference(b *testing.B) {
	product := 0

	// b.N, product and require is used to prevent compiler optimizations
	for range b.N {
		product = b.N * 3
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMul(b *testing.B) {
	product := 0

	// b.N, product and require is used to prevent compiler optimizations
	for range b.N {
		product, _ = Mul(b.N, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkDivReference(b *testing.B) {
	quotient := 0

	// b.N, quotient and require is used to prevent compiler optimizations
	for range b.N {
		quotient = b.N / 3
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, quotient)
}

func BenchmarkDiv(b *testing.B) {
	quotient := 0

	// b.N, quotient and require is used to prevent compiler optimizations
	for range b.N {
		quotient, _ = Div(b.N, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, quotient)
}

func BenchmarkNegateReference(b *testing.B) {
	negated := 0

	// b.N, negated and require is used to prevent compiler optimizations
	for range b.N {
		negated = -b.N
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, negated)
}

func BenchmarkNegate(b *testing.B) {
	negated := 0

	// b.N, negated and require is used to prevent compiler optimizations
	for range b.N {
		negated, _ = Negate(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, negated)
}

func BenchmarkIToIReference(b *testing.B) {
	converted := uint(0)

	// b.N, converted and require is used to prevent compiler optimizations
	for range b.N {
		converted = uint(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkIToI(b *testing.B) {
	converted := uint(0)

	// b.N, converted and require is used to prevent compiler optimizations
	for range b.N {
		converted, _ = IToI[int, uint](b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkUToSReference(b *testing.B) {
	converted := 0

	// b.N, converted and require is used to prevent compiler optimizations
	for range b.N {
		converted = int(uint(b.N))
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkUToS(b *testing.B) {
	converted := 0

	// b.N, converted and require is used to prevent compiler optimizations
	for range b.N {
		converted, _ = UToS[uint, int](uint(b.N))
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkPow10Reference(b *testing.B) {
	product := float64(0)

	// b.N, product and require is used to prevent compiler optimizations
	for range b.N {
		product = math.Pow10(b.N % 19)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkPow10(b *testing.B) {
	product := uint64(0)

	// b.N, product and require is used to prevent compiler optimizations
	for range b.N {
		product, _ = Pow10[uint64](b.N % 19)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkIToFReference(b *testing.B) {
	converted := float64(0)

	// b.N, converted and require is used to prevent compiler optimizations
	for range b.N {
		converted = float64(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkIToF(b *testing.B) {
	converted := float64(0)

	// b.N, converted and require is used to prevent compiler optimizations
	for range b.N {
		converted, _ = IToF[int, float64](b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkFToIReference(b *testing.B) {
	converted := 0

	// b.N, converted and require is used to prevent compiler optimizations
	for range b.N {
		converted = int(float64(b.N))
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkFToI(b *testing.B) {
	converted := 0

	// b.N, converted and require is used to prevent compiler optimizations
	for range b.N {
		converted, _ = FToI[float64, int](float64(b.N))
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkPowReference(b *testing.B) {
	product := float64(0)

	// b.N, product and require is used to prevent compiler optimizations
	for range b.N {
		product = math.Pow(float64(b.N), float64(b.N))
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkPow(b *testing.B) {
	product := uint64(0)

	// b.N, product and require is used to prevent compiler optimizations
	for range b.N {
		product, _ = Pow(uint64(b.N), b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}
