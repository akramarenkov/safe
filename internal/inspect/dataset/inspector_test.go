package dataset

import (
	"bytes"
	"math"
	"path/filepath"
	"testing"

	"github.com/akramarenkov/safe/internal/inspect"

	"github.com/stretchr/testify/require"
)

func TestInspectorIsValid(t *testing.T) {
	inspector := Inspector[int8]{
		Inspected: func(...int8) (int8, error) { return 0, nil },
		Reader:    bytes.NewBuffer(nil),
	}

	require.NoError(t, inspector.IsValid())

	inspector = Inspector[int8]{
		Inspected: func(...int8) (int8, error) { return 0, nil },
	}

	require.Error(t, inspector.IsValid())

	inspector = Inspector[int8]{
		Reader: bytes.NewBuffer(nil),
	}

	require.Error(t, inspector.IsValid())
}

func TestInspector(t *testing.T) {
	t.Run(
		"int",
		func(t *testing.T) {
			t.Parallel()
			TestInspectorInt(t)
		},
	)

	t.Run(
		"uint",
		func(t *testing.T) {
			t.Parallel()
			TestInspectorUint(t)
		},
	)
}

func TestInspectorInt(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	for first := math.MinInt8; first <= math.MaxInt8; first++ {
		for second := math.MinInt8; second <= math.MaxInt8; second++ {
			require.NoError(t, WriteItem(buffer, testReference, int8(first), int8(second)))
		}
	}

	inspector := Inspector[int8]{
		Inspected: testInspectedInt,
		Reader:    buffer,
	}

	result, err := inspector.Inspect()
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

func TestInspectorUint(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	for first := 0; first <= math.MaxUint8; first++ {
		for second := 0; second <= math.MaxUint8; second++ {
			require.NoError(t, WriteItem(buffer, testReference, uint8(first), uint8(second)))
		}
	}

	inspector := Inspector[uint8]{
		Inspected: testInspectedUint,
		Reader:    buffer,
	}

	result, err := inspector.Inspect()
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

func TestInspectorError(t *testing.T) {
	inspector := Inspector[int8]{}

	_, err := inspector.Inspect()
	require.Error(t, err)

	TestInspectorConvertErrorInt(t)
	TestInspectorConvertErrorUint(t)
}

func TestInspectorConvertErrorInt(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	inspector := Inspector[int8]{
		Inspected: testInspectedInt,
		Reader:    buffer,
	}

	buffer.Reset()
	buffer.Write([]byte("false 2"))

	_, err := inspector.Inspect()
	require.Error(t, err)

	buffer.Reset()
	buffer.Write([]byte("flase 2 1 1"))

	_, err = inspector.Inspect()
	require.Error(t, err)

	buffer.Reset()
	buffer.Write([]byte("false true 1 1"))

	_, err = inspector.Inspect()
	require.Error(t, err)

	buffer.Reset()
	buffer.Write([]byte("false 2 true 1"))

	_, err = inspector.Inspect()
	require.Error(t, err)
}

func TestInspectorConvertErrorUint(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	inspector := Inspector[uint8]{
		Inspected: testInspectedUint,
		Reader:    buffer,
	}

	buffer.Reset()
	buffer.Write([]byte("false 2 true 1"))

	_, err := inspector.Inspect()
	require.Error(t, err)
}

func TestInspectorNegativeConclusion(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	errorExpected := func(args ...int8) (int8, error) {
		return int8(testReference8(args...)), nil
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

	collect := func(reference Reference) {
		buffer.Reset()

		for first := math.MinInt8; first <= math.MaxInt8; first++ {
			for second := math.MinInt8; second <= math.MaxInt8; second++ {
				require.NoError(t, WriteItem(buffer, reference, int8(first), int8(second)))
			}
		}
	}

	collect(testReference)

	inspector := Inspector[int8]{
		Inspected: errorExpected,
		Reader:    buffer,
	}

	result, err := inspector.Inspect()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotZero(t, len(result.Args))

	collect(testReference)

	inspector.Inspected = unexpectedError

	result, err = inspector.Inspect()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotZero(t, len(result.Args))

	collect(testReference)

	inspector.Inspected = notEqual

	result, err = inspector.Inspect()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotZero(t, len(result.Args))

	collect(referenceFault)

	inspector.Inspected = testInspectedInt

	result, err = inspector.Inspect()
	require.NoError(t, err)
	require.Error(t, result.Conclusion)
	require.NotZero(t, len(result.Args))
}

func TestInspectorFileError(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "dataset")

	_, err := InspectFromFile(filePath, testInspectedInt)
	require.Error(t, err)
}
