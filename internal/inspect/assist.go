package inspect

import (
	"github.com/akramarenkov/safe/internal/inspect/types"

	"github.com/akramarenkov/intspec"
)

// Converts maximum and minimum values for specified type to reference type.
func ConvSpan[Type types.UpToUSI32, TypeRef types.SIF64]() (TypeRef, TypeRef) {
	minimum, maximum := intspec.Range[Type]()
	return TypeRef(minimum), TypeRef(maximum)
}
