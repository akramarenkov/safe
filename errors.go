package safe

import "errors"

var (
	ErrDivisionByZero   = errors.New("division by zero")
	ErrMissingArguments = errors.New("missing arguments")
	ErrNaN              = errors.New("number is NaN")
	ErrNegativeShift    = errors.New("shift count is negative")
	ErrOverflow         = errors.New("integer overflow")
	ErrPrecisionLoss    = errors.New("loss of precision")
	ErrStepNegative     = errors.New("iterator step is negative")
	ErrStepZero         = errors.New("iterator step is zero")
)
