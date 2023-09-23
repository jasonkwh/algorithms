package others

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	for k, v := range array {
		for k2, v2 := range array {
			if k != k2 && v+v2 == target {
				return []int{v, v2}
			}
		}
	}
	return []int{}
}

func TwoNumberSum2(array []int, target int) []int {
	// Write your code here.
	for k, v := range array {
		if len(array[:k]) != 0 {
			for _, v2 := range array[:k] {
				if v+v2 == target {
					return []int{v, v2}
				}
			}
		}
		if k+1 == len(array) {
			break
		}
		if len(array[k+1:len(array)-1]) != 0 {
			for _, v2 := range array[k+1 : len(array)-1] {
				if v+v2 == target {
					return []int{v, v2}
				}
			}
		}
	}
	return []int{}
}
