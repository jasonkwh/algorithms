package others

import "math"

// O(1)
func algorithm_one(arr []int) int {
	return 1 + arr[0]
}

// O(n)
func algorithm_two(arr []int) int {
	output := 0
	for _, val := range arr {
		output += val
	}
	return output
}

// O(n^2)
func algorithm_three(arr []int) []int {
	output := make([]int, 0)
	for _, val := range arr {
		for _, val2 := range arr {
			output = append(output, val)
			output = append(output, val2)
		}
	}
	return output
}

// theoretically a O(log(n)) algorithm
func find_the_clost(input int, arr []int) int {
	if input > arr[len(arr)-1] {
		return arr[len(arr)-1]
	}
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) > 1 {
		val1, val2 := 0, 0
		if len(arr) == 2 {
			val1, val2 = arr[0], arr[1]
		} else {
			val1, val2 = arr[int(len(arr)/2)-1], arr[int(len(arr)/2)]
		}
		output1 := math.Abs(float64(val1 - input))
		output2 := math.Abs(float64(val2 - input))
		if output1 < output2 {
			return find_the_clost(input, arr[:int(len(arr)/2)])
		} else {
			return find_the_clost(input, arr[int(len(arr)/2):])
		}
	}
	return 0
}
