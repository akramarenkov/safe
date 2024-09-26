package inspect

import (
	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/intspan"
)

// Converts maximum and minimum values for specified type to reference type.
func ConvSpan[Type types.UpToUSI32, TypeRef types.SIF64]() (TypeRef, TypeRef) {
	minimum, maximum := intspan.Get[Type]()
	return TypeRef(minimum), TypeRef(maximum)
}
