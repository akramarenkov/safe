package safe

import (
	"iter"

	"github.com/akramarenkov/safe/internal/iterator"
	"golang.org/x/exp/constraints"
)

// A range iterator for safely (without infinite loops due to counter overflow)
// iterating over integer values from begin to end inclusive with a step one.
//
// If begin is greater than end, the return value will be decremented, otherwise it
// will be incremented.
func Iter[Type constraints.Integer](begin, end Type) iter.Seq[Type] {
	return iterator.Iter(begin, end)
}

// A range iterator for safely (without infinite loops due to counter overflow)
// iterating over integer values from begin to end inclusive with the ability to
// specify the iteration step.
//
// If begin is greater than end, the return value will be decremented, otherwise it
// will be incremented.
//
// If a zero or negative step is specified, the iterator will panic.
func IterStep[Type constraints.Integer](begin, end, step Type) iter.Seq[Type] {
	return iterator.IterStep(begin, end, step, ErrIterStepNegative, ErrIterStepZero)
}
