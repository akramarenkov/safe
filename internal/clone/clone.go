// Internal package with slightly faster, for this module, cloning function.
package clone

// Creates a shallow copy of a slice via the copy function. Faster on 5-30% than append
// on the number of items from 1 to about 128 of type int.
func Slice[Type any](slice []Type) []Type {
	copied := make([]Type, len(slice))

	copy(copied, slice)

	return copied
}
