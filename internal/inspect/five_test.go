package inspect

import (
	"math"
	"os"
	"testing"

	"github.com/akramarenkov/safe"
	"github.com/akramarenkov/safe/internal/consts"
	"github.com/stretchr/testify/require"
)

func TestIsValid5(t *testing.T) {
	opts := Opts5[int8]{
		Inspected: func(first, second, third, fourth, fifth int8) (int8, error) {
			return first + second + third + fourth + fifth, nil
		},
		Reference: func(first, second, third, fourth, fifth int64) (int64, error) {
			return first + second + third + fourth + fifth, nil
		},
	}

	require.NoError(t, opts.IsValid())

	opts = Opts5[int8]{
		Inspected: func(first, second, third, fourth, fifth int8) (int8, error) {
			return first + second + third + fourth + fifth, nil
		},
	}

	require.Error(t, opts.IsValid())

	opts = Opts5[int8]{
		Reference: func(first, second, third, fourth, fifth int64) (int64, error) {
			return first + second + third + fourth + fifth, nil
		},
	}

	require.Error(t, opts.IsValid())
}

func TestDo5(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(consts.EnvEnableLongTest) == "" {
		t.SkipNow()
	}

	testDo5Int(t)
	testDo5Uint(t)
}

func testDo5Int(t *testing.T) {
	opts := Opts5[int8]{
		Inspected: testInspected5Int,
		Reference: testReference5,
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

func testDo5Uint(t *testing.T) {
	opts := Opts5[uint8]{
		Inspected: testInspected5Uint,
		Reference: testReference5,
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

func TestDo5Error(t *testing.T) {
	opts := Opts5[int8]{}

	_, err := opts.Do()
	require.Error(t, err)
}

func TestDo5NegativeConclusion(t *testing.T) {
	testDo5NegativeConclusionInt(t)
	testDo5NegativeConclusionUint(t)
}

func testDo5NegativeConclusionInt(t *testing.T) {
	errorExpected := func(first, second, third, fourth, fifth int8) (int8, error) {
		return first + second + third + fourth + fifth, nil
	}

	unexpectedError := func(_, _, _, _, _ int8) (int8, error) {
		return 0, safe.ErrOverflow
	}

	notEqual := func(first, second, third, fourth, fifth int8) (int8, error) {
		reference := int64(first) + int64(second) + int64(third) + int64(fourth) +
			int64(fifth)

		if reference > math.MaxInt8 || reference < math.MinInt8 {
			return 0, safe.ErrOverflow
		}

		return 0, nil
	}

	referenceFault := func(_, _, _, _, _ int64) (int64, error) {
		return 0, safe.ErrOverflow
	}

	opts := Opts5[int8]{
		Inspected: errorExpected,
		Reference: testReference5,
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

	opts.Inspected = testInspected5Int
	opts.Reference = referenceFault

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
}

func testDo5NegativeConclusionUint(t *testing.T) {
	errorExpected := func(first, second, third, fourth, fifth uint8) (uint8, error) {
		return first + second + third + fourth + fifth, nil
	}

	unexpectedError := func(_, _, _, _, _ uint8) (uint8, error) {
		return 0, safe.ErrOverflow
	}

	notEqual := func(first, second, third, fourth, fifth uint8) (uint8, error) {
		reference := int64(first) + int64(second) + int64(third) + int64(fourth) +
			int64(fifth)

		if reference > math.MaxUint8 || reference < 0 {
			return 0, safe.ErrOverflow
		}

		return 0, nil
	}

	referenceFault := func(_, _, _, _, _ int64) (int64, error) {
		return 0, safe.ErrOverflow
	}

	opts := Opts5[uint8]{
		Inspected: errorExpected,
		Reference: testReference5,
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

	opts.Inspected = testInspected5Uint
	opts.Reference = referenceFault

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
}

func BenchmarkDo5(b *testing.B) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(consts.EnvEnableLongTest) == "" {
		b.SkipNow()
	}

	opts := Opts5[int8]{
		Inspected: testInspected5Int,
		Reference: testReference5,
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

func testReference5(first, second, third, fourth, fifth int64) (int64, error) {
	return first + second + third + fourth + fifth, nil
}

func testInspected5Int(first, second, third, fourth, fifth int8) (int8, error) {
	reference := int64(first) + int64(second) + int64(third) + int64(fourth) +
		int64(fifth)

	if reference > math.MaxInt8 || reference < math.MinInt8 {
		return 0, safe.ErrOverflow
	}

	return int8(reference), nil
}

func testInspected5Uint(first, second, third, fourth, fifth uint8) (uint8, error) {
	reference := int64(first) + int64(second) + int64(third) + int64(fourth) +
		int64(fifth)

	if reference > math.MaxUint8 || reference < 0 {
		return 0, safe.ErrOverflow
	}

	return uint8(reference), nil
}
