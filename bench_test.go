package safe

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkAddReference(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanAdd()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result = first + second
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAdd(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanAdd()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = Add(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddUReference(b *testing.B) {
	result := uint8(0)

	level1, level2 := benchSpanAddU()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result = first + second
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddOnlyUnsigned(b *testing.B) {
	result := uint8(0)

	level1, level2 := benchSpanAddU()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = Add(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddU(b *testing.B) {
	result := uint8(0)

	level1, level2 := benchSpanAddU()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = AddU(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAdd3Reference(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanAdd3()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result = first + second + third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAdd3(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanAdd3()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = Add3(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddM2Args(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanAdd()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = AddM(false, first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddM3Args(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanAdd3()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = AddM(false, first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddMReference(b *testing.B) {
	result := int8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanAddM()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result = first + second + third + fourth + fifth + sixth
							}
						}
					}
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddM(b *testing.B) {
	result := int8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanAddM()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result, _ = AddM(
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

	require.NotNil(b, result)
}

func BenchmarkAddUM2Args(b *testing.B) {
	result := uint8(0)

	level1, level2 := benchSpanAddU()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = AddUM(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddUMReference(b *testing.B) {
	result := uint8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanAddUM()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result = first + second + third + fourth + fifth + sixth
							}
						}
					}
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddMOnlyUnsigned(b *testing.B) {
	result := uint8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanAddUM()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result, _ = AddM(
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

	require.NotNil(b, result)
}

func BenchmarkAddUM(b *testing.B) {
	result := uint8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanAddUM()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result, _ = AddUM(
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

	require.NotNil(b, result)
}

func BenchmarkSubReference(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanSub()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result = first - second
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSub(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanSub()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = Sub(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubUReference(b *testing.B) {
	result := uint8(0)

	level1, level2 := benchSpanSubU()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result = first - second
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubOnlyUnsigned(b *testing.B) {
	result := uint8(0)

	level1, level2 := benchSpanSubU()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = Sub(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubU(b *testing.B) {
	result := uint8(0)

	level1, level2 := benchSpanSubU()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = SubU(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSub3Reference(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanSub3()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result = first - second - third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSub3(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanSub3()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = Sub3(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubUM2Args(b *testing.B) {
	result := uint8(0)

	level1, level2 := benchSpanSubU()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = SubUM(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubUMReference(b *testing.B) {
	result := uint8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanSubUM()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result = first - second - third - fourth - fifth - sixth
							}
						}
					}
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubUM(b *testing.B) {
	result := uint8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanSubUM()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result, _ = SubUM(
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

	require.NotNil(b, result)
}

func BenchmarkMulReference(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanMul()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result = first * second
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMul(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanMul()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = Mul(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMul3Reference(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanMul3()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result = first * second * third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMul3(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanMul3()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = Mul3(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulM2Args(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanMul()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = MulM(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulM3Args(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanMul3()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = MulM(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulMReference(b *testing.B) {
	result := int8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanMulM()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result = first * second * third * fourth * fifth * sixth
							}
						}
					}
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulM(b *testing.B) {
	result := int8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanMulM()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result, _ = MulM(
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

	require.NotNil(b, result)
}

func BenchmarkMulUMReference(b *testing.B) {
	result := uint8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanMulUM()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result = first * second * third * fourth * fifth * sixth
							}
						}
					}
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulMOnlyUnsigned(b *testing.B) {
	result := uint8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanMulUM()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result, _ = MulM(
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

	require.NotNil(b, result)
}

func BenchmarkMulUM(b *testing.B) {
	result := uint8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanMulUM()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result, _ = MulUM(
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

	require.NotNil(b, result)
}

func BenchmarkDivReference(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanDiv()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				if second == 0 {
					continue
				}

				result = first / second
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkDiv(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanDiv()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = Div(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkDivM2Args(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanDiv()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = DivM(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkDivMReference(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanDivM()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					if second == 0 || third == 0 {
						continue
					}

					result = first / second / third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkDivM(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanDivM()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = DivM(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkNegateReference(b *testing.B) {
	result := int8(0)
	resultU := uint8(0)

	signed, unsigned := benchSpanNegate()

	for range b.N {
		for _, number := range signed {
			result = -number
		}

		for _, number := range unsigned {
			resultU = -number
		}
	}

	require.NotNil(b, result)
	require.NotNil(b, resultU)
}

func BenchmarkNegate(b *testing.B) {
	result := int8(0)
	resultU := uint8(0)

	signed, unsigned := benchSpanNegate()

	for range b.N {
		for _, number := range signed {
			result, _ = Negate(number)
		}

		for _, number := range unsigned {
			resultU, _ = Negate(number)
		}
	}

	require.NotNil(b, result)
	require.NotNil(b, resultU)
}

func BenchmarkIToIReference(b *testing.B) {
	result := int8(0)
	resultU := uint8(0)

	s8, u8, u16 := benchSpanIToI()

	for range b.N {
		for _, number := range s8 {
			resultU = uint8(number)
		}

		for _, number := range u8 {
			result = int8(number)
		}

		for _, number := range u16 {
			resultU = uint8(number)
		}
	}

	require.NotNil(b, result)
	require.NotNil(b, resultU)
}

func BenchmarkIToI(b *testing.B) {
	result := int8(0)
	resultU := uint8(0)

	s8, u8, u16 := benchSpanIToI()

	for range b.N {
		for _, number := range s8 {
			resultU, _ = IToI[uint8](number)
		}

		for _, number := range u8 {
			result, _ = IToI[int8](number)
		}

		for _, number := range u16 {
			resultU, _ = IToI[uint8](number)
		}
	}

	require.NotNil(b, result)
	require.NotNil(b, resultU)
}

func BenchmarkIToFReference(b *testing.B) {
	result := float64(0)

	span := benchSpanIToF()

	for range b.N {
		for _, number := range span {
			result = float64(number)
		}
	}

	require.NotNil(b, result)
}

func BenchmarkIToF(b *testing.B) {
	result := float64(0)

	span := benchSpanIToF()

	for range b.N {
		for _, number := range span {
			result, _ = IToF[float64](number)
		}
	}

	require.NotNil(b, result)
}

func BenchmarkFToIReference(b *testing.B) {
	result := 0

	span := benchSpanFToI()

	for range b.N {
		for _, number := range span {
			result = int(number)
		}
	}

	require.NotNil(b, result)
}

func BenchmarkFToI(b *testing.B) {
	result := 0

	span := benchSpanFToI()

	for range b.N {
		for _, number := range span {
			result, _ = FToI[int](number)
		}
	}

	require.NotNil(b, result)
}

func BenchmarkPow10Reference(b *testing.B) {
	result := float64(0)

	for range b.N {
		result = math.Pow10(19)
	}

	require.NotNil(b, result)
}

func BenchmarkPow10(b *testing.B) {
	result := uint64(0)

	for range b.N {
		result, _ = Pow10[uint64](19)
	}

	require.NotNil(b, result)
}

func BenchmarkPowReference(b *testing.B) {
	result := float64(0)

	for range b.N {
		result = math.Pow(14, 14)
	}

	require.NotNil(b, result)
}

func BenchmarkPow(b *testing.B) {
	result := uint64(0)

	for range b.N {
		result, _ = Pow(uint64(14), 14)
	}

	require.NotNil(b, result)
}
