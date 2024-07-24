package inspect

// Constraints by 8-bit integer types.
type EightBits interface {
	~int8 | ~uint8
}

// Inspection result.
type Result[Type EightBits] struct {
	// Value returned by the inspected function. Filled in if its value is not
	// equal to the reference value or the inspected function incorrectly reports
	// the absence of an error
	Actual Type
	// Arguments passed to the reference and inspected functions. Filled in case of a
	// negative inspection conclusion
	Args []Type
	// Inspection conclusion. Filled in case of incorrect error return by the inspected
	// function or discrepancy between the value returned by the inspected function
	// and the reference value
	Conclusion error
	// Error returned by the inspected function. Filled in case of false error return by
	// the inspected function.
	Err error
	// The number of correct returns of no error by the inspected function (mostly
	// errors is overflow)
	NoOverflows int
	// Number of correct error returns by the inspected function (mostly errors is
	// overflow)
	Overflows int
	// Value returned by the reference function. Filled in case of a negative inspection
	// conclusion
	Reference int64
	// Number of correct error returns by the inspected function if the reference
	// function returns an error
	ReferenceFaults int
}
