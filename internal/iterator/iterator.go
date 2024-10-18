// Internal package with a copy of the iterator from the main code used in internal
// packages. The copy is created to resolve import cycling of packages.
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
