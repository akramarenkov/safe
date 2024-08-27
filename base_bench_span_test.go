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

func benchSpanNegate() ([]int8, []uint8) {
	return []int8{-128, 0, 1, 2}, []uint8{0, 1, 2}
}

func TestBenchSpanNegate(*testing.T) {
	signed, unsigned := benchSpanNegate()

	for _, number := range signed {
		_, _ = Negate(number)
	}

	for _, number := range unsigned {
		_, _ = Negate(number)
	}
}

func benchSpanNegateS() []int8 {
	return []int8{-128, -128, 0, 1}
}

func TestBenchSpanNegateS(*testing.T) {
	span := benchSpanNegateS()

	for _, number := range span {
		_, _ = NegateS(number)
	}
}

func benchSpanIToI() ([]int8, []uint8, []uint16) {
	return []int8{-2, -1, 1, 2}, []uint8{0, 1, 128, 129}, []uint16{256, 257}
}

func TestBenchSpanIToI(*testing.T) {
	s8, u8, u16 := benchSpanIToI()

	for _, number := range s8 {
		_, _ = IToI[uint8](number)
	}

	for _, number := range u8 {
		_, _ = IToI[int8](number)
	}

	for _, number := range u16 {
		_, _ = IToI[uint8](number)
	}
}
