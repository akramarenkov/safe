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
	testIterForwardPartial(t)
	testIterBackwardPartial(t)
}

func testIterForwardSig(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference++
	}

	require.Equal(t, int(end)+1, reference)
}

func testIterBackwardSig(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference--
	}

	require.Equal(t, int(end)-1, reference)
}

func testIterForwardUns(t *testing.T) {
	begin := uint8(0)
	end := uint8(math.MaxUint8)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference++
	}

	require.Equal(t, int(end)+1, reference)
}

func testIterBackwardUns(t *testing.T) {
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

func testIterBackwardPartial(t *testing.T) {
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

func TestIterSize(t *testing.T) {
	testIterSizeMan(t)
	testIterSizeSig(t)
	testIterSizeUns(t)
	testIterSizeMax(t)
}

func testIterSizeMan(t *testing.T) {
	require.Equal(t, uint64(3), IterSize[int8](-1, 1))
	require.Equal(t, uint64(3), IterSize[int8](1, -1))
	require.Equal(t, uint64(6), IterSize[int8](-2, 3))
	require.Equal(t, uint64(6), IterSize[int8](3, -2))
	require.Equal(t, uint64(256), IterSize[int8](-128, 127))

	require.Equal(t, uint64(3), IterSize[uint8](1, 3))
	require.Equal(t, uint64(3), IterSize[uint8](3, 1))
	require.Equal(t, uint64(6), IterSize[uint8](1, 6))
	require.Equal(t, uint64(6), IterSize[uint8](6, 1))
	require.Equal(t, uint64(256), IterSize[uint8](0, 255))
}

func testIterSizeSig(t *testing.T) {
	for begin := range Iter[int8](math.MinInt8, math.MaxInt8) {
		for end := range Iter[int8](math.MinInt8, math.MaxInt8) {
			reference := int(end) - int(begin) + 1

			if begin > end {
				reference = int(begin) - int(end) + 1
			}

			//nolint:gosec // Is the result of subtracting the smaller value
			// from the larger value for int8 values
			referenceU := uint64(reference)

			require.Equal(
				t,
				referenceU,
				IterSize(begin, end),
				"begin: %v, end: %v",
				begin,
				end,
			)
		}
	}
}

func testIterSizeUns(t *testing.T) {
	for begin := range Iter[uint8](0, math.MaxUint8) {
		for end := range Iter[uint8](0, math.MaxUint8) {
			reference := uint64(end) - uint64(begin) + 1

			if begin > end {
				reference = uint64(begin) - uint64(end) + 1
			}

			require.Equal(
				t,
				reference,
				IterSize(begin, end),
				"begin: %v, end: %v",
				begin,
				end,
			)
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
		testIterStepForwardSig(t, step)
		testIterStepBackwardSig(t, step)
		testIterStepBackwardSigNotOnEntireRange(t, step)
	}

	for step := range Iter[uint8](1, math.MaxUint8) {
		testIterStepForwardUns(t, step)
		testIterStepBackwardUns(t, step)
		testIterStepBackwardUnsNotOnEntireRange(t, step)
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

func testIterStepForwardSig(t *testing.T, step int8) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(end)-int(begin))/int(step) + 1
	final := int(begin) + iterations*int(step)

	for id, number := range IterStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference += int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testIterStepBackwardSig(t *testing.T, step int8) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(begin)-int(end))/int(step) + 1
	final := int(begin) - iterations*int(step)

	for id, number := range IterStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testIterStepBackwardSigNotOnEntireRange(t *testing.T, step int8) {
	begin := int8(math.MaxInt8)
	end := int8(0)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(begin)-int(end))/int(step) + 1
	final := int(begin) - iterations*int(step)

	for id, number := range IterStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testIterStepForwardUns(t *testing.T, step uint8) {
	begin := uint8(0)
	end := uint8(math.MaxUint8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(end)-int(begin))/int(step) + 1
	final := int(begin) + iterations*int(step)

	for id, number := range IterStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference += int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testIterStepBackwardUns(t *testing.T, step uint8) {
	begin := uint8(math.MaxUint8)
	end := uint8(0)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(begin)-int(end))/int(step) + 1
	final := int(begin) - iterations*int(step)

	for id, number := range IterStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testIterStepBackwardUnsNotOnEntireRange(t *testing.T, step uint8) {
	begin := uint8(math.MaxUint8)
	end := uint8(math.MaxUint8 / 2)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(begin)-int(end))/int(step) + 1
	final := int(begin) - iterations*int(step)

	for id, number := range IterStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testIterStepForwardPartial(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	breakAt := 3
	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range IterStep(begin, end, 1, nil, nil) {
		require.Equal(t, reference, int(number))
		require.Equal(t, referenceID, id)

		if int(number) == breakAt {
			break
		}

		reference++
		referenceID++
	}

	require.Equal(t, breakAt, reference)
}

func testIterStepBackwardPartial(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	breakAt := 3
	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range IterStep(begin, end, 1, nil, nil) {
		require.Equal(t, reference, int(number))
		require.Equal(t, referenceID, id)

		if int(number) == breakAt {
			break
		}

		reference--
		referenceID++
	}

	require.Equal(t, breakAt, reference)
}

func TestIterStepSize(t *testing.T) {
	testIterStepSizeMan(t)
	testIterStepSizeSig(t)
	testIterStepSizeUns(t)
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

func testIterStepSizeMan(t *testing.T) {
	require.Equal(t, uint64(3), IterStepSize[int8](-1, 1, 1, nil, nil))
	require.Equal(t, uint64(3), IterStepSize[int8](1, -1, 1, nil, nil))
	require.Equal(t, uint64(6), IterStepSize[int8](-2, 3, 1, nil, nil))
	require.Equal(t, uint64(6), IterStepSize[int8](3, -2, 1, nil, nil))
	require.Equal(t, uint64(256), IterStepSize[int8](-128, 127, 1, nil, nil))

	require.Equal(t, uint64(3), IterStepSize[int8](-2, 5, 3, nil, nil))
	require.Equal(t, uint64(3), IterStepSize[int8](-2, 6, 3, nil, nil))
	require.Equal(t, uint64(4), IterStepSize[int8](-2, 7, 3, nil, nil))
	require.Equal(t, uint64(3), IterStepSize[int8](-128, 127, 86, nil, nil))
	require.Equal(t, uint64(86), IterStepSize[int8](-128, 127, 3, nil, nil))
	require.Equal(t, uint64(128), IterStepSize[int8](-128, 127, 2, nil, nil))

	require.Equal(t, uint64(3), IterStepSize[uint8](1, 3, 1, nil, nil))
	require.Equal(t, uint64(3), IterStepSize[uint8](3, 1, 1, nil, nil))
	require.Equal(t, uint64(6), IterStepSize[uint8](1, 6, 1, nil, nil))
	require.Equal(t, uint64(6), IterStepSize[uint8](6, 1, 1, nil, nil))
	require.Equal(t, uint64(256), IterStepSize[uint8](0, 255, 1, nil, nil))

	require.Equal(t, uint64(3), IterStepSize[uint8](2, 9, 3, nil, nil))
	require.Equal(t, uint64(3), IterStepSize[uint8](2, 10, 3, nil, nil))
	require.Equal(t, uint64(4), IterStepSize[uint8](2, 11, 3, nil, nil))
	require.Equal(t, uint64(3), IterStepSize[uint8](0, 255, 86, nil, nil))
	require.Equal(t, uint64(86), IterStepSize[uint8](0, 255, 3, nil, nil))
	require.Equal(t, uint64(128), IterStepSize[uint8](0, 255, 2, nil, nil))
}

func testIterStepSizeSig(t *testing.T) {
	for begin := range Iter[int8](math.MinInt8, math.MaxInt8) {
		for end := range Iter[int8](math.MinInt8, math.MaxInt8) {
			for step := range Iter[int8](1, math.MaxInt8) {
				reference := (int(end)-int(begin))/int(step) + 1

				if begin > end {
					reference = (int(begin)-int(end))/int(step) + 1
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

func testIterStepSizeUns(t *testing.T) {
	for begin := range Iter[uint8](0, math.MaxUint8) {
		for end := range Iter[uint8](0, math.MaxUint8) {
			for step := range Iter[uint8](1, math.MaxUint8) {
				reference := (uint64(end)-uint64(begin))/uint64(step) + 1

				if begin > end {
					reference = (uint64(begin)-uint64(end))/uint64(step) + 1
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
		uint64(9223372036854775808),
		IterStepSize[int64](math.MinInt64, math.MaxInt64, 2, nil, nil),
	)
	require.Equal(
		t,
		uint64(9223372036854775808),
		IterStepSize[int64](math.MaxInt64, math.MinInt64, 2, nil, nil),
	)

	require.Equal(
		t,
		uint64(9223372036854775808),
		IterStepSize[uint64](0, math.MaxUint64, 2, nil, nil),
	)
	require.Equal(
		t,
		uint64(9223372036854775808),
		IterStepSize[uint64](math.MaxUint64, 0, 2, nil, nil),
	)
	require.Equal(
		t,
		uint64(2),
		IterStepSize[uint64](0, math.MaxUint64, 9223372036854775808, nil, nil),
	)
	require.Equal(
		t,
		uint64(2),
		IterStepSize[uint64](math.MaxUint64, 0, 9223372036854775808, nil, nil),
	)

	require.Equal(
		t,
		uint64(6148914691236517206),
		IterStepSize[int64](math.MinInt64, math.MaxInt64, 3, nil, nil),
	)
	require.Equal(
		t,
		uint64(6148914691236517206),
		IterStepSize[int64](math.MaxInt64, math.MinInt64, 3, nil, nil),
	)
	require.Equal(
		t,
		uint64(3),
		IterStepSize[int64](math.MinInt64, math.MaxInt64, 6148914691236517206, nil, nil),
	)
	require.Equal(
		t,
		uint64(3),
		IterStepSize[int64](math.MaxInt64, math.MinInt64, 6148914691236517206, nil, nil),
	)

	require.Equal(
		t,
		uint64(6148914691236517206),
		IterStepSize[uint64](0, math.MaxUint64, 3, nil, nil),
	)
	require.Equal(
		t,
		uint64(6148914691236517206),
		IterStepSize[uint64](math.MaxUint64, 0, 3, nil, nil),
	)
	require.Equal(
		t,
		uint64(3),
		IterStepSize[uint64](0, math.MaxUint64, 6148914691236517206, nil, nil),
	)
	require.Equal(
		t,
		uint64(3),
		IterStepSize[uint64](math.MaxUint64, 0, 6148914691236517206, nil, nil),
	)

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
		uint64(4),
		IterStepSize[int64](math.MinInt64, math.MaxInt64, 4611686018427387904, nil, nil),
	)
	require.Equal(
		t,
		uint64(4),
		IterStepSize[int64](math.MaxInt64, math.MinInt64, 4611686018427387904, nil, nil),
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
		uint64(4),
		IterStepSize[uint64](0, math.MaxUint64, 4611686018427387904, nil, nil),
	)
	require.Equal(
		t,
		uint64(4),
		IterStepSize[uint64](math.MaxUint64, 0, 4611686018427387904, nil, nil),
	)

	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		IterStepSize[int64](math.MinInt64+2, math.MaxInt64, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		IterStepSize[int64](math.MaxInt64, math.MinInt64+2, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		IterStepSize[uint64](2, math.MaxUint64, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		IterStepSize[uint64](math.MaxUint64, 2, 1, nil, nil),
	)

	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterStepSize[int64](math.MinInt64+1, math.MaxInt64, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterStepSize[int64](math.MaxInt64, math.MinInt64+1, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterStepSize[uint64](1, math.MaxUint64, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterStepSize[uint64](math.MaxUint64, 1, 1, nil, nil),
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

func TestIterStrict(t *testing.T) {
	for step := range Iter[int8](1, math.MaxInt8) {
		testIterStrictForwardSig(t, step)
		testIterStrictBackwardSig(t, step)
		testIterStrictBackwardSigNotOnEntireRange(t, step)
	}

	for step := range Iter[uint8](1, math.MaxUint8) {
		testIterStrictForwardUns(t, step)
		testIterStrictBackwardUns(t, step)
		testIterStrictBackwardUnsNotOnEntireRange(t, step)
	}

	testIterStrictForwardPartial(t)
	testIterStrictBackwardPartial(t)

	testIterStrictIterations(t, 1, 0, 1, -1, 2)
	testIterStrictIterations(t, 1, 1, 1, -1, 1)
	testIterStrictIterations(t, 1, 2, 1, -1, 0)
	testIterStrictIterations(t, 1, 0, 1, 0, 0)
	testIterStrictIterations(t, 1, 1, 1, 0, 1)
	testIterStrictIterations(t, 1, 2, 1, 0, 2)

	require.Panics(
		t,
		func() {
			for number := range IterStrict(1, 2, -1, 0, nil, nil) {
				_ = number
			}
		},
	)

	require.Panics(
		t,
		func() {
			for number := range IterStrict(1, 2, 0, 0, nil, nil) {
				_ = number
			}
		},
	)
}

func testIterStrictForwardSig(t *testing.T, step int8) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(end)-int(begin))/int(step) + 1
	final := int(begin) + iterations*int(step)

	for id, number := range IterStrict(begin, end, step, 0, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference += int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testIterStrictBackwardSig(t *testing.T, step int8) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(begin)-int(end))/int(step) + 1
	final := int(begin) - iterations*int(step)

	for id, number := range IterStrict(begin, end, step, -1, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testIterStrictBackwardSigNotOnEntireRange(t *testing.T, step int8) {
	begin := int8(math.MaxInt8)
	end := int8(0)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(begin)-int(end))/int(step) + 1
	final := int(begin) - iterations*int(step)

	for id, number := range IterStrict(begin, end, step, -1, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testIterStrictForwardUns(t *testing.T, step uint8) {
	begin := uint8(0)
	end := uint8(math.MaxUint8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(end)-int(begin))/int(step) + 1
	final := int(begin) + iterations*int(step)

	for id, number := range IterStrict(begin, end, step, 0, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference += int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testIterStrictBackwardUns(t *testing.T, step uint8) {
	begin := uint8(math.MaxUint8)
	end := uint8(0)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(begin)-int(end))/int(step) + 1
	final := int(begin) - iterations*int(step)

	for id, number := range IterStrict(begin, end, step, -1, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testIterStrictBackwardUnsNotOnEntireRange(t *testing.T, step uint8) {
	begin := uint8(math.MaxUint8)
	end := uint8(math.MaxUint8 / 2)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(begin)-int(end))/int(step) + 1
	final := int(begin) - iterations*int(step)

	for id, number := range IterStrict(begin, end, step, -1, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testIterStrictForwardPartial(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	breakAt := 3
	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range IterStrict(begin, end, 1, 0, nil, nil) {
		require.Equal(t, reference, int(number))
		require.Equal(t, referenceID, id)

		if int(number) == breakAt {
			break
		}

		reference++
		referenceID++
	}

	require.Equal(t, breakAt, reference)
}

func testIterStrictBackwardPartial(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	breakAt := 3
	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range IterStrict(begin, end, 1, -1, nil, nil) {
		require.Equal(t, reference, int(number))
		require.Equal(t, referenceID, id)

		if int(number) == breakAt {
			break
		}

		reference--
		referenceID++
	}

	require.Equal(t, breakAt, reference)
}

func testIterStrictIterations(
	t *testing.T,
	begin int8,
	end int8,
	step int8,
	direction int,
	expected int,
) {
	actual := 0

	for range IterStrict(begin, end, step, direction, nil, nil) {
		actual++
	}

	require.Equal(t, expected, actual)
}

func TestIterStrictSize(t *testing.T) {
	require.Equal(t, uint64(2), IterStrictSize(1, 0, 1, -1, nil, nil))
	require.Equal(t, uint64(1), IterStrictSize(1, 1, 1, -1, nil, nil))
	require.Equal(t, uint64(0), IterStrictSize(1, 2, 1, -1, nil, nil))
	require.Equal(t, uint64(0), IterStrictSize(1, 0, 1, 0, nil, nil))
	require.Equal(t, uint64(1), IterStrictSize(1, 1, 1, 0, nil, nil))
	require.Equal(t, uint64(2), IterStrictSize(1, 2, 1, 0, nil, nil))

	require.Panics(
		t,
		func() {
			_ = IterStrictSize(1, 2, -1, 0, nil, nil)
		},
	)

	require.Panics(
		t,
		func() {
			_ = IterStrictSize(1, 2, 0, 0, nil, nil)
		},
	)
}

func BenchmarkIterReference(b *testing.B) {
	number := 0

	for value := 1; value <= b.N; value++ {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkIterTwoLevelReference(b *testing.B) {
	number := 0

	for range b.N {
		for value := 1; value <= 1; value++ {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkIter(b *testing.B) {
	number := 0

	for value := range Iter(1, b.N) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkIterTwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for value := range Iter(1, 1) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkIterSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = IterSize(1, b.N)
	}

	require.NotZero(b, size)
}

func BenchmarkIterStep(b *testing.B) {
	number := 0

	for _, value := range IterStep(1, b.N, 1, nil, nil) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkIterStepTwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for _, value := range IterStep(1, 1, 1, nil, nil) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkIterStepSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = IterStepSize(1, b.N, 1, nil, nil)
	}

	require.NotZero(b, size)
}

func BenchmarkIterStrict(b *testing.B) {
	number := 0

	for _, value := range IterStrict(1, b.N, 1, 0, nil, nil) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkIterStrictTwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for _, value := range IterStrict(1, 1, 1, 0, nil, nil) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkIterStrictSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = IterStrictSize(1, b.N, 1, 0, nil, nil)
	}

	require.NotZero(b, size)
}
