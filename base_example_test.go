package safe_test

import (
	"fmt"

	"github.com/akramarenkov/safe"
)

func ExampleAdd() {
	sum, err := safe.Add[int8](124, 3)
	if err != nil {
		panic(err)
	}

	fmt.Println(sum)

	_, err = safe.Add[int8](125, 3)
	if err == nil {
		panic("expected overflow")
	}

	// Output: 127
}
