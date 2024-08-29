package safe

import (
	"math"
	"testing"

	"github.com/akramarenkov/safe/internal/inspect"

	"github.com/stretchr/testify/require"
)

const (
	benchMinInt = -10
	benchMaxInt = 10

	benchMinUint = uint(0)
	benchMaxUint = uint(20)
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
	inspected := func(number uint8) (int8, error) {
		return IToI[int8](number)
	}

	result, err := inspect.Conversion(inspected)
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

func testIToIS8ToU8(t *testing.T) {
	inspected := func(number int8) (uint8, error) {
		return IToI[uint8](number)
	}

	result, err := inspect.Conversion(inspected)
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

func testIToIS8ToU16(t *testing.T) {
	inspected := func(number int8) (uint16, error) {
		return IToI[uint16](number)
	}

	result, err := inspect.Conversion(inspected)
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

func testIToIU16ToS8(t *testing.T) {
	inspected := func(number uint16) (int8, error) {
		return IToI[int8](number)
	}

	result, err := inspect.Conversion(inspected)
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

func testIToIU16ToU8(t *testing.T) {
	inspected := func(number uint16) (uint8, error) {
		return IToI[uint8](number)
	}

	result, err := inspect.Conversion(inspected)
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

func testIToIS16ToS8(t *testing.T) {
	inspected := func(number int16) (int8, error) {
		return IToI[int8](number)
	}

	result, err := inspect.Conversion(inspected)
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

func testIToIS16ToU8(t *testing.T) {
	inspected := func(number int16) (uint8, error) {
		return IToI[uint8](number)
	}

	result, err := inspect.Conversion(inspected)
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

func TestUToS(t *testing.T) {
	testUToS8To8(t)
	testUToS16To8(t)
}

func testUToS8To8(t *testing.T) {
	inspected := func(number uint8) (int8, error) {
		return UToS[int8](number)
	}

	result, err := inspect.Conversion(inspected)
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

func testUToS16To8(t *testing.T) {
	inspected := func(number uint16) (int8, error) {
		return UToS[int8](number)
	}

	result, err := inspect.Conversion(inspected)
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
	additions := []float64{
		math.SmallestNonzeroFloat64,
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
		0.99999999999998,
	}

	for _, addition := range additions {
		testFToI(t, addition)
	}
}

func testFToI(t *testing.T, addition float64) {
	inspected := func(number int16) (int8, error) {
		if number < 0 {
			return FToI[int8](float64(number) - addition)
		}

		return FToI[int8](float64(number) + addition)
	}

	effective := func(result inspect.Result[int16, int8]) float64 {
		if len(result.Args) == 0 {
			return 0
		}

		if result.Args[0] < 0 {
			return float64(result.Args[0]) - addition
		}

		return float64(result.Args[0]) + addition
	}

	result, err := inspect.Conversion(inspected)
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"addition: %v, effective: %v, reference: %v, actual: %v, args: %v, err: %v",
		addition,
		effective(result),
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func TestFToISpecial(t *testing.T) {
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

func BenchmarkIdle(b *testing.B) {
	for range b.N {
		_ = b.N
	}
}

func BenchmarkAddReference(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	b.ResetTimer()

	for range b.N {
		sum = b.N + 1
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAdd(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	b.ResetTimer()

	for range b.N {
		sum, _ = Add(b.N, 1)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddSpanIdle(b *testing.B) {
	// one, two and require is used to prevent compiler optimizations
	one := int8(0)
	two := int8(0)

	span := benchSpanAdd()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				one = first
				two = second
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
	require.NotNil(b, two)
}

func BenchmarkAddSpanReference(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := int8(0)

	span := benchSpanAdd()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				sum = first + second
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddSpan(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := int8(0)

	span := benchSpanAdd()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				sum, _ = Add(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddUReference(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := uint(0)

	b.ResetTimer()

	for range b.N {
		sum = uint(b.N) + 1
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddU(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := uint(0)

	b.ResetTimer()

	for range b.N {
		sum, _ = AddU(uint(b.N), 1)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddUSpanIdle(b *testing.B) {
	// one, two and require is used to prevent compiler optimizations
	one := uint8(0)
	two := uint8(0)

	span := benchSpanAddU()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				one = first
				two = second
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
	require.NotNil(b, two)
}

func BenchmarkAddUSpanReference(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := uint8(0)

	span := benchSpanAddU()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				sum = first + second
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddUSpan(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := uint8(0)

	span := benchSpanAddU()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				sum, _ = AddU(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkSubReference(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := 0

	b.ResetTimer()

	for range b.N {
		diff = b.N - 1
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSub(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := 0

	b.ResetTimer()

	for range b.N {
		diff, _ = Sub(b.N, 1)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubSpanIdle(b *testing.B) {
	// one, two and require is used to prevent compiler optimizations
	one := int8(0)
	two := int8(0)

	span := benchSpanSub()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				one = first
				two = second
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
	require.NotNil(b, two)
}

func BenchmarkSubSpanReference(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := int8(0)

	span := benchSpanSub()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				diff = first - second
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubSpan(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := int8(0)

	span := benchSpanSub()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				diff, _ = Sub(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubUReference(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := uint(0)

	b.ResetTimer()

	for range b.N {
		diff = uint(b.N) - 1
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubU(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := uint(0)

	b.ResetTimer()

	for range b.N {
		diff, _ = SubU(uint(b.N), 1)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubUSpanIdle(b *testing.B) {
	// one, two and require is used to prevent compiler optimizations
	one := uint8(0)
	two := uint8(0)

	span := benchSpanSubU()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				one = first
				two = second
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
	require.NotNil(b, two)
}

func BenchmarkSubUSpanReference(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := uint8(0)

	span := benchSpanSubU()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				diff = first - second
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubUSpan(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := uint8(0)

	span := benchSpanSubU()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				diff, _ = SubU(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkMulReference(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := 0

	b.ResetTimer()

	for range b.N {
		product = 2 * b.N
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMul(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := 0

	b.ResetTimer()

	for range b.N {
		product, _ = Mul(2, b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMulSpanIdle(b *testing.B) {
	// one, two and require is used to prevent compiler optimizations
	one := int8(0)
	two := int8(0)

	span := benchSpanMul()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				one = first
				two = second
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
	require.NotNil(b, two)
}

func BenchmarkMulSpanReference(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := int8(0)

	span := benchSpanMul()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				product = first * second
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMulSpan(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := int8(0)

	span := benchSpanMul()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				product, _ = Mul(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkDivReference(b *testing.B) {
	// quotient and require is used to prevent compiler optimizations
	quotient := 0

	b.ResetTimer()

	for range b.N {
		quotient = b.N / 2
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, quotient)
}

func BenchmarkDiv(b *testing.B) {
	// quotient and require is used to prevent compiler optimizations
	quotient := 0

	b.ResetTimer()

	for range b.N {
		quotient, _ = Div(b.N, 2)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, quotient)
}

func BenchmarkDivSpanIdle(b *testing.B) {
	// one, two and require is used to prevent compiler optimizations
	one := int8(0)
	two := int8(0)

	span := benchSpanDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				one = first
				two = second
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
	require.NotNil(b, two)
}

func BenchmarkDivSpanReference(b *testing.B) {
	// quotient and require is used to prevent compiler optimizations
	quotient := int8(0)

	span := benchSpanDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				if second == 0 {
					continue
				}

				quotient = first / second
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, quotient)
}

func BenchmarkDivSpan(b *testing.B) {
	// quotient and require is used to prevent compiler optimizations
	quotient := int8(0)

	span := benchSpanDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				quotient, _ = Div(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, quotient)
}

func BenchmarkNegateReference(b *testing.B) {
	// negated and require is used to prevent compiler optimizations
	negated := 0

	b.ResetTimer()

	for range b.N {
		negated = -b.N
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, negated)
}

func BenchmarkNegate(b *testing.B) {
	// negated and require is used to prevent compiler optimizations
	negated := 0

	b.ResetTimer()

	for range b.N {
		negated, _ = Negate(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, negated)
}

func BenchmarkNegateSpanIdle(b *testing.B) {
	// one, two and require is used to prevent compiler optimizations
	one := int8(0)
	two := uint8(0)

	signed, unsigned := benchSpanNegate()

	b.ResetTimer()

	for range b.N {
		for _, number := range signed {
			one = number
		}

		for _, number := range unsigned {
			two = number
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
	require.NotNil(b, two)
}

func BenchmarkNegateSpanReference(b *testing.B) {
	// negated and require is used to prevent compiler optimizations
	negated := int8(0)
	negatedU := uint8(0)

	signed, unsigned := benchSpanNegate()

	b.ResetTimer()

	for range b.N {
		for _, number := range signed {
			negated = -number
		}

		for _, number := range unsigned {
			negatedU = -number
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, negated)
	require.NotNil(b, negatedU)
}

func BenchmarkNegateSpan(b *testing.B) {
	// negated and require is used to prevent compiler optimizations
	negated := int8(0)
	negatedU := uint8(0)

	signed, unsigned := benchSpanNegate()

	b.ResetTimer()

	for range b.N {
		for _, number := range signed {
			negated, _ = Negate(number)
		}

		for _, number := range unsigned {
			negatedU, _ = Negate(number)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, negated)
	require.NotNil(b, negatedU)
}

func BenchmarkNegateS(b *testing.B) {
	// negated and require is used to prevent compiler optimizations
	negated := 0

	b.ResetTimer()

	for range b.N {
		negated, _ = NegateS(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, negated)
}

func BenchmarkNegateSSpanIdle(b *testing.B) {
	// one and require is used to prevent compiler optimizations
	one := int8(0)

	span := benchSpanNegateS()

	b.ResetTimer()

	for range b.N {
		for _, number := range span {
			one = number
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
}

func BenchmarkNegateSSpanReference(b *testing.B) {
	// negated and require is used to prevent compiler optimizations
	negated := int8(0)

	span := benchSpanNegateS()

	b.ResetTimer()

	for range b.N {
		for _, number := range span {
			negated = -number
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, negated)
}

func BenchmarkNegateSSpan(b *testing.B) {
	// negated and require is used to prevent compiler optimizations
	negated := int8(0)

	span := benchSpanNegateS()

	b.ResetTimer()

	for range b.N {
		for _, number := range span {
			negated, _ = NegateS(number)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, negated)
}

func BenchmarkIToIReference(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := uint(0)

	b.ResetTimer()

	for range b.N {
		converted = uint(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkIToI(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := uint(0)

	b.ResetTimer()

	for range b.N {
		converted, _ = IToI[uint](b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkIToISpanIdle(b *testing.B) {
	// one and require is used to prevent compiler optimizations
	one := int8(0)
	two := uint8(0)
	three := uint16(0)

	s8, u8, u16 := benchSpanIToI()

	b.ResetTimer()

	for range b.N {
		for _, number := range s8 {
			one = number
		}

		for _, number := range u8 {
			two = number
		}

		for _, number := range u16 {
			three = number
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
	require.NotNil(b, two)
	require.NotNil(b, three)
}

func BenchmarkIToISpanReference(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := int8(0)
	convertedU := uint8(0)

	s8, u8, u16 := benchSpanIToI()

	b.ResetTimer()

	for range b.N {
		for _, number := range s8 {
			convertedU = uint8(number)
		}

		for _, number := range u8 {
			converted = int8(number)
		}

		for _, number := range u16 {
			convertedU = uint8(number)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
	require.NotNil(b, convertedU)
}

func BenchmarkIToISpan(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := int8(0)
	convertedU := uint8(0)

	s8, u8, u16 := benchSpanIToI()

	b.ResetTimer()

	for range b.N {
		for _, number := range s8 {
			convertedU, _ = IToI[uint8](number)
		}

		for _, number := range u8 {
			converted, _ = IToI[int8](number)
		}

		for _, number := range u16 {
			convertedU, _ = IToI[uint8](number)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
	require.NotNil(b, convertedU)
}

func BenchmarkUToSReference(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := 0

	b.ResetTimer()

	for range b.N {
		converted = int(uint(b.N))
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkUToS(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := 0

	b.ResetTimer()

	for range b.N {
		converted, _ = UToS[int](uint(b.N))
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkUToSSpanIdle(b *testing.B) {
	// one and require is used to prevent compiler optimizations
	one := uint8(0)
	two := uint16(0)

	u8, u16 := benchSpanUToS()

	b.ResetTimer()

	for range b.N {
		for _, number := range u8 {
			one = number
		}

		for _, number := range u16 {
			two = number
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
	require.NotNil(b, two)
}

func BenchmarkUToSSpanReference(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := int8(0)

	u8, u16 := benchSpanUToS()

	b.ResetTimer()

	for range b.N {
		for _, number := range u8 {
			converted = int8(number)
		}

		for _, number := range u16 {
			converted = int8(number)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkUToSSpan(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := int8(0)

	u8, u16 := benchSpanUToS()

	b.ResetTimer()

	for range b.N {
		for _, number := range u8 {
			converted, _ = UToS[int8](number)
		}

		for _, number := range u16 {
			converted, _ = UToS[int8](number)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkIToFReference(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := float64(0)

	b.ResetTimer()

	for range b.N {
		converted = float64(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkIToF(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := float64(0)

	b.ResetTimer()

	for range b.N {
		converted, _ = IToF[float64](b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkIToFSpanIdle(b *testing.B) {
	// one and require is used to prevent compiler optimizations
	one := int64(0)

	span := benchSpanIToF()

	b.ResetTimer()

	for range b.N {
		for _, number := range span {
			one = number
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
}

func BenchmarkIToFSpanReference(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := float64(0)

	span := benchSpanIToF()

	b.ResetTimer()

	for range b.N {
		for _, number := range span {
			converted = float64(number)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkIToFSpan(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := float64(0)

	span := benchSpanIToF()

	b.ResetTimer()

	for range b.N {
		for _, number := range span {
			converted, _ = IToF[float64](number)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkFToIReference(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := 0

	b.ResetTimer()

	for range b.N {
		converted = int(float64(b.N))
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkFToI(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := 0

	b.ResetTimer()

	for range b.N {
		converted, _ = FToI[int](float64(b.N))
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkFToISpanIdle(b *testing.B) {
	// one and require is used to prevent compiler optimizations
	one := float64(0)

	span := benchSpanFToI()

	b.ResetTimer()

	for range b.N {
		for _, number := range span {
			one = number
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
}

func BenchmarkFToISpanReference(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := 0

	span := benchSpanFToI()

	b.ResetTimer()

	for range b.N {
		for _, number := range span {
			converted = int(number)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}

func BenchmarkFToISpan(b *testing.B) {
	// converted and require is used to prevent compiler optimizations
	converted := 0

	span := benchSpanFToI()

	b.ResetTimer()

	for range b.N {
		for _, number := range span {
			converted, _ = FToI[int](number)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, converted)
}
