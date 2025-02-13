package dataset

import (
	"bytes"
	"math"
	"path/filepath"
	"testing"

	"github.com/akramarenkov/safe/internal/inspect"
	"github.com/akramarenkov/safe/internal/inspect/dataset/filler"
	"github.com/akramarenkov/safe/internal/inspect/types"

	"github.com/stretchr/testify/require"
)

func TestFile(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "dataset")

	opts := Opts[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 10,
		OverflowedItemsQuantity:    10,
		Reference:                  testReference,
	}

	err := CollectToFile(opts, filePath)
	require.NoError(t, err)

	result, err := InspectFromFile(testInspectedSig, filePath)
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

func BenchmarkDataset(b *testing.B) {
	const (
		argsQuantity  = 10
		itemsQuantity = 1 << 16
	)

	buffer := bytes.NewBuffer(nil)

	buffer.Grow(2 * itemsQuantity * calcMaxItemLength[int8](argsQuantity))

	opts := Opts[int8]{
		ArgsQuantity:               argsQuantity,
		NotOverflowedItemsQuantity: itemsQuantity,
		OverflowedItemsQuantity:    itemsQuantity,
		Reference:                  testReference,
		Fillers: []filler.Filler[int8]{
			filler.NewSet[int8](),
		},
	}

	for range b.N {
		if err := Collect(opts, buffer); err != nil {
			require.NoError(b, err)
		}

		result, err := Inspect(testInspectedSig, buffer)
		if err != nil {
			require.NoError(b, err)
		}

		if result.Conclusion != nil {
			require.NoError(b, result.Conclusion)
		}
	}
}

func testReference(args ...int64) (int64, error) {
	reference := int64(0)

	for _, arg := range args {
		reference += arg
	}

	return reference, nil
}

func testReference8[Type types.USI8](args ...Type) int64 {
	reference := int64(0)

	for _, arg := range args {
		reference += int64(arg)
	}

	return reference
}

func testInspectedSig(args ...int8) (int8, error) {
	reference := testReference8(args...)

	if reference > math.MaxInt8 || reference < math.MinInt8 {
		return 0, inspect.ErrOverflow
	}

	return int8(reference), nil
}

func testInspectedUns(args ...uint8) (uint8, error) {
	reference := testReference8(args...)

	if reference > math.MaxUint8 || reference < 0 {
		return 0, inspect.ErrOverflow
	}

	return uint8(reference), nil
}

func testInspectedUnsafe[Type types.USI8](args ...Type) Type {
	sum := Type(0)

	for _, arg := range args {
		sum += arg
	}

	return sum
}
