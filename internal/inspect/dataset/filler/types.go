package filler

import "github.com/akramarenkov/safe/internal/inspect/types"

// Fill arguments of dataset item by values.
type Filler[Type types.UpToI32] interface {
	Fill(args []Type, args64 []int64) (bool, error)
}
