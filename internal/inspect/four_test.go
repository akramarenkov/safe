package inspect

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDo4Sig(t *testing.T) {
	result, err := Do4(testInspected4Sig, testReference4)
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

func TestDo4Uns(t *testing.T) {
	result, err := Do4(testInspected4Uns, testReference4)
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
	inspected := func(_, _, _, _ int8) (int8, error) {
		return 0, nil
	}

	reference := func(_, _, _, _ int64) (int64, error) {
		return 0, nil
	}

	result, err := Do4(inspected, nil)
	require.Error(t, err)
	require.Equal(t, Result[int8, int8, int64]{}, result)

	result, err = Do4[int8](nil, reference)
	require.Error(t, err)
	require.Equal(t, Result[int8, int8, int64]{}, result)

	result, err = Do4[int8](nil, nil)
	require.Error(t, err)
	require.Equal(t, Result[int8, int8, int64]{}, result)
}

func TestDo4NegativeConclusionSig(t *testing.T) {
	errorExpected := func(first, second, third, fourth int8) (int8, error) {
		return first + second + third + fourth, nil
	}

	unexpectedError := func(_, _, _, _ int8) (int8, error) {
		return 0, errOverflow
	}

	notEqual := func(first, second, third, fourth int8) (int8, error) {
		reference := int64(first) + int64(second) + int64(third) + int64(fourth)

		if reference > math.MaxInt8 || reference < math.MinInt8 {
			return 0, errOverflow
		}

		return 0, nil
	}

	referenceFault := func(_, _, _, _ int64) (int64, error) {
		return 0, errOverflow
	}

	result, err := Do4(errorExpected, testReference4)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	result, err = Do4(unexpectedError, testReference4)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	result, err = Do4(notEqual, testReference4)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	result, err = Do4(testInspected4Sig, referenceFault)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)
}

func TestDo4NegativeConclusionUns(t *testing.T) {
	errorExpected := func(first, second, third, fourth uint8) (uint8, error) {
		return first + second + third + fourth, nil
	}

	unexpectedError := func(_, _, _, _ uint8) (uint8, error) {
		return 0, errOverflow
	}

	notEqual := func(first, second, third, fourth uint8) (uint8, error) {
		reference := int64(first) + int64(second) + int64(third) + int64(fourth)

		if reference > math.MaxUint8 || reference < 0 {
			return 0, errOverflow
		}

		return 0, nil
	}

	referenceFault := func(_, _, _, _ int64) (int64, error) {
		return 0, errOverflow
	}

	result, err := Do4(errorExpected, testReference4)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	result, err = Do4(unexpectedError, testReference4)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	result, err = Do4(notEqual, testReference4)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	result, err = Do4(testInspected4Uns, referenceFault)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)
}

func BenchmarkDo4(b *testing.B) {
	var (
		result Result[int8, int8, int64]
		err    error
	)

	for range b.N {
		result, err = Do4(testInspected4Sig, testReference4)
	}

	require.NoError(b, err)
	require.NoError(b, result.Conclusion)
}

func testReference4(first, second, third, fourth int64) (int64, error) {
	return first + second + third + fourth, nil
}

func testInspected4Sig(first, second, third, fourth int8) (int8, error) {
	reference := int64(first) + int64(second) + int64(third) + int64(fourth)

	if reference > math.MaxInt8 || reference < math.MinInt8 {
		return 0, errOverflow
	}

	return int8(reference), nil
}

func testInspected4Uns(first, second, third, fourth uint8) (uint8, error) {
	reference := int64(first) + int64(second) + int64(third) + int64(fourth)

	if reference > math.MaxUint8 || reference < 0 {
		return 0, errOverflow
	}

	return uint8(reference), nil
}
