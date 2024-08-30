package inspect

import (
	"math"

	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/is"
)

// Picks up maximum and minimum values for specified type.
func PickUpSpan[Type types.UpToUSI32, TypeRef types.SIF64]() (TypeRef, TypeRef) {
	if is.Signed[Type]() {
		return pickUpSpanSigned[Type, TypeRef]()
	}

	return pickUpSpanUnsigned[Type, TypeRef]()
}

func pickUpSpanSigned[Type types.UpToUSI32, TypeRef types.SIF64]() (TypeRef, TypeRef) {
	reference := TypeRef(math.MaxInt32)

	if TypeRef(Type(reference)) == math.MaxInt32 {
		return math.MinInt32, math.MaxInt32
	}

	reference = TypeRef(math.MaxInt16)

	if TypeRef(Type(reference)) == math.MaxInt16 {
		return math.MinInt16, math.MaxInt16
	}

	return math.MinInt8, math.MaxInt8
}

func pickUpSpanUnsigned[Type types.UpToUSI32, TypeRef types.SIF64]() (TypeRef, TypeRef) {
	reference := TypeRef(math.MaxUint32)

	if TypeRef(Type(reference)) == math.MaxUint32 {
		return 0, math.MaxUint32
	}

	reference = TypeRef(math.MaxUint16)

	if TypeRef(Type(reference)) == math.MaxUint16 {
		return 0, math.MaxUint16
	}

	return 0, math.MaxUint8
}
