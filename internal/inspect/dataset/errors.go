package dataset

import "errors"

var (
	ErrNotEnoughDataInElement = errors.New("not enough data in dataset element")
	ErrNotEnoughDataInFillers = errors.New("not enough data from the fillers to form a dataset")
	ErrReaderNotSpecified     = errors.New("reader is not specified")
	ErrWriterNotSpecified     = errors.New("writer is not specified")
)
