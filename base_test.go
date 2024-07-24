package safe

import (
	"math"
	"strconv"
	"testing"

	"github.com/akramarenkov/safe/internal/inspect"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	testAddInt(t)
	testAddUint(t)
}

func testAddInt(t *testing.T) {
	opts := inspect.Opts[int8]{
		LoopsQuantity: 2,

		Inspected: func(args ...int8) (int8, error) {
			return Add(args[0], args[1])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func testAddUint(t *testing.T) {
	opts := inspect.Opts[uint8]{
		LoopsQuantity: 2,

		Inspected: func(args ...uint8) (uint8, error) {
			return Add(args[0], args[1])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func TestAddU(t *testing.T) {
	opts := inspect.Opts[uint8]{
		LoopsQuantity: 2,

		Inspected: func(args ...uint8) (uint8, error) {
			return AddU(args[0], args[1])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func TestSub(t *testing.T) {
	testSubInt(t)
	testSubUint(t)
}

func testSubInt(t *testing.T) {
	opts := inspect.Opts[int8]{
		LoopsQuantity: 2,

		Inspected: func(args ...int8) (int8, error) {
			return Sub(args[0], args[1])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func testSubUint(t *testing.T) {
	opts := inspect.Opts[uint8]{
		LoopsQuantity: 2,

		Inspected: func(args ...uint8) (uint8, error) {
			return Sub(args[0], args[1])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func TestSubU(t *testing.T) {
	opts := inspect.Opts[uint8]{
		LoopsQuantity: 2,

		Inspected: func(args ...uint8) (uint8, error) {
			return SubU(args[0], args[1])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func TestMul(t *testing.T) {
	testMulInt(t)
	testMulUint(t)
}

func testMulInt(t *testing.T) {
	opts := inspect.Opts[int8]{
		LoopsQuantity: 2,

		Inspected: func(args ...int8) (int8, error) {
			return Mul(args[0], args[1])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func testMulUint(t *testing.T) {
	opts := inspect.Opts[uint8]{
		LoopsQuantity: 2,

		Inspected: func(args ...uint8) (uint8, error) {
			return Mul(args[0], args[1])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func TestDiv(t *testing.T) {
	testDivInt(t)
	testDivUint(t)
}

func testDivInt(t *testing.T) {
	opts := inspect.Opts[int8]{
		LoopsQuantity: 2,

		Inspected: func(args ...int8) (int8, error) {
			return Div(args[0], args[1])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[1] == 0 {
				return 0, ErrDivisionByZero
			}

			return args[0] / args[1], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
	require.NotZero(t, result.ReferenceFaults)
}

func testDivUint(t *testing.T) {
	opts := inspect.Opts[uint8]{
		LoopsQuantity: 2,

		Inspected: func(args ...uint8) (uint8, error) {
			return Div(args[0], args[1])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[1] == 0 {
				return 0, ErrDivisionByZero
			}

			return args[0] / args[1], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.Zero(t, result.Overflows)
	require.NotZero(t, result.ReferenceFaults)
}

func TestNegate(t *testing.T) {
	testNegateInt(t)
	testNegateUint(t)
}

func testNegateInt(t *testing.T) {
	opts := inspect.Opts[int8]{
		LoopsQuantity: 1,

		Inspected: func(args ...int8) (int8, error) {
			return Negate(args[0])
		},
		Reference: func(args ...int64) (int64, error) {
			return -args[0], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func testNegateUint(t *testing.T) {
	opts := inspect.Opts[uint8]{
		LoopsQuantity: 1,

		Inspected: func(args ...uint8) (uint8, error) {
			return Negate(args[0])
		},
		Reference: func(...int64) (int64, error) {
			return 0, ErrOverflow
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
}

func TestNegateS(t *testing.T) {
	opts := inspect.Opts[int8]{
		LoopsQuantity: 1,

		Inspected: func(args ...int8) (int8, error) {
			return NegateS(args[0])
		},
		Reference: func(args ...int64) (int64, error) {
			return -args[0], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
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
		converted, err := IToI[int8](uint8(number))

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
		converted, err := IToI[uint8](int8(number))

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
		converted, err := IToI[uint16](int8(number))

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
		converted, err := IToI[int8](uint16(number))

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
		converted, err := IToI[uint8](uint16(number))

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
		converted, err := IToI[int8](int16(number))

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
		converted, err := IToI[uint8](int16(number))

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
		converted, err := UToS[int8](uint8(number))

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
		converted, err := UToS[int8](uint16(number))

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

func TestIToF(t *testing.T) {
	testIToF32(t)
	testIToF64(t)
}

func testIToF32(t *testing.T) {
	_, err := IToF[float32](-1 << 24)
	require.NoError(t, err)

	_, err = IToF[float32](1 << 24)
	require.NoError(t, err)

	_, err = IToF[float32](-1<<24 - 1)
	require.Error(t, err)

	_, err = IToF[float32](1<<24 + 1)
	require.Error(t, err)

	_, err = IToF[float32](math.MinInt32)
	require.NoError(t, err)

	_, err = IToF[float32](math.MaxInt32)
	require.Error(t, err)

	_, err = IToF[float32](uint32(math.MaxUint32))
	require.Error(t, err)
}

func testIToF64(t *testing.T) {
	_, err := IToF[float64](math.MinInt32)
	require.NoError(t, err)

	_, err = IToF[float64](math.MaxInt32)
	require.NoError(t, err)

	_, err = IToF[float64](int64(-1 << 53))
	require.NoError(t, err)

	_, err = IToF[float64](int64(1 << 53))
	require.NoError(t, err)

	_, err = IToF[float64](int64(-1<<53 - 1))
	require.Error(t, err)

	_, err = IToF[float64](int64(1<<53 + 1))
	require.Error(t, err)

	_, err = IToF[float64](int64(math.MinInt64))
	require.NoError(t, err)

	_, err = IToF[float64](int64(math.MaxInt64))
	require.Error(t, err)

	_, err = IToF[float64](uint64(math.MaxUint64))
	require.Error(t, err)
}

func TestFToI(t *testing.T) {
	steps := []float64{
		0.1,
		0.2,
		0.25,
		0.333,
		0.444,
		0.5,
		0.555,
		0.666,
		0.75,
		0.777,
		0.999,
		// with max fractional value for 0.0
		// 0b0011111111101111111111111111111111111111111111111111111111111111
		0.9999999999999999,
		1,
		1.001,
		// with max fractional value for 1.0
		// 0b0011111111111111111111111111111111111111111111111111111111111111
		1.9999999999999997,
	}

	for _, step := range steps {
		t.Run(
			"step="+strconv.FormatFloat(step, 'f', -1, 64),
			func(t *testing.T) {
				t.Parallel()
				testFToIInt(t, step)
				testFToIUint(t, step)
			},
		)
	}

	_, err := FToI[int64](math.Inf(-1))
	require.Error(t, err)

	_, err = FToI[uint64](math.Inf(-1))
	require.Error(t, err)

	_, err = FToI[int64](math.Inf(0))
	require.Error(t, err)

	_, err = FToI[uint64](math.Inf(0))
	require.Error(t, err)

	_, err = FToI[int64](math.NaN())
	require.Error(t, err)

	_, err = FToI[uint64](math.NaN())
	require.Error(t, err)
}

func testFToIInt(t *testing.T, step float64) {
	faults := 0
	successful := 0

	begin := float64(2 * math.MinInt16)
	end := float64(2 * math.MaxInt16)

	// imprecision accumulation is acceptable
	for number := begin; number <= end; number += step {
		converted, err := FToI[int16](number)

		reference := int(number)

		if reference > math.MaxInt16 || reference < math.MinInt16 {
			require.Error(
				t,
				err,
				"converted: %v, reference: %v",
				converted,
				number,
			)

			faults++

			continue
		}

		require.NoError(
			t,
			err,
			"converted: %v, reference: %v",
			converted,
			number,
		)

		require.Equal(t, reference, int(converted))

		successful++
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func testFToIUint(t *testing.T, step float64) {
	faults := 0
	successful := 0

	begin := 0.0
	end := float64(2 * math.MaxUint16)

	// imprecision accumulation is acceptable
	for number := begin; number <= end; number += step {
		converted, err := FToI[uint16](number)

		reference := int(number)

		if reference > math.MaxUint16 {
			require.Error(
				t,
				err,
				"converted: %v, reference: %v",
				converted,
				number,
			)

			faults++

			continue
		}

		require.NoError(
			t,
			err,
			"converted: %v, reference: %v",
			converted,
			number,
		)

		require.Equal(t, reference, int(converted))

		successful++
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func BenchmarkIdle(b *testing.B) {
	for range b.N {
		_ = b.N
	}
}

func BenchmarkAddReference(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	for range b.N {
		for first := -3; first <= 3; first++ {
			for second := -3; second <= 3; second++ {
				sum = first + second
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAdd(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	for range b.N {
		for first := -3; first <= 3; first++ {
			for second := -3; second <= 3; second++ {
				sum, _ = Add(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddU(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := uint(0)

	for range b.N {
		for first := uint(0); first <= 6; first++ {
			for second := uint(0); second <= 6; second++ {
				sum, _ = AddU(first, second)
			}
		}
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

func BenchmarkNegateS(b *testing.B) {
	negated := 0

	// b.N, negated and require is used to prevent compiler optimizations
	for range b.N {
		negated, _ = NegateS(b.N)
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
		converted, _ = IToI[uint](b.N)
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
		converted, _ = UToS[int](uint(b.N))
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
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
		converted, _ = IToF[float64](b.N)
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
		converted, _ = FToI[int](float64(b.N))
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}
