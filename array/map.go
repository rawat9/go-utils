package array

// Map
//
// Applies a function `fn` to each element of an array and returns a new array with the results.
//
// Example:
//
//	nums := []{1, 2, 3}
//	res := array.Map(nums, func(num int) int {
//	  return num * 2
//	})
//	fmt.Println(res) // [2, 4, 6]
func Map[T Type](arr []T, fn func(item T, index int) T) []T {
	result := make([]T, len(arr))

	for index := range arr {
		result[index] = fn(arr[index], index)
	}

	return result
}
