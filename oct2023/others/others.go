package others

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
