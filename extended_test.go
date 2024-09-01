package safe

import (
	"math"
	"os"
	"slices"
	"testing"

	"github.com/akramarenkov/safe/internal/consts"
	"github.com/akramarenkov/safe/internal/inspect"
	"github.com/akramarenkov/safe/internal/inspect/dataset"

	"github.com/stretchr/testify/require"
)

func TestAddM(t *testing.T) {
	_, err := AddM[int](false)
	require.Error(t, err)

	testAddMInt(t)
	testAddMUint(t)
}

func testAddMInt(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...int8) (int8, error) {
			return AddM(false, args...)
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
	require.Zero(t, result.Overflows)
	require.Zero(t, result.ReferenceFaults)

	opts = inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 2,

		Inspected: func(args ...int8) (int8, error) {
			return AddM(false, args...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1], nil
		},
	}

	result, err = opts.Do()
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

	opts = inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return AddM(false, args...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] + args[2], nil
		},
	}

	result, err = opts.Do()
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

func testAddMUint(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...uint8) (uint8, error) {
			return AddM(false, args...)
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
	require.Zero(t, result.Overflows)
	require.Zero(t, result.ReferenceFaults)

	opts = inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 2,

		Inspected: func(args ...uint8) (uint8, error) {
			return AddM(false, args...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1], nil
		},
	}

	result, err = opts.Do()
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

	opts = inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return AddM(false, args...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] + args[2], nil
		},
	}

	result, err = opts.Do()
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

func TestAddMDataSet(t *testing.T) {
	inspected := func(args ...int8) (int8, error) {
		return AddM(false, args...)
	}

	result, err := dataset.InspectFromFile("dataset/addm", inspected)
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

func TestAddMDataSetUnmodify(t *testing.T) {
	inspected := func(args ...int8) (int8, error) {
		return AddM(true, args...)
	}

	result, err := dataset.InspectFromFile("dataset/addm", inspected)
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

func TestAddMCollectDataSet(t *testing.T) {
	if os.Getenv(consts.EnvCollectDataSet) == "" {
		t.SkipNow()
	}

	reference := func(args ...int64) (int64, error) {
		reference := int64(0)

		for _, arg := range args {
			reference += arg
		}

		return reference, nil
	}

	collector := dataset.Collector[int8]{
		ArgsQuantity:               6,
		NotOverflowedItemsQuantity: 1 << 15,
		OverflowedItemsQuantity:    1 << 15,
		Reference:                  reference,
	}

	err := collector.CollectToFile("dataset/addm")
	require.NoError(t, err)
}

func TestAddM4Args(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(consts.EnvEnableLongTest) == "" {
		t.SkipNow()
	}

	testAddM4ArgsInt(t, false)
	testAddM4ArgsUint(t, false)
}

func TestAddM4ArgsUnmodify(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(consts.EnvEnableLongTest) == "" {
		t.SkipNow()
	}

	testAddM4ArgsInt(t, true)
	testAddM4ArgsUint(t, true)
}

func testAddM4ArgsInt(t *testing.T, unmodify bool) {
	opts := inspect.Opts4[int8]{
		Inspected: func(first, second, third, fourth int8) (int8, error) {
			return AddM(unmodify, first, second, third, fourth)
		},
		Reference: func(first, second, third, fourth int64) (int64, error) {
			return first + second + third + fourth, nil
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

func testAddM4ArgsUint(t *testing.T, unmodify bool) {
	opts := inspect.Opts4[uint8]{
		Inspected: func(first, second, third, fourth uint8) (uint8, error) {
			return AddM(unmodify, first, second, third, fourth)
		},
		Reference: func(first, second, third, fourth int64) (int64, error) {
			return first + second + third + fourth, nil
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

func TestAddM5Args(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(consts.EnvEnableLongTest) == "" {
		t.SkipNow()
	}

	testAddM5ArgsInt(t, false)
	testAddM5ArgsUint(t, false)
}

func TestAddM5ArgsUnmodify(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(consts.EnvEnableLongTest) == "" {
		t.SkipNow()
	}

	testAddM5ArgsInt(t, true)
	testAddM5ArgsUint(t, true)
}

func testAddM5ArgsInt(t *testing.T, unmodify bool) {
	opts := inspect.Opts5[int8]{
		Inspected: func(first, second, third, fourth, fifth int8) (int8, error) {
			return AddM(unmodify, first, second, third, fourth, fifth)
		},
		Reference: func(first, second, third, fourth, fifth int64) (int64, error) {
			return first + second + third + fourth + fifth, nil
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

func testAddM5ArgsUint(t *testing.T, unmodify bool) {
	opts := inspect.Opts5[uint8]{
		Inspected: func(first, second, third, fourth, fifth uint8) (uint8, error) {
			return AddM(unmodify, first, second, third, fourth, fifth)
		},
		Reference: func(first, second, third, fourth, fifth int64) (int64, error) {
			return first + second + third + fourth + fifth, nil
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

func TestAdd3(t *testing.T) {
	testAdd3Int(t)
	testAdd3Uint(t)
}

func testAdd3Int(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return Add3(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] + args[2], nil
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

func testAdd3Uint(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return Add3(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] + args[2], nil
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

func TestAddUM(t *testing.T) {
	_, err := AddUM[uint]()
	require.Error(t, err)

	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: AddUM[uint8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] + args[2], nil
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

func TestSub3(t *testing.T) {
	testSub3Int(t)
	testSub3Uint(t)
}

func testSub3Int(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return Sub3(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1] - args[2], nil
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

func testSub3Uint(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return Sub3(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1] - args[2], nil
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

func TestSubUM(t *testing.T) {
	diff, err := SubUM(uint(math.MaxUint))
	require.NoError(t, err)
	require.Equal(t, uint(math.MaxUint), diff)

	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return SubUM(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1] - args[2], nil
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

func TestMulM(t *testing.T) {
	_, err := MulM[int]()
	require.Error(t, err)

	testMulMInt(t)
	testMulMUint(t)
}

func testMulMInt(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: MulM[int8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
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

func testMulMUint(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: MulM[uint8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
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

func TestCmpMulM(t *testing.T) {
	factors := []int{15, 0, 27, -1, -5}

	slices.SortFunc(factors, cmpMulM)
	require.Equal(t, []int{-1, -5, 0, 15, 27}, factors)
}

func TestMulT(t *testing.T) {
	testMulTInt(t)
	testMulTUint(t)
}

func testMulTInt(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return MulT(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
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

func testMulTUint(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return MulT(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
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

func TestMulUM(t *testing.T) {
	_, err := MulUM[uint]()
	require.Error(t, err)

	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: MulUM[uint8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
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

func TestDivM(t *testing.T) {
	quotient, err := DivM(uint(math.MaxUint))
	require.NoError(t, err)
	require.Equal(t, uint(math.MaxUint), quotient)

	testDivMInt(t)
	testDivMUint(t)
}

func testDivMInt(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return DivM(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[1] == 0 || args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return args[0] / args[1] / args[2], nil
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

func testDivMUint(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return DivM(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[1] == 0 || args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return args[0] / args[1] / args[2], nil
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

		require.Equal(t, uint64(consts.DecimalBase), current/previous, "power: %v", power)
		require.Equal(t, uint64(0), current%previous, "power: %v", power)
	}
}

func TestPow(t *testing.T) {
	opts := inspect.Opts[int32, int32, float64]{
		LoopsQuantity: 2,

		Inspected: func(args ...int32) (int32, error) {
			return Pow(args[0], args[1])
		},
		Reference: func(args ...float64) (float64, error) {
			reference := math.Pow(args[0], args[1])
			require.False(t, math.IsNaN(reference))

			return reference, nil
		},
		Span: func() (float64, float64) {
			return math.MinInt8, math.MaxInt8
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

func BenchmarkAdd3Reference(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	b.ResetTimer()

	for range b.N {
		sum = b.N + b.N + 1
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAdd3(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	b.ResetTimer()

	for range b.N {
		sum, _ = Add3(b.N, b.N, 1)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAdd3SpanIdle(b *testing.B) {
	// one, two, three and require is used to prevent compiler optimizations
	one := int8(0)
	two := int8(0)
	three := int8(0)

	span := benchSpanAdd3()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					one = first
					two = second
					three = third
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
	require.NotNil(b, two)
	require.NotNil(b, three)
}

func BenchmarkAdd3SpanReference(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := int8(0)

	span := benchSpanAdd3()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					sum = first + second + third
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAdd3Span(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := int8(0)

	span := benchSpanAdd3()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					sum, _ = Add3(first, second, third)
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddM2Args(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	b.ResetTimer()

	for range b.N {
		sum, _ = AddM(false, b.N, 1)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddM2ArgsSpan(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := int8(0)

	span := benchSpanAdd()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				sum, _ = AddM(false, first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddM3Args(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	b.ResetTimer()

	for range b.N {
		sum, _ = AddM(false, b.N, b.N, 1)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddM3ArgsSpan(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := int8(0)

	span := benchSpanAdd3()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					sum, _ = AddM(false, first, second, third)
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddMReference(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	b.ResetTimer()

	for range b.N {
		sum = b.N + b.N + b.N + b.N + b.N + 1
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddM(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	b.ResetTimer()

	for range b.N {
		sum, _ = AddM(false, b.N, b.N, b.N, b.N, b.N, 1)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddMSpanIdle(b *testing.B) {
	// one, two ... six and require is used to prevent compiler optimizations
	one := int8(0)
	two := int8(0)
	three := int8(0)
	four := int8(0)
	five := int8(0)
	six := int8(0)

	span := benchSpanAddM()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					for _, fourth := range span {
						for _, fifth := range span {
							for _, sixth := range span {
								one = first
								two = second
								three = third
								four = fourth
								five = fifth
								six = sixth
							}
						}
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
	require.NotNil(b, two)
	require.NotNil(b, three)
	require.NotNil(b, four)
	require.NotNil(b, five)
	require.NotNil(b, six)
}

func BenchmarkAddMSpanReference(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := int8(0)

	span := benchSpanAddM()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					for _, fourth := range span {
						for _, fifth := range span {
							for _, sixth := range span {
								sum = first + second + third + fourth + fifth + sixth
							}
						}
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddMSpan(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := int8(0)

	span := benchSpanAddM()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					for _, fourth := range span {
						for _, fifth := range span {
							for _, sixth := range span {
								sum, _ = AddM(
									false,
									first,
									second,
									third,
									fourth,
									fifth,
									sixth,
								)
							}
						}
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddUM2Args(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := uint(0)

	b.ResetTimer()

	for range b.N {
		sum, _ = AddUM(uint(b.N), 1)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddUM2ArgsSpan(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := uint8(0)

	span := benchSpanAddU()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				sum, _ = AddUM(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddUMSpanIdle(b *testing.B) {
	// one, two ... six and require is used to prevent compiler optimizations
	one := uint8(0)
	two := uint8(0)
	three := uint8(0)
	four := uint8(0)
	five := uint8(0)
	six := uint8(0)

	span := benchSpanAddUM()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					for _, fourth := range span {
						for _, fifth := range span {
							for _, sixth := range span {
								one = first
								two = second
								three = third
								four = fourth
								five = fifth
								six = sixth
							}
						}
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
	require.NotNil(b, two)
	require.NotNil(b, three)
	require.NotNil(b, four)
	require.NotNil(b, five)
	require.NotNil(b, six)
}

func BenchmarkAddUMSpanReference(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := uint8(0)

	span := benchSpanAddUM()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					for _, fourth := range span {
						for _, fifth := range span {
							for _, sixth := range span {
								sum = first + second + third + fourth + fifth + sixth
							}
						}
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddUMSpan(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := uint8(0)

	span := benchSpanAddUM()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					for _, fourth := range span {
						for _, fifth := range span {
							for _, sixth := range span {
								sum, _ = AddUM(
									first,
									second,
									third,
									fourth,
									fifth,
									sixth,
								)
							}
						}
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkSub3Reference(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := 0

	b.ResetTimer()

	for range b.N {
		diff = b.N - 2 - 1
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSub3(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := 0

	b.ResetTimer()

	for range b.N {
		diff, _ = Sub3(b.N, 2, 1)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSub3SpanIdle(b *testing.B) {
	// one, two, three and require is used to prevent compiler optimizations
	one := int8(0)
	two := int8(0)
	three := int8(0)

	span := benchSpanSub3()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					one = first
					two = second
					three = third
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
	require.NotNil(b, two)
	require.NotNil(b, three)
}

func BenchmarkSub3SpanReference(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := int8(0)

	span := benchSpanSub3()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					diff = first - second - third
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSub3Span(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := int8(0)

	span := benchSpanSub3()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					diff, _ = Sub3(first, second, third)
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubUM2Args(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := uint(0)

	b.ResetTimer()

	for range b.N {
		diff, _ = SubUM(uint(b.N), 1)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubUM2ArgsSpan(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := uint8(0)

	span := benchSpanSubU()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				diff, _ = SubUM(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubUMSpanIdle(b *testing.B) {
	// one, two ... six and require is used to prevent compiler optimizations
	one := uint8(0)
	two := uint8(0)
	three := uint8(0)
	four := uint8(0)
	five := uint8(0)
	six := uint8(0)

	span := benchSpanSubUM()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					for _, fourth := range span {
						for _, fifth := range span {
							for _, sixth := range span {
								one = first
								two = second
								three = third
								four = fourth
								five = fifth
								six = sixth
							}
						}
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, one)
	require.NotNil(b, two)
	require.NotNil(b, three)
	require.NotNil(b, four)
	require.NotNil(b, five)
	require.NotNil(b, six)
}

func BenchmarkSubUMSpanReference(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := uint8(0)

	span := benchSpanSubUM()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					for _, fourth := range span {
						for _, fifth := range span {
							for _, sixth := range span {
								diff = first + second + third + fourth + fifth + sixth
							}
						}
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubUMSpan(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := uint8(0)

	span := benchSpanSubUM()

	b.ResetTimer()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					for _, fourth := range span {
						for _, fifth := range span {
							for _, sixth := range span {
								diff, _ = SubUM(
									first,
									second,
									third,
									fourth,
									fifth,
									sixth,
								)
							}
						}
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkMulReference3Args(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					product = first * second * third
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMulReference4Args(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					for fourth := benchMinInt; fourth <= benchMaxInt; fourth++ {
						product = first * second * third * fourth
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMulM1Args(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			product, _ = MulM(first)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMulM2Args(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				product, _ = MulM(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMulM3Args(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					product, _ = MulM(first, second, third)
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMulM4Args(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					for fourth := benchMinInt; fourth <= benchMaxInt; fourth++ {
						product, _ = MulM(first, second, third, fourth)
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMulT(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					product, _ = MulT(first, second, third)
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMulUM1Args(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			product, _ = MulUM(first)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMulUM2Args(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			for second := benchMinUint; second <= benchMaxUint; second++ {
				product, _ = MulUM(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMulUM3Args(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			for second := benchMinUint; second <= benchMaxUint; second++ {
				for third := benchMinUint; third <= benchMaxUint; third++ {
					product, _ = MulUM(first, second, third)
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMulUM4Args(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			for second := benchMinUint; second <= benchMaxUint; second++ {
				for third := benchMinUint; third <= benchMaxUint; third++ {
					for fourth := benchMinUint; fourth <= benchMaxUint; fourth++ {
						product, _ = MulUM(first, second, third, fourth)
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkDivReference3Args(b *testing.B) {
	// quotient and require is used to prevent compiler optimizations
	quotient := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					if second == 0 || third == 0 {
						continue
					}

					quotient = first / second / third
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, quotient)
}

func BenchmarkDivReference4Args(b *testing.B) {
	// quotient and require is used to prevent compiler optimizations
	quotient := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					for fourth := benchMinInt; fourth <= benchMaxInt; fourth++ {
						if second == 0 || third == 0 || fourth == 0 {
							continue
						}

						quotient = first / second / third / fourth
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, quotient)
}

func BenchmarkDivM1Args(b *testing.B) {
	// quotient and require is used to prevent compiler optimizations
	quotient := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			quotient, _ = DivM(first)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, quotient)
}

func BenchmarkDivM2Args(b *testing.B) {
	// quotient and require is used to prevent compiler optimizations
	quotient := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				quotient, _ = DivM(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, quotient)
}

func BenchmarkDivM3Args(b *testing.B) {
	// quotient and require is used to prevent compiler optimizations
	quotient := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					quotient, _ = DivM(first, second, third)
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, quotient)
}

func BenchmarkDivM4Args(b *testing.B) {
	// quotient and require is used to prevent compiler optimizations
	quotient := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					for fourth := benchMinInt; fourth <= benchMaxInt; fourth++ {
						quotient, _ = DivM(first, second, third, fourth)
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, quotient)
}

func BenchmarkPow10Reference(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := float64(0)

	for range b.N {
		product = math.Pow10(19)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkPow10(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := uint64(0)

	for range b.N {
		product, _ = Pow10[uint64](19)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkPowReference(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := float64(0)

	for range b.N {
		product = math.Pow(14, 14)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkPow(b *testing.B) {
	// product and require is used to prevent compiler optimizations
	product := uint64(0)

	for range b.N {
		product, _ = Pow(uint64(14), 14)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}
