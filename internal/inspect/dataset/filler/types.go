package filler

import "github.com/akramarenkov/safe/internal/inspect/constraints"

// Fill arguments of dataset item by values.
type Filler[Type constraints.UpToI32] interface {
	Fill(args []Type, args64 []int64) (bool, error)
}
