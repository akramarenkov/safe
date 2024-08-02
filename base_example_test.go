package safe_test

import (
	"fmt"

	"github.com/akramarenkov/safe"
)

func ExampleAdd() {
	sum, err := safe.Add[int8](124, 3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sum)

	sum, err = safe.Add[int8](125, 3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sum)

	// Output:
	// 127
	// integer overflow
	// 0
}
