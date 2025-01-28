package inspect

import (
	"math"
	"os"
	"testing"

	"github.com/akramarenkov/safe/internal/env"
	"github.com/akramarenkov/safe/internal/inspect/incrementor"
	"github.com/akramarenkov/safe/internal/inspect/types"

	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	opts := Opts[int8, int8, int64]{
		Inspected: func(...int8) (int8, error) { return 0, nil },
		Reference: func(...int64) (int64, error) { return 0, nil },
	}

	require.NoError(t, opts.IsValid())

	opts = Opts[int8, int8, int64]{
		Inspected: func(...int8) (int8, error) { return 0, nil },
	}

	require.Error(t, opts.IsValid())

	opts = Opts[int8, int8, int64]{
		Reference: func(...int64) (int64, error) { return 0, nil },
	}

	require.Error(t, opts.IsValid())
}

func TestDo(t *testing.T) {
	testDoSig(t)
	testDoUns(t)
}

func testDoSig(t *testing.T) {
	opts := Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: testInspected3Sig,
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

func testDoUns(t *testing.T) {
	opts := Opts[uint8, uint8, int64]{
		LoopsQuantity: 3,

		Inspected: testInspected3Uns,
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

func TestDoError(t *testing.T) {
	opts := Opts[int8, int8, int64]{
		LoopsQuantity: 3,
	}

	_, err := opts.Do()
	require.Error(t, err)
}

func TestDoNegativeConclusion(t *testing.T) {
	testDoNegativeConclusionSig(t)
	testDoNegativeConclusionUns(t)
}

func testDoNegativeConclusionSig(t *testing.T) {
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

	opts := Opts[int8, int8, int64]{
		LoopsQuantity: 2,

		Inspected: errorExpected,
		Reference: testReference2,
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	opts.Inspected = unexpectedError

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	opts.Inspected = notEqual

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	opts.Inspected = testInspected2Sig
	opts.Reference = referenceFault

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)
}

func testDoNegativeConclusionUns(t *testing.T) {
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

	opts := Opts[uint8, uint8, int64]{
		LoopsQuantity: 2,

		Inspected: errorExpected,
		Reference: testReference2,
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	opts.Inspected = unexpectedError

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	opts.Inspected = notEqual

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	opts.Inspected = testInspected2Uns
	opts.Reference = referenceFault

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)
}

func TestLoop(t *testing.T) {
	testLoopSig(t)
	testLoopUns(t)
	testLoopSpanSig(t)
	testLoopSpanUns(t)
	testLoopFloatU16(t)
}

func testLoopSig(t *testing.T) {
	const levels = 3

	incrementor, err := incrementor.New[int8](levels, math.MinInt8, math.MaxInt8)
	require.NoError(t, err)

	do := func(args ...int8) bool {
		// duplication of conditions is done for performance reasons
		if err := incrementor.Test(args...); err != nil {
			require.NoError(t, err)
		}

		return false
	}

	stop := loop[int64](levels, nil, do)
	require.False(t, stop)
}

func testLoopUns(t *testing.T) {
	const levels = 3

	incrementor, err := incrementor.New[uint8](levels, 0, math.MaxUint8)
	require.NoError(t, err)

	do := func(args ...uint8) bool {
		// duplication of conditions is done for performance reasons
		if err := incrementor.Test(args...); err != nil {
			require.NoError(t, err)
		}

		return false
	}

	stop := loop[int64](levels, nil, do)
	require.False(t, stop)
}

func testLoopSpanSig(t *testing.T) {
	const (
		levels = 3
		begin  = -1
		end    = 1
	)

	incrementor, err := incrementor.New[int8](levels, begin, end)
	require.NoError(t, err)

	do := func(args ...int8) bool {
		// duplication of conditions is done for performance reasons
		if err := incrementor.Test(args...); err != nil {
			require.NoError(t, err)
		}

		return false
	}

	span := func() (int8, int8) {
		return begin, end
	}

	stop := loop[int64](levels, span, do)
	require.False(t, stop)
}

func testLoopSpanUns(t *testing.T) {
	const (
		levels = 3
		begin  = 1
		end    = 3
	)

	incrementor, err := incrementor.New[uint8](levels, begin, end)
	require.NoError(t, err)

	do := func(args ...uint8) bool {
		// duplication of conditions is done for performance reasons
		if err := incrementor.Test(args...); err != nil {
			require.NoError(t, err)
		}

		return false
	}

	span := func() (uint8, uint8) {
		return begin, end
	}

	stop := loop[int64](levels, span, do)
	require.False(t, stop)
}

func testLoopFloatU16(t *testing.T) {
	const levels = 1

	incrementor, err := incrementor.New[uint16](levels, 0, math.MaxUint16)
	require.NoError(t, err)

	do := func(args ...uint16) bool {
		// duplication of conditions is done for performance reasons
		if err := incrementor.Test(args...); err != nil {
			require.NoError(t, err)
		}

		return false
	}

	stop := loop[float64](levels, nil, do)
	require.False(t, stop)
}

func TestLoopFloatU32(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	const levels = 1

	incrementor, err := incrementor.New[uint32](levels, 0, math.MaxUint32)
	require.NoError(t, err)

	do := func(args ...uint32) bool {
		// duplication of conditions is done for performance reasons
		if err := incrementor.Test(args...); err != nil {
			require.NoError(t, err)
		}

		return false
	}

	stop := loop[float64](levels, nil, do)
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

	stop := loop[int64](0, nil, do, 1)
	require.False(t, stop)

	stop = loop[int64](0, nil, doU, 1)
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

	stop := loop[int64](2, nil, do)
	require.True(t, stop)
	require.Equal(t, expected, actual)

	stop = loop[int64](2, nil, doU)
	require.True(t, stop)
	require.Equal(t, expectedU, actualU)
}

func BenchmarkDo(b *testing.B) {
	opts := Opts[int8, int8, int64]{
		LoopsQuantity: 3,

		Inspected: testInspected3Sig,
		Reference: testReference3,
	}

	var (
		result types.Result[int8, int8, int64]
		err    error
	)

	for range b.N {
		result, err = opts.Do()
	}

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
		_ = loop[int64](3, nil, do)
	}

	require.NotZero(b, quantity)
}

func BenchmarkLoopFixed(b *testing.B) {
	quantity := 0

	do := func(args ...int) bool {
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
					if stop := do(first, second, third); stop {
						return
					}
				}
			}
		}
	}

	require.NotZero(b, quantity)
}

func BenchmarkLoopFixedArgsFixed(b *testing.B) {
	quantity := 0

	do := func(first, second, third int) bool {
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
					if stop := do(first, second, third); stop {
						return
					}
				}
			}
		}
	}

	require.NotZero(b, quantity)
}

func testReference2(args ...int64) (int64, error) {
	return args[0] + args[1], nil
}

func testReference3(args ...int64) (int64, error) {
	return args[0] + args[1] + args[2], nil
}

func testInspected2Sig(args ...int8) (int8, error) {
	reference := int64(args[0]) + int64(args[1])

	if reference > math.MaxInt8 || reference < math.MinInt8 {
		return 0, ErrOverflow
	}

	return int8(reference), nil
}

func testInspected2Uns(args ...uint8) (uint8, error) {
	reference := int64(args[0]) + int64(args[1])

	if reference > math.MaxUint8 || reference < 0 {
		return 0, ErrOverflow
	}

	return uint8(reference), nil
}

func testInspected3Sig(args ...int8) (int8, error) {
	reference := int64(args[0]) + int64(args[1]) + int64(args[2])

	if reference > math.MaxInt8 || reference < math.MinInt8 {
		return 0, ErrOverflow
	}

	return int8(reference), nil
}

func testInspected3Uns(args ...uint8) (uint8, error) {
	reference := int64(args[0]) + int64(args[1]) + int64(args[2])

	if reference > math.MaxUint8 || reference < 0 {
		return 0, ErrOverflow
	}

	return uint8(reference), nil
}
