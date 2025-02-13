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

func TestAdd3Sig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return Add3(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] + args[2], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestAdd3Uns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return Add3(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] + args[2], nil
		},
	}

	result, err := inspect.Do(opts)
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

	result, err := inspect.Do(opts)
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

func TestAddM1ArgsSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 1,

		Inspected: AddM[int8],
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
		},
	}

	result, err := inspect.Do(opts)
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
}

func TestAddM2ArgsSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 2,

		Inspected: AddM[int8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestAddM3ArgsSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: AddM[int8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] + args[2], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestAddM1ArgsUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 1,

		Inspected: AddM[uint8],
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
		},
	}

	result, err := inspect.Do(opts)
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
}

func TestAddM2ArgsUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 2,

		Inspected: AddM[uint8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestAddM3ArgsUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: AddM[uint8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] + args[2], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestAddMErrorMissingArguments(t *testing.T) {
	_, err := AddM[int]()
	require.Error(t, err)
}

func TestAddMUnmodify(t *testing.T) {
	expected := []int8{126, 2, 1, 0, -127, -128}
	unmodified := []int8{126, 2, 1, 0, -127, -128}

	_, err := AddM(unmodified...)
	require.NoError(t, err)
	require.Equal(t, expected, unmodified)
}

func TestAddMDataset(t *testing.T) {
	result, err := dataset.InspectFromFile("testdata/addm", AddM[int8])
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

	opts := dataset.Opts[int8]{
		ArgsQuantity:               6,
		NotOverflowedItemsQuantity: 1 << 15,
		OverflowedItemsQuantity:    1 << 15,
		Reference:                  reference,
	}

	err := dataset.CollectToFile(opts, "testdata/addm")
	require.NoError(t, err)
}

func TestAddM4ArgsSig(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth int8) (int8, error) {
		return AddM(first, second, third, fourth)
	}

	reference := func(first, second, third, fourth int64) (int64, error) {
		return first + second + third + fourth, nil
	}

	result, err := inspect.Do4(inspected, reference)
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

func TestAddM4ArgsUns(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth uint8) (uint8, error) {
		return AddM(first, second, third, fourth)
	}

	reference := func(first, second, third, fourth int64) (int64, error) {
		return first + second + third + fourth, nil
	}

	result, err := inspect.Do4(inspected, reference)
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

func TestAddM5ArgsSig(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth, fifth int8) (int8, error) {
		return AddM(first, second, third, fourth, fifth)
	}

	reference := func(first, second, third, fourth, fifth int64) (int64, error) {
		return first + second + third + fourth + fifth, nil
	}

	result, err := inspect.Do5(inspected, reference)
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

func TestAddM5ArgsUns(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth, fifth uint8) (uint8, error) {
		return AddM(first, second, third, fourth, fifth)
	}

	reference := func(first, second, third, fourth, fifth int64) (int64, error) {
		return first + second + third + fourth + fifth, nil
	}

	result, err := inspect.Do5(inspected, reference)
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

	result, err := inspect.Do(opts)
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

func TestAddMUErrorMissingArguments(t *testing.T) {
	_, err := AddMU[uint]()
	require.Error(t, err)
}

func TestSub3Sig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return Sub3(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1] - args[2], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestSub3Uns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return Sub3(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1] - args[2], nil
		},
	}

	result, err := inspect.Do(opts)
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

	result, err := inspect.Do(opts)
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

func TestSubM1ArgsSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...int8) (int8, error) {
			return SubM(args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
		},
	}

	result, err := inspect.Do(opts)
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
}

func TestSubM2ArgsSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 2,

		Inspected: func(args ...int8) (int8, error) {
			return SubM(args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestSubM3ArgsSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return SubM(args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1] - args[2], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestSubM1ArgsUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...uint8) (uint8, error) {
			return SubM(args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
		},
	}

	result, err := inspect.Do(opts)
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
}

func TestSubM2ArgsUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 2,

		Inspected: func(args ...uint8) (uint8, error) {
			return SubM(args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestSubM3ArgsUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return SubM(args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1] - args[2], nil
		},
	}

	result, err := inspect.Do(opts)
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
	unmodified := []int8{-128, 127, 127, -2, -126, -128}

	_, err := SubM(unmodified[0], unmodified[1:]...)
	require.NoError(t, err)
	require.Equal(t, expected, unmodified)
}

func TestSubMDataset(t *testing.T) {
	inspected := func(args ...int8) (int8, error) {
		return SubM(args[0], args[1:]...)
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

	opts := dataset.Opts[int8]{
		ArgsQuantity:               6,
		NotOverflowedItemsQuantity: 1 << 15,
		OverflowedItemsQuantity:    1 << 15,
		Reference:                  reference,
	}

	err := dataset.CollectToFile(opts, "testdata/subm")
	require.NoError(t, err)
}

func TestSubM4ArgsSig(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth int8) (int8, error) {
		return SubM(first, second, third, fourth)
	}

	reference := func(first, second, third, fourth int64) (int64, error) {
		return first - second - third - fourth, nil
	}

	result, err := inspect.Do4(inspected, reference)
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

func TestSubM4ArgsUns(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth uint8) (uint8, error) {
		return SubM(first, second, third, fourth)
	}

	reference := func(first, second, third, fourth int64) (int64, error) {
		return first - second - third - fourth, nil
	}

	result, err := inspect.Do4(inspected, reference)
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

func TestSubM5ArgsSig(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth, fifth int8) (int8, error) {
		return SubM(first, second, third, fourth, fifth)
	}

	reference := func(first, second, third, fourth, fifth int64) (int64, error) {
		return first - second - third - fourth - fifth, nil
	}

	result, err := inspect.Do5(inspected, reference)
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

func TestSubM5ArgsUns(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth, fifth uint8) (uint8, error) {
		return SubM(first, second, third, fourth, fifth)
	}

	reference := func(first, second, third, fourth, fifth int64) (int64, error) {
		return first - second - third - fourth - fifth, nil
	}

	result, err := inspect.Do5(inspected, reference)
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

	result, err := inspect.Do(opts)
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

func TestMul3Sig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return Mul3(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestMul3Uns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return Mul3(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestMul3U(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return Mul3U(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestMulM1ArgsSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 1,

		Inspected: MulM[int8],
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
		},
	}

	result, err := inspect.Do(opts)
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
}

func TestMulM2ArgsSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 2,

		Inspected: MulM[int8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestMulM3ArgsSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: MulM[int8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestMulM1ArgsUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 1,

		Inspected: MulM[uint8],
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
		},
	}

	result, err := inspect.Do(opts)
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
}

func TestMulM2ArgsUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 2,

		Inspected: MulM[uint8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestMulM3ArgsUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: MulM[uint8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
		},
	}

	result, err := inspect.Do(opts)
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

func TestMulMErrorMissingArguments(t *testing.T) {
	_, err := MulM[int]()
	require.Error(t, err)
}

func TestMulMUnmodify(t *testing.T) {
	expected := []int8{0, -126, -127, -128, -128, -128}
	unmodified := []int8{0, -126, -127, -128, -128, -128}

	_, err := MulM(unmodified...)
	require.NoError(t, err)
	require.Equal(t, expected, unmodified)
}

func TestMulMDataset(t *testing.T) {
	result, err := dataset.InspectFromFile("testdata/mulm", MulM[int8])
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

	opts := dataset.Opts[int8]{
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

	err := dataset.CollectToFile(opts, "testdata/mulm")
	require.NoError(t, err)
}

func TestMulM4ArgsSig(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth int8) (int8, error) {
		return MulM(first, second, third, fourth)
	}

	reference := func(first, second, third, fourth int64) (int64, error) {
		return first * second * third * fourth, nil
	}

	result, err := inspect.Do4(inspected, reference)
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

func TestMulM4ArgsUns(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth uint8) (uint8, error) {
		return MulM(first, second, third, fourth)
	}

	reference := func(first, second, third, fourth int64) (int64, error) {
		return first * second * third * fourth, nil
	}

	result, err := inspect.Do4(inspected, reference)
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

func TestMulM5ArgsSig(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth, fifth int8) (int8, error) {
		return MulM(first, second, third, fourth, fifth)
	}

	reference := func(first, second, third, fourth, fifth int64) (int64, error) {
		return first * second * third * fourth * fifth, nil
	}

	result, err := inspect.Do5(inspected, reference)
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

func TestMulM5ArgsUns(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth, fifth uint8) (uint8, error) {
		return MulM(first, second, third, fourth, fifth)
	}

	reference := func(first, second, third, fourth, fifth int64) (int64, error) {
		return first * second * third * fourth * fifth, nil
	}

	result, err := inspect.Do5(inspected, reference)
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

	result, err := inspect.Do(opts)
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

func TestDivM1ArgsSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...int8) (int8, error) {
			return DivM(args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
		},
	}

	result, err := inspect.Do(opts)
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
}

func TestDivM2ArgsSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
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

	result, err := inspect.Do(opts)
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

func TestDivM3ArgsSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
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

	result, err := inspect.Do(opts)
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

func TestDivM1ArgsUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 1,

		Inspected: func(args ...uint8) (uint8, error) {
			return DivM(args[0], args[1:]...)
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0], nil
		},
	}

	result, err := inspect.Do(opts)
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
}

func TestDivM2ArgsUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
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

	result, err := inspect.Do(opts)
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

func TestDivM3ArgsUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
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

	result, err := inspect.Do(opts)
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

	opts := dataset.Opts[int8]{
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

	err := dataset.CollectToFile(opts, "testdata/divm")
	require.NoError(t, err)
}

func TestDivM4ArgsSig(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth int8) (int8, error) {
		return DivM(first, second, third, fourth)
	}

	reference := func(first, second, third, fourth int64) (int64, error) {
		if second == 0 || third == 0 || fourth == 0 {
			return 0, ErrDivisionByZero
		}

		return first / second / third / fourth, nil
	}

	result, err := inspect.Do4(inspected, reference)
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

func TestDivM4ArgsUns(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth uint8) (uint8, error) {
		return DivM(first, second, third, fourth)
	}

	reference := func(first, second, third, fourth int64) (int64, error) {
		if second == 0 || third == 0 || fourth == 0 {
			return 0, ErrDivisionByZero
		}

		return first / second / third / fourth, nil
	}

	result, err := inspect.Do4(inspected, reference)
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

func TestDivM5ArgsSig(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth, fifth int8) (int8, error) {
		return DivM(first, second, third, fourth, fifth)
	}

	reference := func(first, second, third, fourth, fifth int64) (int64, error) {
		if second == 0 || third == 0 || fourth == 0 || fifth == 0 {
			return 0, ErrDivisionByZero
		}

		return first / second / third / fourth / fifth, nil
	}

	result, err := inspect.Do5(inspected, reference)
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

func TestDivM5ArgsUns(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	inspected := func(first, second, third, fourth, fifth uint8) (uint8, error) {
		return DivM(first, second, third, fourth, fifth)
	}

	reference := func(first, second, third, fourth, fifth int64) (int64, error) {
		if second == 0 || third == 0 || fourth == 0 || fifth == 0 {
			return 0, ErrDivisionByZero
		}

		return first / second / third / fourth / fifth, nil
	}

	result, err := inspect.Do5(inspected, reference)
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

func TestPow10Manually(t *testing.T) {
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

func TestPow10Diff(t *testing.T) {
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
		Span: func() (int32, int32) {
			return math.MinInt8, math.MaxInt8
		},
	}

	result, err := inspect.Do(opts)
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
