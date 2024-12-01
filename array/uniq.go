package array

import (
	"slices"
)

type Type interface {
	~string | ~int | ~float64
}

// Uniq returns a new array with all duplicate values removed.
func Uniq[T Type, Slice ~[]T](slice Slice) Slice {
	slices.Sort(slice)
	return slices.Compact(slice)
}
