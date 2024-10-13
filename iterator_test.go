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

func TestIter2(t *testing.T) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)
	referenceID := uint64(0)

	for id, number := range Iter2(begin, end) {
		require.Equal(t, reference, int(number))
		require.Equal(t, referenceID, id)

		reference++
		referenceID++
	}

	require.Equal(t, int(end)+1, reference)
}

func TestIterSize(t *testing.T) {
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

func TestIterStep(t *testing.T) {
	for step := range Iter[int8](1, math.MaxInt8) {
		testIterStep(t, step)
	}
}

func testIterStep(t *testing.T, step int8) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)

	iterations := (int(end)-int(begin))/int(step) + 1
	final := int(begin) + iterations*int(step)

	for number := range IterStep(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)

		reference += int(step)
	}

	require.Equal(t, final, reference, "step: %v", step)
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

func TestIterStep2(t *testing.T) {
	for step := range Iter[int8](1, math.MaxInt8) {
		testIterStep2(t, step)
	}
}

func testIterStep2(t *testing.T, step int8) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(end)-int(begin))/int(step) + 1
	final := int(begin) + iterations*int(step)

	for id, number := range IterStep2(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference += int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func TestIterStep2Panic(t *testing.T) {
	func() {
		defer func() {
			require.Equal(t, ErrIterStepNegative, recover())
		}()

		for number := range IterStep2(1, 2, -1) {
			_ = number
		}
	}()

	func() {
		defer func() {
			require.Equal(t, ErrIterStepZero, recover())
		}()

		for number := range IterStep2(1, 2, 0) {
			_ = number
		}
	}()
}

func TestIterStepSize(t *testing.T) {
	require.Equal(t, uint64(3), IterStepSize[int8](-1, 1, 1))
	require.Equal(t, uint64(3), IterStepSize[int8](1, -1, 1))
	require.Equal(t, uint64(6), IterStepSize[int8](-2, 3, 1))
	require.Equal(t, uint64(6), IterStepSize[int8](3, -2, 1))
	require.Equal(t, uint64(256), IterStepSize[int8](-128, 127, 1))

	require.Equal(t, uint64(3), IterStepSize[int8](-2, 5, 3))
	require.Equal(t, uint64(3), IterStepSize[int8](-2, 6, 3))
	require.Equal(t, uint64(4), IterStepSize[int8](-2, 7, 3))
	require.Equal(t, uint64(3), IterStepSize[int8](-128, 127, 86))
	require.Equal(t, uint64(86), IterStepSize[int8](-128, 127, 3))
	require.Equal(t, uint64(128), IterStepSize[int8](-128, 127, 2))

	require.Equal(t, uint64(3), IterStepSize[uint8](1, 3, 1))
	require.Equal(t, uint64(3), IterStepSize[uint8](3, 1, 1))
	require.Equal(t, uint64(6), IterStepSize[uint8](1, 6, 1))
	require.Equal(t, uint64(6), IterStepSize[uint8](6, 1, 1))
	require.Equal(t, uint64(256), IterStepSize[uint8](0, 255, 1))

	require.Equal(t, uint64(3), IterStepSize[uint8](2, 9, 3))
	require.Equal(t, uint64(3), IterStepSize[uint8](2, 10, 3))
	require.Equal(t, uint64(4), IterStepSize[uint8](2, 11, 3))
	require.Equal(t, uint64(3), IterStepSize[uint8](0, 255, 86))
	require.Equal(t, uint64(86), IterStepSize[uint8](0, 255, 3))
	require.Equal(t, uint64(128), IterStepSize[uint8](0, 255, 2))
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

func BenchmarkIterReference(b *testing.B) {
	number := 0

	for value := 1; value <= b.N; value++ {
		number = value
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
		for value := range Iter(1, b.N) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkIter2(b *testing.B) {
	number := 0

	for _, value := range Iter2(1, b.N) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkIter2TwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for _, value := range Iter2(1, b.N) {
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

	for value := range IterStep(1, b.N, 1) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkIterStepTwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for value := range IterStep(1, b.N, 1) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkIterStep2(b *testing.B) {
	number := 0

	for _, value := range IterStep2(1, b.N, 1) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkIterStep2TwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for _, value := range IterStep2(1, b.N, 1) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkIterStepSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = IterStepSize(1, b.N, 1)
	}

	require.NotZero(b, size)
}
