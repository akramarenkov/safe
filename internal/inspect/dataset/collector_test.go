package dataset

import (
	"bufio"
	"bytes"
	"io"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/akramarenkov/safe/internal/consts"
	"github.com/akramarenkov/safe/internal/inspect/dataset/filler"
	"github.com/akramarenkov/wrecker"
	"github.com/stretchr/testify/require"
)

func TestCollectorIsValid(t *testing.T) {
	collector := Collector[int8]{
		Reference: func(...int64) (int64, error) { return 0, nil },
		Writer:    bytes.NewBuffer(nil),
	}

	require.NoError(t, collector.IsValid())

	collector = Collector[int8]{
		Reference: func(...int64) (int64, error) { return 0, nil },
	}

	require.Error(t, collector.IsValid())

	collector = Collector[int8]{
		Writer: bytes.NewBuffer(nil),
	}

	require.Error(t, collector.IsValid())
}

func TestCollector(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	collector := Collector[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 1,
		OverflowedItemsQuantity:    1,
		Reference:                  testReference,
		Writer:                     buffer,
		Fillers: []filler.Filler[int8]{
			filler.NewSet(
				func() []int8 {
					return filler.Span[int8](0, 1)
				},
				func() []int8 {
					return filler.Span[int8](126, 127)
				},
			),
		},
	}

	expected := "false 0 0 0 0 0 0\nfalse 128 127 1 0 0 0\n"

	err := collector.Collect()
	require.NoError(t, err)
	require.Equal(t, expected, buffer.String())
}

func TestCollectorDefaultFillers(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	collector := Collector[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 1,
		OverflowedItemsQuantity:    1,
		Reference:                  testReference,
		Writer:                     buffer,
	}

	err := collector.Collect()
	require.NoError(t, err)

	// two dataset items + empty string after last separator
	require.Len(t, strings.Split(buffer.String(), "\n"), 3)
}

func TestCollectorReferenceLimits(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	limited := int64(127)

	limits := map[int64]uint{
		limited: 1,
	}

	expected := map[int64]uint{
		limited: 1,
	}

	accounter := func(desired int64) int {
		quantity := 0

		scanner := bufio.NewScanner(buffer)

		for scanner.Scan() {
			items := strings.Split(scanner.Text(), " ")

			reference, err := strconv.ParseInt(items[1], consts.DecimalBase, 64)
			require.NoError(t, err)

			if reference == desired {
				quantity++
			}
		}

		return quantity
	}

	testCollectorReferenceLimits(t, buffer, nil)
	require.Equal(t, 4, accounter(limited))

	testCollectorReferenceLimits(t, buffer, limits)
	require.Equal(t, 1, accounter(limited))
	require.Equal(t, expected, limits)
}

func testCollectorReferenceLimits(
	t *testing.T,
	buffer *bytes.Buffer,
	limits map[int64]uint,
) {
	buffer.Reset()

	collector := Collector[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 10,
		OverflowedItemsQuantity:    10,
		Reference:                  testReference,
		ReferenceLimits:            limits,
		Writer:                     buffer,
		Fillers: []filler.Filler[int8]{
			filler.NewSet(
				func() []int8 {
					return filler.Span[int8](0, 1)
				},
				func() []int8 {
					return filler.Span[int8](126, 127)
				},
			),
		},
	}

	err := collector.Collect()
	require.NoError(t, err)
}

func TestCollectorUniqueness(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	collector := Collector[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 10,
		OverflowedItemsQuantity:    10,
		Reference:                  testReference,
		Writer:                     buffer,
		Fillers: []filler.Filler[int8]{
			filler.NewSame[int8](1, 100),
			filler.NewSame[int8](127, 100),
		},
	}

	err := collector.Collect()
	require.Equal(t, ErrNotEnoughDataInFillers, err)

	items := strings.Split(buffer.String(), "\n")

	slices.Sort(items)

	require.Equal(t, slices.Compact(slices.Clone(items)), items)

	// two dataset items + empty string after last separator
	require.Len(t, items, 3)
}

func TestCollectorCalcDatasetLength(t *testing.T) {
	collector := Collector[int8]{
		NotOverflowedItemsQuantity: 20,
		OverflowedItemsQuantity:    -1,
	}

	require.Equal(t, collector.NotOverflowedItemsQuantity, collector.calcDatasetLength())

	collector = Collector[int8]{
		NotOverflowedItemsQuantity: -1,
		OverflowedItemsQuantity:    10,
	}

	require.Equal(t, collector.OverflowedItemsQuantity, collector.calcDatasetLength())

	collector = Collector[int8]{
		NotOverflowedItemsQuantity: 20,
		OverflowedItemsQuantity:    10,
	}

	require.Equal(
		t,
		collector.NotOverflowedItemsQuantity+collector.OverflowedItemsQuantity,
		collector.calcDatasetLength(),
	)
}

func TestCollectorError(t *testing.T) {
	collector := Collector[int8]{}

	err := collector.Collect()
	require.Error(t, err)

	collector = Collector[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 1,
		OverflowedItemsQuantity:    1,
		Reference:                  testReference,
		Writer:                     wrecker.New(wrecker.Opts{Error: io.ErrUnexpectedEOF}),
	}

	err = collector.Collect()
	require.Equal(t, io.ErrUnexpectedEOF, err)

	collector = Collector[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 1,
		OverflowedItemsQuantity:    1,
		Reference:                  testReference,
		Writer:                     bytes.NewBuffer(nil),
		Fillers: []filler.Filler[int8]{
			filler.NewFaulty[int8](),
		},
	}

	err = collector.Collect()
	require.Equal(t, filler.ErrFaulty, err)

	collector = Collector[int8]{
		ArgsQuantity:               2,
		NotOverflowedItemsQuantity: 1 << 16,
		OverflowedItemsQuantity:    1 << 16,
		Reference:                  testReference,
		Writer:                     bytes.NewBuffer(nil),
		Fillers: []filler.Filler[int8]{
			filler.NewSet(
				func() []int8 {
					return filler.Span[int8](0, 1)
				},
				func() []int8 {
					return filler.Span[int8](126, 127)
				},
			),
		},
	}

	err = collector.Collect()
	require.Equal(t, ErrNotEnoughDataInFillers, err)
}

func TestCollectorFileError(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "")

	collector := Collector[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 10,
		OverflowedItemsQuantity:    10,
		Reference:                  testReference,
	}

	err := collector.CollectToFile(filePath)
	require.Error(t, err)
}

func TestWriteItem(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	err := WriteItem[int8](buffer, testReference, 1, 1, 1, 1, 1)
	require.NoError(t, err)

	err = WriteItem[int8](buffer, testReference, -128, -128, 127, 127, 1)
	require.NoError(t, err)

	require.Equal(t, "false 5 1 1 1 1 1\nfalse -1 -128 -128 127 127 1\n", buffer.String())
}

func TestWriteItemError(t *testing.T) {
	err := WriteItem[int8](wrecker.New(wrecker.Opts{}), testReference, 1, 1, 1, 1, 1)
	require.Error(t, err)
}

func TestItem(t *testing.T) {
	expected := "false -9223372036854775808 -128 -128 -128 -128 -128\n"

	actual := calcMaxItemLength[int8](5)
	require.Len(t, expected, actual)

	buffer := prepareItem[int8](
		make([]byte, actual),
		-9223372036854775808,
		nil,
		-128,
		-128,
		-128,
		-128,
		-128,
	)
	require.Equal(t, expected, string(buffer))
}
