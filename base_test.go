package safe

import (
	"math"
	"testing"

	"github.com/akramarenkov/safe/internal/inspect"
	"github.com/akramarenkov/safe/internal/inspect/types"

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
	opts := inspect.Opts[int8, int8, int64]{
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
	require.Zero(t, result.ReferenceFaults)
}

func testAddUint(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
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
	require.Zero(t, result.ReferenceFaults)
}

func TestAddU(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
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
	require.Zero(t, result.ReferenceFaults)
}

func TestSub(t *testing.T) {
	testSubInt(t)
	testSubUint(t)
}

func testSubInt(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
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
	require.Zero(t, result.ReferenceFaults)
}

func testSubUint(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
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
	require.Zero(t, result.ReferenceFaults)
}

func TestSubU(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
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
	require.Zero(t, result.ReferenceFaults)
}

func TestMul(t *testing.T) {
	testMulInt(t)
	testMulUint(t)
}

func testMulInt(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
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
	require.Zero(t, result.ReferenceFaults)
}

func testMulUint(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
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
	require.Zero(t, result.ReferenceFaults)
}

func TestDiv(t *testing.T) {
	testDivInt(t)
	testDivUint(t)
}

func testDivInt(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
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
	opts := inspect.Opts[uint8, uint8, int64]{
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
	opts := inspect.Opts[int8, int8, int64]{
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
	require.Zero(t, result.ReferenceFaults)
}

func testNegateUint(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
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
	opts := inspect.Opts[int8, int8, int64]{
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
	require.Zero(t, result.ReferenceFaults)
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
	opts := inspect.Opts[uint8, int8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...uint8) (int8, error) {
			return IToI[int8](args[0])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
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
	require.Zero(t, result.ReferenceFaults)
}

func testIToIS8ToU8(t *testing.T) {
	opts := inspect.Opts[int8, uint8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...int8) (uint8, error) {
			return IToI[uint8](args[0])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
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
	require.Zero(t, result.ReferenceFaults)
}

func testIToIS8ToU16(t *testing.T) {
	opts := inspect.Opts[int8, uint16, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...int8) (uint16, error) {
			return IToI[uint16](args[0])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
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
	require.Zero(t, result.ReferenceFaults)
}

func testIToIU16ToS8(t *testing.T) {
	opts := inspect.Opts[uint16, int8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...uint16) (int8, error) {
			return IToI[int8](args[0])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
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
	require.Zero(t, result.ReferenceFaults)
}

func testIToIU16ToU8(t *testing.T) {
	opts := inspect.Opts[uint16, uint8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...uint16) (uint8, error) {
			return IToI[uint8](args[0])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
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
	require.Zero(t, result.ReferenceFaults)
}

func testIToIS16ToS8(t *testing.T) {
	opts := inspect.Opts[int16, int8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...int16) (int8, error) {
			return IToI[int8](args[0])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
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
	require.Zero(t, result.ReferenceFaults)
}

func testIToIS16ToU8(t *testing.T) {
	opts := inspect.Opts[int16, uint8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...int16) (uint8, error) {
			return IToI[uint8](args[0])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
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
	require.Zero(t, result.ReferenceFaults)
}

func TestUToS(t *testing.T) {
	testUToS8To8(t)
	testUToS16To8(t)
}

func testUToS8To8(t *testing.T) {
	opts := inspect.Opts[uint8, int8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...uint8) (int8, error) {
			return UToS[int8](args[0])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
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
	require.Zero(t, result.ReferenceFaults)
}

func testUToS16To8(t *testing.T) {
	opts := inspect.Opts[uint16, int8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...uint16) (int8, error) {
			return UToS[int8](args[0])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
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
	require.Zero(t, result.ReferenceFaults)
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
	effective := func(result types.Result[int16, int8, int64]) float64 {
		if len(result.Args) == 0 {
			return 0
		}

		if result.Args[0] < 0 {
			return float64(result.Args[0]) - addition
		}

		return float64(result.Args[0]) + addition
	}

	opts := inspect.Opts[int16, int8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...int16) (int8, error) {
			if args[0] < 0 {
				return FToI[int8](float64(args[0]) - addition)
			}

			return FToI[int8](float64(args[0]) + addition)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
		},
	}

	result, err := opts.Do()
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
	require.Zero(t, result.ReferenceFaults)
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

func BenchmarkAddReference(b *testing.B) {
	result := int8(0)

	span := benchSpanAdd()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				result = first + second
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAdd(b *testing.B) {
	result := int8(0)

	span := benchSpanAdd()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				result, _ = Add(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddUReference(b *testing.B) {
	result := uint8(0)

	span := benchSpanAddU()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				result = first + second
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddU(b *testing.B) {
	result := uint8(0)

	span := benchSpanAddU()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				result, _ = AddU(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubReference(b *testing.B) {
	result := int8(0)

	span := benchSpanSub()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				result = first - second
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSub(b *testing.B) {
	result := int8(0)

	span := benchSpanSub()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				result, _ = Sub(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubUReference(b *testing.B) {
	result := uint8(0)

	span := benchSpanSubU()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				result = first - second
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubU(b *testing.B) {
	result := uint8(0)

	span := benchSpanSubU()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				result, _ = SubU(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulReference(b *testing.B) {
	result := int8(0)

	span := benchSpanMul()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				result = first * second
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMul(b *testing.B) {
	result := int8(0)

	span := benchSpanMul()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				result, _ = Mul(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkDivReference(b *testing.B) {
	result := int8(0)

	span := benchSpanDiv()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				if second == 0 {
					continue
				}

				result = first / second
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkDiv(b *testing.B) {
	result := int8(0)

	span := benchSpanDiv()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				result, _ = Div(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkNegateReference(b *testing.B) {
	result := int8(0)
	resultU := uint8(0)

	signed, unsigned := benchSpanNegate()

	for range b.N {
		for _, number := range signed {
			result = -number
		}

		for _, number := range unsigned {
			resultU = -number
		}
	}

	require.NotNil(b, result)
	require.NotNil(b, resultU)
}

func BenchmarkNegate(b *testing.B) {
	result := int8(0)
	resultU := uint8(0)

	signed, unsigned := benchSpanNegate()

	for range b.N {
		for _, number := range signed {
			result, _ = Negate(number)
		}

		for _, number := range unsigned {
			resultU, _ = Negate(number)
		}
	}

	require.NotNil(b, result)
	require.NotNil(b, resultU)
}

func BenchmarkNegateSReference(b *testing.B) {
	result := int8(0)

	span := benchSpanNegateS()

	for range b.N {
		for _, number := range span {
			result = -number
		}
	}

	require.NotNil(b, result)
}

func BenchmarkNegateS(b *testing.B) {
	result := int8(0)

	span := benchSpanNegateS()

	for range b.N {
		for _, number := range span {
			result, _ = NegateS(number)
		}
	}

	require.NotNil(b, result)
}

func BenchmarkIToIReference(b *testing.B) {
	result := int8(0)
	resultU := uint8(0)

	s8, u8, u16 := benchSpanIToI()

	for range b.N {
		for _, number := range s8 {
			resultU = uint8(number)
		}

		for _, number := range u8 {
			result = int8(number)
		}

		for _, number := range u16 {
			resultU = uint8(number)
		}
	}

	require.NotNil(b, result)
	require.NotNil(b, resultU)
}

func BenchmarkIToI(b *testing.B) {
	result := int8(0)
	resultU := uint8(0)

	s8, u8, u16 := benchSpanIToI()

	for range b.N {
		for _, number := range s8 {
			resultU, _ = IToI[uint8](number)
		}

		for _, number := range u8 {
			result, _ = IToI[int8](number)
		}

		for _, number := range u16 {
			resultU, _ = IToI[uint8](number)
		}
	}

	require.NotNil(b, result)
	require.NotNil(b, resultU)
}

func BenchmarkUToSReference(b *testing.B) {
	result := int8(0)

	u8, u16 := benchSpanUToS()

	for range b.N {
		for _, number := range u8 {
			result = int8(number)
		}

		for _, number := range u16 {
			result = int8(number)
		}
	}

	require.NotNil(b, result)
}

func BenchmarkUToS(b *testing.B) {
	result := int8(0)

	u8, u16 := benchSpanUToS()

	for range b.N {
		for _, number := range u8 {
			result, _ = UToS[int8](number)
		}

		for _, number := range u16 {
			result, _ = UToS[int8](number)
		}
	}

	require.NotNil(b, result)
}

func BenchmarkIToFReference(b *testing.B) {
	result := float64(0)

	span := benchSpanIToF()

	for range b.N {
		for _, number := range span {
			result = float64(number)
		}
	}

	require.NotNil(b, result)
}

func BenchmarkIToF(b *testing.B) {
	result := float64(0)

	span := benchSpanIToF()

	for range b.N {
		for _, number := range span {
			result, _ = IToF[float64](number)
		}
	}

	require.NotNil(b, result)
}

func BenchmarkFToIReference(b *testing.B) {
	result := 0

	span := benchSpanFToI()

	for range b.N {
		for _, number := range span {
			result = int(number)
		}
	}

	require.NotNil(b, result)
}

func BenchmarkFToI(b *testing.B) {
	result := 0

	span := benchSpanFToI()

	for range b.N {
		for _, number := range span {
			result, _ = FToI[int](number)
		}
	}

	require.NotNil(b, result)
}
