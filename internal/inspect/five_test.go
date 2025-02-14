package inspect

import (
	"math"
	"os"
	"testing"

	"github.com/akramarenkov/safe/internal/env"
	"github.com/akramarenkov/safe/internal/inspect/types"

	"github.com/stretchr/testify/require"
)

func TestDo5Sig(t *testing.T) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		t.SkipNow()
	}

	result, err := Do5(testInspected5Sig, testReference5)
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

	result, err := Do5(testInspected5Uns, testReference5)
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
	inspected := func(_, _, _, _, _ int8) (int8, error) {
		return 0, nil
	}

	reference := func(_, _, _, _, _ int64) (int64, error) {
		return 0, nil
	}

	result, err := Do5(inspected, nil)
	require.Error(t, err)
	require.Equal(t, types.Result[int8, int8, int64]{}, result)

	result, err = Do5[int8](nil, reference)
	require.Error(t, err)
	require.Equal(t, types.Result[int8, int8, int64]{}, result)

	result, err = Do5[int8](nil, nil)
	require.Error(t, err)
	require.Equal(t, types.Result[int8, int8, int64]{}, result)
}

func TestDo5NegativeConclusionSig(t *testing.T) {
	errorExpected := func(first, second, third, fourth, fifth int8) (int8, error) {
		return first + second + third + fourth + fifth, nil
	}

	unexpectedError := func(_, _, _, _, _ int8) (int8, error) {
		return 0, errOverflow
	}

	notEqual := func(first, second, third, fourth, fifth int8) (int8, error) {
		reference := int64(first) + int64(second) + int64(third) + int64(fourth) +
			int64(fifth)

		if reference > math.MaxInt8 || reference < math.MinInt8 {
			return 0, errOverflow
		}

		return 0, nil
	}

	referenceFault := func(_, _, _, _, _ int64) (int64, error) {
		return 0, errOverflow
	}

	result, err := Do5(errorExpected, testReference5)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	result, err = Do5(unexpectedError, testReference5)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	result, err = Do5(notEqual, testReference5)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	result, err = Do5(testInspected5Sig, referenceFault)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)
}

func TestDo5NegativeConclusionUns(t *testing.T) {
	errorExpected := func(first, second, third, fourth, fifth uint8) (uint8, error) {
		return first + second + third + fourth + fifth, nil
	}

	unexpectedError := func(_, _, _, _, _ uint8) (uint8, error) {
		return 0, errOverflow
	}

	notEqual := func(first, second, third, fourth, fifth uint8) (uint8, error) {
		reference := int64(first) + int64(second) + int64(third) + int64(fourth) +
			int64(fifth)

		if reference > math.MaxUint8 || reference < 0 {
			return 0, errOverflow
		}

		return 0, nil
	}

	referenceFault := func(_, _, _, _, _ int64) (int64, error) {
		return 0, errOverflow
	}

	result, err := Do5(errorExpected, testReference5)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	result, err = Do5(unexpectedError, testReference5)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	result, err = Do5(notEqual, testReference5)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	result, err = Do5(testInspected5Uns, referenceFault)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)
}

func BenchmarkDo5(b *testing.B) {
	// It is impossible to test in automatic mode in an acceptable time
	if os.Getenv(env.EnableLongTest) == "" {
		b.SkipNow()
	}

	var (
		result types.Result[int8, int8, int64]
		err    error
	)

	for range b.N {
		result, err = Do5(testInspected5Sig, testReference5)
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
		return 0, errOverflow
	}

	return int8(reference), nil
}

func testInspected5Uns(first, second, third, fourth, fifth uint8) (uint8, error) {
	reference := int64(first) + int64(second) + int64(third) + int64(fourth) +
		int64(fifth)

	if reference > math.MaxUint8 || reference < 0 {
		return 0, errOverflow
	}

	return uint8(reference), nil
}
