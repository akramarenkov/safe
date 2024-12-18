package safe

import (
	"math"
	"math/big"
	"testing"

	"github.com/akramarenkov/safe/internal/inspect"
	"github.com/akramarenkov/safe/internal/inspect/types"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	testAddSig(t)
	testAddUns(t)
}

func testAddSig(t *testing.T) {
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

func testAddUns(t *testing.T) {
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
	testSubSig(t)
	testSubUns(t)
}

func testSubSig(t *testing.T) {
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

func testSubUns(t *testing.T) {
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
	testMulSig(t)
	testMulUns(t)
}

func testMulSig(t *testing.T) {
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

func testMulUns(t *testing.T) {
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

func TestMulU(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 2,

		Inspected: func(args ...uint8) (uint8, error) {
			return MulU(args[0], args[1])
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
	testDivSig(t)
	testDivUns(t)
}

func testDivSig(t *testing.T) {
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

func testDivUns(t *testing.T) {
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
	testNegateSig(t)
	testNegateUns(t)
}

func testNegateSig(t *testing.T) {
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

func testNegateUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...uint8) (uint8, error) {
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
		testFToISig(t, addition)
		testFToIUns(t, addition)
	}
}

func testFToISig(t *testing.T, addition float64) {
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

func testFToIUns(t *testing.T, addition float64) {
	effective := func(result types.Result[int16, uint8, int64]) float64 {
		if len(result.Args) == 0 {
			return 0
		}

		if result.Args[0] < 0 {
			return float64(result.Args[0]) - addition
		}

		return float64(result.Args[0]) + addition
	}

	opts := inspect.Opts[int16, uint8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...int16) (uint8, error) {
			if args[0] < 0 {
				return FToI[uint8](float64(args[0]) - addition)
			}

			return FToI[uint8](float64(args[0]) + addition)
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

func TestShift(t *testing.T) {
	testShiftSig(t)
	testShiftUns(t)
	testShiftIntViaFloat(t)
	testShiftUintViaFloat(t)
}

func testShiftSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 2,

		Inspected: func(args ...int8) (int8, error) {
			return Shift(args[0], args[1])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[1] < 0 {
				return 0, ErrNegativeShift
			}

			shift, err := IToI[uint](args[1])
			require.NoError(t, err)

			shifted := new(big.Int).Lsh(big.NewInt(args[0]), shift)

			if !shifted.IsInt64() {
				return 0, ErrOverflow
			}

			return shifted.Int64(), nil
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

func testShiftUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 2,

		Inspected: func(args ...uint8) (uint8, error) {
			return Shift(args[0], args[1])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[1] < 0 {
				return 0, ErrNegativeShift
			}

			shift, err := IToI[uint](args[1])
			require.NoError(t, err)

			shifted := new(big.Int).Lsh(big.NewInt(args[0]), shift)

			if !shifted.IsInt64() {
				return 0, ErrOverflow
			}

			return shifted.Int64(), nil
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

func testShiftIntViaFloat(t *testing.T) {
	opts := inspect.Opts[int8, int8, float64]{
		LoopsQuantity: 2,

		Inspected: func(args ...int8) (int8, error) {
			return Shift(args[0], args[1])
		},
		Reference: func(args ...float64) (float64, error) {
			if args[1] < 0 {
				return 0, ErrNegativeShift
			}

			reference := args[0] * math.Pow(2, args[1])
			require.False(t, math.IsNaN(reference))

			return reference, nil
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

func testShiftUintViaFloat(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, float64]{
		LoopsQuantity: 2,

		Inspected: func(args ...uint8) (uint8, error) {
			return Shift(args[0], args[1])
		},
		Reference: func(args ...float64) (float64, error) {
			if args[1] < 0 {
				return 0, ErrNegativeShift
			}

			reference := args[0] * math.Pow(2, args[1])
			require.False(t, math.IsNaN(reference))

			return reference, nil
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

func TestAbs(t *testing.T) {
	require.Equal(t, uint64(math.MaxInt8), Abs[int8](math.MaxInt8))
	require.Equal(t, uint64(-(math.MinInt8 + 1)), Abs[int8](math.MinInt8+1))
	require.Equal(t, uint64(-math.MinInt8), Abs[int8](math.MinInt8))

	require.Equal(t, uint64(math.MaxInt64), Abs[int64](math.MaxInt64))
	require.Equal(t, uint64(-(math.MinInt64 + 1)), Abs[int64](math.MinInt64+1))
	require.Equal(t, uint64(-math.MinInt64), Abs[int64](math.MinInt64))
}
