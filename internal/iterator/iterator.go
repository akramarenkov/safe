// Internal package with iterators, used to resolve import cycling.
package iterator

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
	forward := func(yield func(Type) bool) {
		for number := begin; number < end; number++ {
			if !yield(number) {
				return
			}
		}

		yield(end)
	}

	backward := func(yield func(Type) bool) {
		for number := begin; number > end; number-- {
			if !yield(number) {
				return
			}
		}

		yield(end)
	}

	if begin > end {
		return backward
	}

	return forward
}

// A range iterator for safely (without infinite loops due to counter overflow)
// iterating over integer values from begin to end inclusive with a step one.
//
// If begin is greater than end, the return value will be decremented, otherwise it
// will be incremented.
//
// Unlike [Iter], it returns, in addition to the main integer, its index in the
// begin-end sequence.
func Iter2[Type constraints.Integer](begin, end Type) iter.Seq2[uint64, Type] {
	forward := func(yield func(uint64, Type) bool) {
		id := uint64(0)

		for number := begin; number < end; number++ {
			if !yield(id, number) {
				return
			}

			id++
		}

		yield(id, end)
	}

	backward := func(yield func(uint64, Type) bool) {
		id := uint64(0)

		for number := begin; number > end; number-- {
			if !yield(id, number) {
				return
			}

			id++
		}

		yield(id, end)
	}

	if begin > end {
		return backward
	}

	return forward
}

// Calculates the number of iterations when using [Iter]. The return value is
// intended to be used as the size parameter in the make call, so, and because the
// maximum possible number of iterations is one more than the maximum value for uint64,
// the return value is truncated to the maximum value for uint64 if the calculated
// value exceeds it.
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
// iterating over integer values from begin to end inclusive with the ability to
// specify the iteration step.
//
// As in a regular loop, if the step is not a multiple of the begin-end range, the end
// value will not be returned.
//
// If begin is greater than end, the return value will be decremented, otherwise it
// will be incremented.
//
// If a zero or negative step is specified, the iterator will panic.
func IterStep[Type constraints.Integer](
	begin Type,
	end Type,
	step Type,
	stepNegative error,
	stepZero error,
) iter.Seq[Type] {
	if step < 0 {
		panic(stepNegative)
	}

	if step == 0 {
		panic(stepZero)
	}

	forward := func(yield func(Type) bool) {
		previous := begin

		for number := begin; number <= end; number += step {
			// integer overflow
			if number < previous {
				return
			}

			previous = number

			if !yield(number) {
				return
			}
		}
	}

	backward := func(yield func(Type) bool) {
		previous := begin

		for number := begin; number >= end; number -= step {
			// integer overflow
			if number > previous {
				return
			}

			previous = number

			if !yield(number) {
				return
			}
		}
	}

	if begin > end {
		return backward
	}

	return forward
}

// Calculates the number of iterations when using [IterStep]. The return value is
// intended to be used as the size parameter in the make call, so, and because the
// maximum possible number of iterations is one more than the maximum value for uint64,
// the return value is truncated to the maximum value for uint64 if the calculated
// value exceeds it.
//
// Like [IterStep] this function panics if a zero or negative step is specified.
func IterStepSize[Type constraints.Integer](
	begin Type,
	end Type,
	step Type,
	stepNegative error,
	stepZero error,
) uint64 {
	if step < 0 {
		panic(stepNegative)
	}

	if step == 0 {
		panic(stepZero)
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

	if size == intspec.MaxUint64 {
		if stepU64 == 1 {
			return size
		}

		quotient := size / stepU64
		remainder := size % stepU64

		return quotient + (remainder+1)/stepU64
	}

	return (size + 1) / stepU64
}
