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

func TestInc(t *testing.T) {
	testIncSig(t)
	testIncUns(t)
	testIncPart(t)

	testIncIterations(t, 1, 0, 0)
	testIncIterations(t, 1, 1, 1)
	testIncIterations(t, 1, 2, 2)
}

func testIncSig(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)

	for number := range Inc(begin, end) {
		require.Equal(t, reference, int(number))

		reference++
	}

	require.Equal(t, int(end)+1, reference)
}

func testIncUns(t *testing.T) {
	begin := uint8(0)
	end := uint8(math.MaxUint8)

	reference := int(begin)

	for number := range Inc(begin, end) {
		require.Equal(t, reference, int(number))

		reference++
	}

	require.Equal(t, int(end)+1, reference)
}

func testIncPart(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	breakAt := 3
	reference := int(begin)

	for number := range Inc(begin, end) {
		require.Equal(t, reference, int(number))

		if int(number) == breakAt {
			break
		}

		reference++
	}

	require.Equal(t, breakAt, reference)
}

func testIncIterations(t *testing.T, begin, end, expected int) {
	actual := 0

	for range Inc(begin, end) {
		actual++
	}

	require.Equal(t, expected, actual)
}

func TestIncSize(t *testing.T) {
	require.Equal(t, uint64(0), IncSize(1, 0))
	require.Equal(t, uint64(1), IncSize(1, 1))
	require.Equal(t, uint64(2), IncSize(1, 2))
}

func TestDec(t *testing.T) {
	testDecSig(t)
	testDecUns(t)
	testDecPart(t)

	testDecIterations(t, 1, 0, 2)
	testDecIterations(t, 1, 1, 1)
	testDecIterations(t, 1, 2, 0)
}

func testDecSig(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	reference := int(begin)

	for number := range Dec(begin, end) {
		require.Equal(t, reference, int(number))

		reference--
	}

	require.Equal(t, int(end)-1, reference)
}

func testDecUns(t *testing.T) {
	begin := uint8(math.MaxUint8)
	end := uint8(0)

	reference := int(begin)

	for number := range Dec(begin, end) {
		require.Equal(t, reference, int(number))

		reference--
	}

	require.Equal(t, int(end)-1, reference)
}

func testDecPart(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	breakAt := 3
	reference := int(begin)

	for number := range Dec(begin, end) {
		require.Equal(t, reference, int(number))

		if int(number) == breakAt {
			break
		}

		reference--
	}

	require.Equal(t, breakAt, reference)
}

func testDecIterations(t *testing.T, begin, end, expected int) {
	actual := 0

	for range Dec(begin, end) {
		actual++
	}

	require.Equal(t, expected, actual)
}

func TestDecSize(t *testing.T) {
	require.Equal(t, uint64(2), DecSize(1, 0))
	require.Equal(t, uint64(1), DecSize(1, 1))
	require.Equal(t, uint64(0), DecSize(1, 2))
}

func TestStep(t *testing.T) {
	for step := range Iter[int8](1, math.MaxInt8) {
		testStepForwardSig(t, step)
		testStepBackwardSig(t, step)
		testStepBackwardSigNotOnEntireRange(t, step)
	}

	for step := range Iter[uint8](1, math.MaxUint8) {
		testStepForwardUns(t, step)
		testStepBackwardUns(t, step)
		testStepBackwardUnsNotOnEntireRange(t, step)
	}

	testStepForwardPart(t)
	testStepBackwardPart(t)

	require.Panics(
		t,
		func() {
			for number := range Step(1, 2, -1, nil, nil) {
				_ = number
			}
		},
	)

	require.Panics(
		t,
		func() {
			for number := range Step(1, 2, 0, nil, nil) {
				_ = number
			}
		},
	)
}

func testStepForwardSig(t *testing.T, step int8) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(end)-int(begin))/int(step) + 1
	final := int(begin) + iterations*int(step)

	for id, number := range Step(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference += int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testStepBackwardSig(t *testing.T, step int8) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(begin)-int(end))/int(step) + 1
	final := int(begin) - iterations*int(step)

	for id, number := range Step(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testStepBackwardSigNotOnEntireRange(t *testing.T, step int8) {
	begin := int8(math.MaxInt8)
	end := int8(0)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(begin)-int(end))/int(step) + 1
	final := int(begin) - iterations*int(step)

	for id, number := range Step(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testStepForwardUns(t *testing.T, step uint8) {
	begin := uint8(0)
	end := uint8(math.MaxUint8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(end)-int(begin))/int(step) + 1
	final := int(begin) + iterations*int(step)

	for id, number := range Step(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference += int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testStepBackwardUns(t *testing.T, step uint8) {
	begin := uint8(math.MaxUint8)
	end := uint8(0)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(begin)-int(end))/int(step) + 1
	final := int(begin) - iterations*int(step)

	for id, number := range Step(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testStepBackwardUnsNotOnEntireRange(t *testing.T, step uint8) {
	begin := uint8(math.MaxUint8)
	end := uint8(math.MaxUint8 / 2)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(begin)-int(end))/int(step) + 1
	final := int(begin) - iterations*int(step)

	for id, number := range Step(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testStepForwardPart(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	breakAt := 3
	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range Step(begin, end, 1, nil, nil) {
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

func testStepBackwardPart(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	breakAt := 3
	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range Step(begin, end, 1, nil, nil) {
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

func TestStepSize(t *testing.T) {
	testStepSizeMan(t)
	testStepSizeSig(t)
	testStepSizeUns(t)
	testStepSizeMax(t)

	require.Panics(
		t,
		func() {
			_ = StepSize(1, 2, -1, nil, nil)
		},
	)

	require.Panics(
		t,
		func() {
			_ = StepSize(1, 2, 0, nil, nil)
		},
	)
}

func testStepSizeMan(t *testing.T) {
	require.Equal(t, uint64(3), StepSize[int8](-1, 1, 1, nil, nil))
	require.Equal(t, uint64(3), StepSize[int8](1, -1, 1, nil, nil))
	require.Equal(t, uint64(6), StepSize[int8](-2, 3, 1, nil, nil))
	require.Equal(t, uint64(6), StepSize[int8](3, -2, 1, nil, nil))
	require.Equal(t, uint64(256), StepSize[int8](-128, 127, 1, nil, nil))

	require.Equal(t, uint64(3), StepSize[int8](-2, 5, 3, nil, nil))
	require.Equal(t, uint64(3), StepSize[int8](-2, 6, 3, nil, nil))
	require.Equal(t, uint64(4), StepSize[int8](-2, 7, 3, nil, nil))
	require.Equal(t, uint64(3), StepSize[int8](-128, 127, 86, nil, nil))
	require.Equal(t, uint64(86), StepSize[int8](-128, 127, 3, nil, nil))
	require.Equal(t, uint64(128), StepSize[int8](-128, 127, 2, nil, nil))

	require.Equal(t, uint64(3), StepSize[uint8](1, 3, 1, nil, nil))
	require.Equal(t, uint64(3), StepSize[uint8](3, 1, 1, nil, nil))
	require.Equal(t, uint64(6), StepSize[uint8](1, 6, 1, nil, nil))
	require.Equal(t, uint64(6), StepSize[uint8](6, 1, 1, nil, nil))
	require.Equal(t, uint64(256), StepSize[uint8](0, 255, 1, nil, nil))

	require.Equal(t, uint64(3), StepSize[uint8](2, 9, 3, nil, nil))
	require.Equal(t, uint64(3), StepSize[uint8](2, 10, 3, nil, nil))
	require.Equal(t, uint64(4), StepSize[uint8](2, 11, 3, nil, nil))
	require.Equal(t, uint64(3), StepSize[uint8](0, 255, 86, nil, nil))
	require.Equal(t, uint64(86), StepSize[uint8](0, 255, 3, nil, nil))
	require.Equal(t, uint64(128), StepSize[uint8](0, 255, 2, nil, nil))
}

func testStepSizeSig(t *testing.T) {
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

				actual := StepSize(begin, end, step, nil, nil)

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

func testStepSizeUns(t *testing.T) {
	for begin := range Iter[uint8](0, math.MaxUint8) {
		for end := range Iter[uint8](0, math.MaxUint8) {
			for step := range Iter[uint8](1, math.MaxUint8) {
				reference := (uint64(end)-uint64(begin))/uint64(step) + 1

				if begin > end {
					reference = (uint64(begin)-uint64(end))/uint64(step) + 1
				}

				actual := StepSize(begin, end, step, nil, nil)

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

func testStepSizeMax(t *testing.T) {
	require.Equal(
		t,
		uint64(9223372036854775808),
		StepSize[int64](math.MinInt64, math.MaxInt64, 2, nil, nil),
	)
	require.Equal(
		t,
		uint64(9223372036854775808),
		StepSize[int64](math.MaxInt64, math.MinInt64, 2, nil, nil),
	)

	require.Equal(
		t,
		uint64(9223372036854775808),
		StepSize[uint64](0, math.MaxUint64, 2, nil, nil),
	)
	require.Equal(
		t,
		uint64(9223372036854775808),
		StepSize[uint64](math.MaxUint64, 0, 2, nil, nil),
	)
	require.Equal(
		t,
		uint64(2),
		StepSize[uint64](0, math.MaxUint64, 9223372036854775808, nil, nil),
	)
	require.Equal(
		t,
		uint64(2),
		StepSize[uint64](math.MaxUint64, 0, 9223372036854775808, nil, nil),
	)

	require.Equal(
		t,
		uint64(6148914691236517206),
		StepSize[int64](math.MinInt64, math.MaxInt64, 3, nil, nil),
	)
	require.Equal(
		t,
		uint64(6148914691236517206),
		StepSize[int64](math.MaxInt64, math.MinInt64, 3, nil, nil),
	)
	require.Equal(
		t,
		uint64(3),
		StepSize[int64](math.MinInt64, math.MaxInt64, 6148914691236517206, nil, nil),
	)
	require.Equal(
		t,
		uint64(3),
		StepSize[int64](math.MaxInt64, math.MinInt64, 6148914691236517206, nil, nil),
	)

	require.Equal(
		t,
		uint64(6148914691236517206),
		StepSize[uint64](0, math.MaxUint64, 3, nil, nil),
	)
	require.Equal(
		t,
		uint64(6148914691236517206),
		StepSize[uint64](math.MaxUint64, 0, 3, nil, nil),
	)
	require.Equal(
		t,
		uint64(3),
		StepSize[uint64](0, math.MaxUint64, 6148914691236517206, nil, nil),
	)
	require.Equal(
		t,
		uint64(3),
		StepSize[uint64](math.MaxUint64, 0, 6148914691236517206, nil, nil),
	)

	require.Equal(
		t,
		uint64(4611686018427387904),
		StepSize[int64](math.MinInt64, math.MaxInt64, 4, nil, nil),
	)
	require.Equal(
		t,
		uint64(4611686018427387904),
		StepSize[int64](math.MaxInt64, math.MinInt64, 4, nil, nil),
	)
	require.Equal(
		t,
		uint64(4),
		StepSize[int64](math.MinInt64, math.MaxInt64, 4611686018427387904, nil, nil),
	)
	require.Equal(
		t,
		uint64(4),
		StepSize[int64](math.MaxInt64, math.MinInt64, 4611686018427387904, nil, nil),
	)

	require.Equal(
		t,
		uint64(4611686018427387904),
		StepSize[uint64](0, math.MaxUint64, 4, nil, nil),
	)
	require.Equal(
		t,
		uint64(4611686018427387904),
		StepSize[uint64](math.MaxUint64, 0, 4, nil, nil),
	)
	require.Equal(
		t,
		uint64(4),
		StepSize[uint64](0, math.MaxUint64, 4611686018427387904, nil, nil),
	)
	require.Equal(
		t,
		uint64(4),
		StepSize[uint64](math.MaxUint64, 0, 4611686018427387904, nil, nil),
	)

	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		StepSize[int64](math.MinInt64+2, math.MaxInt64, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		StepSize[int64](math.MaxInt64, math.MinInt64+2, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		StepSize[uint64](2, math.MaxUint64, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		StepSize[uint64](math.MaxUint64, 2, 1, nil, nil),
	)

	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[int64](math.MinInt64+1, math.MaxInt64, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[int64](math.MaxInt64, math.MinInt64+1, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[uint64](1, math.MaxUint64, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[uint64](math.MaxUint64, 1, 1, nil, nil),
	)

	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[int64](math.MinInt64, math.MaxInt64, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[int64](math.MaxInt64, math.MinInt64, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[uint64](0, math.MaxUint64, 1, nil, nil),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[uint64](math.MaxUint64, 0, 1, nil, nil),
	)
}

func TestIncStep(t *testing.T) {
	for step := range Iter[int8](1, math.MaxInt8) {
		testIncStepSig(t, step)
	}

	for step := range Iter[uint8](1, math.MaxUint8) {
		testIncStepUns(t, step)
	}

	testIncStepPart(t)

	testIncStepIterations(t, 1, 0, 1, 0)
	testIncStepIterations(t, 1, 1, 1, 1)
	testIncStepIterations(t, 1, 2, 1, 2)

	require.Panics(
		t,
		func() {
			for number := range IncStep(1, 2, -1, nil, nil) {
				_ = number
			}
		},
	)

	require.Panics(
		t,
		func() {
			for number := range IncStep(1, 2, 0, nil, nil) {
				_ = number
			}
		},
	)
}

func testIncStepSig(t *testing.T, step int8) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(end)-int(begin))/int(step) + 1
	final := int(begin) + iterations*int(step)

	for id, number := range IncStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference += int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testIncStepUns(t *testing.T, step uint8) {
	begin := uint8(0)
	end := uint8(math.MaxUint8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(end)-int(begin))/int(step) + 1
	final := int(begin) + iterations*int(step)

	for id, number := range IncStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference += int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testIncStepPart(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	breakAt := 3
	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range IncStep(begin, end, 1, nil, nil) {
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

func testIncStepIterations(t *testing.T, begin, end, step int8, expected int) {
	actual := 0

	for range IncStep(begin, end, step, nil, nil) {
		actual++
	}

	require.Equal(t, expected, actual)
}

func TestIncStepSize(t *testing.T) {
	require.Equal(t, uint64(0), IncStepSize(1, 0, 1, nil, nil))
	require.Equal(t, uint64(1), IncStepSize(1, 1, 1, nil, nil))
	require.Equal(t, uint64(2), IncStepSize(1, 2, 1, nil, nil))

	require.Panics(
		t,
		func() {
			_ = IncStepSize(1, 2, -1, nil, nil)
		},
	)
	require.Panics(
		t,
		func() {
			_ = IncStepSize(2, 1, -1, nil, nil)
		},
	)

	require.Panics(
		t,
		func() {
			_ = IncStepSize(1, 2, 0, nil, nil)
		},
	)
	require.Panics(
		t,
		func() {
			_ = IncStepSize(2, 1, 0, nil, nil)
		},
	)
}

func TestDecStep(t *testing.T) {
	for step := range Iter[int8](1, math.MaxInt8) {
		testDecStepSig(t, step)
	}

	for step := range Iter[uint8](1, math.MaxUint8) {
		testDecStepUns(t, step)
	}

	testDecStepPart(t)

	testDecStepIterations(t, 1, 0, 1, 2)
	testDecStepIterations(t, 1, 1, 1, 1)
	testDecStepIterations(t, 1, 2, 1, 0)

	require.Panics(
		t,
		func() {
			for number := range DecStep(2, 1, -1, nil, nil) {
				_ = number
			}
		},
	)

	require.Panics(
		t,
		func() {
			for number := range DecStep(2, 1, 0, nil, nil) {
				_ = number
			}
		},
	)
}

func testDecStepSig(t *testing.T, step int8) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(begin)-int(end))/int(step) + 1
	final := int(begin) - iterations*int(step)

	for id, number := range DecStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testDecStepUns(t *testing.T, step uint8) {
	begin := uint8(math.MaxUint8)
	end := uint8(0)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(begin)-int(end))/int(step) + 1
	final := int(begin) - iterations*int(step)

	for id, number := range DecStep(begin, end, step, nil, nil) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func testDecStepPart(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	breakAt := 3
	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range DecStep(begin, end, 1, nil, nil) {
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

func testDecStepIterations(t *testing.T, begin, end, step int8, expected int) {
	actual := 0

	for range DecStep(begin, end, step, nil, nil) {
		actual++
	}

	require.Equal(t, expected, actual)
}

func TestDecStepSize(t *testing.T) {
	require.Equal(t, uint64(2), DecStepSize(1, 0, 1, nil, nil))
	require.Equal(t, uint64(1), DecStepSize(1, 1, 1, nil, nil))
	require.Equal(t, uint64(0), DecStepSize(1, 2, 1, nil, nil))

	require.Panics(
		t,
		func() {
			_ = DecStepSize(1, 2, -1, nil, nil)
		},
	)

	require.Panics(
		t,
		func() {
			_ = DecStepSize(2, 1, -1, nil, nil)
		},
	)

	require.Panics(
		t,
		func() {
			_ = DecStepSize(1, 2, 0, nil, nil)
		},
	)

	require.Panics(
		t,
		func() {
			_ = DecStepSize(2, 1, 0, nil, nil)
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

func BenchmarkInc(b *testing.B) {
	number := 0

	for value := range Inc(1, b.N) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkIncTwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for value := range Inc(1, 1) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkIncSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = IncSize(1, b.N)
	}

	require.NotZero(b, size)
}

func BenchmarkDec(b *testing.B) {
	number := 0

	for value := range Dec(b.N, 1) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkDecTwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for value := range Dec(1, 1) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkDecSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = DecSize(b.N, 1)
	}

	require.NotZero(b, size)
}

func BenchmarkStep(b *testing.B) {
	number := 0

	for _, value := range Step(1, b.N, 1, nil, nil) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkStepTwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for _, value := range Step(1, 1, 1, nil, nil) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkStepSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = StepSize(1, b.N, 1, nil, nil)
	}

	require.NotZero(b, size)
}

func BenchmarkIncStep(b *testing.B) {
	number := 0

	for _, value := range IncStep(1, b.N, 1, nil, nil) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkIncStepTwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for _, value := range IncStep(1, 1, 1, nil, nil) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkIncStepSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = IncStepSize(1, b.N, 1, nil, nil)
	}

	require.NotZero(b, size)
}

func BenchmarkDecStep(b *testing.B) {
	number := 0

	for _, value := range DecStep(b.N, 1, 1, nil, nil) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkDecStepTwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for _, value := range DecStep(1, 1, 1, nil, nil) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkDecStepSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = DecStepSize(b.N, 1, 1, nil, nil)
	}

	require.NotZero(b, size)
}
