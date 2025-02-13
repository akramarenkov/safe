package safe

import (
	"slices"

	"golang.org/x/exp/constraints"
)

// Faster on 5-60% than direct using of [slices.SortFunc] on the number of items
// from 1 to 16 due to simplicity and low nesting.
func sortMulM[Type constraints.Integer](items []Type) {
	// The number is chosen based on benchmark readings
	const advantageLimit = 16

	if len(items) > advantageLimit {
		slices.SortFunc(items, compareMulM)
		return
	}

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
