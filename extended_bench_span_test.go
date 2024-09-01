package safe

import "testing"

// For benchmarking with full coverage.
func benchSpanAdd3() []int8 {
	return []int8{127, 126, -128, -127}
}

// These test functions are used to determine the coverage level of the functions
// being benchmarked.
func TestBenchSpanAdd3(*testing.T) {
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

func TestBenchSpanAddM(*testing.T) {
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

func TestBenchSpanAddUM(*testing.T) {
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

func TestBenchSpanSub3(*testing.T) {
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

func TestBenchSpanSubUM(*testing.T) {
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
