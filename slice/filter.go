package slice

// Filter
//
// Filters an array based on a predicate func.
//
// Example:
//
//	nums := []{1, 2, 3}
//	res := array.Filter(nums, func(num int) bool {
//	  return num % 2 == 0
//	})
//	fmt.Println(res) // [2]
func Filter[T any, Slice ~[]T](arr Slice, predicate func(item T, index int) bool) Slice {
	filtered := make(Slice, 0, len(arr))
	for index := range arr {
		if predicate(arr[index], index) {
			filtered = append(filtered, arr[index])
		}
	}
	return filtered
}
