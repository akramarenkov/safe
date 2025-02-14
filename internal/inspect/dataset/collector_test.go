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

func TestOptsDatasetLength(t *testing.T) {
	opts := Opts[int8]{
		NotOverflowedItemsQuantity: 20,
		OverflowedItemsQuantity:    -1,
	}

	require.Equal(t, opts.NotOverflowedItemsQuantity, opts.datasetLength())

	opts = Opts[int8]{
		NotOverflowedItemsQuantity: -1,
		OverflowedItemsQuantity:    10,
	}

	require.Equal(t, opts.OverflowedItemsQuantity, opts.datasetLength())

	opts = Opts[int8]{
		NotOverflowedItemsQuantity: 20,
		OverflowedItemsQuantity:    10,
	}

	require.Equal(
		t,
		opts.NotOverflowedItemsQuantity+opts.OverflowedItemsQuantity,
		opts.datasetLength(),
	)
}

func TestCollect(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	opts := Opts[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 1,
		OverflowedItemsQuantity:    1,
		Reference:                  testReference,
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

	err := Collect(opts, buffer)
	require.NoError(t, err)
	require.Equal(t, expected, buffer.String())
}

func TestCollectDefaultFillers(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	opts := Opts[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 1,
		OverflowedItemsQuantity:    1,
		Reference:                  testReference,
	}

	err := Collect(opts, buffer)
	require.NoError(t, err)

	// two dataset items + empty string after last separator
	require.Len(t, strings.Split(buffer.String(), "\n"), 3)
}

func TestCollectReferenceLimits(t *testing.T) {
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

	testCollectReferenceLimits(t, buffer, nil)
	require.Equal(t, 4, accounter(limited))

	testCollectReferenceLimits(t, buffer, limits)
	require.Equal(t, 1, accounter(limited))
	require.Equal(t, expected, limits)
}

func testCollectReferenceLimits(
	t *testing.T,
	buffer *bytes.Buffer,
	limits map[int64]uint,
) {
	buffer.Reset()

	opts := Opts[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 10,
		OverflowedItemsQuantity:    10,
		Reference:                  testReference,
		ReferenceLimits:            limits,
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

	err := Collect(opts, buffer)
	require.NoError(t, err)
}

func TestCollectUniqueness(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	opts := Opts[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 10,
		OverflowedItemsQuantity:    10,
		Reference:                  testReference,
		Fillers: []filler.Filler[int8]{
			filler.NewSame[int8](1, 100),
			filler.NewSame[int8](127, 100),
		},
	}

	err := Collect(opts, buffer)
	require.Equal(t, ErrNotEnoughDataInFillers, err)

	items := strings.Split(buffer.String(), "\n")

	slices.Sort(items)

	require.Equal(t, slices.Compact(slices.Clone(items)), items)

	// two dataset items + empty string after last separator
	require.Len(t, items, 3)
}

func TestCollectError(t *testing.T) {
	opts := Opts[int8]{}

	err := Collect(opts, bytes.NewBuffer(nil))
	require.Error(t, err)

	opts = Opts[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 1,
		OverflowedItemsQuantity:    1,
		Reference:                  testReference,
	}

	err = Collect(opts, nil)
	require.Error(t, err)

	err = Collect(opts, wrecker.New(wrecker.Opts{Error: io.ErrUnexpectedEOF}))
	require.Equal(t, io.ErrUnexpectedEOF, err)

	opts = Opts[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 1,
		OverflowedItemsQuantity:    1,
		Reference:                  testReference,
		Fillers: []filler.Filler[int8]{
			filler.NewFaulty[int8](),
		},
	}

	err = Collect(opts, bytes.NewBuffer(nil))
	require.Equal(t, filler.ErrFaulty, err)

	opts = Opts[int8]{
		ArgsQuantity:               2,
		NotOverflowedItemsQuantity: 1 << 16,
		OverflowedItemsQuantity:    1 << 16,
		Reference:                  testReference,
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

	err = Collect(opts, bytes.NewBuffer(nil))
	require.Equal(t, ErrNotEnoughDataInFillers, err)
}

func TestCollectToFileError(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "")

	opts := Opts[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 10,
		OverflowedItemsQuantity:    10,
		Reference:                  testReference,
	}

	err := CollectToFile(opts, filePath)
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

func TestItemSig(t *testing.T) {
	expected := "false -9223372036854775808 -128 -128 -128 -128 -128\n"

	actual := calcMaxItemLength[int8](5)
	require.Equal(t, len(expected), actual)

	item := prepareItem[int8](
		make([]byte, actual),
		-9223372036854775808,
		nil,
		-128,
		-128,
		-128,
		-128,
		-128,
	)
	require.Equal(t, expected, string(item))
}

func TestItemUns(t *testing.T) {
	expected := "false -9223372036854775808 255 255 255 255 255\n"

	actual := calcMaxItemLength[uint8](5)
	require.Equal(t, len(expected), actual)

	item := prepareItem[uint8](
		make([]byte, actual),
		-9223372036854775808,
		nil,
		255,
		255,
		255,
		255,
		255,
	)
	require.Equal(t, expected, string(item))
}
