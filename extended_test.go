package safe

import (
	"math"
	"os"
	"testing"

	"github.com/akramarenkov/safe/internal/consts"
	"github.com/akramarenkov/safe/internal/env"
	"github.com/akramarenkov/safe/internal/inspect"
	"github.com/akramarenkov/safe/internal/inspect/dataset"
	"github.com/akramarenkov/safe/internal/inspect/dataset/filler"

	"github.com/stretchr/testify/require"
)

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

func TestAdd3U(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return Add3U(args[0], args[1], args[2])
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

func TestAddM(t *testing.T) {
	testAddMInt(t, false)
	testAddMInt(t, true)
	testAddMUint(t, false)
	testAddMUint(t, true)
}

func testAddMInt(t *testing.T, unmodify bool) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...int8) (int8, error) {
			return AddM(unmodify, args...)
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
			return AddM(unmodify, args...)
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
			return AddM(unmodify, args...)
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

func testAddMUint(t *testing.T, unmodify bool) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...uint8) (uint8, error) {
			return AddM(unmodify, args...)
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
			return AddM(unmodify, args...)
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
			return AddM(unmodify, args...)
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

func TestAddMError(t *testing.T) {
	_, err := AddM[int](false)
	require.Error(t, err)
}

func TestAddMUnmodify(t *testing.T) {
	expected := []int8{126, 2, 1, 0, -127, -128}
	modified := []int8{126, 2, 1, 0, -127, -128}
	unmodified := []int8{126, 2, 1, 0, -127, -128}

	_, err := AddM(false, modified...)
	require.NoError(t, err)
	require.NotEqual(t, expected, modified)

	_, err = AddM(true, unmodified...)
	require.NoError(t, err)
	require.Equal(t, expected, unmodified)
}

func TestAddMDataset(t *testing.T) {
	testAddMDataset(t, false)
	testAddMDataset(t, true)
}

func testAddMDataset(t *testing.T, unmodify bool) {
	inspected := func(args ...int8) (int8, error) {
		return AddM(unmodify, args...)
	}

	result, err := dataset.InspectFromFile("testdata/addm", inspected)
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

func TestAddMCollectDataset(t *testing.T) {
	if os.Getenv(env.CollectDataset) == "" {
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

	err := collector.CollectToFile("testdata/addm")
	require.NoError(t, err)
}

func TestAddM4Args(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	testAddM4ArgsInt(t, false)
	testAddM4ArgsInt(t, true)
	testAddM4ArgsUint(t, false)
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
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	testAddM5ArgsInt(t, false)
	testAddM5ArgsInt(t, true)
	testAddM5ArgsUint(t, false)
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

func TestAddMU(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: AddMU[uint8],
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

func TestAddMUError(t *testing.T) {
	_, err := AddMU[uint]()
	require.Error(t, err)
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

func TestSub3U(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return Sub3U(args[0], args[1], args[2])
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

func TestSubM(t *testing.T) {
	testSubMInt(t, false)
	testSubMInt(t, true)
	testSubMUint(t, false)
	testSubMUint(t, true)
}

func testSubMInt(t *testing.T, unmodify bool) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...int8) (int8, error) {
			return SubM(unmodify, args[0], args[1:]...)
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
			return SubM(unmodify, args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1], nil
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
			return SubM(unmodify, args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1] - args[2], nil
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

func testSubMUint(t *testing.T, unmodify bool) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...uint8) (uint8, error) {
			return SubM(unmodify, args[0], args[1:]...)
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
			return SubM(unmodify, args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1], nil
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
			return SubM(unmodify, args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1] - args[2], nil
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

func TestSubMUnmodify(t *testing.T) {
	expected := []int8{-128, 127, 127, -2, -126, -128}
	modified := []int8{-128, 127, 127, -2, -126, -128}
	unmodified := []int8{-128, 127, 127, -2, -126, -128}

	_, err := SubM(false, modified[0], modified[1:]...)
	require.NoError(t, err)
	require.NotEqual(t, expected, modified)

	_, err = SubM(true, unmodified[0], unmodified[1:]...)
	require.NoError(t, err)
	require.Equal(t, expected, unmodified)
}

func TestSubMDataset(t *testing.T) {
	testSubMDataset(t, false)
	testSubMDataset(t, true)
}

func testSubMDataset(t *testing.T, unmodify bool) {
	inspected := func(args ...int8) (int8, error) {
		return SubM(unmodify, args[0], args[1:]...)
	}

	result, err := dataset.InspectFromFile("testdata/subm", inspected)
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

func TestSubMCollectDataset(t *testing.T) {
	if os.Getenv(env.CollectDataset) == "" {
		t.SkipNow()
	}

	reference := func(args ...int64) (int64, error) {
		reference := args[0]

		for _, arg := range args[1:] {
			reference -= arg
		}

		return reference, nil
	}

	collector := dataset.Collector[int8]{
		ArgsQuantity:               6,
		NotOverflowedItemsQuantity: 1 << 15,
		OverflowedItemsQuantity:    1 << 15,
		Reference:                  reference,
	}

	err := collector.CollectToFile("testdata/subm")
	require.NoError(t, err)
}

func TestSubM4Args(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	testSubM4ArgsInt(t, false)
	testSubM4ArgsInt(t, true)
	testSubM4ArgsUint(t, false)
	testSubM4ArgsUint(t, true)
}

func testSubM4ArgsInt(t *testing.T, unmodify bool) {
	opts := inspect.Opts4[int8]{
		Inspected: func(first, second, third, fourth int8) (int8, error) {
			return SubM(unmodify, first, second, third, fourth)
		},
		Reference: func(first, second, third, fourth int64) (int64, error) {
			return first - second - third - fourth, nil
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

func testSubM4ArgsUint(t *testing.T, unmodify bool) {
	opts := inspect.Opts4[uint8]{
		Inspected: func(first, second, third, fourth uint8) (uint8, error) {
			return SubM(unmodify, first, second, third, fourth)
		},
		Reference: func(first, second, third, fourth int64) (int64, error) {
			return first - second - third - fourth, nil
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

func TestSubM5Args(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	testSubM5ArgsInt(t, false)
	testSubM5ArgsInt(t, true)
	testSubM5ArgsUint(t, false)
	testSubM5ArgsUint(t, true)
}

func testSubM5ArgsInt(t *testing.T, unmodify bool) {
	opts := inspect.Opts5[int8]{
		Inspected: func(first, second, third, fourth, fifth int8) (int8, error) {
			return SubM(unmodify, first, second, third, fourth, fifth)
		},
		Reference: func(first, second, third, fourth, fifth int64) (int64, error) {
			return first - second - third - fourth - fifth, nil
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

func testSubM5ArgsUint(t *testing.T, unmodify bool) {
	opts := inspect.Opts5[uint8]{
		Inspected: func(first, second, third, fourth, fifth uint8) (uint8, error) {
			return SubM(unmodify, first, second, third, fourth, fifth)
		},
		Reference: func(first, second, third, fourth, fifth int64) (int64, error) {
			return first - second - third - fourth - fifth, nil
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

func TestSubMU(t *testing.T) {
	diff, err := SubMU(uint(math.MaxUint))
	require.NoError(t, err)
	require.Equal(t, uint(math.MaxUint), diff)

	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return SubMU(args[0], args[1], args[2])
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

func TestMul3(t *testing.T) {
	testMul3Int(t)
	testMul3Uint(t)
}

func testMul3Int(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return Mul3(args[0], args[1], args[2])
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

func testMul3Uint(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return Mul3(args[0], args[1], args[2])
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

func TestMulM(t *testing.T) {
	testMulMInt(t, false)
	testMulMInt(t, true)
	testMulMUint(t, false)
	testMulMUint(t, true)
}

func testMulMInt(t *testing.T, unmodify bool) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...int8) (int8, error) {
			return MulM(unmodify, args...)
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
			return MulM(unmodify, args...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1], nil
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
			return MulM(unmodify, args...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
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

func testMulMUint(t *testing.T, unmodify bool) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...uint8) (uint8, error) {
			return MulM(unmodify, args...)
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
			return MulM(unmodify, args...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1], nil
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
			return MulM(unmodify, args...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
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

func TestMulMError(t *testing.T) {
	_, err := MulM[int](false)
	require.Error(t, err)
}

func TestMulMUnmodify(t *testing.T) {
	expected := []int8{0, -126, -127, -128, -128, -128}
	modified := []int8{0, -126, -127, -128, -128, -128}
	unmodified := []int8{0, -126, -127, -128, -128, -128}

	_, err := MulM(false, modified...)
	require.NoError(t, err)
	require.NotEqual(t, expected, modified)

	_, err = MulM(true, unmodified...)
	require.NoError(t, err)
	require.Equal(t, expected, unmodified)
}

func TestMulMDataset(t *testing.T) {
	testMulMDataset(t, false)
	testMulMDataset(t, true)
}

func testMulMDataset(t *testing.T, unmodify bool) {
	inspected := func(args ...int8) (int8, error) {
		return MulM(unmodify, args...)
	}

	result, err := dataset.InspectFromFile("testdata/mulm", inspected)
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

func TestMulMCollectDataset(t *testing.T) {
	if os.Getenv(env.CollectDataset) == "" {
		t.SkipNow()
	}

	reference := func(args ...int64) (int64, error) {
		reference := args[0]

		for _, arg := range args[1:] {
			reference *= arg
		}

		return reference, nil
	}

	collector := dataset.Collector[int8]{
		ArgsQuantity:               6,
		NotOverflowedItemsQuantity: 1 << 15,
		OverflowedItemsQuantity:    1 << 15,
		Reference:                  reference,
		ReferenceLimits: map[int64]uint{
			0: 6,
		},
		Fillers: []filler.Filler[int8]{
			filler.NewSet[int8](),
			filler.NewSet(
				func() []int8 {
					return filler.Span[int8](-20, 20)
				},
			),
			filler.NewRand[int8](),
		},
	}

	err := collector.CollectToFile("testdata/mulm")
	require.NoError(t, err)
}

func TestMulM4Args(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	testMulM4ArgsInt(t, false)
	testMulM4ArgsInt(t, true)
	testMulM4ArgsUint(t, false)
	testMulM4ArgsUint(t, true)
}

func testMulM4ArgsInt(t *testing.T, unmodify bool) {
	opts := inspect.Opts4[int8]{
		Inspected: func(first, second, third, fourth int8) (int8, error) {
			return MulM(unmodify, first, second, third, fourth)
		},
		Reference: func(first, second, third, fourth int64) (int64, error) {
			return first * second * third * fourth, nil
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

func testMulM4ArgsUint(t *testing.T, unmodify bool) {
	opts := inspect.Opts4[uint8]{
		Inspected: func(first, second, third, fourth uint8) (uint8, error) {
			return MulM(unmodify, first, second, third, fourth)
		},
		Reference: func(first, second, third, fourth int64) (int64, error) {
			return first * second * third * fourth, nil
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

func TestMulM5Args(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	testMulM5ArgsInt(t, false)
	testMulM5ArgsInt(t, true)
	testMulM5ArgsUint(t, false)
	testMulM5ArgsUint(t, true)
}

func testMulM5ArgsInt(t *testing.T, unmodify bool) {
	opts := inspect.Opts5[int8]{
		Inspected: func(first, second, third, fourth, fifth int8) (int8, error) {
			return MulM(unmodify, first, second, third, fourth, fifth)
		},
		Reference: func(first, second, third, fourth, fifth int64) (int64, error) {
			return first * second * third * fourth * fifth, nil
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

func testMulM5ArgsUint(t *testing.T, unmodify bool) {
	opts := inspect.Opts5[uint8]{
		Inspected: func(first, second, third, fourth, fifth uint8) (uint8, error) {
			return MulM(unmodify, first, second, third, fourth, fifth)
		},
		Reference: func(first, second, third, fourth, fifth int64) (int64, error) {
			return first * second * third * fourth * fifth, nil
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

func TestMulMU(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: MulMU[uint8],
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

func TestMulMUError(t *testing.T) {
	_, err := MulMU[uint]()
	require.Error(t, err)
}

func TestDivM(t *testing.T) {
	testDivMInt(t)
	testDivMUint(t)
}

func testDivMInt(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...int8) (int8, error) {
			return DivM(args[0], args[1:]...)
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
			return DivM(args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			if args[1] == 0 {
				return 0, ErrDivisionByZero
			}

			return args[0] / args[1], nil
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
	require.NotZero(t, result.ReferenceFaults)

	opts = inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return DivM(args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			if args[1] == 0 || args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return args[0] / args[1] / args[2], nil
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
	require.NotZero(t, result.ReferenceFaults)
}

func testDivMUint(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...uint8) (uint8, error) {
			return DivM(args[0], args[1:]...)
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
			return DivM(args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			if args[1] == 0 {
				return 0, ErrDivisionByZero
			}

			return args[0] / args[1], nil
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
	require.Zero(t, result.Overflows)
	require.NotZero(t, result.ReferenceFaults)

	opts = inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return DivM(args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			if args[1] == 0 || args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return args[0] / args[1] / args[2], nil
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
	require.Zero(t, result.Overflows)
	require.NotZero(t, result.ReferenceFaults)
}

func TestDivMDataset(t *testing.T) {
	inspected := func(args ...int8) (int8, error) {
		return DivM(args[0], args[1:]...)
	}

	result, err := dataset.InspectFromFile("testdata/divm", inspected)
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

func TestDivMCollectDataset(t *testing.T) {
	if os.Getenv(env.CollectDataset) == "" {
		t.SkipNow()
	}

	reference := func(args ...int64) (int64, error) {
		reference := args[0]

		for _, arg := range args[1:] {
			if arg == 0 {
				return 0, ErrDivisionByZero
			}

			reference /= arg
		}

		return reference, nil
	}

	collector := dataset.Collector[int8]{
		ArgsQuantity:               6,
		NotOverflowedItemsQuantity: 1<<16 - 1<<4,
		OverflowedItemsQuantity:    1 << 4,
		Reference:                  reference,
		ReferenceLimits: map[int64]uint{
			0: 1 << 15,
		},
		Fillers: []filler.Filler[int8]{
			filler.NewSet[int8](),
			filler.NewSet(
				func() []int8 {
					return filler.Span[int8](-10, 10)
				},
			),
			filler.NewRand[int8](),
		},
	}

	err := collector.CollectToFile("testdata/divm")
	require.NoError(t, err)
}

func TestDivM4Args(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	testDivM4ArgsInt(t)
	testDivM4ArgsUint(t)
}

func testDivM4ArgsInt(t *testing.T) {
	opts := inspect.Opts4[int8]{
		Inspected: func(first, second, third, fourth int8) (int8, error) {
			return DivM(first, second, third, fourth)
		},
		Reference: func(first, second, third, fourth int64) (int64, error) {
			if second == 0 || third == 0 || fourth == 0 {
				return 0, ErrDivisionByZero
			}

			return first / second / third / fourth, nil
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

func testDivM4ArgsUint(t *testing.T) {
	opts := inspect.Opts4[uint8]{
		Inspected: func(first, second, third, fourth uint8) (uint8, error) {
			return DivM(first, second, third, fourth)
		},
		Reference: func(first, second, third, fourth int64) (int64, error) {
			if second == 0 || third == 0 || fourth == 0 {
				return 0, ErrDivisionByZero
			}

			return first / second / third / fourth, nil
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

func TestDivM5Args(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	testDivM5ArgsInt(t)
	testDivM5ArgsUint(t)
}

func testDivM5ArgsInt(t *testing.T) {
	opts := inspect.Opts5[int8]{
		Inspected: func(first, second, third, fourth, fifth int8) (int8, error) {
			return DivM(first, second, third, fourth, fifth)
		},
		Reference: func(first, second, third, fourth, fifth int64) (int64, error) {
			if second == 0 || third == 0 || fourth == 0 || fifth == 0 {
				return 0, ErrDivisionByZero
			}

			return first / second / third / fourth / fifth, nil
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

func testDivM5ArgsUint(t *testing.T) {
	opts := inspect.Opts5[uint8]{
		Inspected: func(first, second, third, fourth, fifth uint8) (uint8, error) {
			return DivM(first, second, third, fourth, fifth)
		},
		Reference: func(first, second, third, fourth, fifth int64) (int64, error) {
			if second == 0 || third == 0 || fourth == 0 || fifth == 0 {
				return 0, ErrDivisionByZero
			}

			return first / second / third / fourth / fifth, nil
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
