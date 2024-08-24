package inspect

import (
	"math"
	"testing"

	"github.com/akramarenkov/safe"
	"github.com/stretchr/testify/require"
)

func TestIsValid4(t *testing.T) {
	opts := Opts4[int8]{
		Inspected: func(first, second, third, fourth int8) (int8, error) {
			return first + second + third + fourth, nil
		},
		Reference: func(first, second, third, fourth int64) (int64, error) {
			return first + second + third + fourth, nil
		},
	}

	require.NoError(t, opts.IsValid())

	opts = Opts4[int8]{
		Inspected: func(first, second, third, fourth int8) (int8, error) {
			return first + second + third + fourth, nil
		},
	}

	require.Error(t, opts.IsValid())

	opts = Opts4[int8]{
		Reference: func(first, second, third, fourth int64) (int64, error) {
			return first + second + third + fourth, nil
		},
	}

	require.Error(t, opts.IsValid())
}

func TestDo4(t *testing.T) {
	testDo4Int(t)
	testDo4Uint(t)
}

func testDo4Int(t *testing.T) {
	opts := Opts4[int8]{
		Inspected: testInspected4Int,
		Reference: testReference4,
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

func testDo4Uint(t *testing.T) {
	opts := Opts4[uint8]{
		Inspected: testInspected4Uint,
		Reference: testReference4,
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

func TestDo4Error(t *testing.T) {
	opts := Opts4[int8]{}

	_, err := opts.Do()
	require.Error(t, err)
}

func TestDo4NegativeConclusion(t *testing.T) {
	testDo4NegativeConclusionInt(t)
	testDo4NegativeConclusionUint(t)
}

func testDo4NegativeConclusionInt(t *testing.T) {
	errorExpected := func(first, second, third, fourth int8) (int8, error) {
		return first + second + third + fourth, nil
	}

	unexpectedError := func(_, _, _, _ int8) (int8, error) {
		return 0, safe.ErrOverflow
	}

	notEqual := func(first, second, third, fourth int8) (int8, error) {
		reference := int64(first) + int64(second) + int64(third) + int64(fourth)

		if reference > math.MaxInt8 || reference < math.MinInt8 {
			return 0, safe.ErrOverflow
		}

		return 0, nil
	}

	referenceFault := func(_, _, _, _ int64) (int64, error) {
		return 0, safe.ErrOverflow
	}

	opts := Opts4[int8]{
		Inspected: errorExpected,
		Reference: testReference4,
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

	opts.Inspected = testInspected4Int
	opts.Reference = referenceFault

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
}

func testDo4NegativeConclusionUint(t *testing.T) {
	errorExpected := func(first, second, third, fourth uint8) (uint8, error) {
		return first + second + third + fourth, nil
	}

	unexpectedError := func(_, _, _, _ uint8) (uint8, error) {
		return 0, safe.ErrOverflow
	}

	notEqual := func(first, second, third, fourth uint8) (uint8, error) {
		reference := int64(first) + int64(second) + int64(third) + int64(fourth)

		if reference > math.MaxUint8 || reference < 0 {
			return 0, safe.ErrOverflow
		}

		return 0, nil
	}

	referenceFault := func(_, _, _, _ int64) (int64, error) {
		return 0, safe.ErrOverflow
	}

	opts := Opts4[uint8]{
		Inspected: errorExpected,
		Reference: testReference4,
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

	opts.Inspected = testInspected4Uint
	opts.Reference = referenceFault

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
}

func BenchmarkDo4(b *testing.B) {
	opts := Opts4[int8]{
		Inspected: testInspected4Int,
		Reference: testReference4,
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

func testReference4(first, second, third, fourth int64) (int64, error) {
	return first + second + third + fourth, nil
}

func testInspected4Int(first, second, third, fourth int8) (int8, error) {
	reference := int64(first) + int64(second) + int64(third) + int64(fourth)

	if reference > math.MaxInt8 || reference < math.MinInt8 {
		return 0, safe.ErrOverflow
	}

	return int8(reference), nil
}

func testInspected4Uint(first, second, third, fourth uint8) (uint8, error) {
	reference := int64(first) + int64(second) + int64(third) + int64(fourth)

	if reference > math.MaxUint8 || reference < 0 {
		return 0, safe.ErrOverflow
	}

	return uint8(reference), nil
}
