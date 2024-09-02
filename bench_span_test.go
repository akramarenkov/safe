package safe

import (
	"math"
	"os"
	"testing"

	"github.com/akramarenkov/safe/internal/consts"
)

// Benchmark spans used for benchmarking with more or less full and uniform coverage.

// In order not to create a false full coverage in which there are no error checks.
func testBenchSpanSkip(t *testing.T) {
	if os.Getenv(consts.EnvEnableBenchSpanTest) == "" {
		t.SkipNow()
	}
}

func benchSpanAdd() ([]int8, []int8) {
	span := []int8{127, 126, -128, -127}
	return span, span
}

func TestBenchSpanAdd(t *testing.T) {
	testBenchSpanSkip(t)

	level1, level2 := benchSpanAdd()

	for _, first := range level1 {
		for _, second := range level2 {
			_, _ = Add(first, second)
		}
	}
}

func benchSpanAddU() ([]uint8, []uint8) {
	span := []uint8{255, 1, 2, 3}
	return span, span
}

func TestBenchSpanAddU(t *testing.T) {
	testBenchSpanSkip(t)

	level1, level2 := benchSpanAddU()

	for _, first := range level1 {
		for _, second := range level2 {
			_, _ = AddU(first, second)
		}
	}
}

func benchSpanAdd3() ([]int8, []int8, []int8) {
	span := []int8{127, 126, -128, -127}
	return span, span, span
}

func TestBenchSpanAdd3(t *testing.T) {
	testBenchSpanSkip(t)

	level1, level2, level3 := benchSpanAdd3()

	for _, first := range level1 {
		for _, second := range level2 {
			for _, third := range level3 {
				_, _ = Add3(first, second, third)
			}
		}
	}
}

func benchSpanAddM() ([]int8, []int8, []int8, []int8, []int8, []int8) {
	span := []int8{127, 126, -128, -127}
	return span, span, span, span, span, span
}

func TestBenchSpanAddM(t *testing.T) {
	testBenchSpanSkip(t)

	level1, level2, level3, level4, level5, level6 := benchSpanAddM()

	for _, first := range level1 {
		for _, second := range level2 {
			for _, third := range level3 {
				for _, fourth := range level4 {
					for _, fifth := range level5 {
						for _, sixth := range level6 {
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

func benchSpanAddUM() ([]uint8, []uint8, []uint8, []uint8, []uint8, []uint8) {
	span := []uint8{255, 1, 2, 3}
	return span, span, span, span, span, span
}

func TestBenchSpanAddUM(t *testing.T) {
	testBenchSpanSkip(t)

	level1, level2, level3, level4, level5, level6 := benchSpanAddUM()

	for _, first := range level1 {
		for _, second := range level2 {
			for _, third := range level3 {
				for _, fourth := range level4 {
					for _, fifth := range level5 {
						for _, sixth := range level6 {
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

func benchSpanSub() ([]int8, []int8) {
	span := []int8{127, 126, -128, -127}
	return span, span
}

func TestBenchSpanSub(t *testing.T) {
	testBenchSpanSkip(t)

	level1, level2 := benchSpanSub()

	for _, first := range level1 {
		for _, second := range level2 {
			_, _ = Sub(first, second)
		}
	}
}

func benchSpanSubU() ([]uint8, []uint8) {
	span := []uint8{255, 1, 2, 3}
	return span, span
}

func TestBenchSpanSubU(t *testing.T) {
	testBenchSpanSkip(t)

	level1, level2 := benchSpanSubU()

	for _, first := range level1 {
		for _, second := range level2 {
			_, _ = SubU(first, second)
		}
	}
}

func benchSpanSub3() ([]int8, []int8, []int8) {
	span := []int8{127, 126, -128, -127}
	return span, span, span
}

func TestBenchSpanSub3(t *testing.T) {
	testBenchSpanSkip(t)

	level1, level2, level3 := benchSpanSub3()

	for _, first := range level1 {
		for _, second := range level2 {
			for _, third := range level3 {
				_, _ = Sub3(first, second, third)
			}
		}
	}
}

func benchSpanSubUM() ([]uint8, []uint8, []uint8, []uint8, []uint8, []uint8) {
	span := []uint8{255, 1, 2, 3}
	return span, span, span, span, span, span
}

func TestBenchSpanSubUM(t *testing.T) {
	testBenchSpanSkip(t)

	level1, level2, level3, level4, level5, level6 := benchSpanSubUM()

	for _, first := range level1 {
		for _, second := range level2 {
			for _, third := range level3 {
				for _, fourth := range level4 {
					for _, fifth := range level5 {
						for _, sixth := range level6 {
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

func benchSpanMul() ([]int8, []int8) {
	span := []int8{0, -128, -128, -1, 127}
	return span, span
}

func TestBenchSpanMul(t *testing.T) {
	testBenchSpanSkip(t)

	level1, level2 := benchSpanMul()

	for _, first := range level1 {
		for _, second := range level2 {
			_, _ = Mul(first, second)
		}
	}
}

func benchSpanMul3() ([]int8, []int8, []int8) {
	span := []int8{0, -128, -128, -1, 127}
	return span, span, span
}

func TestBenchSpanMul3(t *testing.T) {
	testBenchSpanSkip(t)

	level1, level2, level3 := benchSpanMul3()

	for _, first := range level1 {
		for _, second := range level2 {
			for _, third := range level3 {
				_, _ = Mul3(first, second, third)
			}
		}
	}
}

func benchSpanMulM() ([]int8, []int8, []int8, []int8, []int8, []int8) {
	span := []int8{0, 2, -128, -1, 3, 1, 1}
	return span, span, span, span, span, span
}

func TestBenchSpanMulM(t *testing.T) {
	testBenchSpanSkip(t)

	level1, level2, level3, level4, level5, level6 := benchSpanMulM()

	for _, first := range level1 {
		for _, second := range level2 {
			for _, third := range level3 {
				for _, fourth := range level4 {
					for _, fifth := range level5 {
						for _, sixth := range level6 {
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

func benchSpanMulUM() ([]uint8, []uint8, []uint8, []uint8, []uint8, []uint8) {
	span := []uint8{255, 1, 2, 3}
	return span, span, span, span, span, span
}

func TestBenchSpanMulUM(t *testing.T) {
	testBenchSpanSkip(t)

	level1, level2, level3, level4, level5, level6 := benchSpanMulUM()

	for _, first := range level1 {
		for _, second := range level2 {
			for _, third := range level3 {
				for _, fourth := range level4 {
					for _, fifth := range level5 {
						for _, sixth := range level6 {
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

func benchSpanDiv() ([]int8, []int8) {
	span := []int8{0, -128, -1, -1, -1, 127}
	return span, span
}

func TestBenchSpanDiv(t *testing.T) {
	testBenchSpanSkip(t)

	level1, level2 := benchSpanDiv()

	for _, first := range level1 {
		for _, second := range level2 {
			_, _ = Div(first, second)
		}
	}
}

func benchSpanDivM() ([]int8, []int8, []int8) {
	span := []int8{0, -128, -1, -1, -1, 127}
	return span, span, span
}

func TestBenchSpanDivM(t *testing.T) {
	testBenchSpanSkip(t)

	level1, level2, level3 := benchSpanDivM()

	for _, first := range level1 {
		for _, second := range level2 {
			for _, third := range level3 {
				_, _ = DivM(first, second, third)
			}
		}
	}
}

func benchSpanNegate() ([]int8, []uint8) {
	signed := []int8{-128, -128, 0, 1, 2, 3}
	unsigned := []uint8{0, 1, 2, 3}

	return signed, unsigned
}

func TestBenchSpanNegate(t *testing.T) {
	testBenchSpanSkip(t)

	signed, unsigned := benchSpanNegate()

	for _, number := range signed {
		_, _ = Negate(number)
	}

	for _, number := range unsigned {
		_, _ = Negate(number)
	}
}

func benchSpanIToI() ([]int8, []uint8, []uint16) {
	s8 := []int8{-2, -1, 1, 2}
	u8 := []uint8{0, 1, 2, 128, 129, 130, 131}
	u16 := []uint16{256, 257, 258}

	return s8, u8, u16
}

func TestBenchSpanIToI(t *testing.T) {
	testBenchSpanSkip(t)

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

func benchSpanIToF() []int64 {
	return []int64{1, 2, 9007199254740993, 9007199254740995}
}

func TestBenchSpanIToF(t *testing.T) {
	testBenchSpanSkip(t)

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

func TestBenchSpanFToI(t *testing.T) {
	testBenchSpanSkip(t)

	span := benchSpanFToI()

	for _, number := range span {
		_, _ = FToI[int](number)
	}
}
