package safe

import "errors"

var (
	ErrDivisionByZero   = errors.New("division by zero")
	ErrMissingArguments = errors.New("missing arguments")
	ErrNaN              = errors.New("number is NaN")
	ErrOverflow         = errors.New("integer overflow")
	ErrPrecisionLoss    = errors.New("loss of precision")
)
