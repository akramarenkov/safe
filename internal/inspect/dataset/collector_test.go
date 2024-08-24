package dataset

import (
	"bytes"
	"path/filepath"
	"testing"

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
			filler.NewBoundary[int8](),
		},
	}

	expected := "false -640 -128 -128 -128 -128 -128\n" +
		"false -128 127 127 -126 -128 -128\n"

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
			filler.NewBoundary[int8](),
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
			filler.NewBoundary[int8](),
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
			filler.NewBoundary[int8](),
		},
	}

	err := collector.CollectToFile(filePath)
	require.Error(t, err)
}

func TestCalcItemLength(t *testing.T) {
	expected := len("false 18446744073709551615 -127 -127 -127 -127 -127\n")

	require.Equal(t, expected, calcMaxItemLength(5))
}
