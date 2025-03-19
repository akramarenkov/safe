// Internal package with constraints used when inspecting.
package constraints

// Constraints by 8-bit integer types.
type I8 interface {
	~int8 | ~uint8
}

// Constraints by up to 32-bit integer types.
type UpToI32 interface {
	I8 | ~int16 | ~uint16 | ~int32 | ~uint32
}

// Constraints by 64-bit signed integer and floating point types.
type IF64 interface {
	~int64 | ~float64
}
