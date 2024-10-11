package safe

import (
	"testing"

	"github.com/akramarenkov/safe/internal/inspect"
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
