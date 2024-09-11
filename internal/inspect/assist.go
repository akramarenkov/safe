package inspect

import (
	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/intspan"
)

// Converts maximum and minimum values for specified type to reference type.
func ConvSpan[Type types.UpToUSI32, TypeRef types.SIF64]() (TypeRef, TypeRef) {
	min, max, _ := intspan.Get[Type]()
	return TypeRef(min), TypeRef(max)
}
