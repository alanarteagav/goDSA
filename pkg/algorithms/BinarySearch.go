package algorithms

import ds "goDSA/pkg/dataStructures"

// Binary Search algorithm
func BinarySearch[T ds.Ordered](s []T, item T) int {
	low := 0
	high := len(s) - 1
	for low <= high {
		mid := (low + high) / 2
		if s[mid] == item {
			return mid
		} else if item < s[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
