package safe_test

import (
	"fmt"

	"github.com/akramarenkov/safe"
)

func ExampleIter() {
	for number := range safe.Iter[int8](126, 127) {
		fmt.Println(number)
	}
	// Output:
	// 126
	// 127
}

func ExampleIterStep() {
	for _, number := range safe.IterStep[int8](126, 127, 2) {
		fmt.Println(number)
	}
	// Output:
	// 126
}
