package safe

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkAddReference(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanAdd()

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

func BenchmarkAdd3UReference(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanAdd3U()

	b.ResetTimer()

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

func BenchmarkAdd3OnlyUnsigned(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanAdd3U()

	b.ResetTimer()

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

func BenchmarkAdd3U(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanAdd3U()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = Add3U(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddM2Args(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanAdd()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = AddM(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddM3Args(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanAdd3()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = AddM(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddMReference(b *testing.B) {
	result := int8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanAddM()

	b.ResetTimer()

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

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result, _ = AddM(
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

func BenchmarkAddMU2Args(b *testing.B) {
	result := uint8(0)

	level1, level2 := benchSpanAddU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = AddMU(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddMU3Args(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanAdd3U()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = AddMU(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddMUReference(b *testing.B) {
	result := uint8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanAddMU()

	b.ResetTimer()

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

	level1, level2, level3, level4, level5, level6 := benchSpanAddMU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result, _ = AddM(
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

func BenchmarkAddMU(b *testing.B) {
	result := uint8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanAddMU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result, _ = AddMU(
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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

func BenchmarkSub3UReference(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanSub3U()

	b.ResetTimer()

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

func BenchmarkSub3OnlyUnsigned(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanSub3U()

	b.ResetTimer()

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

func BenchmarkSub3U(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanSub3U()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = Sub3U(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubM2Args(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanSub()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = SubM(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubM3Args(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanSub3()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = SubM(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubMReference(b *testing.B) {
	result := int8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanSubM()

	b.ResetTimer()

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

func BenchmarkSubM(b *testing.B) {
	result := int8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanSubM()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result, _ = SubM(
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

func BenchmarkSubMU2Args(b *testing.B) {
	result := uint8(0)

	level1, level2 := benchSpanSubU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = SubMU(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubMU3Args(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanSub3U()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = SubMU(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubMUReference(b *testing.B) {
	result := uint8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanSubMU()

	b.ResetTimer()

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

func BenchmarkSubMOnlyUnsigned(b *testing.B) {
	result := uint8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanSubMU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result, _ = SubM(
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

func BenchmarkSubMU(b *testing.B) {
	result := uint8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanSubMU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result, _ = SubMU(
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

	b.ResetTimer()

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

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = Mul(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulUReference(b *testing.B) {
	result := uint8(0)

	level1, level2 := benchSpanMulU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result = first * second
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulOnlyUnsigned(b *testing.B) {
	result := uint8(0)

	level1, level2 := benchSpanMulU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = Mul(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulU(b *testing.B) {
	result := uint8(0)

	level1, level2 := benchSpanMulU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = MulU(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMul3Reference(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanMul3()

	b.ResetTimer()

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

	b.ResetTimer()

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

func BenchmarkMul3UReference(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanMul3U()

	b.ResetTimer()

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

func BenchmarkMul3OnlyUnsigned(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanMul3U()

	b.ResetTimer()

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

func BenchmarkMul3U(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanMul3U()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = Mul3U(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulM2Args(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanMul()

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

func BenchmarkMulMU3Args(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanMul3U()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = MulMU(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkMulMUReference(b *testing.B) {
	result := uint8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanMulMU()

	b.ResetTimer()

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

	level1, level2, level3, level4, level5, level6 := benchSpanMulMU()

	b.ResetTimer()

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

func BenchmarkMulMU(b *testing.B) {
	result := uint8(0)

	level1, level2, level3, level4, level5, level6 := benchSpanMulMU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						for _, fifth := range level5 {
							for _, sixth := range level6 {
								result, _ = MulMU(
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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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
	resultU2 := uint8(0)

	s8, u8, u16 := benchSpanIToI()

	b.ResetTimer()

	for range b.N {
		for _, number := range s8 {
			//nolint:gosec // Unsafe conversion used to compare with safe
			resultU = uint8(number)
		}

		for _, number := range u8 {
			//nolint:gosec // Unsafe conversion used to compare with safe
			result = int8(number)
		}

		for _, number := range u16 {
			//nolint:gosec // Unsafe conversion used to compare with safe
			resultU2 = uint8(number)
		}
	}

	require.NotNil(b, result)
	require.NotNil(b, resultU)
	require.NotNil(b, resultU2)
}

func BenchmarkIToI(b *testing.B) {
	result := int8(0)
	resultU := uint8(0)
	resultU2 := uint8(0)

	s8, u8, u16 := benchSpanIToI()

	b.ResetTimer()

	for range b.N {
		for _, number := range s8 {
			resultU, _ = IToI[uint8](number)
		}

		for _, number := range u8 {
			result, _ = IToI[int8](number)
		}

		for _, number := range u16 {
			resultU2, _ = IToI[uint8](number)
		}
	}

	require.NotNil(b, result)
	require.NotNil(b, resultU)
	require.NotNil(b, resultU2)
}

func BenchmarkIToFReference(b *testing.B) {
	result := float64(0)

	span := benchSpanIToF()

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

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

	b.ResetTimer()

	for range b.N {
		for _, number := range span {
			result, _ = FToI[int](number)
		}
	}

	require.NotNil(b, result)
}

func BenchmarkPow10Reference(b *testing.B) {
	result := float64(0)

	b.ResetTimer()

	for range b.N {
		result = math.Pow10(19)
	}

	require.InEpsilon(b, float64(1e19), result, 0)
}

func BenchmarkPow10(b *testing.B) {
	result := uint64(0)

	b.ResetTimer()

	for range b.N {
		result, _ = Pow10[uint64](19)
	}

	require.Equal(b, uint64(1e19), result)
}

func BenchmarkPowReference(b *testing.B) {
	result := float64(0)

	b.ResetTimer()

	for range b.N {
		result = math.Pow(14, 14)
	}

	require.InEpsilon(b, float64(11112006825558016), result, 0)
}

func BenchmarkPow(b *testing.B) {
	result := uint64(0)

	b.ResetTimer()

	for range b.N {
		result, _ = Pow(uint64(14), 14)
	}

	require.Equal(b, uint64(11112006825558016), result)
}

func BenchmarkAddSubReference(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanAddSub()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result = first + second - third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddSub(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanAddSub()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = AddSub(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddDivReference(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanAddDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					if third == 0 {
						continue
					}

					result = (first + second) / third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddDiv(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanAddDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = AddDiv(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddDivRemReference(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanAddDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					if third == 0 {
						continue
					}

					result = (first + second) % third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddDivRem(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanAddDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = AddDivRem(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddDivUReference(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanAddDivU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					if third == 0 {
						continue
					}

					result = (first + second) / third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddDivOnlyUnsigned(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanAddDivU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = AddDiv(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddDivU(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanAddDivU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = AddDivU(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubDivReference(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanSubDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					if third == 0 {
						continue
					}

					result = (first + second) / third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubDiv(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanSubDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = SubDiv(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubDivRemReference(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanSubDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					if third == 0 {
						continue
					}

					result = (first + second) % third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubDivRem(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanSubDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = SubDivRem(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubDivUReference(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanSubDivU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					if third == 0 {
						continue
					}

					result = (first + second) / third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubDivOnlyUnsigned(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanSubDivU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = SubDiv(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkSubDivU(b *testing.B) {
	result := uint8(0)

	level1, level2, level3 := benchSpanSubDivU()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = SubDivU(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkShiftReference(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanShift()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				if second < 0 {
					continue
				}

				result = first << second
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkShift(b *testing.B) {
	result := int8(0)

	level1, level2 := benchSpanShift()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				result, _ = Shift(first, second)
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddSubDivReference(b *testing.B) {
	result := int8(0)

	level1, level2, level3, level4 := benchSpanAddSubDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						if fourth == 0 {
							continue
						}

						result = (first + second - third) / fourth
					}
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddSubDiv(b *testing.B) {
	result := int8(0)

	level1, level2, level3, level4 := benchSpanAddSubDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					for _, fourth := range level4 {
						result, _ = AddSubDiv(first, second, third, fourth)
					}
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddOneSubDivReference(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanAddOneSubDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					if third == 0 {
						continue
					}

					result = (first + 1 - second) / third
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddSubDivOnlyOne(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanAddOneSubDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = AddSubDiv(first, 1, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}

func BenchmarkAddOneSubDiv(b *testing.B) {
	result := int8(0)

	level1, level2, level3 := benchSpanAddOneSubDiv()

	b.ResetTimer()

	for range b.N {
		for _, first := range level1 {
			for _, second := range level2 {
				for _, third := range level3 {
					result, _ = AddOneSubDiv(first, second, third)
				}
			}
		}
	}

	require.NotNil(b, result)
}
