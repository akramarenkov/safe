package safe

import (
	"os"
	"testing"

	"github.com/akramarenkov/safe/internal/env"
	"github.com/akramarenkov/safe/internal/inspect"
	"github.com/akramarenkov/safe/internal/inspect/dataset"

	"github.com/stretchr/testify/require"
)

func TestAddSubSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return AddSub(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] - args[2], nil
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

func TestAddSubUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return AddSub(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] - args[2], nil
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

func TestAddDivSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return AddDiv(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return (args[0] + args[1]) / args[2], nil
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

func TestAddDivUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return AddDiv(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return (args[0] + args[1]) / args[2], nil
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

func TestAddDivRemSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return AddDivRem(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return (args[0] + args[1]) % args[2], nil
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

func TestAddDivRemUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return AddDivRem(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return (args[0] + args[1]) % args[2], nil
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

func TestAddDivU(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return AddDivU(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return (args[0] + args[1]) / args[2], nil
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

func TestSubDivSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return SubDiv(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return (args[0] - args[1]) / args[2], nil
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

func TestSubDivUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return SubDiv(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return (args[0] - args[1]) / args[2], nil
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

func TestSubDivRemSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return SubDivRem(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return (args[0] - args[1]) % args[2], nil
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

func TestSubDivRemUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return SubDivRem(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return (args[0] - args[1]) % args[2], nil
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

func TestSubDivU(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return SubDivU(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return (args[0] - args[1]) / args[2], nil
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

func TestAddSubDivDatasetSig(t *testing.T) {
	inspected := func(args ...int8) (int8, error) {
		return AddSubDiv(args[0], args[1], args[2], args[3])
	}

	result, err := dataset.InspectFromFile("testdata/addsubdiv/signed", inspected)
	require.NoError(t, err)
	require.NoError(
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

func TestAddSubDivDatasetUns(t *testing.T) {
	inspected := func(args ...uint8) (uint8, error) {
		return AddSubDiv(args[0], args[1], args[2], args[3])
	}

	result, err := dataset.InspectFromFile("testdata/addsubdiv/unsigned", inspected)
	require.NoError(t, err)
	require.NoError(
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

func TestAddSubDivCollectDatasetSig(t *testing.T) {
	if os.Getenv(env.CollectDataset) == "" {
		t.SkipNow()
	}

	reference := func(args ...int64) (int64, error) {
		if args[3] == 0 {
			return 0, ErrDivisionByZero
		}

		return (args[0] + args[1] - args[2]) / args[3], nil
	}

	collector := dataset.Collector[int8]{
		ArgsQuantity:               4,
		NotOverflowedItemsQuantity: 1 << 14,
		OverflowedItemsQuantity:    1 << 14,
		Reference:                  reference,
	}

	err := collector.CollectToFile("testdata/addsubdiv/signed")
	require.NoError(t, err)
}

func TestAddSubDivCollectDatasetUns(t *testing.T) {
	if os.Getenv(env.CollectDataset) == "" {
		t.SkipNow()
	}

	reference := func(args ...int64) (int64, error) {
		if args[3] == 0 {
			return 0, ErrDivisionByZero
		}

		return (args[0] + args[1] - args[2]) / args[3], nil
	}

	collector := dataset.Collector[uint8]{
		ArgsQuantity:               4,
		NotOverflowedItemsQuantity: 1 << 14,
		OverflowedItemsQuantity:    1 << 14,
		Reference:                  reference,
	}

	err := collector.CollectToFile("testdata/addsubdiv/unsigned")
	require.NoError(t, err)
}

func TestAddSubDivFullSig(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	opts := inspect.Opts4[int8]{
		Inspected: AddSubDiv[int8],
		Reference: func(first, second, third, fourth int64) (int64, error) {
			if fourth == 0 {
				return 0, ErrDivisionByZero
			}

			return (first + second - third) / fourth, nil
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

func TestAddSubDivFullUns(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	opts := inspect.Opts4[uint8]{
		Inspected: AddSubDiv[uint8],
		Reference: func(first, second, third, fourth int64) (int64, error) {
			if fourth == 0 {
				return 0, ErrDivisionByZero
			}

			return (first + second - third) / fourth, nil
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

func TestAddOneSubDivSig(t *testing.T) {
	opts := inspect.Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return AddOneSubDiv(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return (args[0] + 1 - args[1]) / args[2], nil
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

func TestAddOneSubDivUns(t *testing.T) {
	opts := inspect.Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return AddOneSubDiv(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return (args[0] + 1 - args[1]) / args[2], nil
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
