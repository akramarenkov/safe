// Internal package used to determine integer variables parameters (minimum, maximum
// values ​​and bit size).
package intspan

import (
	"github.com/akramarenkov/safe/internal/is"
	"golang.org/x/exp/constraints"
)

// To avoid including the entire math library because of a few constants.
const (
	bitSize64  = 64
	bitSize32  = 32
	bitSize16  = 16
	bitSize8   = 8
	bitSizeInt = bitSize32 << (^uint(0) >> (bitSize64 - 1))

	maxUint64 = 1<<bitSize64 - 1
	maxUint32 = 1<<bitSize32 - 1
	maxUint16 = 1<<bitSize16 - 1
	maxUint8  = 1<<bitSize8 - 1

	maxInt64 = 1<<(bitSize64-1) - 1
	minInt64 = -1 << (bitSize64 - 1)
	maxInt32 = 1<<(bitSize32-1) - 1
	minInt32 = -1 << (bitSize32 - 1)
	maxInt16 = 1<<(bitSize16-1) - 1
	minInt16 = -1 << (bitSize16 - 1)
	maxInt8  = 1<<(bitSize8-1) - 1
	minInt8  = -1 << (bitSize8 - 1)
)

// Returns minimum and maximum values of specified type.
func Get[Type constraints.Integer]() (Type, Type) {
	if !is.Signed[Type]() {
		return 0, ^Type(0) // return 0, Type(0)-1
	}

	refMaxInt64 := int64(maxInt64)
	refMinInt64 := int64(minInt64)
	refMaxInt32 := int32(maxInt32)
	refMinInt32 := int32(minInt32)
	refMaxInt16 := int16(maxInt16)
	refMinInt16 := int16(minInt16)
	refMaxInt8 := int8(maxInt8)
	refMinInt8 := int8(minInt8)

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
	refMaxInt64 := int64(maxInt64)
	refMaxInt32 := int32(maxInt32)
	refMaxInt16 := int16(maxInt16)

	convMaxInt64 := Type(refMaxInt64)
	convMaxInt32 := Type(refMaxInt32)
	convMaxInt16 := Type(refMaxInt16)

	refMaxUint64 := uint64(maxUint64)
	refMaxUint32 := uint32(maxUint32)
	refMaxUint16 := uint16(maxUint16)

	convMaxUint64 := Type(refMaxUint64)
	convMaxUint32 := Type(refMaxUint32)
	convMaxUint16 := Type(refMaxUint16)

	if is.Signed[Type]() {
		switch {
		case int64(convMaxInt64) == refMaxInt64:
			return bitSize64
		case int32(convMaxInt32) == refMaxInt32:
			return bitSize32
		case int16(convMaxInt16) == refMaxInt16:
			return bitSize16
		}

		return bitSize8
	}

	switch {
	case uint64(convMaxUint64) == refMaxUint64:
		return bitSize64
	case uint32(convMaxUint32) == refMaxUint32:
		return bitSize32
	case uint16(convMaxUint16) == refMaxUint16:
		return bitSize16
	}

	return bitSize8
}
