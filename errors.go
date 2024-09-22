package safe

import "errors"

var (
	ErrDivisionByZero   = errors.New("division by zero")
	ErrIterStepNegative = errors.New("iterator step is negative")
	ErrIterStepZero     = errors.New("iterator step is zero")
	ErrMissingArguments = errors.New("missing arguments")
	ErrNaN              = errors.New("number is NaN")
	ErrOverflow         = errors.New("integer overflow")
	ErrPrecisionLoss    = errors.New("loss of precision")
)
