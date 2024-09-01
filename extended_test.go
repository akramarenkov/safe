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
	result := int8(0)

	span := benchSpanAdd3()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					result = first + second + third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAdd3(b *testing.B) {
	result := int8(0)

	span := benchSpanAdd3()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					result, _ = Add3(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddM2Args(b *testing.B) {
	result := int8(0)

	span := benchSpanAdd()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				result, _ = AddM(false, first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddM3Args(b *testing.B) {
	result := int8(0)

	span := benchSpanAdd3()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					result, _ = AddM(false, first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddMReference(b *testing.B) {
	result := int8(0)

	span := benchSpanAddM()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					for _, fourth := range span {
						for _, fifth := range span {
							for _, sixth := range span {
								result = first + second + third + fourth + fifth + sixth
							}
						}
					}
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddM(b *testing.B) {
	result := int8(0)

	span := benchSpanAddM()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					for _, fourth := range span {
						for _, fifth := range span {
							for _, sixth := range span {
								result, _ = AddM(
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

	require.NotNil(b, result)
}

func BenchmarkAddUM2Args(b *testing.B) {
	result := uint8(0)

	span := benchSpanAddU()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				result, _ = AddUM(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddUMReference(b *testing.B) {
	result := uint8(0)

	span := benchSpanAddUM()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					for _, fourth := range span {
						for _, fifth := range span {
							for _, sixth := range span {
								result = first + second + third + fourth + fifth + sixth
							}
						}
					}
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddUM(b *testing.B) {
	result := uint8(0)

	span := benchSpanAddUM()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					for _, fourth := range span {
						for _, fifth := range span {
							for _, sixth := range span {
								result, _ = AddUM(
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

	require.NotNil(b, result)
}

func BenchmarkSub3Reference(b *testing.B) {
	result := int8(0)

	span := benchSpanSub3()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					result = first - second - third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSub3(b *testing.B) {
	result := int8(0)

	span := benchSpanSub3()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					result, _ = Sub3(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubUM2Args(b *testing.B) {
	result := uint8(0)

	span := benchSpanSubU()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				result, _ = SubUM(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubUMReference(b *testing.B) {
	result := uint8(0)

	span := benchSpanSubUM()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					for _, fourth := range span {
						for _, fifth := range span {
							for _, sixth := range span {
								result = first + second + third + fourth + fifth + sixth
							}
						}
					}
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubUM(b *testing.B) {
	result := uint8(0)

	span := benchSpanSubUM()

	for range b.N {
		for _, first := range span {
			for _, second := range span {
				for _, third := range span {
					for _, fourth := range span {
						for _, fifth := range span {
							for _, sixth := range span {
								result, _ = SubUM(
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

	require.NotNil(b, result)
}

func BenchmarkMulReference3Args(b *testing.B) {
	result := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					result = first * second * third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulReference4Args(b *testing.B) {
	result := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					for fourth := benchMinInt; fourth <= benchMaxInt; fourth++ {
						result = first * second * third * fourth
					}
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulM1Args(b *testing.B) {
	result := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			result, _ = MulM(first)
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulM2Args(b *testing.B) {
	result := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				result, _ = MulM(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulM3Args(b *testing.B) {
	result := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					result, _ = MulM(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulM4Args(b *testing.B) {
	result := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					for fourth := benchMinInt; fourth <= benchMaxInt; fourth++ {
						result, _ = MulM(first, second, third, fourth)
					}
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulT(b *testing.B) {
	result := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					result, _ = MulT(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulUM1Args(b *testing.B) {
	result := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			result, _ = MulUM(first)
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulUM2Args(b *testing.B) {
	result := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			for second := benchMinUint; second <= benchMaxUint; second++ {
				result, _ = MulUM(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulUM3Args(b *testing.B) {
	result := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			for second := benchMinUint; second <= benchMaxUint; second++ {
				for third := benchMinUint; third <= benchMaxUint; third++ {
					result, _ = MulUM(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulUM4Args(b *testing.B) {
	result := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			for second := benchMinUint; second <= benchMaxUint; second++ {
				for third := benchMinUint; third <= benchMaxUint; third++ {
					for fourth := benchMinUint; fourth <= benchMaxUint; fourth++ {
						result, _ = MulUM(first, second, third, fourth)
					}
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkDivReference3Args(b *testing.B) {
	result := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					if second == 0 || third == 0 {
						continue
					}

					result = first / second / third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkDivReference4Args(b *testing.B) {
	result := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					for fourth := benchMinInt; fourth <= benchMaxInt; fourth++ {
						if second == 0 || third == 0 || fourth == 0 {
							continue
						}

						result = first / second / third / fourth
					}
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkDivM1Args(b *testing.B) {
	result := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			result, _ = DivM(first)
		}
	}

	require.NotNil(b, result)
}

func BenchmarkDivM2Args(b *testing.B) {
	result := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				result, _ = DivM(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkDivM3Args(b *testing.B) {
	result := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					result, _ = DivM(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkDivM4Args(b *testing.B) {
	result := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					for fourth := benchMinInt; fourth <= benchMaxInt; fourth++ {
						result, _ = DivM(first, second, third, fourth)
					}
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkPow10Reference(b *testing.B) {
	result := float64(0)

	for range b.N {
		result = math.Pow10(19)
	}

	require.NotNil(b, result)
}

func BenchmarkPow10(b *testing.B) {
	result := uint64(0)

	for range b.N {
		result, _ = Pow10[uint64](19)
	}

	require.NotNil(b, result)
}

func BenchmarkPowReference(b *testing.B) {
	result := float64(0)

	for range b.N {
		result = math.Pow(14, 14)
	}

	require.NotNil(b, result)
}

func BenchmarkPow(b *testing.B) {
	result := uint64(0)

	for range b.N {
		result, _ = Pow(uint64(14), 14)
	}

	require.NotNil(b, result)
}
