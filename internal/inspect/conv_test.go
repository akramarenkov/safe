package inspect

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConversion(t *testing.T) {
	inspected := func(number int8) (uint8, error) {
		if number < 0 {
			return 0, ErrOverflow
		}

		return uint8(number), nil
	}

	result, err := Conversion[int8, uint8](inspected)
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

func TestConversionError(t *testing.T) {
	_, err := Conversion[int8, uint8](nil)
	require.Error(t, err)
}

func TestConversionNegativeConclusion(t *testing.T) {
	errorExpected := func(number int8) (uint8, error) {
		return uint8(number), nil
	}

	unexpectedError := func(int8) (uint8, error) {
		return 0, ErrOverflow
	}

	notEqual := func(number int8) (uint8, error) {
		if number < 0 {
			return 0, ErrOverflow
		}

		return 0, nil
	}

	result, err := Conversion[int8, uint8](errorExpected)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)

	result, err = Conversion[int8, uint8](unexpectedError)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)

	result, err = Conversion[int8, uint8](notEqual)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
}
