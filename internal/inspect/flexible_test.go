package inspect

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	opts := Opts[int8]{
		Inspected: func(...int8) (int8, error) { return 0, nil },
		Reference: func(...int64) (int64, error) { return 0, nil },
	}

	require.NoError(t, opts.IsValid())

	opts = Opts[int8]{
		Inspected: func(...int8) (int8, error) { return 0, nil },
	}

	require.Error(t, opts.IsValid())

	opts = Opts[int8]{
		Reference: func(...int64) (int64, error) { return 0, nil },
	}

	require.Error(t, opts.IsValid())
}

func TestDo(t *testing.T) {
	t.Run(
		"int",
		func(t *testing.T) {
			t.Parallel()
			testDoInt(t)
		},
	)

	t.Run(
		"uint",
		func(t *testing.T) {
			t.Parallel()
			testDoUint(t)
		},
	)
}

func testDoInt(t *testing.T) {
	opts := Opts[int8]{
		LoopsQuantity: 3,

		Inspected: testInspected3Int,
		Reference: testReference3,
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

func testDoUint(t *testing.T) {
	opts := Opts[uint8]{
		LoopsQuantity: 3,

		Inspected: testInspected3Uint,
		Reference: testReference3,
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

func TestArgs(t *testing.T) {
	quantity := uint64(0)

	expected := make([]int8, 3)

	for id := range expected {
		expected[id] = math.MinInt8
	}

	for first := math.MinInt8; first <= math.MaxInt8; first++ {
		for second := math.MinInt8; second <= math.MaxInt8; second++ {
			for third := math.MinInt8; third <= math.MaxInt8; third++ {
				quantity = testArgs(
					t,
					quantity,
					expected,
					int8(first),
					int8(second),
					int8(third),
				)
			}
		}
	}
}

// checks that the arguments change as if they were incremented by nested loops.
func testArgs[Type EightBits](
	t *testing.T,
	quantity uint64,
	expected []Type,
	args ...Type,
) uint64 {
	// uint64 contains eight uint8
	const maxLoopsLevels = 8

	// duplication of conditions is done for performance reasons
	if len(expected) > maxLoopsLevels {
		require.LessOrEqual(t, len(expected), maxLoopsLevels)
	}

	// duplication of conditions is done for performance reasons
	if len(args) > maxLoopsLevels {
		require.LessOrEqual(t, len(args), maxLoopsLevels)
	}

	// duplication of conditions is done for performance reasons
	if len(args) != len(expected) {
		require.Equal(t, len(args), len(expected))
	}

	for id := range args {
		// duplication of conditions is done for performance reasons
		if args[id] != expected[id] {
			require.Equal(t, expected[id], args[id])
		}
	}

	quantity++

	for id := range expected {
		// highest byte corresponds to the first argument, lowest byte
		// corresponds to the last argument
		multiplicity := uint64(1 << (8 * (len(expected) - id - 1)))

		// if the condition is met, it means that all bits of all bytes lower
		// than the current one are set to 0 i.e. there was an overflow into the
		// current byte and the corresponding argument should have increased
		if quantity%multiplicity == 0 {
			expected[id]++
		}
	}

	return quantity
}

func TestDoArgs(t *testing.T) {
	t.Run(
		"int",
		func(t *testing.T) {
			t.Parallel()
			testDoArgsInt(t)
		},
	)

	t.Run(
		"uint",
		func(t *testing.T) {
			t.Parallel()
			testDoArgsUint(t)
		},
	)
}

func testDoArgsInt(t *testing.T) {
	const levels = 3

	quantity := uint64(0)

	expected := make([]int8, levels)

	for id := range expected {
		expected[id] = math.MinInt8
	}

	inspected := func(args ...int8) (int8, error) {
		quantity = testArgs(t, quantity, expected, args...)
		return 0, nil
	}

	reference := func(...int64) (int64, error) {
		return 0, nil
	}

	opts := Opts[int8]{
		LoopsQuantity: levels,

		Inspected: inspected,
		Reference: reference,
	}

	_, err := opts.Do()
	require.NoError(t, err)
	require.Equal(t, uint64(1<<(8*levels)), quantity)
}

func testDoArgsUint(t *testing.T) {
	const levels = 3

	quantity := uint64(0)

	expected := make([]uint8, levels)

	for id := range expected {
		expected[id] = 0
	}

	inspected := func(args ...uint8) (uint8, error) {
		quantity = testArgs(t, quantity, expected, args...)
		return 0, nil
	}

	reference := func(...int64) (int64, error) {
		return 0, nil
	}

	opts := Opts[uint8]{
		LoopsQuantity: levels,

		Inspected: inspected,
		Reference: reference,
	}

	_, err := opts.Do()
	require.NoError(t, err)
	require.Equal(t, uint64(1<<(8*levels)), quantity)
}

func TestDoError(t *testing.T) {
	opts := Opts[int8]{
		LoopsQuantity: 2,
	}

	_, err := opts.Do()
	require.Error(t, err)
}

func TestDoNegativeConclusion(t *testing.T) {
	testDoNegativeConclusionInt(t)
	testDoNegativeConclusionUint(t)
}

func testDoNegativeConclusionInt(t *testing.T) {
	errorExpected := func(args ...int8) (int8, error) {
		return args[0] + args[1], nil
	}

	unexpectedError := func(...int8) (int8, error) {
		return 0, ErrOverflow
	}

	notEqual := func(args ...int8) (int8, error) {
		reference := int64(args[0]) + int64(args[1])

		if reference > math.MaxInt8 || reference < math.MinInt8 {
			return 0, ErrOverflow
		}

		return 0, nil
	}

	referenceFault := func(...int64) (int64, error) {
		return 0, ErrOverflow
	}

	opts := Opts[int8]{
		LoopsQuantity: 2,

		Inspected: errorExpected,
		Reference: testReference2,
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)

	opts.Inspected = unexpectedError

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)

	opts.Inspected = notEqual

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)

	opts.Inspected = testInspected2Int
	opts.Reference = referenceFault

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
}

func testDoNegativeConclusionUint(t *testing.T) {
	errorExpected := func(args ...uint8) (uint8, error) {
		return args[0] + args[1], nil
	}

	unexpectedError := func(...uint8) (uint8, error) {
		return 0, ErrOverflow
	}

	notEqual := func(args ...uint8) (uint8, error) {
		reference := int64(args[0]) + int64(args[1])

		if reference > math.MaxUint8 || reference < 0 {
			return 0, ErrOverflow
		}

		return 0, nil
	}

	referenceFault := func(...int64) (int64, error) {
		return 0, ErrOverflow
	}

	opts := Opts[uint8]{
		LoopsQuantity: 2,

		Inspected: errorExpected,
		Reference: testReference2,
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)

	opts.Inspected = unexpectedError

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)

	opts.Inspected = notEqual

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)

	opts.Inspected = testInspected2Uint
	opts.Reference = referenceFault

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
}

func TestLoop(t *testing.T) {
	t.Run(
		"int",
		func(t *testing.T) {
			t.Parallel()
			testLoopInt(t)
		},
	)

	t.Run(
		"uint",
		func(t *testing.T) {
			t.Parallel()
			testLoopUint(t)
		},
	)
}

func testLoopInt(t *testing.T) {
	const levels = 3

	quantity := uint64(0)

	expected := make([]int8, levels)

	for id := range expected {
		expected[id] = math.MinInt8
	}

	do := func(args ...int8) bool {
		quantity = testArgs(t, quantity, expected, args...)
		return false
	}

	stop := loop(levels, do)
	require.False(t, stop)
}

func testLoopUint(t *testing.T) {
	const levels = 3

	quantity := uint64(0)

	expected := make([]uint8, levels)

	for id := range expected {
		expected[id] = 0
	}

	do := func(args ...uint8) bool {
		quantity = testArgs(t, quantity, expected, args...)
		return false
	}

	stop := loop(levels, do)
	require.False(t, stop)
}

func TestLoopZero(t *testing.T) {
	do := func(args ...int8) bool {
		require.Equal(t, []int8{1}, args)
		return false
	}

	doU := func(args ...uint8) bool {
		require.Equal(t, []uint8{1}, args)
		return false
	}

	stop := loop(0, do, 1)
	require.False(t, stop)

	stop = loop(0, doU, 1)
	require.False(t, stop)
}

func TestLoopStop(t *testing.T) {
	expected := []int8{-128, -128, -128, -127}
	expectedU := []uint8{0, 0, 0, 1}

	actual := make([]int8, 0, 1)
	actualU := make([]uint8, 0, 1)

	do := func(args ...int8) bool {
		actual = append(actual, args...)
		return args[1] == -127
	}

	doU := func(args ...uint8) bool {
		actualU = append(actualU, args...)
		return args[1] == 1
	}

	stop := loop(2, do)
	require.True(t, stop)
	require.Equal(t, expected, actual)

	stop = loop(2, doU)
	require.True(t, stop)
	require.Equal(t, expectedU, actualU)
}

func BenchmarkDo(b *testing.B) {
	opts := Opts[int8]{
		LoopsQuantity: 3,

		Inspected: testInspected3Int,
		Reference: testReference3,
	}

	var (
		result Result[int8, int8]
		err    error
	)

	for range b.N {
		result, err = opts.Do()
	}

	b.StopTimer()

	require.NoError(b, err)
	require.NoError(b, result.Conclusion)
}

func BenchmarkLoop(b *testing.B) {
	quantity := 0

	do := func(args ...int8) bool {
		sum := args[0] + args[1] + args[2]

		if sum == 0 {
			quantity++
		}

		return false
	}

	for range b.N {
		_ = loop(3, do)
	}

	b.StopTimer()

	require.NotZero(b, quantity)
}

func BenchmarkLoopFixed(b *testing.B) {
	quantity := 0

	do := func(args ...int8) bool {
		sum := args[0] + args[1] + args[2]

		if sum == 0 {
			quantity++
		}

		return false
	}

	for range b.N {
		for first := math.MinInt8; first <= math.MaxInt8; first++ {
			for second := math.MinInt8; second <= math.MaxInt8; second++ {
				for third := math.MinInt8; third <= math.MaxInt8; third++ {
					if stop := do(int8(first), int8(second), int8(third)); stop {
						return
					}
				}
			}
		}
	}

	b.StopTimer()

	require.NotZero(b, quantity)
}

func BenchmarkLoopFixedArgsFixed(b *testing.B) {
	quantity := 0

	do := func(first, second, third int8) bool {
		sum := first + second + third

		if sum == 0 {
			quantity++
		}

		return false
	}

	for range b.N {
		for first := math.MinInt8; first <= math.MaxInt8; first++ {
			for second := math.MinInt8; second <= math.MaxInt8; second++ {
				for third := math.MinInt8; third <= math.MaxInt8; third++ {
					if stop := do(int8(first), int8(second), int8(third)); stop {
						return
					}
				}
			}
		}
	}

	b.StopTimer()

	require.NotZero(b, quantity)
}

func testReference2(args ...int64) (int64, error) {
	return args[0] + args[1], nil
}

func testReference3(args ...int64) (int64, error) {
	return args[0] + args[1] + args[2], nil
}

func testInspected2Int(args ...int8) (int8, error) {
	reference := int64(args[0]) + int64(args[1])

	if reference > math.MaxInt8 || reference < math.MinInt8 {
		return 0, ErrOverflow
	}

	return int8(reference), nil
}

func testInspected2Uint(args ...uint8) (uint8, error) {
	reference := int64(args[0]) + int64(args[1])

	if reference > math.MaxUint8 || reference < 0 {
		return 0, ErrOverflow
	}

	return uint8(reference), nil
}

func testInspected3Int(args ...int8) (int8, error) {
	reference := int64(args[0]) + int64(args[1]) + int64(args[2])

	if reference > math.MaxInt8 || reference < math.MinInt8 {
		return 0, ErrOverflow
	}

	return int8(reference), nil
}

func testInspected3Uint(args ...uint8) (uint8, error) {
	reference := int64(args[0]) + int64(args[1]) + int64(args[2])

	if reference > math.MaxUint8 || reference < 0 {
		return 0, ErrOverflow
	}

	return uint8(reference), nil
}
