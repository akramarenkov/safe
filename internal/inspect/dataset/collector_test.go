package dataset

import (
	"bufio"
	"bytes"
	"path/filepath"
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

	expected := "false 0 0 0 0 0 0\n" +
		"false 128 127 1 0 0 0\n"

	err := collector.Collect()
	require.NoError(t, err)
	require.Equal(t, expected, buffer.String())
}

func TestCollectorNotOverflowedReaching(t *testing.T) {
	collector := Collector[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 1 << 16,
		OverflowedItemsQuantity:    1 << 16,
		Reference:                  testReference,
		Writer:                     bytes.NewBuffer(nil),
		Fillers: []filler.Filler[int8]{
			filler.NewSet[int8](),
		},
	}

	err := collector.Collect()
	require.NoError(t, err)
}

func TestCollectorReferenceLimits(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	limits := map[int64]uint{
		-6: 1,
	}

	expected := map[int64]uint{
		-6: 1,
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
	require.Equal(t, 3, accounter(-6))

	testCollectorReferenceLimits(t, buffer, limits)
	require.Equal(t, 1, accounter(-6))
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
			filler.NewSet[int8](),
		},
	}

	err := collector.Collect()
	require.NoError(t, err)
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
		Writer:                     wrecker.New(wrecker.Opts{}),
	}

	err = collector.Collect()
	require.Error(t, err)

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
	require.Error(t, err)

	collector = Collector[int8]{
		ArgsQuantity:               2,
		NotOverflowedItemsQuantity: 1 << 16,
		OverflowedItemsQuantity:    1 << 16,
		Reference:                  testReference,
		Writer:                     bytes.NewBuffer(nil),
		Fillers: []filler.Filler[int8]{
			filler.NewSet[int8](),
		},
	}

	err = collector.Collect()
	require.Error(t, err)
}

func TestCollectorFileError(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "")

	collector := Collector[int8]{
		ArgsQuantity:               5,
		NotOverflowedItemsQuantity: 10,
		OverflowedItemsQuantity:    10,
		Reference:                  testReference,
		Fillers: []filler.Filler[int8]{
			filler.NewSet[int8](),
		},
	}

	err := collector.CollectToFile(filePath)
	require.Error(t, err)
}

func TestCalcItemLength(t *testing.T) {
	expected := len("false 18446744073709551615 -127 -127 -127 -127 -127\n")

	require.Equal(t, expected, calcMaxItemLength(5))
}
