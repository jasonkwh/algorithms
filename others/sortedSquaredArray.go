package others

import "sort"

func SortedSquaredArray(array []int) []int {
	for k, v := range array {
		array[k] = v * v
	}

	sort.Ints(array)
	return array
}
