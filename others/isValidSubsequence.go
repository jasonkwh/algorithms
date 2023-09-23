package others

func IsValidSubsequence(array []int, sequence []int) bool {
	for _, v := range sequence {
		if !contains(array, v) {
			return false
		}
	}
	return true
}

func contains(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func IsValidSubsequence2(array []int, sequence []int) bool {
	testArr := array
	for i := 0; i < len(sequence); i++ {
		index, contain := contains2(testArr, sequence[i], len(array)-len(testArr))
		if contain {
			testArr = array[index+1:]
			continue
		} else {
			return false
		}
	}
	return true
}

func contains2(arr []int, target, arrayDiff int) (int, bool) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i + arrayDiff, true
		}
	}
	return -1, false
}
