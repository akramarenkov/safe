package safe

import (
	"testing"

	"github.com/akramarenkov/safe/internal/inspect"
	"github.com/stretchr/testify/require"
)

func TestAddDiv(t *testing.T) {
	testAddDivInt(t)
	testAddDivUint(t)
}

func testAddDivInt(t *testing.T) {
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

func testAddDivUint(t *testing.T) {
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
	testAddDivRemInt(t)
	testAddDivRemUint(t)
}

func testAddDivRemInt(t *testing.T) {
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

func testAddDivRemUint(t *testing.T) {
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
	testSubDivInt(t)
	testSubDivUint(t)
}

func testSubDivInt(t *testing.T) {
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

func testSubDivUint(t *testing.T) {
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
