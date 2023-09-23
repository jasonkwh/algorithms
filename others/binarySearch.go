package others

// theoretically a O(log(n)) algorithm
func BinarySearch(array []int, target int) int {
	return binarySearch(array, target, 0)
}

func binarySearch(array []int, target, base int) int {
	if len(array) == 1 && array[0] == target {
		return base
	}
	if len(array) > 1 {
		midNum := array[int(len(array)/2)]
		if target < midNum {
			return binarySearch(array[:int(len(array)/2)], target, base+0)
		} else if target > midNum {
			return binarySearch(array[int(len(array)/2):], target, base+int(len(array)/2))
		} else {
			return base + int(len(array)/2)
		}
	}
	return -1
}
