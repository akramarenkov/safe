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

func benchSpanAddU() []uint8 {
	return []uint8{255, 1, 2, 3}
}

func TestBenchSpanAddU(*testing.T) {
	span := benchSpanAddU()

	for _, first := range span {
		for _, second := range span {
			_, _ = AddU(first, second)
		}
	}
}

func benchSpanSub() []int8 {
	return []int8{127, 126, -128, -127}
}

func TestBenchSpanSub(*testing.T) {
	span := benchSpanSub()

	for _, first := range span {
		for _, second := range span {
			_, _ = Sub(first, second)
		}
	}
}

func benchSpanSubU() []uint8 {
	return []uint8{255, 1, 2, 3}
}

func TestBenchSpanSubU(*testing.T) {
	span := benchSpanSubU()

	for _, first := range span {
		for _, second := range span {
			_, _ = SubU(first, second)
		}
	}
}

func benchSpanMul() []int8 {
	return []int8{0, -128, -128, -1, 127}
}

func TestBenchSpanMul(*testing.T) {
	span := benchSpanMul()

	for _, first := range span {
		for _, second := range span {
			_, _ = Mul(first, second)
		}
	}
}

func benchSpanDiv() []int8 {
	return []int8{0, -128, -1, -1, -1, 127}
}

func TestBenchSpanDiv(*testing.T) {
	span := benchSpanDiv()

	for _, first := range span {
		for _, second := range span {
			_, _ = Div(first, second)
		}
	}
}
