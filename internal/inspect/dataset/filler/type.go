package filler

import (
	"github.com/akramarenkov/safe/internal/inspect"
)

// Fill arguments of dataset item by values.
type Filler[Type inspect.EightBits] interface {
	Fill(args []Type, args64 []int64) (bool, error)
}
