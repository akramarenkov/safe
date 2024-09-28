package safe

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIter(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)

	for number := range Iter(begin, end) {
		require.Equal(t, reference, int(number))

		reference++
	}

	require.Equal(t, int(end)+1, reference)
}

func TestIterSize(t *testing.T) {
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
		testIterStep(t, step)
	}
}

func testIterStep(t *testing.T, step int8) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	iterations := 0
	reference := int(begin)

	expectedIterations := (int(end)-int(begin))/int(step) + 1
	expectedFinalReference := int(begin) + expectedIterations*int(step)

	for number := range IterStep(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)

		iterations++
		reference += int(step)
	}

	require.Equal(t, expectedIterations, iterations, "step: %v", step)
	require.Equal(t, expectedFinalReference, reference, "step: %v", step)
}

func TestIterStepPanic(t *testing.T) {
	func() {
		defer func() {
			require.Equal(t, ErrIterStepNegative, recover())
		}()

		for number := range IterStep(1, 2, -1) {
			_ = number
		}
	}()

	func() {
		defer func() {
			require.Equal(t, ErrIterStepZero, recover())
		}()

		for number := range IterStep(1, 2, 0) {
			_ = number
		}
	}()
}

func TestIterStepSize(t *testing.T) {
	require.Equal(
		t,
		uint64(4611686018427387904),
		IterStepSize[int64](math.MinInt64, math.MaxInt64, 4),
	)
	require.Equal(
		t,
		uint64(4611686018427387904),
		IterStepSize[int64](math.MaxInt64, math.MinInt64, 4),
	)

	require.Equal(
		t,
		uint64(4611686018427387904),
		IterStepSize[uint64](0, math.MaxUint64, 4),
	)
	require.Equal(
		t,
		uint64(4611686018427387904),
		IterStepSize[uint64](math.MaxUint64, 0, 4),
	)

	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterStepSize[int64](math.MinInt64, math.MaxInt64, 1),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterStepSize[int64](math.MaxInt64, math.MinInt64, 1),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterStepSize[uint64](0, math.MaxUint64, 1),
	)
	require.Equal(
		t,
		uint64(math.MaxUint64),
		IterStepSize[uint64](math.MaxUint64, 0, 1),
	)
}

func TestIterStepSizePanic(t *testing.T) {
	func() {
		defer func() {
			require.Equal(t, ErrIterStepNegative, recover())
		}()

		_ = IterStepSize(1, 2, -1)
	}()

	func() {
		defer func() {
			require.Equal(t, ErrIterStepZero, recover())
		}()

		_ = IterStepSize(1, 2, 0)
	}()
}

func BenchmarkIter(b *testing.B) {
	number := 0

	for value := range Iter(1, b.N) {
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

	for value := range IterStep(1, b.N, 1) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkIterStepSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = IterStepSize(0, b.N, 1)
	}

	require.NotZero(b, size)
}
