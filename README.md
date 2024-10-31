# Safe

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/safe.svg)](https://pkg.go.dev/github.com/akramarenkov/safe)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/safe)](https://goreportcard.com/report/github.com/akramarenkov/safe)
[![Coverage Status](https://coveralls.io/repos/github/akramarenkov/safe/badge.svg)](https://coveralls.io/github/akramarenkov/safe)

## Purpose

Library that allows you to detect and avoid overflows in operations with integer numbers

## Usage

Example:

```go
package main

import (
    "fmt"

    "github.com/akramarenkov/safe"
)

func main() {
    sum, err := safe.Add[int8](124, 3)
    fmt.Println(err)
    fmt.Println(sum)

    sum, err = safe.Add[int8](125, 3)
    fmt.Println(err)
    fmt.Println(sum)
    // Output:
    // <nil>
    // 127
    // integer overflow
    // 0
}
```
