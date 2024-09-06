package dataset

import "errors"

var (
	ErrNotEnoughDataInFillers = errors.New("not enough data from the fillers to form a dataset")
	ErrNotEnoughDataInItem    = errors.New("not enough data in dataset item")
	ErrReaderNotSpecified     = errors.New("reader is not specified")
	ErrWriterNotSpecified     = errors.New("writer is not specified")
)
