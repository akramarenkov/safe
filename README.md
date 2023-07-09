# Safe

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/safe.svg)](https://pkg.go.dev/github.com/akramarenkov/safe)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/safe)](https://goreportcard.com/report/github.com/akramarenkov/safe)
[![codecov](https://codecov.io/gh/akramarenkov/safe/branch/master/graph/badge.svg?token=C1AZ5V2ZT7)](https://codecov.io/gh/akramarenkov/safe)

## Purpose

Library that allows you to detect overflows in operations with integer numbers

## Usage

Example:

```go
package main

import (
    "fmt"

    "github.com/akramarenkov/safe"
)

func main() {
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
```
