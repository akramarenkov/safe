package safe

import (
	"os"
	"testing"

	"github.com/akramarenkov/safe/internal/env"
	"github.com/akramarenkov/safe/internal/inspect"
	"github.com/akramarenkov/safe/internal/inspect/dataset"

	"github.com/stretchr/testify/require"
)

func TestAddSub(t *testing.T) {
	testAddSubSig(t)
	testAddSubUns(t)
}

func testAddSubSig(t *testing.T) {
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

func testAddSubUns(t *testing.T) {
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

func TestAddDiv(t *testing.T) {
	testAddDivSig(t)
	testAddDivUns(t)
}

func testAddDivSig(t *testing.T) {
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

func testAddDivUns(t *testing.T) {
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

func TestAddDivRem(t *testing.T) {
	testAddDivRemSig(t)
	testAddDivRemUns(t)
}

func testAddDivRemSig(t *testing.T) {
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

func testAddDivRemUns(t *testing.T) {
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

func TestSubDiv(t *testing.T) {
	testSubDivSig(t)
	testSubDivUns(t)
}

func testSubDivSig(t *testing.T) {
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

func testSubDivUns(t *testing.T) {
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

func TestSubDivRem(t *testing.T) {
	testSubDivRemSig(t)
	testSubDivRemUns(t)
}

func testSubDivRemSig(t *testing.T) {
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

func testSubDivRemUns(t *testing.T) {
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

func TestAddSubDivDataset(t *testing.T) {
	testAddSubDivDatasetSig(t)
	testAddSubDivDatasetUns(t)
}

func testAddSubDivDatasetSig(t *testing.T) {
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

func testAddSubDivDatasetUns(t *testing.T) {
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

func TestAddSubDivCollectDataset(t *testing.T) {
	if os.Getenv(env.CollectDataset) == "" {
		t.SkipNow()
	}

	testAddSubDivCollectDatasetSig(t)
	testAddSubDivCollectDatasetUns(t)
}

func testAddSubDivCollectDatasetSig(t *testing.T) {
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

func testAddSubDivCollectDatasetUns(t *testing.T) {
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

func TestAddSubDivFull(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	testAddSubDivFullSig(t)
	testAddSubDivFullUns(t)
}

func testAddSubDivFullSig(t *testing.T) {
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

func testAddSubDivFullUns(t *testing.T) {
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

func TestAddOneSubDiv(t *testing.T) {
	testAddOneSubDivSig(t)
	testAddOneSubDivUns(t)
}

func testAddOneSubDivSig(t *testing.T) {
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

func testAddOneSubDivUns(t *testing.T) {
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
