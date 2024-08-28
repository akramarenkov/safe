package inspect

import "errors"

var (
	ErrErrorExpected         = errors.New("an error is expected but got nil")
	ErrInspectedNotSpecified = errors.New("inspected function is not specified")
	ErrNotEqual              = errors.New("actual value is not equal to reference value")
	ErrReferenceNotSpecified = errors.New("reference function is not specified")
	ErrUnexpectedError       = errors.New("received unexpected error")
)

var (
	ErrOverflow = errors.New("overflow")
)
