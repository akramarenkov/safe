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
	expectedReference := int(begin) + expectedIterations*int(step)

	for number := range IterStep(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)

		iterations++
		reference += int(step)
	}

	require.Equal(t, expectedIterations, iterations, "step: %v", step)
	require.Equal(t, expectedReference, reference, "step: %v", step)
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

func BenchmarkIter(b *testing.B) {
	number := 0

	for value := range Iter(1, b.N) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkIterStep(b *testing.B) {
	number := 0

	for value := range IterStep(1, b.N, 1) {
		number = value
	}

	require.NotZero(b, number)
}
