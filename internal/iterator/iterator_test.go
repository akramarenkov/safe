package iterator

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIter(t *testing.T) {
	testIterForwardInt(t)
	testIterBackwardInt(t)

	testIterForwardUint(t)
	testIterBackwardUint(t)

	testIterForwardPartial(t)
	testIterBackwardPartial(t)
}

func testIterForwardInt(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference++
	}

	require.Equal(t, int(end)+1, reference)
}

func testIterBackwardInt(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference--
	}

	require.Equal(t, int(end)-1, reference)
}

func testIterForwardUint(t *testing.T) {
	begin := uint8(0)
	end := uint8(math.MaxUint8)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference++
	}

	require.Equal(t, int(end)+1, reference)
}

func testIterBackwardUint(t *testing.T) {
	begin := uint8(math.MaxUint8)
	end := uint8(0)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference--
	}

	require.Equal(t, int(end)-1, reference)
}

func testIterForwardPartial(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		if number == 0 {
			break
		}

		reference++
	}

	require.Equal(t, 0, reference)
}

func testIterBackwardPartial(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		if number == 0 {
			break
		}

		reference--
	}

	require.Equal(t, 0, reference)
}

func TestIterStep(t *testing.T) {
	for step := range Iter[int8](1, math.MaxInt8) {
		testIterStepForwardInt(t, step)
		testIterStepBackwardInt(t, step)
	}

	for step := range Iter[uint8](1, math.MaxUint8) {
		testIterStepForwardUint(t, step)
		testIterStepBackwardUint(t, step)
	}

	testIterStepForwardPartial(t)
	testIterStepBackwardPartial(t)

	require.Panics(
		t,
		func() {
			for number := range IterStep(1, 2, -1, nil, nil) {
				_ = number
			}
		},
	)

	require.Panics(
		t,
		func() {
			for number := range IterStep(1, 2, 0, nil, nil) {
				_ = number
			}
		},
	)
}

func testIterStepForwardInt(t *testing.T, step int8) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	iterations := 0
	reference := int(begin)

	expectedIterations := (int(end)-int(begin))/int(step) + 1
	expectedReference := int(begin) + expectedIterations*int(step)

	for number := range IterStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)

		iterations++
		reference += int(step)
	}

	require.Equal(t, expectedIterations, iterations, "step: %v", step)
	require.Equal(t, expectedReference, reference, "step: %v", step)
}

func testIterStepBackwardInt(t *testing.T, step int8) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	iterations := 0
	reference := int(begin)

	expectedIterations := (int(begin)-int(end))/int(step) + 1
	expectedReference := int(begin) - expectedIterations*int(step)

	for number := range IterStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)

		iterations++
		reference -= int(step)
	}

	require.Equal(t, expectedIterations, iterations, "step: %v", step)
	require.Equal(t, expectedReference, reference, "step: %v", step)
}

func testIterStepForwardUint(t *testing.T, step uint8) {
	begin := uint8(0)
	end := uint8(math.MaxUint8)

	iterations := 0
	reference := int(begin)

	expectedIterations := (int(end)-int(begin))/int(step) + 1
	expectedReference := int(begin) + expectedIterations*int(step)

	for number := range IterStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)

		iterations++
		reference += int(step)
	}

	require.Equal(t, expectedIterations, iterations, "step: %v", step)
	require.Equal(t, expectedReference, reference, "step: %v", step)
}

func testIterStepBackwardUint(t *testing.T, step uint8) {
	begin := uint8(math.MaxUint8)
	end := uint8(0)

	iterations := 0
	reference := int(begin)

	expectedIterations := (int(begin)-int(end))/int(step) + 1
	expectedReference := int(begin) - expectedIterations*int(step)

	for number := range IterStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)

		iterations++
		reference -= int(step)
	}

	require.Equal(t, expectedIterations, iterations, "step: %v", step)
	require.Equal(t, expectedReference, reference, "step: %v", step)
}

func testIterStepForwardPartial(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)

	for number := range IterStep(begin, end, 1, nil, nil) {
		require.Equal(t, reference, int(number))

		if number == 0 {
			break
		}

		reference++
	}

	require.Equal(t, 0, reference)
}

func testIterStepBackwardPartial(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	reference := int(begin)

	for number := range IterStep(begin, end, 1, nil, nil) {
		require.Equal(t, reference, int(number))

		if number == 0 {
			break
		}

		reference--
	}

	require.Equal(t, 0, reference)
}

func BenchmarkIter(b *testing.B) {
	number := 0

	for value := range Iter(0, b.N) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkIterStep(b *testing.B) {
	number := 0

	for value := range IterStep(0, b.N, 1, nil, nil) {
		number = value
	}

	require.NotZero(b, number)
}
