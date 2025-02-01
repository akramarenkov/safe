package inspect

import (
	"math"
	"os"
	"testing"

	"github.com/akramarenkov/safe/internal/env"
	"github.com/akramarenkov/safe/internal/inspect/types"

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

func TestDo5Sig(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	opts := Opts5[int8]{
		Inspected: testInspected5Sig,
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

func TestDo5Uns(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	opts := Opts5[uint8]{
		Inspected: testInspected5Uns,
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

func TestDo5NegativeConclusionSig(t *testing.T) {
	errorExpected := func(first, second, third, fourth, fifth int8) (int8, error) {
		return first + second + third + fourth + fifth, nil
	}

	unexpectedError := func(_, _, _, _, _ int8) (int8, error) {
		return 0, ErrOverflow
	}

	notEqual := func(first, second, third, fourth, fifth int8) (int8, error) {
		reference := int64(first) + int64(second) + int64(third) + int64(fourth) +
			int64(fifth)

		if reference > math.MaxInt8 || reference < math.MinInt8 {
			return 0, ErrOverflow
		}

		return 0, nil
	}

	referenceFault := func(_, _, _, _, _ int64) (int64, error) {
		return 0, ErrOverflow
	}

	opts := Opts5[int8]{
		Inspected: errorExpected,
		Reference: testReference5,
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

	opts.Inspected = testInspected5Sig
	opts.Reference = referenceFault

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)
}

func TestDo5NegativeConclusionUns(t *testing.T) {
	errorExpected := func(first, second, third, fourth, fifth uint8) (uint8, error) {
		return first + second + third + fourth + fifth, nil
	}

	unexpectedError := func(_, _, _, _, _ uint8) (uint8, error) {
		return 0, ErrOverflow
	}

	notEqual := func(first, second, third, fourth, fifth uint8) (uint8, error) {
		reference := int64(first) + int64(second) + int64(third) + int64(fourth) +
			int64(fifth)

		if reference > math.MaxUint8 || reference < 0 {
			return 0, ErrOverflow
		}

		return 0, nil
	}

	referenceFault := func(_, _, _, _, _ int64) (int64, error) {
		return 0, ErrOverflow
	}

	opts := Opts5[uint8]{
		Inspected: errorExpected,
		Reference: testReference5,
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

	opts.Inspected = testInspected5Uns
	opts.Reference = referenceFault

	result, err = opts.Do()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)
}

func BenchmarkDo5(b *testing.B) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		b.SkipNow()
	}

	opts := Opts5[int8]{
		Inspected: testInspected5Sig,
		Reference: testReference5,
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

func testReference5(first, second, third, fourth, fifth int64) (int64, error) {
	return first + second + third + fourth + fifth, nil
}

func testInspected5Sig(first, second, third, fourth, fifth int8) (int8, error) {
	reference := int64(first) + int64(second) + int64(third) + int64(fourth) +
		int64(fifth)

	if reference > math.MaxInt8 || reference < math.MinInt8 {
		return 0, ErrOverflow
	}

	return int8(reference), nil
}

func testInspected5Uns(first, second, third, fourth, fifth uint8) (uint8, error) {
	reference := int64(first) + int64(second) + int64(third) + int64(fourth) +
		int64(fifth)

	if reference > math.MaxUint8 || reference < 0 {
		return 0, ErrOverflow
	}

	return uint8(reference), nil
}
