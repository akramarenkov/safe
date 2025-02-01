package safe

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIterForwardSig(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference++
	}

	require.NotEqual(t, int(begin), reference)
}

func TestIterBackwardSig(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference--
	}

	require.NotEqual(t, int(begin), reference)
}

func TestIterForwardUns(t *testing.T) {
	begin := uint8(0)
	end := uint8(math.MaxUint8)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference++
	}

	require.NotEqual(t, int(begin), reference)
}

func TestIterBackwardUns(t *testing.T) {
	begin := uint8(math.MaxUint8)
	end := uint8(0)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference--
	}

	require.NotEqual(t, int(begin), reference)
}

func TestIterForwardPart(t *testing.T) {
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

func TestIterBackwardPart(t *testing.T) {
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

func TestIterSizeManually(t *testing.T) {
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

func TestIterSizeSig(t *testing.T) {
	for begin := range Iter[int8](math.MinInt8, math.MaxInt8) {
		for end := range Iter[int8](math.MinInt8, math.MaxInt8) {
			reference := uint64(0)

			for range Iter(begin, end) {
				reference++
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

func TestIterSizeUns(t *testing.T) {
	for begin := range Iter[uint8](0, math.MaxUint8) {
		for end := range Iter[uint8](0, math.MaxUint8) {
			reference := uint64(0)

			for range Iter(begin, end) {
				reference++
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

func TestIterSizeMax(t *testing.T) {
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

func TestIncSig(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)

	for number := range Inc(begin, end) {
		require.Equal(t, reference, int(number))

		reference++
	}

	require.NotEqual(t, int(begin), reference)
}

func TestIncUns(t *testing.T) {
	begin := uint8(0)
	end := uint8(math.MaxUint8)

	reference := int(begin)

	for number := range Inc(begin, end) {
		require.Equal(t, reference, int(number))

		reference++
	}

	require.NotEqual(t, int(begin), reference)
}

func TestIncPart(t *testing.T) {
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

func TestIncIterations(t *testing.T) {
	testIncIterations(t, 1, 0, 0)
	testIncIterations(t, 1, 1, 1)
	testIncIterations(t, 1, 2, 2)
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

func TestDecSig(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	reference := int(begin)

	for number := range Dec(begin, end) {
		require.Equal(t, reference, int(number))

		reference--
	}

	require.NotEqual(t, int(begin), reference)
}

func TestDecUns(t *testing.T) {
	begin := uint8(math.MaxUint8)
	end := uint8(0)

	reference := int(begin)

	for number := range Dec(begin, end) {
		require.Equal(t, reference, int(number))

		reference--
	}

	require.NotEqual(t, int(begin), reference)
}

func TestDecPart(t *testing.T) {
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

func TestDecIterations(t *testing.T) {
	testDecIterations(t, 1, 0, 2)
	testDecIterations(t, 1, 1, 1)
	testDecIterations(t, 1, 2, 0)
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
}

func testStepForwardSig(t *testing.T, step int8) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range Step(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference += int(step)
		referenceID++
	}

	require.NotZero(t, referenceID, "step: %v", step)
}

func testStepBackwardSig(t *testing.T, step int8) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range Step(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.NotZero(t, referenceID, "step: %v", step)
}

func testStepBackwardSigNotOnEntireRange(t *testing.T, step int8) {
	begin := int8(math.MaxInt8)
	end := int8(0)

	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range Step(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.NotZero(t, referenceID, "step: %v", step)
}

func testStepForwardUns(t *testing.T, step uint8) {
	begin := uint8(0)
	end := uint8(math.MaxUint8)

	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range Step(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference += int(step)
		referenceID++
	}

	require.NotZero(t, referenceID, "step: %v", step)
}

func testStepBackwardUns(t *testing.T, step uint8) {
	begin := uint8(math.MaxUint8)
	end := uint8(0)

	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range Step(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.NotZero(t, referenceID, "step: %v", step)
}

func testStepBackwardUnsNotOnEntireRange(t *testing.T, step uint8) {
	begin := uint8(math.MaxUint8)
	end := uint8(math.MaxUint8 / 2)

	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range Step(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.NotZero(t, referenceID, "step: %v", step)
}

func TestStepForwardPart(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	breakAt := 3
	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range Step(begin, end, 1) {
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

func TestStepBackwardPart(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	breakAt := 3
	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range Step(begin, end, 1) {
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

func TestStepPanic(t *testing.T) {
	func() {
		defer func() {
			require.Equal(t, ErrStepNegative, recover())
		}()

		for number := range Step(1, 2, -1) {
			_ = number
		}
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepZero, recover())
		}()

		for number := range Step(1, 2, 0) {
			_ = number
		}
	}()
}

func TestStepSizeManually(t *testing.T) {
	require.Equal(t, uint64(3), StepSize[int8](-1, 1, 1))
	require.Equal(t, uint64(3), StepSize[int8](1, -1, 1))
	require.Equal(t, uint64(6), StepSize[int8](-2, 3, 1))
	require.Equal(t, uint64(6), StepSize[int8](3, -2, 1))
	require.Equal(t, uint64(256), StepSize[int8](-128, 127, 1))

	require.Equal(t, uint64(3), StepSize[int8](-2, 5, 3))
	require.Equal(t, uint64(3), StepSize[int8](-2, 6, 3))
	require.Equal(t, uint64(4), StepSize[int8](-2, 7, 3))
	require.Equal(t, uint64(3), StepSize[int8](-128, 127, 86))
	require.Equal(t, uint64(86), StepSize[int8](-128, 127, 3))
	require.Equal(t, uint64(128), StepSize[int8](-128, 127, 2))

	require.Equal(t, uint64(3), StepSize[uint8](1, 3, 1))
	require.Equal(t, uint64(3), StepSize[uint8](3, 1, 1))
	require.Equal(t, uint64(6), StepSize[uint8](1, 6, 1))
	require.Equal(t, uint64(6), StepSize[uint8](6, 1, 1))
	require.Equal(t, uint64(256), StepSize[uint8](0, 255, 1))

	require.Equal(t, uint64(3), StepSize[uint8](2, 9, 3))
	require.Equal(t, uint64(3), StepSize[uint8](2, 10, 3))
	require.Equal(t, uint64(4), StepSize[uint8](2, 11, 3))
	require.Equal(t, uint64(3), StepSize[uint8](0, 255, 86))
	require.Equal(t, uint64(86), StepSize[uint8](0, 255, 3))
	require.Equal(t, uint64(128), StepSize[uint8](0, 255, 2))
}

func TestStepSizeSig(t *testing.T) {
	for begin := range Iter[int8](math.MinInt8, math.MaxInt8) {
		for end := range Iter[int8](math.MinInt8, math.MaxInt8) {
			for step := range Iter[int8](1, math.MaxInt8) {
				reference := uint64(0)

				for range Step(begin, end, step) {
					reference++
				}

				actual := StepSize(begin, end, step)

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

func TestStepSizeUns(t *testing.T) {
	for begin := range Iter[uint8](0, math.MaxUint8) {
		for end := range Iter[uint8](0, math.MaxUint8) {
			for step := range Iter[uint8](1, math.MaxUint8) {
				reference := uint64(0)

				for range Step(begin, end, step) {
					reference++
				}

				actual := StepSize(begin, end, step)

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

func TestStepSizeMax(t *testing.T) {
	require.Equal(
		t,
		uint64(9223372036854775808),
		StepSize[int64](math.MinInt64, math.MaxInt64, 2),
	)
	require.Equal(
		t,
		uint64(9223372036854775808),
		StepSize[int64](math.MaxInt64, math.MinInt64, 2),
	)

	require.Equal(
		t,
		uint64(9223372036854775808),
		StepSize[uint64](0, math.MaxUint64, 2),
	)
	require.Equal(
		t,
		uint64(9223372036854775808),
		StepSize[uint64](math.MaxUint64, 0, 2),
	)
	require.Equal(
		t,
		uint64(2),
		StepSize[uint64](0, math.MaxUint64, 9223372036854775808),
	)
	require.Equal(
		t,
		uint64(2),
		StepSize[uint64](math.MaxUint64, 0, 9223372036854775808),
	)

	require.Equal(
		t,
		uint64(6148914691236517206),
		StepSize[int64](math.MinInt64, math.MaxInt64, 3),
	)
	require.Equal(
		t,
		uint64(6148914691236517206),
		StepSize[int64](math.MaxInt64, math.MinInt64, 3),
	)
	require.Equal(
		t,
		uint64(3),
		StepSize[int64](math.MinInt64, math.MaxInt64, 6148914691236517206),
	)
	require.Equal(
		t,
		uint64(3),
		StepSize[int64](math.MaxInt64, math.MinInt64, 6148914691236517206),
	)

	require.Equal(
		t,
		uint64(6148914691236517206),
		StepSize[uint64](0, math.MaxUint64, 3),
	)
	require.Equal(
		t,
		uint64(6148914691236517206),
		StepSize[uint64](math.MaxUint64, 0, 3),
	)
	require.Equal(
		t,
		uint64(3),
		StepSize[uint64](0, math.MaxUint64, 6148914691236517206),
	)
	require.Equal(
		t,
		uint64(3),
		StepSize[uint64](math.MaxUint64, 0, 6148914691236517206),
	)

	require.Equal(
		t,
		uint64(4611686018427387904),
		StepSize[int64](math.MinInt64, math.MaxInt64, 4),
	)
	require.Equal(
		t,
		uint64(4611686018427387904),
		StepSize[int64](math.MaxInt64, math.MinInt64, 4),
	)
	require.Equal(
		t,
		uint64(4),
		StepSize[int64](math.MinInt64, math.MaxInt64, 4611686018427387904),
	)
	require.Equal(
		t,
		uint64(4),
		StepSize[int64](math.MaxInt64, math.MinInt64, 4611686018427387904),
	)

	require.Equal(
		t,
		uint64(4611686018427387904),
		StepSize[uint64](0, math.MaxUint64, 4),
	)
	require.Equal(
		t,
		uint64(4611686018427387904),
		StepSize[uint64](math.MaxUint64, 0, 4),
	)
	require.Equal(
		t,
		uint64(4),
		StepSize[uint64](0, math.MaxUint64, 4611686018427387904),
	)
	require.Equal(
		t,
		uint64(4),
		StepSize[uint64](math.MaxUint64, 0, 4611686018427387904),
	)

	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		StepSize[int64](math.MinInt64+2, math.MaxInt64, 1),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		StepSize[int64](math.MaxInt64, math.MinInt64+2, 1),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		StepSize[uint64](2, math.MaxUint64, 1),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64-1),
		StepSize[uint64](math.MaxUint64, 2, 1),
	)

	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[int64](math.MinInt64+1, math.MaxInt64, 1),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[int64](math.MaxInt64, math.MinInt64+1, 1),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[uint64](1, math.MaxUint64, 1),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[uint64](math.MaxUint64, 1, 1),
	)

	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[int64](math.MinInt64, math.MaxInt64, 1),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[int64](math.MaxInt64, math.MinInt64, 1),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[uint64](0, math.MaxUint64, 1),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		StepSize[uint64](math.MaxUint64, 0, 1),
	)
}

func TestStepSizePanic(t *testing.T) {
	func() {
		defer func() {
			require.Equal(t, ErrStepNegative, recover())
		}()

		_ = StepSize(1, 2, -1)
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepZero, recover())
		}()

		_ = StepSize(1, 2, 0)
	}()
}

func TestIncStep(t *testing.T) {
	for step := range Iter[int8](1, math.MaxInt8) {
		testIncStepSig(t, step)
	}

	for step := range Iter[uint8](1, math.MaxUint8) {
		testIncStepUns(t, step)
	}
}

func testIncStepSig(t *testing.T, step int8) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range IncStep(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference += int(step)
		referenceID++
	}

	require.NotZero(t, referenceID, "step: %v", step)
}

func testIncStepUns(t *testing.T, step uint8) {
	begin := uint8(0)
	end := uint8(math.MaxUint8)

	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range IncStep(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference += int(step)
		referenceID++
	}

	require.NotZero(t, referenceID, "step: %v", step)
}

func TestIncStepPart(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	breakAt := 3
	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range IncStep(begin, end, 1) {
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

func TestIncStepIterations(t *testing.T) {
	testIncStepIterations(t, 1, 0, 1, 0)
	testIncStepIterations(t, 1, 1, 1, 1)
	testIncStepIterations(t, 1, 2, 1, 2)
}

func testIncStepIterations(t *testing.T, begin, end, step int8, expected int) {
	actual := 0

	for range IncStep(begin, end, step) {
		actual++
	}

	require.Equal(t, expected, actual)
}

func TestIncStepPanic(t *testing.T) {
	func() {
		defer func() {
			require.Equal(t, ErrStepNegative, recover())
		}()

		for number := range IncStep(2, 1, -1) {
			_ = number
		}
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepZero, recover())
		}()

		for number := range IncStep(2, 1, 0) {
			_ = number
		}
	}()
}

func TestIncStepSize(t *testing.T) {
	require.Equal(t, uint64(0), IncStepSize(1, 0, 1))
	require.Equal(t, uint64(1), IncStepSize(1, 1, 1))
	require.Equal(t, uint64(2), IncStepSize(1, 2, 1))
}

func TestIncStepSizePanic(t *testing.T) {
	func() {
		defer func() {
			require.Equal(t, ErrStepNegative, recover())
		}()

		_ = IncStepSize(2, 1, -1)
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepNegative, recover())
		}()

		_ = IncStepSize(1, 2, -1)
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepZero, recover())
		}()

		_ = IncStepSize(2, 1, 0)
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepZero, recover())
		}()

		_ = IncStepSize(1, 2, 0)
	}()
}

func TestDecStep(t *testing.T) {
	for step := range Iter[int8](1, math.MaxInt8) {
		testDecStepSig(t, step)
	}

	for step := range Iter[uint8](1, math.MaxUint8) {
		testDecStepUns(t, step)
	}
}

func testDecStepSig(t *testing.T, step int8) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range DecStep(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.NotZero(t, referenceID, "step: %v", step)
}

func testDecStepUns(t *testing.T, step uint8) {
	begin := uint8(math.MaxUint8)
	end := uint8(0)

	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range DecStep(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.NotZero(t, referenceID, "step: %v", step)
}

func TestDecStepPart(t *testing.T) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	breakAt := 3
	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range DecStep(begin, end, 1) {
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

func TestDecStepIterations(t *testing.T) {
	testDecStepIterations(t, 1, 0, 1, 2)
	testDecStepIterations(t, 1, 1, 1, 1)
	testDecStepIterations(t, 1, 2, 1, 0)
}

func testDecStepIterations(t *testing.T, begin, end, step int8, expected int) {
	actual := 0

	for range DecStep(begin, end, step) {
		actual++
	}

	require.Equal(t, expected, actual)
}

func TestDecStepPanic(t *testing.T) {
	func() {
		defer func() {
			require.Equal(t, ErrStepNegative, recover())
		}()

		for number := range DecStep(1, 2, -1) {
			_ = number
		}
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepZero, recover())
		}()

		for number := range DecStep(1, 2, 0) {
			_ = number
		}
	}()
}

func TestDecStepSize(t *testing.T) {
	require.Equal(t, uint64(2), DecStepSize(1, 0, 1))
	require.Equal(t, uint64(1), DecStepSize(1, 1, 1))
	require.Equal(t, uint64(0), DecStepSize(1, 2, 1))
}

func TestDecStepSizePanic(t *testing.T) {
	func() {
		defer func() {
			require.Equal(t, ErrStepNegative, recover())
		}()

		_ = DecStepSize(1, 2, -1)
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepNegative, recover())
		}()

		_ = DecStepSize(2, 1, -1)
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepZero, recover())
		}()

		_ = DecStepSize(1, 2, 0)
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepZero, recover())
		}()

		_ = DecStepSize(2, 1, 0)
	}()
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

	for _, value := range Step(1, b.N, 1) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkStepTwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for _, value := range Step(1, 1, 1) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkStepSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = StepSize(1, b.N, 1)
	}

	require.NotZero(b, size)
}

func BenchmarkIncStep(b *testing.B) {
	number := 0

	for _, value := range IncStep(1, b.N, 1) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkIncStepTwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for _, value := range IncStep(1, 1, 1) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkIncStepSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = IncStepSize(1, b.N, 1)
	}

	require.NotZero(b, size)
}

func BenchmarkDecStep(b *testing.B) {
	number := 0

	for _, value := range DecStep(b.N, 1, 1) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkDecStepTwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for _, value := range DecStep(1, 1, 1) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkDecStepSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = DecStepSize(b.N, 1, 1)
	}

	require.NotZero(b, size)
}
