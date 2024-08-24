package filler

import (
	"github.com/akramarenkov/safe/internal/inspect"
)

type Filler[Type inspect.EightBits] interface {
	Fill(args []Type, args64 []int64) (bool, error)
}
