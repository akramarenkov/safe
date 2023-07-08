package safe_test

import (
	"fmt"

	"github.com/akramarenkov/safe"
)

func ExampleSumInt() {
	sum, err := safe.SumInt[int8](3, 124)
	if err != nil {
		panic(err)
	}

	fmt.Println(sum)

	_, err = safe.SumInt[int8](3, 125)
	if err == nil {
		panic("expected overflow")
	}
	// Output: 127
}
