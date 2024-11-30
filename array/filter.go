package array

/*
Filter function

Usage:

	nums := []{1, 3, 4}
	A.Filter(nums)
*/
func Filter(arr []int, predicate func(int) bool) []int {
	filtered := make([]int, 0)
	for _, num := range arr {
		if predicate(num) {
			filtered = append(filtered, num)
		}
	}
	return filtered
}
