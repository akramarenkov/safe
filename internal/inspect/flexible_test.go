package inspect

import (
	"math"
	"testing"

	"github.com/akramarenkov/safe"

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
	inspected := func(args ...int8) (int8, error) {
		reference := int64(args[0]) + int64(args[1]) + int64(args[2])

		if reference > math.MaxInt8 || reference < math.MinInt8 {
			return 0, safe.ErrOverflow
		}

		return int8(reference), nil
	}

	reference := func(args ...int64) (int64, error) {
		return args[0] + args[1] + args[2], nil
	}

	opts := Opts[int8]{
		LoopsQuantity: 3,

		Inspected: inspected,
		Reference: reference,
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
	inspected := func(args ...uint8) (uint8, error) {
		reference := int64(args[0]) + int64(args[1]) + int64(args[2])

		if reference > math.MaxUint8 || reference < 0 {
			return 0, safe.ErrOverflow
		}

		return uint8(reference), nil
	}

	reference := func(args ...int64) (int64, error) {
		return args[0] + args[1] + args[2], nil
	}

	opts := Opts[uint8]{
		LoopsQuantity: 3,

		Inspected: inspected,
		Reference: reference,
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
	quantity := uint(0)

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

func testArgs[Type EightBits](
	t *testing.T,
	quantity uint,
	expected []Type,
	args ...Type,
) uint {
	for id := range args {
		if args[id] != expected[id] {
			require.Equal(t, expected[id], args[id])
		}
	}

	quantity++

	for id := range expected {
		multiplicity := uint(1 << (8 * (len(expected) - id - 1)))

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

	quantity := uint(0)

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
	require.Equal(t, uint(1<<(8*levels)), quantity)
}

func testDoArgsUint(t *testing.T) {
	const levels = 3

	quantity := uint(0)

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
	require.Equal(t, uint(1<<(8*levels)), quantity)
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
	inspected := func(args ...int8) (int8, error) {
		reference := int64(args[0]) + int64(args[1])

		if reference > math.MaxInt8 || reference < math.MinInt8 {
			return 0, safe.ErrOverflow
		}

		return int8(reference), nil
	}

	errorExpected := func(args ...int8) (int8, error) {
		return args[0] + args[1], nil
	}

	unexpectedError := func(args ...int8) (int8, error) {
		reference := int64(args[0]) + int64(args[1])

		if reference > math.MaxInt8 || reference < math.MinInt8 {
			return 0, safe.ErrOverflow
		}

		return int8(reference), safe.ErrOverflow
	}

	notEqual := func(args ...int8) (int8, error) {
		reference := int64(args[0]) + int64(args[1])

		if reference > math.MaxInt8 || reference < math.MinInt8 {
			return 0, safe.ErrOverflow
		}

		return 0, nil
	}

	reference := func(args ...int64) (int64, error) {
		return args[0] + args[1], nil
	}

	referenceFault := func(args ...int64) (int64, error) {
		return args[0] + args[1], safe.ErrOverflow
	}

	opts := Opts[int8]{
		LoopsQuantity: 2,

		Inspected: errorExpected,
		Reference: reference,
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

	opts.Inspected = inspected
	opts.Reference = referenceFault

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
}

func testDoNegativeConclusionUint(t *testing.T) {
	inspected := func(args ...uint8) (uint8, error) {
		reference := int64(args[0]) + int64(args[1])

		if reference > math.MaxInt8 || reference < math.MinInt8 {
			return 0, safe.ErrOverflow
		}

		return uint8(reference), nil
	}

	errorExpected := func(args ...uint8) (uint8, error) {
		return args[0] + args[1], nil
	}

	unexpectedError := func(args ...uint8) (uint8, error) {
		reference := int64(args[0]) + int64(args[1])

		if reference > math.MaxUint8 || reference < 0 {
			return 0, safe.ErrOverflow
		}

		return uint8(reference), safe.ErrOverflow
	}

	notEqual := func(args ...uint8) (uint8, error) {
		reference := int64(args[0]) + int64(args[1])

		if reference > math.MaxUint8 || reference < 0 {
			return 0, safe.ErrOverflow
		}

		return 0, nil
	}

	reference := func(args ...int64) (int64, error) {
		return args[0] + args[1], nil
	}

	referenceFault := func(args ...int64) (int64, error) {
		return args[0] + args[1], safe.ErrOverflow
	}

	opts := Opts[uint8]{
		LoopsQuantity: 2,

		Inspected: errorExpected,
		Reference: reference,
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

	opts.Inspected = inspected
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

	quantity := uint(0)

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

	quantity := uint(0)

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
	inspected := func(args ...int8) (int8, error) {
		reference := int64(args[0]) + int64(args[1]) + int64(args[2])

		if reference > math.MaxInt8 || reference < math.MinInt8 {
			return 0, safe.ErrOverflow
		}

		return int8(reference), nil
	}

	reference := func(args ...int64) (int64, error) {
		return args[0] + args[1] + args[2], nil
	}

	opts := Opts[int8]{
		LoopsQuantity: 3,

		Inspected: inspected,
		Reference: reference,
	}

	var (
		result Result[int8]
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
