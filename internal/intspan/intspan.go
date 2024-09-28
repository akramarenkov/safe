// Internal package used to determine integer variables parameters (minimum, maximum
// values ​​and bit size).
package intspan

import (
	"github.com/akramarenkov/safe/internal/intspec"
	"github.com/akramarenkov/safe/internal/is"
	"golang.org/x/exp/constraints"
)

// Returns minimum and maximum values of specified type.
func Get[Type constraints.Integer]() (Type, Type) {
	if !is.Signed[Type]() {
		return 0, ^Type(0) // return 0, Type(0)-1
	}

	refMaxInt64 := int64(intspec.MaxInt64)
	refMinInt64 := int64(intspec.MinInt64)
	refMaxInt32 := int32(intspec.MaxInt32)
	refMinInt32 := int32(intspec.MinInt32)
	refMaxInt16 := int16(intspec.MaxInt16)
	refMinInt16 := int16(intspec.MinInt16)
	refMaxInt8 := int8(intspec.MaxInt8)
	refMinInt8 := int8(intspec.MinInt8)

	convMaxInt64 := Type(refMaxInt64)
	convMinInt64 := Type(refMinInt64)
	convMaxInt32 := Type(refMaxInt32)
	convMinInt32 := Type(refMinInt32)
	convMaxInt16 := Type(refMaxInt16)
	convMinInt16 := Type(refMinInt16)
	convMaxInt8 := Type(refMaxInt8)
	convMinInt8 := Type(refMinInt8)

	switch {
	case int64(convMaxInt64) == refMaxInt64:
		return convMinInt64, convMaxInt64
	case int32(convMaxInt32) == refMaxInt32:
		return convMinInt32, convMaxInt32
	case int16(convMaxInt16) == refMaxInt16:
		return convMinInt16, convMaxInt16
	}

	return convMinInt8, convMaxInt8
}

// Returns bit size for specified type.
func BitSize[Type constraints.Integer]() int {
	refMaxInt64 := int64(intspec.MaxInt64)
	refMaxInt32 := int32(intspec.MaxInt32)
	refMaxInt16 := int16(intspec.MaxInt16)

	convMaxInt64 := Type(refMaxInt64)
	convMaxInt32 := Type(refMaxInt32)
	convMaxInt16 := Type(refMaxInt16)

	refMaxUint64 := uint64(intspec.MaxUint64)
	refMaxUint32 := uint32(intspec.MaxUint32)
	refMaxUint16 := uint16(intspec.MaxUint16)

	convMaxUint64 := Type(refMaxUint64)
	convMaxUint32 := Type(refMaxUint32)
	convMaxUint16 := Type(refMaxUint16)

	if is.Signed[Type]() {
		switch {
		case int64(convMaxInt64) == refMaxInt64:
			return intspec.BitSize64
		case int32(convMaxInt32) == refMaxInt32:
			return intspec.BitSize32
		case int16(convMaxInt16) == refMaxInt16:
			return intspec.BitSize16
		}

		return intspec.BitSize8
	}

	switch {
	case uint64(convMaxUint64) == refMaxUint64:
		return intspec.BitSize64
	case uint32(convMaxUint32) == refMaxUint32:
		return intspec.BitSize32
	case uint16(convMaxUint16) == refMaxUint16:
		return intspec.BitSize16
	}

	return intspec.BitSize8
}
