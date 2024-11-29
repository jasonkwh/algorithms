package twosum

func twoSum(nums []int, target int) []int {
	output := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if i+1 == len(nums) {
			break
		}

		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				output = append(output, i)
				output = append(output, j)
			}
		}
	}

	return output
}
