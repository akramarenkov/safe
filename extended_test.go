package safe

import (
	"math"
	"slices"
	"testing"

	"github.com/akramarenkov/safe/internal/inspect"

	"github.com/stretchr/testify/require"
)

func TestAddM(t *testing.T) {
	_, err := AddM[int]()
	require.Error(t, err)

	sum, err := AddM(1)
	require.NoError(t, err)
	require.Equal(t, 1, sum)

	sum, err = AddM(1, 2)
	require.NoError(t, err)
	require.Equal(t, 3, sum)

	sum, err = AddM(1, 2, 3)
	require.NoError(t, err)
	require.Equal(t, 6, sum)

	sum, err = AddM(1, 2, 3, 4)
	require.NoError(t, err)
	require.Equal(t, 10, sum)

	sum, err = AddM(1, 2, 3, 4, 5)
	require.NoError(t, err)
	require.Equal(t, 15, sum)

	testAddMInt(t)
	testAddMUint(t)
}

func testAddMInt(t *testing.T) {
	opts := inspect.Opts[int8]{
		LoopsQuantity: 2,

		Inspected: AddM[int8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)

	opts = inspect.Opts[int8]{
		LoopsQuantity: 3,

		Inspected: AddM[int8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] + args[2], nil
		},
	}

	result, err = opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func testAddMUint(t *testing.T) {
	opts := inspect.Opts[uint8]{
		LoopsQuantity: 2,

		Inspected: AddM[uint8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)

	opts = inspect.Opts[uint8]{
		LoopsQuantity: 3,

		Inspected: AddM[uint8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] + args[2], nil
		},
	}

	result, err = opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func TestAddM4(t *testing.T) {
	testAddM4Int(t)
	testAddM4Uint(t)
}

func testAddM4Int(t *testing.T) {
	opts := inspect.Opts4[int8]{
		Inspected: func(first, second, third, fourth int8) (int8, error) {
			return AddM(first, second, third, fourth)
		},
		Reference4: func(first, second, third, fourth int64) (int64, error) {
			return first + second + third + fourth, nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func testAddM4Uint(t *testing.T) {
	opts := inspect.Opts4[uint8]{
		Inspected: func(first, second, third, fourth uint8) (uint8, error) {
			return AddM(first, second, third, fourth)
		},
		Reference4: func(first, second, third, fourth int64) (int64, error) {
			return first + second + third + fourth, nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func TestAddT(t *testing.T) {
	testAddTInt(t)
	testAddTUint(t)
}

func testAddTInt(t *testing.T) {
	opts := inspect.Opts[int8]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return AddT(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] + args[2], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func testAddTUint(t *testing.T) {
	opts := inspect.Opts[uint8]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return AddT(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] + args[2], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func TestAddUM(t *testing.T) {
	_, err := AddUM[uint]()
	require.Error(t, err)

	opts := inspect.Opts[uint8]{
		LoopsQuantity: 3,

		Inspected: AddUM[uint8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] + args[1] + args[2], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func TestSubT(t *testing.T) {
	testSubTInt(t)
	testSubTUint(t)
}

func testSubTInt(t *testing.T) {
	opts := inspect.Opts[int8]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return SubT(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1] - args[2], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func testSubTUint(t *testing.T) {
	opts := inspect.Opts[uint8]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return SubT(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1] - args[2], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func TestSubUM(t *testing.T) {
	diff, err := SubUM(uint(math.MaxUint))
	require.NoError(t, err)
	require.Equal(t, uint(math.MaxUint), diff)

	opts := inspect.Opts[uint8]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return SubUM(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] - args[1] - args[2], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func TestMulM(t *testing.T) {
	_, err := MulM[int]()
	require.Error(t, err)

	testMulMInt(t)
	testMulMUint(t)
}

func testMulMInt(t *testing.T) {
	opts := inspect.Opts[int8]{
		LoopsQuantity: 3,

		Inspected: MulM[int8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func testMulMUint(t *testing.T) {
	opts := inspect.Opts[uint8]{
		LoopsQuantity: 3,

		Inspected: MulM[uint8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func TestCmpMulM(t *testing.T) {
	factors := []int{15, 0, 27, -1, -5}

	slices.SortFunc(factors, cmpMulM)
	require.Equal(t, []int{-1, -5, 0, 15, 27}, factors)
}

func TestMulT(t *testing.T) {
	testMulTInt(t)
	testMulTUint(t)
}

func testMulTInt(t *testing.T) {
	opts := inspect.Opts[int8]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return MulT(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func testMulTUint(t *testing.T) {
	opts := inspect.Opts[uint8]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return MulT(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func TestMulUM(t *testing.T) {
	_, err := MulUM[uint]()
	require.Error(t, err)

	opts := inspect.Opts[uint8]{
		LoopsQuantity: 3,

		Inspected: MulUM[uint8],
		Reference: func(args ...int64) (int64, error) {
			return args[0] * args[1] * args[2], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
}

func TestDivM(t *testing.T) {
	quotient, err := DivM(uint(math.MaxUint))
	require.NoError(t, err)
	require.Equal(t, uint(math.MaxUint), quotient)

	testDivMInt(t)
	testDivMUint(t)
}

func testDivMInt(t *testing.T) {
	opts := inspect.Opts[int8]{
		LoopsQuantity: 3,

		Inspected: func(args ...int8) (int8, error) {
			return DivM(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[1] == 0 || args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return args[0] / args[1] / args[2], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.NotZero(t, result.Overflows)
	require.NotZero(t, result.ReferenceFaults)
}

func testDivMUint(t *testing.T) {
	opts := inspect.Opts[uint8]{
		LoopsQuantity: 3,

		Inspected: func(args ...uint8) (uint8, error) {
			return DivM(args[0], args[1], args[2])
		},
		Reference: func(args ...int64) (int64, error) {
			if args[1] == 0 || args[2] == 0 {
				return 0, ErrDivisionByZero
			}

			return args[0] / args[1] / args[2], nil
		},
	}

	result, err := opts.Do()
	require.NoError(t, err)
	require.NoError(
		t,
		result.Conclusion,
		"reference: %v, actual: %v, args: %v, err: %v",
		result.Reference,
		result.Actual,
		result.Args,
		result.Err,
	)
	require.NotZero(t, result.NoOverflows)
	require.Zero(t, result.Overflows)
	require.NotZero(t, result.ReferenceFaults)
}

func TestPow10(t *testing.T) {
	testPow10Manually(t)
	testPow10Diff(t)
}

func testPow10Manually(t *testing.T) {
	product, err := Pow10[uint64](-3)
	require.NoError(t, err)
	require.Equal(t, uint64(0), product)

	product, err = Pow10[uint64](-2)
	require.NoError(t, err)
	require.Equal(t, uint64(0), product)

	product, err = Pow10[uint64](-1)
	require.NoError(t, err)
	require.Equal(t, uint64(0), product)

	product, err = Pow10[uint64](0)
	require.NoError(t, err)
	require.Equal(t, uint64(1), product)

	product, err = Pow10[uint64](1)
	require.NoError(t, err)
	require.Equal(t, uint64(10), product)

	product, err = Pow10[uint64](2)
	require.NoError(t, err)
	require.Equal(t, uint64(100), product)

	product, err = Pow10[uint64](3)
	require.NoError(t, err)
	require.Equal(t, uint64(1000), product)

	product, err = Pow10[uint64](19)
	require.NoError(t, err)
	require.Equal(t, uint64(1e19), product)

	product, err = Pow10[uint64](20)
	require.Error(t, err)
	require.Equal(t, uint64(0), product)

	product, err = Pow10[uint64](21)
	require.Error(t, err)
	require.Equal(t, uint64(0), product)
}

func testPow10Diff(t *testing.T) {
	for power := 1; power <= 19; power++ {
		previous, err := Pow10[uint64](power - 1)
		require.NoError(t, err, "power: %v", power)

		current, err := Pow10[uint64](power)
		require.NoError(t, err, "power: %v", power)

		require.Equal(t, uint64(10), current/previous, "power: %v", power)
		require.Equal(t, uint64(0), current%previous, "power: %v", power)
	}
}

func TestPow(t *testing.T) {
	faults := 0
	successful := 0

	// Is used int32 because in its value range float64 does not lose the precision of
	// the integer part and comparison of reference and tested values ​​can be done simply
	maxInt32 := float64(math.MaxInt32)
	minInt32 := float64(math.MinInt32)

	for base := int32(math.MinInt8); base <= math.MaxInt8; base++ {
		for power := int32(math.MinInt8); power <= math.MaxInt8; power++ {
			reference := math.Pow(float64(base), float64(power))
			require.False(t, math.IsNaN(reference))

			// To ensure that overflow conditions are satisfied correctly when
			// obtaining reference values ​​with non-zero fractional parts and close
			// to maximum/minimum int32 values, reference values ​​in these areas
			// are checked separately
			if reference >= maxInt32-2 && reference <= maxInt32+2 {
				require.InDelta(t, maxInt32+1, reference, 0)

				// reference > maxInt32 == true
				require.Greater(t, reference, maxInt32)
			}

			if reference >= minInt32-2 && reference <= minInt32+2 {
				require.InDelta(t, minInt32, reference, 0)

				// reference < minInt32 == false
				require.GreaterOrEqual(t, reference, minInt32)
			}

			product, err := Pow(base, power)

			// Converting reference to any integer type can and will overflow,
			// so the comparison is done in float64
			if reference > maxInt32 || reference < minInt32 {
				require.Error(
					t,
					err,
					"base: %v, power: %v, product: %v, reference: %f",
					base,
					power,
					product,
					reference,
				)

				faults++

				continue
			}

			successful++

			require.NoError(
				t,
				err,
				"base: %v, power: %v, product: %v, reference: %f",
				base,
				power,
				product,
				reference,
			)

			require.Equal(
				t,
				int32(reference),
				product,
				"base: %v, power: %v, reference: %f",
				base,
				power,
				reference,
			)

			require.InDelta(
				t,
				reference,
				float64(product),
				0.5,
				"base: %v, power: %v, product: %v",
				base,
				power,
				product,
			)
		}
	}

	require.NotZero(t, faults)
	require.NotZero(t, successful)
}

func BenchmarkAddMReference(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	for range b.N {
		for first := -3; first <= 3; first++ {
			for second := -3; second <= 3; second++ {
				for third := -3; third <= 3; third++ {
					sum = first + second + third
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddM(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	for range b.N {
		for first := -3; first <= 3; first++ {
			for second := -3; second <= 3; second++ {
				for third := -3; third <= 3; third++ {
					sum, _ = AddM(first, second, third)
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddM2Args(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	for range b.N {
		for first := -3; first <= 3; first++ {
			for second := -3; second <= 3; second++ {
				sum, _ = AddM(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddT(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := 0

	for range b.N {
		for first := -3; first <= 3; first++ {
			for second := -3; second <= 3; second++ {
				for third := -3; third <= 3; third++ {
					sum, _ = AddT(first, second, third)
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddUM(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := uint(0)

	for range b.N {
		for first := uint(0); first <= 6; first++ {
			for second := uint(0); second <= 6; second++ {
				for third := uint(0); third <= 6; third++ {
					sum, _ = AddUM(first, second, third)
				}
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkAddUM2Args(b *testing.B) {
	// sum and require is used to prevent compiler optimizations
	sum := uint(0)

	for range b.N {
		for first := uint(0); first <= 6; first++ {
			for second := uint(0); second <= 6; second++ {
				sum, _ = AddUM(first, second)
			}
		}
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkSubT(b *testing.B) {
	diff := 0

	// b.N, diff and require is used to prevent compiler optimizations
	for range b.N {
		diff, _ = SubT(b.N, 3, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkSubUM(b *testing.B) {
	diff := uint(0)

	// b.N, diff and require is used to prevent compiler optimizations
	for range b.N {
		diff, _ = SubUM(uint(b.N), 3, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, diff)
}

func BenchmarkMulT(b *testing.B) {
	product := 0

	// b.N, product and require is used to prevent compiler optimizations
	for range b.N {
		product, _ = MulT(b.N, 3, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMulM(b *testing.B) {
	product := 0

	// b.N, product and require is used to prevent compiler optimizations
	for range b.N {
		product, _ = MulM(b.N, 3, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkMulUM(b *testing.B) {
	product := uint(0)

	// b.N, product and require is used to prevent compiler optimizations
	for range b.N {
		product, _ = MulUM(uint(b.N), 3, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkDivM(b *testing.B) {
	quotient := 0

	// b.N, quotient and require is used to prevent compiler optimizations
	for range b.N {
		quotient, _ = DivM(b.N, 3, 3)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, quotient)
}

func BenchmarkPow10Reference(b *testing.B) {
	product := float64(0)

	// product and require is used to prevent compiler optimizations
	for range b.N {
		product = math.Pow10(19)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkPow10(b *testing.B) {
	product := uint64(0)

	// product and require is used to prevent compiler optimizations
	for range b.N {
		product, _ = Pow10[uint64](19)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkPowReference(b *testing.B) {
	product := float64(0)

	// product and require is used to prevent compiler optimizations
	for range b.N {
		product = math.Pow(14, 14)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}

func BenchmarkPow(b *testing.B) {
	product := uint64(0)

	// product and require is used to prevent compiler optimizations
	for range b.N {
		product, _ = Pow(uint64(14), 14)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, product)
}
