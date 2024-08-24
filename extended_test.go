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
	opts := inspect.Opts[int8]{
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

	opts = inspect.Opts[int8]{
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

	opts = inspect.Opts[int8]{
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
}

func testAddMUint(t *testing.T) {
	opts := inspect.Opts[uint8]{
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

	opts = inspect.Opts[uint8]{
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

	opts = inspect.Opts[uint8]{
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
}

func TestAdd3(t *testing.T) {
	testAdd3Int(t)
	testAdd3Uint(t)
}

func testAdd3Int(t *testing.T) {
	opts := inspect.Opts[int8]{
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
}

func testAdd3Uint(t *testing.T) {
	opts := inspect.Opts[uint8]{
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
}

func TestAddUM(t *testing.T) {
	_, err := AddUM[uint]()
	require.Error(t, err)

	opts := inspect.Opts[uint8]{
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
}

func TestSub3(t *testing.T) {
	testSub3Int(t)
	testSub3Uint(t)
}

func testSub3Int(t *testing.T) {
	opts := inspect.Opts[int8]{
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
}

func testSub3Uint(t *testing.T) {
	opts := inspect.Opts[uint8]{
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
}

func TestSubUM(t *testing.T) {
	diff, err := SubUM(uint(math.MaxUint))
	require.NoError(t, err)
	require.Equal(t, uint(math.MaxUint), diff)

	opts := inspect.Opts[uint8]{
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
}

func TestMulM(t *testing.T) {
	_, err := MulM[int]()
	require.Error(t, err)

	testMulMInt(t)
	testMulMUint(t)
}

func testMulMInt(t *testing.T) {
	opts := inspect.Opts[int8]{
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
}

func testMulMUint(t *testing.T) {
	opts := inspect.Opts[uint8]{
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
	opts := inspect.Opts[int8]{
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
}

func testMulTUint(t *testing.T) {
	opts := inspect.Opts[uint8]{
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
}

func TestMulUM(t *testing.T) {
	_, err := MulUM[uint]()
	require.Error(t, err)

	opts := inspect.Opts[uint8]{
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
}

func TestDivM(t *testing.T) {
	quotient, err := DivM(uint(math.MaxUint))
	require.NoError(t, err)
	require.Equal(t, uint(math.MaxUint), quotient)

	testDivMInt(t)
	testDivMUint(t)
}

func testDivMInt(t *testing.T) {
	opts := inspect.Opts[int8]{
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
	opts := inspect.Opts[uint8]{
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
	faults := 0
	successful := 0

	// Is used int32 because in its value range float64 does not lose the precision of
	// the integer part and comparison of reference and tested values ​​can be done simply
	maxInt32 := float64(math.MaxInt32)
	minInt32 := float64(math.MinInt32)

	for base := int32(math.MinInt8); base <= math.MaxInt8; base++ {
		for power := int32(math.MinInt8); power <= math.MaxInt8; power++ {
			reference := math.Pow(float64(base), float64(power))
			require.False(t, math.IsNaN(reference))

			// To ensure that overflow conditions are satisfied correctly when
			// obtaining reference values ​​with non-zero fractional parts and close
			// to maximum/minimum int32 values, reference values ​​in these areas
			// are checked separately
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

func BenchmarkReference1Args(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			number = first
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkAddReference3Args(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					sum = first + second + third
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddReference4Args(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					for fourth := benchMinInt; fourth <= benchMaxInt; fourth++ {
						sum = first + second + third + fourth
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddM1Args(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			sum, _ = AddM(false, first)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddM2Args(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
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

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					sum, _ = AddM(false, first, second, third)
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddM4Args(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					for fourth := benchMinInt; fourth <= benchMaxInt; fourth++ {
						sum, _ = AddM(false, first, second, third, fourth)
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAdd3(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					sum, _ = Add3(first, second, third)
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddUM1Args(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			sum, _ = AddUM(first)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddUM2Args(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			for second := benchMinUint; second <= benchMaxUint; second++ {
				sum, _ = AddUM(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddUM3Args(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			for second := benchMinUint; second <= benchMaxUint; second++ {
				for third := benchMinUint; third <= benchMaxUint; third++ {
					sum, _ = AddUM(first, second, third)
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddUM4Args(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			for second := benchMinUint; second <= benchMaxUint; second++ {
				for third := benchMinUint; third <= benchMaxUint; third++ {
					for fourth := benchMinUint; fourth <= benchMaxUint; fourth++ {
						sum, _ = AddUM(first, second, third, fourth)
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkSubReference3Args(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					diff = first - second - third
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubReference4Args(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					for fourth := benchMinInt; fourth <= benchMaxInt; fourth++ {
						diff = first - second - third - fourth
					}
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSub3(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := 0

	for range b.N {
		for first := benchMinInt; first <= benchMaxInt; first++ {
			for second := benchMinInt; second <= benchMaxInt; second++ {
				for third := benchMinInt; third <= benchMaxInt; third++ {
					diff, _ = Sub3(first, second, third)
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubUM1Args(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			diff, _ = SubUM(first)
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubUM2Args(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			for second := benchMinUint; second <= benchMaxUint; second++ {
				diff, _ = SubUM(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubUM3Args(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			for second := benchMinUint; second <= benchMaxUint; second++ {
				for third := benchMinUint; third <= benchMaxUint; third++ {
					diff, _ = SubUM(first, second, third)
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubUM4Args(b *testing.B) {
	// diff and require is used to prevent compiler optimizations
	diff := uint(0)

	for range b.N {
		for first := benchMinUint; first <= benchMaxUint; first++ {
			for second := benchMinUint; second <= benchMaxUint; second++ {
				for third := benchMinUint; third <= benchMaxUint; third++ {
					for fourth := benchMinUint; fourth <= benchMaxUint; fourth++ {
						diff, _ = SubUM(first, second, third, fourth)
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
