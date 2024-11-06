package iterator

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIter(t *testing.T) {
	testIterForwardSig(t)
	testIterBackwardSig(t)
	testIterForwardUns(t)
	testIterBackwardUns(t)
	testIterForwardPart(t)
	testIterBackwardPart(t)
}

func testIterForwardSig(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference++
	}

	require.NotEqual(t, int(begin), reference)
}

func testIterBackwardSig(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference--
	}

	require.NotEqual(t, int(begin), reference)
}

func testIterForwardUns(t *testing.T) {
	begin := uint8(0)
	end := uint8(math.MaxUint8)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference++
	}

	require.NotEqual(t, int(begin), reference)
}

func testIterBackwardUns(t *testing.T) {
	begin := uint8(math.MaxUint8)
	end := uint8(0)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference--
	}

	require.NotEqual(t, int(begin), reference)
}

func testIterForwardPart(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	breakAt := 3
	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		if int(number) == breakAt {
			break
		}

		reference++
	}

	require.Equal(t, breakAt, reference)
}

func testIterBackwardPart(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	breakAt := 3
	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		if int(number) == breakAt {
			break
		}

		reference--
	}

	require.Equal(t, breakAt, reference)
}
