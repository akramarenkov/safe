package safe

import (
	"iter"

	"github.com/akramarenkov/safe/intspec"
	"golang.org/x/exp/constraints"
)

// A range iterator for safely (without infinite loops due to counter overflow)
// iterating over integer values from begin to end inclusive with a step one.
//
// If begin is greater than end, the return value will be decremented, otherwise it
// will be incremented.
func Iter[Type constraints.Integer](begin, end Type) iter.Seq[Type] {
	iterator := func(yield func(Type) bool) {
		if begin > end {
			for number := begin; number > end; number-- {
				if !yield(number) {
					return
				}
			}

			yield(end)

			return
		}

		for number := begin; number < end; number++ {
			if !yield(number) {
				return
			}
		}

		yield(end)
	}

	return iterator
}

// Calculates the number of iterations when using [Iter]. The return value
// is intended to be used as the size parameter in the make call, so, and because
// the maximum possible number of iterations is one more than the maximum value for
// uint64, the return value is truncated to the maximum value for uint64 if the
// calculated value exceeds it.
func IterSize[Type constraints.Integer](begin, end Type) uint64 {
	beginU64 := u64(begin)
	endU64 := u64(end)

	size := endU64 - beginU64

	if beginU64 > endU64 {
		size = beginU64 - endU64
	}

	// begin < 0 && end > 0 || begin > 0 && end < 0
	if begin^end < 0 {
		size = endU64 + beginU64
	}

	if size == intspec.MaxUint64 {
		return size
	}

	return size + 1
}

// A range iterator for safely (without infinite loops due to counter overflow)
// iterating over integer values from begin to end inclusive towards increase with
// a step one.
//
// If begin is greater than end, then no one iteration of the loop will occur.
func Inc[Type constraints.Integer](begin, end Type) iter.Seq[Type] {
	iterator := func(yield func(Type) bool) {
		if begin > end {
			return
		}

		for number := begin; number < end; number++ {
			if !yield(number) {
				return
			}
		}

		yield(end)
	}

	return iterator
}

// Calculates the number of iterations when using [Inc]. The return value
// is intended to be used as the size parameter in the make call, so, and because
// the maximum possible number of iterations is one more than the maximum value for
// uint64, the return value is truncated to the maximum value for uint64 if the
// calculated value exceeds it.
func IncSize[Type constraints.Integer](begin, end Type) uint64 {
	if begin > end {
		return 0
	}

	return IterSize(begin, end)
}

// A range iterator for safely (without infinite loops due to counter overflow)
// iterating over integer values from begin to end inclusive towards decrease with
// a step one.
//
// If begin is lesser than end, then no one iteration of the loop will occur.
func Dec[Type constraints.Integer](begin, end Type) iter.Seq[Type] {
	iterator := func(yield func(Type) bool) {
		if begin < end {
			return
		}

		for number := begin; number > end; number-- {
			if !yield(number) {
				return
			}
		}

		yield(end)
	}

	return iterator
}

// Calculates the number of iterations when using [Dec]. The return value
// is intended to be used as the size parameter in the make call, so, and because
// the maximum possible number of iterations is one more than the maximum value for
// uint64, the return value is truncated to the maximum value for uint64 if the
// calculated value exceeds it.
func DecSize[Type constraints.Integer](begin, end Type) uint64 {
	if begin < end {
		return 0
	}

	return IterSize(begin, end)
}

// A range iterator for safely (without infinite loops due to counter overflow)
// iterating over integer values from begin to end inclusive with the ability to
// specify the iteration step.
//
// If begin is greater than end, the return value will be decremented, otherwise it
// will be incremented.
//
// As in a regular loop, if the step is not a multiple of the begin-end range, the end
// value will not be returned.
//
// If a zero or negative step is specified, the iterator will panic.
//
// In addition to the main integer, its index in the begin-end sequence is returned.
func Step[Type constraints.Integer](begin, end, step Type) iter.Seq2[uint64, Type] {
	if step < 0 {
		panic(ErrStepNegative)
	}

	if step == 0 {
		panic(ErrStepZero)
	}

	iterator := func(yield func(uint64, Type) bool) {
		if begin > end {
			id := uint64(0)

			previous := begin

			for number := begin; number >= end; number -= step {
				// integer overflow
				if number > previous {
					return
				}

				previous = number

				if !yield(id, number) {
					return
				}

				id++
			}

			return
		}

		id := uint64(0)

		previous := begin

		for number := begin; number <= end; number += step {
			// integer overflow
			if number < previous {
				return
			}

			previous = number

			if !yield(id, number) {
				return
			}

			id++
		}
	}

	return iterator
}

// Calculates the number of iterations when using [Step]. The return value
// is intended to be used as the size parameter in the make call, so, and because
// the maximum possible number of iterations is one more than the maximum value for
// uint64, the return value is truncated to the maximum value for uint64 if the
// calculated value exceeds it.
//
// Like [Step] this function panics if a zero or negative step is specified.
func StepSize[Type constraints.Integer](begin, end, step Type) uint64 {
	if step < 0 {
		panic(ErrStepNegative)
	}

	if step == 0 {
		panic(ErrStepZero)
	}

	beginU64 := u64(begin)
	endU64 := u64(end)
	stepU64 := uint64(step)

	size := endU64 - beginU64

	if beginU64 > endU64 {
		size = beginU64 - endU64
	}

	// begin < 0 && end > 0 || begin > 0 && end < 0
	if begin^end < 0 {
		size = endU64 + beginU64
	}

	if size == intspec.MaxUint64 && stepU64 == 1 {
		return size
	}

	return size/stepU64 + 1
}

// A range iterator for safely (without infinite loops due to counter overflow)
// iterating over integer values from begin to end inclusive towards increase with
// the ability to specify the iteration step.
//
// If begin is greater than end, then no one iteration of the loop will occur.
//
// As in a regular loop, if the step is not a multiple of the begin-end range, the end
// value will not be returned.
//
// If a zero or negative step is specified, the iterator will panic.
//
// In addition to the main integer, its index in the begin-end sequence is returned.
func IncStep[Type constraints.Integer](begin, end, step Type) iter.Seq2[uint64, Type] {
	if step < 0 {
		panic(ErrStepNegative)
	}

	if step == 0 {
		panic(ErrStepZero)
	}

	iterator := func(yield func(uint64, Type) bool) {
		id := uint64(0)

		previous := begin

		for number := begin; number <= end; number += step {
			// integer overflow
			if number < previous {
				return
			}

			previous = number

			if !yield(id, number) {
				return
			}

			id++
		}
	}

	return iterator
}

// Calculates the number of iterations when using [IncStep]. The return value
// is intended to be used as the size parameter in the make call, so, and because
// the maximum possible number of iterations is one more than the maximum value for
// uint64, the return value is truncated to the maximum value for uint64 if the
// calculated value exceeds it.
//
// Like [IncStep] this function panics if a zero or negative step is specified.
func IncStepSize[Type constraints.Integer](begin, end, step Type) uint64 {
	size := StepSize(begin, end, step)

	if begin > end {
		return 0
	}

	return size
}

// A range iterator for safely (without infinite loops due to counter overflow)
// iterating over integer values from begin to end inclusive towards decrease with
// the ability to specify the iteration step.
//
// If begin is lesser than end, then no one iteration of the loop will occur.
//
// As in a regular loop, if the step is not a multiple of the begin-end range, the end
// value will not be returned.
//
// If a zero or negative step is specified, the iterator will panic.
//
// In addition to the main integer, its index in the begin-end sequence is returned.
func DecStep[Type constraints.Integer](begin, end, step Type) iter.Seq2[uint64, Type] {
	if step < 0 {
		panic(ErrStepNegative)
	}

	if step == 0 {
		panic(ErrStepZero)
	}

	iterator := func(yield func(uint64, Type) bool) {
		id := uint64(0)

		previous := begin

		for number := begin; number >= end; number -= step {
			// integer overflow
			if number > previous {
				return
			}

			previous = number

			if !yield(id, number) {
				return
			}

			id++
		}
	}

	return iterator
}

// Calculates the number of iterations when using [DecStep]. The return value
// is intended to be used as the size parameter in the make call, so, and because
// the maximum possible number of iterations is one more than the maximum value for
// uint64, the return value is truncated to the maximum value for uint64 if the
// calculated value exceeds it.
//
// Like [DecStep] this function panics if a zero or negative step is specified.
func DecStepSize[Type constraints.Integer](begin, end, step Type) uint64 {
	size := StepSize(begin, end, step)

	if begin < end {
		return 0
	}

	return size
}
