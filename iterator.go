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

// Calculates the number of iterations when using [Iter]. The return value
// is intended to be used as the size parameter in the make call, so, and because
// the maximum possible number of iterations is one more than the maximum value for
// uint64, the return value is truncated to the maximum value for uint64 if the
// calculated value exceeds it.
func IterSize[Type constraints.Integer](begin, end Type) uint64 {
	return iterator.IterSize(begin, end)
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
func IterStep[Type constraints.Integer](begin, end, step Type) iter.Seq2[uint64, Type] {
	return iterator.IterStep(begin, end, step, ErrIterStepNegative, ErrIterStepZero)
}

// Calculates the number of iterations when using [IterStep]. The return value
// is intended to be used as the size parameter in the make call, so, and because
// the maximum possible number of iterations is one more than the maximum value for
// uint64, the return value is truncated to the maximum value for uint64 if the
// calculated value exceeds it.
//
// Like [IterStep] this function panics if a zero or negative step is specified.
func IterStepSize[Type constraints.Integer](begin, end, step Type) uint64 {
	return iterator.IterStepSize(begin, end, step, ErrIterStepNegative, ErrIterStepZero)
}

// A range iterator for safely (without infinite loops due to counter overflow)
// iterating over integer values from begin to end inclusive with the ability to
// specify the iteration step and direction.
//
// Unlike [IterStep], the iteration direction (increase or decrease the return value)
// is selected manually by the direction parameter. If direction is less than 0, the
// return value will be decremented, otherwise it will be incremented. If the begin
// and end values ​​do not match the specified iteration direction, then no iteration of
// the loop will occur.
//
// As in a regular loop, if the step is not a multiple of the begin-end range, the end
// value will not be returned.
//
// If a zero or negative step is specified, the iterator will panic.
//
// In addition to the main integer, its index in the begin-end sequence is returned.
func IterStrict[Type constraints.Integer](begin, end, step Type, direction int) iter.Seq2[uint64, Type] {
	return iterator.IterStrict(begin, end, step, direction, ErrIterStepNegative, ErrIterStepZero)
}

// Calculates the number of iterations when using [IterStrict]. The return value
// is intended to be used as the size parameter in the make call, so, and because
// the maximum possible number of iterations is one more than the maximum value for
// uint64, the return value is truncated to the maximum value for uint64 if the
// calculated value exceeds it.
//
// Like [IterStrict] this function panics if a zero or negative step is specified.
func IterStrictSize[Type constraints.Integer](begin, end, step Type, direction int) uint64 {
	return iterator.IterStrictSize(begin, end, step, direction, ErrIterStepNegative, ErrIterStepZero)
}
