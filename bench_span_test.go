package safe

import (
	"math"
)

// Benchmark spans used for benchmarking with more or less full and uniform coverage.

func benchSpanAdd() ([]int8, []int8) {
	span := []int8{-128, -127, 126, 127}
	return span, span
}

func benchSpanAddU() ([]uint8, []uint8) {
	span := []uint8{1, 2, 3, 255}
	return span, span
}

func benchSpanAdd3() ([]int8, []int8, []int8) {
	span := []int8{-128, -127, 126, 127}
	return span, span, span
}

func benchSpanAdd3U() ([]uint8, []uint8, []uint8) {
	span := []uint8{1, 2, 3, 255}
	return span, span, span
}

func benchSpanAddM() ([]int8, []int8, []int8, []int8, []int8, []int8) {
	span := []int8{-128, -127, 126, 127}
	return span, span, span, span, span, span
}

func benchSpanAddMU() ([]uint8, []uint8, []uint8, []uint8, []uint8, []uint8) {
	span := []uint8{1, 2, 3, 255}
	return span, span, span, span, span, span
}

func benchSpanSub() ([]int8, []int8) {
	span := []int8{-128, -127, 126, 127}
	return span, span
}

func benchSpanSubU() ([]uint8, []uint8) {
	span := []uint8{1, 2, 3, 255}
	return span, span
}

func benchSpanSub3() ([]int8, []int8, []int8) {
	span := []int8{-128, -127, 126, 127}
	return span, span, span
}

func benchSpanSub3U() ([]uint8, []uint8, []uint8) {
	span := []uint8{1, 2, 3, 255}
	return span, span, span
}

func benchSpanSubM() ([]int8, []int8, []int8, []int8, []int8, []int8) {
	span := []int8{-128, -127, 1, 2, 126, 127}
	return span, span, span, span, span, span
}

func benchSpanSubMU() ([]uint8, []uint8, []uint8, []uint8, []uint8, []uint8) {
	span := []uint8{1, 2, 3, 4, 5, 255}
	return span, span, span, span, span, span
}

func benchSpanMul() ([]int8, []int8) {
	span := []int8{-128, -128, -1, 0, 127}
	return span, span
}

func benchSpanMul3() ([]int8, []int8, []int8) {
	span := []int8{-128, -128, -1, 0, 127}
	return span, span, span
}

func benchSpanMul3U() ([]uint8, []uint8, []uint8) {
	span := []uint8{0, 1, 2, 3, 255}
	return span, span, span
}

func benchSpanMulM() ([]int8, []int8, []int8, []int8, []int8, []int8) {
	span := []int8{-128, -1, 0, 1, 1, 2, 3}
	return span, span, span, span, span, span
}

func benchSpanMulMU() ([]uint8, []uint8, []uint8, []uint8, []uint8, []uint8) {
	span := []uint8{0, 1, 2, 3, 255}
	return span, span, span, span, span, span
}

func benchSpanDiv() ([]int8, []int8) {
	span := []int8{-128, -1, -1, -1, 0, 127}
	return span, span
}

func benchSpanDivM() ([]int8, []int8, []int8) {
	span := []int8{-128, -1, -1, -1, 0, 127}
	return span, span, span
}

func benchSpanNegate() ([]int8, []uint8) {
	signed := []int8{-128, -128, 0, 1, 2, 3}
	unsigned := []uint8{0, 1, 2, 3}

	return signed, unsigned
}

func benchSpanIToI() ([]int8, []uint8, []uint16) {
	s8 := []int8{-2, -1, 1, 2}
	u8 := []uint8{0, 1, 2, 128, 129, 130, 131}
	u16 := []uint16{256, 257, 258}

	return s8, u8, u16
}

func benchSpanIToF() []int64 {
	return []int64{1, 2, 9007199254740993, 9007199254740995}
}

func benchSpanFToI() []float64 {
	span := []float64{
		-18446744073709551617,
		-18446744073709551616,
		1,
		2,
		18446744073709551616,
		18446744073709551617,
		math.NaN(),
		math.NaN(),
	}

	return span
}

func benchSpanAddSub() ([]uint8, []uint8, []uint8) {
	span := []uint8{0, 1, 2, 254, 255}
	return span, span, span
}

func benchSpanAddDiv() ([]int8, []int8, []int8) {
	span := []int8{-128, -127, -1, 0, 1, 126, 127}
	return span, span, span
}

func benchSpanAddDivU() ([]uint8, []uint8, []uint8) {
	span := []uint8{0, 1, 2, 254, 255}
	return span, span, span
}

func benchSpanSubDiv() ([]int8, []int8, []int8) {
	span := []int8{-128, -127, -1, 0, 1, 126, 127}
	return span, span, span
}

func benchSpanSubDivU() ([]uint8, []uint8, []uint8) {
	span := []uint8{0, 1, 2, 254, 255}
	return span, span, span
}
