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
	require.Equal(t, uint64(3), IterSize[int8](-1, 1))
	require.Equal(t, uint64(3), IterSize[uint8](1, 3))
}

func TestStep(t *testing.T) {
	for step := range Iter[int8](1, math.MaxInt8) {
		testStep(t, step)
	}

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

func testStep(t *testing.T, step int8) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(end)-int(begin))/int(step) + 1
	final := int(begin) + iterations*int(step)

	for id, number := range Step(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference += int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func TestStepSize(t *testing.T) {
	require.Equal(t, uint64(3), StepSize[int8](-1, 1, 1))
	require.Equal(t, uint64(3), StepSize[uint8](1, 3, 1))

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

func TestInc(t *testing.T) {
	for step := range Iter[int8](1, math.MaxInt8) {
		testInc(t, step)
	}

	func() {
		defer func() {
			require.Equal(t, ErrStepNegative, recover())
		}()

		for number := range Inc(2, 1, -1) {
			_ = number
		}
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepZero, recover())
		}()

		for number := range Inc(2, 1, 0) {
			_ = number
		}
	}()
}

func testInc(t *testing.T, step int8) {
	begin := int8(math.MinInt8)
	end := int8(math.MaxInt8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(end)-int(begin))/int(step) + 1
	final := int(begin) + iterations*int(step)

	for id, number := range Inc(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference += int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func TestIncSize(t *testing.T) {
	require.Equal(t, uint64(0), IncSize(1, 0, 1))
	require.Equal(t, uint64(1), IncSize(1, 1, 1))
	require.Equal(t, uint64(2), IncSize(1, 2, 1))

	func() {
		defer func() {
			require.Equal(t, ErrStepNegative, recover())
		}()

		_ = IncSize(2, 1, -1)
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepZero, recover())
		}()

		_ = IncSize(2, 1, 0)
	}()
}

func TestDec(t *testing.T) {
	for step := range Iter[int8](1, math.MaxInt8) {
		testDec(t, step)
	}

	func() {
		defer func() {
			require.Equal(t, ErrStepNegative, recover())
		}()

		for number := range Dec(1, 2, -1) {
			_ = number
		}
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepZero, recover())
		}()

		for number := range Dec(1, 2, 0) {
			_ = number
		}
	}()
}

func testDec(t *testing.T, step int8) {
	begin := int8(math.MaxInt8)
	end := int8(math.MinInt8)

	reference := int(begin)
	referenceID := uint64(0)

	iterations := (int(begin)-int(end))/int(step) + 1
	final := int(begin) - iterations*int(step)

	for id, number := range Dec(begin, end, step) {
		require.Equal(t, reference, int(number), "step: %v", step)
		require.Equal(t, referenceID, id, "step: %v", step)

		reference -= int(step)
		referenceID++
	}

	require.Equal(t, final, reference, "step: %v", step)
}

func TestDecSize(t *testing.T) {
	require.Equal(t, uint64(2), DecSize(1, 0, 1))
	require.Equal(t, uint64(1), DecSize(1, 1, 1))
	require.Equal(t, uint64(0), DecSize(1, 2, 1))

	func() {
		defer func() {
			require.Equal(t, ErrStepNegative, recover())
		}()

		_ = DecSize(1, 2, -1)
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepZero, recover())
		}()

		_ = DecSize(1, 2, 0)
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

func BenchmarkInc(b *testing.B) {
	number := 0

	for _, value := range Inc(1, b.N, 1) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkIncTwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for _, value := range Inc(1, 1, 1) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkIncSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = IncSize(1, b.N, 1)
	}

	require.NotZero(b, size)
}

func BenchmarkDec(b *testing.B) {
	number := 0

	for _, value := range Dec(b.N, 1, 1) {
		number = value
	}

	require.NotZero(b, number)
}

func BenchmarkDecTwoLevel(b *testing.B) {
	number := 0

	for range b.N {
		for _, value := range Dec(1, 1, 1) {
			number = value
		}
	}

	require.NotZero(b, number)
}

func BenchmarkDecSize(b *testing.B) {
	size := uint64(0)

	for range b.N {
		size = DecSize(b.N, 1, 1)
	}

	require.NotZero(b, size)
}
