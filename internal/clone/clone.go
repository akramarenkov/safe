// Internal package with slightly optimized cloning function for this module.
package clone

// Creates a shallow copy of a slice. Slightly faster than append on small number of
// elements.
func Slice[Type any](slice []Type) []Type {
	copied := make([]Type, len(slice))

	copy(copied, slice)

	return copied
}
