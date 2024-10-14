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
func Step[Type constraints.Integer](begin, end, step Type) iter.Seq2[uint64, Type] {
	return iterator.Step(begin, end, step, ErrStepNegative, ErrStepZero)
}

// Calculates the number of iterations when using [Step]. The return value
// is intended to be used as the size parameter in the make call, so, and because
// the maximum possible number of iterations is one more than the maximum value for
// uint64, the return value is truncated to the maximum value for uint64 if the
// calculated value exceeds it.
//
// Like [Step] this function panics if a zero or negative step is specified.
func StepSize[Type constraints.Integer](begin, end, step Type) uint64 {
	return iterator.StepSize(begin, end, step, ErrStepNegative, ErrStepZero)
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
func Inc[Type constraints.Integer](begin, end, step Type) iter.Seq2[uint64, Type] {
	return iterator.Inc(begin, end, step, ErrStepNegative, ErrStepZero)
}

// Calculates the number of iterations when using [Inc]. The return value
// is intended to be used as the size parameter in the make call, so, and because
// the maximum possible number of iterations is one more than the maximum value for
// uint64, the return value is truncated to the maximum value for uint64 if the
// calculated value exceeds it.
//
// Like [Inc] this function panics if a zero or negative step is specified.
func IncSize[Type constraints.Integer](begin, end, step Type) uint64 {
	return iterator.IncSize(begin, end, step, ErrStepNegative, ErrStepZero)
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
func Dec[Type constraints.Integer](begin, end, step Type) iter.Seq2[uint64, Type] {
	return iterator.Dec(begin, end, step, ErrStepNegative, ErrStepZero)
}

// Calculates the number of iterations when using [Dec]. The return value
// is intended to be used as the size parameter in the make call, so, and because
// the maximum possible number of iterations is one more than the maximum value for
// uint64, the return value is truncated to the maximum value for uint64 if the
// calculated value exceeds it.
//
// Like [Dec] this function panics if a zero or negative step is specified.
func DecSize[Type constraints.Integer](begin, end, step Type) uint64 {
	return iterator.DecSize(begin, end, step, ErrStepNegative, ErrStepZero)
}
