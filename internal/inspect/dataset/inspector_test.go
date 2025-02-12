package dataset

import (
	"bytes"
	"math"
	"path/filepath"
	"testing"

	"github.com/akramarenkov/safe/internal/inspect"
	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/iterator"

	"github.com/stretchr/testify/require"
)

func TestInspectSig(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	for first := range iterator.Iter[int8](math.MinInt8, math.MaxInt8) {
		for second := range iterator.Iter[int8](math.MinInt8, math.MaxInt8) {
			require.NoError(t, WriteItem(buffer, testReference, first, second))
		}
	}

	result, err := Inspect(buffer, testInspectedSig)
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

func TestInspectUns(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	for first := range iterator.Iter[uint8](0, math.MaxUint8) {
		for second := range iterator.Iter[uint8](0, math.MaxUint8) {
			require.NoError(t, WriteItem(buffer, testReference, first, second))
		}
	}

	result, err := Inspect(buffer, testInspectedUns)
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

func TestInspectError(t *testing.T) {
	inspected := func(...int8) (int8, error) { return 0, nil }

	result, err := Inspect(bytes.NewBuffer(nil), inspected)
	require.NoError(t, err)
	require.Equal(t, types.Result[int8, int8, int64]{}, result)

	result, err = Inspect(nil, inspected)
	require.Error(t, err)
	require.Equal(t, types.Result[int8, int8, int64]{}, result)

	result, err = Inspect[int8](bytes.NewBuffer(nil), nil)
	require.Error(t, err)
	require.Equal(t, types.Result[int8, int8, int64]{}, result)

	result, err = Inspect[int8](nil, nil)
	require.Error(t, err)
	require.Equal(t, types.Result[int8, int8, int64]{}, result)
}

func TestInspectConvertErrorSig(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	buffer.Reset()
	buffer.WriteString("false 2")

	_, err := Inspect(buffer, testInspectedSig)
	require.Error(t, err)

	buffer.Reset()
	buffer.WriteString("flase 2 1 1")

	_, err = Inspect(buffer, testInspectedSig)
	require.Error(t, err)

	buffer.Reset()
	buffer.WriteString("false true 1 1")

	_, err = Inspect(buffer, testInspectedSig)
	require.Error(t, err)

	buffer.Reset()
	buffer.WriteString("false 2 true 1")

	_, err = Inspect(buffer, testInspectedSig)
	require.Error(t, err)
}

func TestInspectConvertErrorUns(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	buffer.Reset()
	buffer.WriteString("false 2 true 1")

	_, err := Inspect(buffer, testInspectedUns)
	require.Error(t, err)
}

func TestInspectNegativeConclusion(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	errorExpected := func(args ...int8) (int8, error) {
		return testInspectedUnsafe(args...), nil
	}

	unexpectedError := func(...int8) (int8, error) {
		return 0, inspect.ErrOverflow
	}

	notEqual := func(args ...int8) (int8, error) {
		reference := testReference8(args...)

		if reference > math.MaxInt8 || reference < math.MinInt8 {
			return 0, inspect.ErrOverflow
		}

		return 0, nil
	}

	referenceFault := func(...int64) (int64, error) {
		return 0, inspect.ErrOverflow
	}

	collect := func(reference types.Reference[int64]) {
		buffer.Reset()

		for first := range iterator.Iter[int8](math.MinInt8, math.MaxInt8) {
			for second := range iterator.Iter[int8](math.MinInt8, math.MaxInt8) {
				require.NoError(t, WriteItem(buffer, reference, first, second))
			}
		}
	}

	collect(testReference)

	result, err := Inspect(buffer, errorExpected)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	collect(testReference)

	result, err = Inspect(buffer, unexpectedError)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	collect(testReference)

	result, err = Inspect(buffer, notEqual)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)

	collect(referenceFault)

	result, err = Inspect(buffer, testInspectedSig)
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotEmpty(t, result.Args)
}

func TestInspectFromFileError(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "dataset")

	_, err := InspectFromFile(filePath, testInspectedSig)
	require.Error(t, err)
}
