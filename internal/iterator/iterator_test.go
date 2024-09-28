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

func TestIterSize(t *testing.T) {
	testIterSizeInt(t)
	testIterSizeUint(t)
	testIterSizeMax(t)
}

func testIterSizeInt(t *testing.T) {
	for begin := range Iter[int8](math.MinInt8, math.MaxInt8) {
		for end := range Iter[int8](math.MinInt8, math.MaxInt8) {
			reference := int64(end) - int64(begin) + 1

			if begin > end {
				reference = int64(begin) - int64(end) + 1
			}

			require.Equal(
				t,
				reference,
				int64(IterSize(begin, end)),
				"begin: %v, end: %v",
				begin,
				end,
			)
		}
	}
}

func testIterSizeUint(t *testing.T) {
	for begin := range Iter[uint8](0, math.MaxUint8) {
		for end := range Iter[uint8](0, math.MaxUint8) {
			reference := int64(end) - int64(begin) + 1

			if begin > end {
				reference = int64(begin) - int64(end) + 1
			}

			require.Equal(
				t,
				reference,
				int64(IterSize(begin, end)),
				"begin: %v, end: %v",
				begin,
				end,
			)
		}
	}
}

func testIterSizeMax(t *testing.T) {
	require.Equal(t, math.MaxInt-1, IterSize(0, math.MaxInt-2))
	require.Equal(t, math.MaxInt, IterSize(0, math.MaxInt-1))
	require.Equal(t, math.MaxInt, IterSize(0, math.MaxInt))
	require.Equal(t, math.MaxInt, IterSize[int64](math.MinInt64, math.MaxInt64))
	require.Equal(t, math.MaxInt, IterSize[int64](math.MaxInt64, math.MinInt64))
	require.Equal(t, math.MaxInt, IterSize[uint64](0, math.MaxUint64))
	require.Equal(t, math.MaxInt, IterSize[uint64](math.MaxUint64, 0))
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

func BenchmarkIterSize(b *testing.B) {
	size := 0

	for range b.N {
		size = IterSize(0, b.N)
	}

	require.NotZero(b, size)
}

func BenchmarkIterStep(b *testing.B) {
	number := 0

	for value := range IterStep(0, b.N, 1, nil, nil) {
		number = value
	}

	require.NotZero(b, number)
}
