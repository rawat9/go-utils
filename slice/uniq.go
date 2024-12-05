package slice

import (
	"slices"
)

// Uniq returns a new array with all duplicate values removed.
func Uniq[T ~string | ~int | ~float32 | ~float64, Slice ~[]T](slice Slice) Slice {
	slices.Sort(slice)
	return slices.Compact(slice)
}
