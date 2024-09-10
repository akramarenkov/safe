// Internal package used to get integers span (minimum and maximum values).
package intspan

import (
	"github.com/akramarenkov/safe/internal/is"
	"golang.org/x/exp/constraints"
)

// To avoid including the entire math library because of a few constants.
const (
	MaxUint64 = 1<<64 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint8  = 1<<8 - 1

	MinInt64 = -1 << 63
	MaxInt64 = 1<<63 - 1
	MinInt32 = -1 << 31
	MaxInt32 = 1<<31 - 1
	MinInt16 = -1 << 15
	MaxInt16 = 1<<15 - 1
	MinInt8  = -1 << 7
	MaxInt8  = 1<<7 - 1
)

func Get[Type constraints.Integer]() (Type, Type) {
	refMinInt64 := int64(MinInt64)
	refMaxInt64 := int64(MaxInt64)
	refMinInt32 := int64(MinInt32)
	refMaxInt32 := int64(MaxInt32)
	refMinInt16 := int64(MinInt16)
	refMaxInt16 := int64(MaxInt16)
	refMinInt8 := int64(MinInt8)
	refMaxInt8 := int64(MaxInt8)

	convMinInt64 := Type(refMinInt64)
	convMaxInt64 := Type(refMaxInt64)
	convMinInt32 := Type(refMinInt32)
	convMaxInt32 := Type(refMaxInt32)
	convMinInt16 := Type(refMinInt16)
	convMaxInt16 := Type(refMaxInt16)
	convMinInt8 := Type(refMinInt8)
	convMaxInt8 := Type(refMaxInt8)

	refMaxUint64 := uint64(MaxUint64)
	refMaxUint32 := uint64(MaxUint32)
	refMaxUint16 := uint64(MaxUint16)
	refMaxUint8 := uint64(MaxUint8)

	convMaxUint64 := Type(refMaxUint64)
	convMaxUint32 := Type(refMaxUint32)
	convMaxUint16 := Type(refMaxUint16)
	convMaxUint8 := Type(refMaxUint8)

	if is.Signed[Type]() {
		switch {
		case int64(convMaxInt64) == refMaxInt64:
			return convMinInt64, convMaxInt64
		case int64(convMaxInt32) == refMaxInt32:
			return convMinInt32, convMaxInt32
		case int64(convMaxInt16) == refMaxInt16:
			return convMinInt16, convMaxInt16
		}

		return convMinInt8, convMaxInt8
	}

	switch {
	case uint64(convMaxUint64) == refMaxUint64:
		return 0, convMaxUint64
	case uint64(convMaxUint32) == refMaxUint32:
		return 0, convMaxUint32
	case uint64(convMaxUint16) == refMaxUint16:
		return 0, convMaxUint16
	}

	return 0, convMaxUint8
}
