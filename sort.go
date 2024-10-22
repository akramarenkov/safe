package safe

import (
	"slices"

	"golang.org/x/exp/constraints"
)

func sortAddM[Type constraints.Integer](items []Type) {
	// The number is chosen based on benchmark readings
	const smallLimit = 13

	if len(items) < smallLimit {
		sortAddMSmall(items)
		return
	}

	slices.Sort(items)
}

// Faster on 5-30% than direct using of [slices.Sort] on the number of items
// from 1 to 12 due to simplicity.
func sortAddMSmall[Type constraints.Integer](items []Type) {
	for first := 1; first < len(items); first++ {
		for second := first; second > 0 && items[second] < items[second-1]; second-- {
			items[second], items[second-1] = items[second-1], items[second]
		}
	}
}

func sortMulM[Type constraints.Integer](items []Type) {
	// The number is chosen based on benchmark readings
	const smallLimit = 17

	if len(items) < smallLimit {
		sortMulMSmall(items)
		return
	}

	slices.SortFunc(items, compareMulM)
}

// Faster on 5-60% than direct using of [slices.SortFunc] on the number of items
// from 1 to 16 due to simplicity.
func sortMulMSmall[Type constraints.Integer](items []Type) {
	for first := 1; first < len(items); first++ {
		for second := first; second > 0; second-- {
			// items[second] < 0 && items[second-1] < 0
			if items[second]&items[second-1] < 0 {
				if items[second] < items[second-1] {
					break
				}
			} else {
				if items[second] > items[second-1] {
					break
				}
			}

			items[second], items[second-1] = items[second-1], items[second]
		}
	}
}

func compareMulM[Type constraints.Integer](a, b Type) int {
	// a < 0 && b < 0
	if a&b < 0 {
		if a < b {
			return 1
		}

		return -1
	}

	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	}

	return 0
}
