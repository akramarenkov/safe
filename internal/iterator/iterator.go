// Internal package with iterators, used to resolve import cycling.
package iterator

import (
	"iter"

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
