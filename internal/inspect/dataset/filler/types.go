package filler

import "github.com/akramarenkov/safe/internal/inspect/confines"

// Fill arguments of dataset item by values.
type Filler[Type confines.UpToI32] interface {
	Fill(args []Type, args64 []int64) (bool, error)
}
