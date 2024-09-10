package inspect

import (
	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/intspan"
)

// Picks up maximum and minimum values for specified type.
func PickUpSpan[Type types.UpToUSI32, TypeRef types.SIF64]() (TypeRef, TypeRef) {
	min, max := intspan.Get[Type]()
	return TypeRef(min), TypeRef(max)
}
