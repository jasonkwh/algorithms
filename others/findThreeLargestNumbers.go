package others

import "sort"

func FindThreeLargestNumbers(array []int) []int {
	if len(array) < 3 {
		return nil
	}

	output := make([]int, 3)
	for _, num := range array {
		if output[0] == 0 {
			output[0] = num
			continue
		}
		if output[1] == 0 {
			output[1] = num
			continue
		}
		if output[2] == 0 {
			output[2] = num
			continue
		}
		if num > output[0] || num > output[1] || num > output[2] {
			smallIndex := findSmallestInOutputArray(output)
			output[smallIndex] = num
			continue
		}
	}

	sort.Ints(output)
	return output
}

func findSmallestInOutputArray(array []int) int {
	smallest := array[0]
	key := 0
	for k, num := range array {
		if num < smallest {
			smallest = num
			key = k
		}
	}
	return key
}

func FindThreeLargestNumbers2(array []int) []int {
	if len(array) < 3 {
		return nil
	}

	output := make([]int, 3)
	for _, num := range array {
		if output[0] == 0 {
			output[0] = num
			continue
		}
		if output[1] == 0 {
			output[1] = num
			continue
		}
		if output[2] == 0 {
			output[2] = num
			continue
		}
		if num > output[0] || num > output[1] || num > output[2] {
			sort.Ints(output)
			output[0] = num
			continue
		}
	}

	sort.Ints(output)
	return output
}
