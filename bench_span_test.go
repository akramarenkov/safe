package safe

import (
	"math"
	"os"
	"testing"

	"github.com/akramarenkov/safe/internal/consts"
)

// For benchmarking with full and more or less uniform coverage.
func TestBenchSpan(t *testing.T) {
	// In order not to create a false full coverage in which there are no error checks
	if os.Getenv(consts.EnvEnableBenchSpanTest) == "" {
		t.SkipNow()
	}

	testBenchSpanAdd(t)
	testBenchSpanAddU(t)
	testBenchSpanSub(t)
	testBenchSpanSubU(t)
	testBenchSpanMul(t)
	testBenchSpanDiv(t)
	testBenchSpanNegate(t)
	testBenchSpanNegateS(t)
	testBenchSpanIToI(t)
	testBenchSpanUToS(t)
	testBenchSpanIToF(t)
	testBenchSpanFToI(t)
	testBenchSpanAdd3(t)
	testBenchSpanAddM(t)
	testBenchSpanAddUM(t)
	testBenchSpanSub3(t)
	testBenchSpanSubUM(t)
	testBenchSpanMul3(t)
	testBenchSpanMulM(t)
	testBenchSpanMulUM(t)
	testBenchSpanDivM(t)
}

func benchSpanAdd() []int8 {
	return []int8{127, 126, -128, -127}
}

func testBenchSpanAdd(*testing.T) {
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

func testBenchSpanAddU(*testing.T) {
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

func testBenchSpanSub(*testing.T) {
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

func testBenchSpanSubU(*testing.T) {
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

func testBenchSpanMul(*testing.T) {
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

func testBenchSpanDiv(*testing.T) {
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

func testBenchSpanNegate(*testing.T) {
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

func testBenchSpanNegateS(*testing.T) {
	span := benchSpanNegateS()

	for _, number := range span {
		_, _ = NegateS(number)
	}
}

func benchSpanIToI() ([]int8, []uint8, []uint16) {
	return []int8{-2, -1, 1, 2}, []uint8{0, 1, 128, 129}, []uint16{256, 257}
}

func testBenchSpanIToI(*testing.T) {
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

func benchSpanUToS() ([]uint8, []uint16) {
	return []uint8{0, 1, 128, 129}, []uint16{256, 257}
}

func testBenchSpanUToS(*testing.T) {
	u8, u16 := benchSpanUToS()

	for _, number := range u8 {
		_, _ = UToS[int8](number)
	}

	for _, number := range u16 {
		_, _ = UToS[int8](number)
	}
}

func benchSpanIToF() []int64 {
	return []int64{1, 2, 9007199254740993, 9007199254740995}
}

func testBenchSpanIToF(*testing.T) {
	span := benchSpanIToF()

	for _, number := range span {
		_, _ = IToF[float64](number)
	}
}

func benchSpanFToI() []float64 {
	span := []float64{
		math.NaN(),
		math.NaN(),
		1,
		2,
		18446744073709551616,
		18446744073709551617,
		-18446744073709551616,
		-18446744073709551617,
	}

	return span
}

func testBenchSpanFToI(*testing.T) {
	span := benchSpanFToI()

	for _, number := range span {
		_, _ = FToI[int](number)
	}
}

func benchSpanAdd3() []int8 {
	return []int8{127, 126, -128, -127}
}

func testBenchSpanAdd3(*testing.T) {
	span := benchSpanAdd3()

	for _, first := range span {
		for _, second := range span {
			for _, third := range span {
				_, _ = Add3(first, second, third)
			}
		}
	}
}

func benchSpanAddM() []int8 {
	return []int8{127, 126, -128, -127}
}

func testBenchSpanAddM(*testing.T) {
	span := benchSpanAddM()

	for _, first := range span {
		for _, second := range span {
			for _, third := range span {
				for _, fourth := range span {
					for _, fifth := range span {
						for _, sixth := range span {
							_, _ = AddM(
								false,
								first,
								second,
								third,
								fourth,
								fifth,
								sixth,
							)
						}
					}
				}
			}
		}
	}
}

func benchSpanAddUM() []uint8 {
	return []uint8{255, 1, 2, 3}
}

func testBenchSpanAddUM(*testing.T) {
	span := benchSpanAddUM()

	for _, first := range span {
		for _, second := range span {
			for _, third := range span {
				for _, fourth := range span {
					for _, fifth := range span {
						for _, sixth := range span {
							_, _ = AddUM(
								first,
								second,
								third,
								fourth,
								fifth,
								sixth,
							)
						}
					}
				}
			}
		}
	}
}

func benchSpanSub3() []int8 {
	return []int8{127, 126, -128, -127}
}

func testBenchSpanSub3(*testing.T) {
	span := benchSpanSub3()

	for _, first := range span {
		for _, second := range span {
			for _, third := range span {
				_, _ = Sub3(first, second, third)
			}
		}
	}
}

func benchSpanSubUM() []uint8 {
	return []uint8{255, 1, 2, 3}
}

func testBenchSpanSubUM(*testing.T) {
	span := benchSpanSubUM()

	for _, first := range span {
		for _, second := range span {
			for _, third := range span {
				for _, fourth := range span {
					for _, fifth := range span {
						for _, sixth := range span {
							_, _ = SubUM(
								first,
								second,
								third,
								fourth,
								fifth,
								sixth,
							)
						}
					}
				}
			}
		}
	}
}

func benchSpanMul3() []int8 {
	return []int8{0, -128, -128, -1, 127}
}

func testBenchSpanMul3(*testing.T) {
	span := benchSpanMul3()

	for _, first := range span {
		for _, second := range span {
			for _, third := range span {
				_, _ = Mul3(first, second, third)
			}
		}
	}
}

func benchSpanMulM() []int8 {
	return []int8{0, 2, -128, -1, 3, 1, 1}
}

func testBenchSpanMulM(*testing.T) {
	span := benchSpanMulM()

	for _, first := range span {
		for _, second := range span {
			for _, third := range span {
				for _, fourth := range span {
					for _, fifth := range span {
						for _, sixth := range span {
							_, _ = MulM(
								first,
								second,
								third,
								fourth,
								fifth,
								sixth,
							)
						}
					}
				}
			}
		}
	}
}

func benchSpanMulUM() []uint8 {
	return []uint8{255, 1, 2, 3}
}

func testBenchSpanMulUM(*testing.T) {
	span := benchSpanMulUM()

	for _, first := range span {
		for _, second := range span {
			for _, third := range span {
				for _, fourth := range span {
					for _, fifth := range span {
						for _, sixth := range span {
							_, _ = MulUM(
								first,
								second,
								third,
								fourth,
								fifth,
								sixth,
							)
						}
					}
				}
			}
		}
	}
}

func benchSpanDivM() []int8 {
	return []int8{0, -128, -1, -1, -1, 127}
}

func testBenchSpanDivM(*testing.T) {
	span := benchSpanDivM()

	for _, first := range span {
		for _, second := range span {
			_, _ = DivM(first, second)
		}
	}
}
