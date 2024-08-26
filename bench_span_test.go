package safe

import (
	"testing"
)

// For benchmarking with full coverage.
func benchSpanAdd() []int8 {
	return []int8{127, 126, -128, -127}
}

// These test functions are used to determine the coverage level of the functions
// being benchmarked.
func TestBenchSpanAdd(*testing.T) {
	span := benchSpanAdd()

	for _, first := range span {
		for _, second := range span {
			_, _ = Add(first, second)
		}
	}
}
