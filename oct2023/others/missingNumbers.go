package others

func MissingNumbers(nums []int) []int {
	output := make([]int, 0)

	for i := 1; i <= len(nums)+2; i++ {
		found := false

		for j := 0; j < len(nums); j++ {
			if i == nums[j] {
				found = true
				break
			}
		}

		if !found {
			output = append(output, i)
		}
	}
	return output
}
