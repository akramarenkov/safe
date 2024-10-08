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
			reference := int(end) - int(begin) + 1

			if begin > end {
				reference = int(begin) - int(end) + 1
			}

			//nolint:gosec // Is the result of subtracting the smaller value
			// from the larger value for int8 values
			referenceU := uint64(reference)

			require.Equal(t,
				referenceU,
				IterSize(begin, end),
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
			reference := uint64(end) - uint64(begin) + 1

			if begin > end {
				reference = uint64(begin) - uint64(end) + 1
			}

			actual := IterSize(begin, end)

			require.Equal(t, reference, actual, "begin: %v, end: %v", begin, end)
		}
	}
}

func testIterSizeMax(t *testing.T) {
	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		IterSize[int64](math.MinInt64+2, math.MaxInt64),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		IterSize[int64](math.MaxInt64, math.MinInt64+2),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		IterSize[uint64](2, math.MaxUint64),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		IterSize[uint64](math.MaxUint64, 2),
	)

	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterSize[int64](math.MinInt64+1, math.MaxInt64),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterSize[int64](math.MaxInt64, math.MinInt64+1),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterSize[uint64](1, math.MaxUint64),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterSize[uint64](math.MaxUint64, 1),
	)

	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterSize[int64](math.MinInt64, math.MaxInt64),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterSize[int64](math.MaxInt64, math.MinInt64),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterSize[uint64](0, math.MaxUint64),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterSize[uint64](math.MaxUint64, 0),
	)
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
	expectedFinalReference := int(begin) + expectedIterations*int(step)

	for number := range IterStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)

		iterations++
		reference += int(step)
	}

	require.Equal(t, expectedIterations, iterations, "step: %v", step)
	require.Equal(t, expectedFinalReference, reference, "step: %v", step)
}

func testIterStepBackwardInt(t *testing.T, step int8) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	iterations := 0
	reference := int(begin)

	expectedIterations := (int(begin)-int(end))/int(step) + 1
	expectedFinalReference := int(begin) - expectedIterations*int(step)

	for number := range IterStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)

		iterations++
		reference -= int(step)
	}

	require.Equal(t, expectedIterations, iterations, "step: %v", step)
	require.Equal(t, expectedFinalReference, reference, "step: %v", step)
}

func testIterStepForwardUint(t *testing.T, step uint8) {
	begin := uint8(0)
	end := uint8(math.MaxUint8)

	iterations := 0
	reference := int(begin)

	expectedIterations := (int(end)-int(begin))/int(step) + 1
	expectedFinalReference := int(begin) + expectedIterations*int(step)

	for number := range IterStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)

		iterations++
		reference += int(step)
	}

	require.Equal(t, expectedIterations, iterations, "step: %v", step)
	require.Equal(t, expectedFinalReference, reference, "step: %v", step)
}

func testIterStepBackwardUint(t *testing.T, step uint8) {
	begin := uint8(math.MaxUint8)
	end := uint8(0)

	iterations := 0
	reference := int(begin)

	expectedIterations := (int(begin)-int(end))/int(step) + 1
	expectedFinalReference := int(begin) - expectedIterations*int(step)

	for number := range IterStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)

		iterations++
		reference -= int(step)
	}

	require.Equal(t, expectedIterations, iterations, "step: %v", step)
	require.Equal(t, expectedFinalReference, reference, "step: %v", step)
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

func TestIterStepSize(t *testing.T) {
	testIterStepSizeInt(t)
	testIterStepSizeUint(t)
	testIterStepSizeMax(t)

	require.Panics(
		t,
		func() {
			_ = IterStepSize(1, 2, -1, nil, nil)
		},
	)

	require.Panics(
		t,
		func() {
			_ = IterStepSize(1, 2, 0, nil, nil)
		},
	)
}

func testIterStepSizeInt(t *testing.T) {
	for begin := range Iter[int8](math.MinInt8, math.MaxInt8) {
		for end := range Iter[int8](math.MinInt8, math.MaxInt8) {
			for step := range Iter[int8](1, math.MaxInt8) {
				reference := (int(end) - int(begin) + 1) / int(step)

				if begin > end {
					reference = (int(begin) - int(end) + 1) / int(step)
				}

				//nolint:gosec // Is the result of subtracting the smaller value
				// from the larger value for int8 values
				referenceU := uint64(reference)

				actual := IterStepSize(begin, end, step, nil, nil)

				// duplication of conditions is done for performance reasons
				if actual != referenceU {
					require.Equal(
						t,
						referenceU,
						actual,
						"begin: %v, end: %v, step: %v",
						begin,
						end,
						step,
					)
				}
			}
		}
	}
}

func testIterStepSizeUint(t *testing.T) {
	for begin := range Iter[uint8](0, math.MaxUint8) {
		for end := range Iter[uint8](0, math.MaxUint8) {
			for step := range Iter[uint8](1, math.MaxUint8) {
				reference := (uint64(end) - uint64(begin) + 1) / uint64(step)

				if begin > end {
					reference = (uint64(begin) - uint64(end) + 1) / uint64(step)
				}

				actual := IterStepSize(begin, end, step, nil, nil)

				// duplication of conditions is done for performance reasons
				if actual != reference {
					require.Equal(
						t,
						reference,
						actual,
						"begin: %v, end: %v, step: %v",
						begin,
						end,
						step,
					)
				}
			}
		}
	}
}

func testIterStepSizeMax(t *testing.T) {
	require.Equal(
		t,
		uint64(4611686018427387904),
		IterStepSize[int64](math.MinInt64, math.MaxInt64, 4, nil, nil),
	)
	require.Equal(
		t,
		uint64(4611686018427387904),
		IterStepSize[int64](math.MaxInt64, math.MinInt64, 4, nil, nil),
	)

	require.Equal(
		t,
		uint64(4611686018427387904),
		IterStepSize[uint64](0, math.MaxUint64, 4, nil, nil),
	)
	require.Equal(
		t,
		uint64(4611686018427387904),
		IterStepSize[uint64](math.MaxUint64, 0, 4, nil, nil),
	)

	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterStepSize[int64](math.MinInt64, math.MaxInt64, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterStepSize[int64](math.MaxInt64, math.MinInt64, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterStepSize[uint64](0, math.MaxUint64, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterStepSize[uint64](math.MaxUint64, 0, 1, nil, nil),
	)
}

func BenchmarkIter(b *testing.B) {
	number := 0

	for value := range Iter(0, b.N) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkIterSize(b *testing.B) {
	size := uint64(0)

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

func BenchmarkIterStepSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = IterStepSize(0, b.N, 1, nil, nil)
	}

	require.NotZero(b, size)
}
