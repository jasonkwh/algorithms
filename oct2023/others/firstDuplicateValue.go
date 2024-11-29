package others

func FirstDuplicateValue(array []int) int {
	if len(array) <= 1 {
		return -1
	}
	interactMap := make(map[int]bool, len(array))
	for i := 0; i < len(array); i++ {
		if interactMap[i] || i == len(array)-1 {
			continue
		}
		for j := 0; j < len(array); j++ {
			if i == j || interactMap[j] {
				continue
			}
			if array[i] == array[j] && !interactMap[j] {
				interactMap[j] = true
			}
		}
	}
	if len(interactMap) != 0 {
		smallest := 0
		count := 0
		for k := range interactMap {
			if count == 0 || k < smallest {
				smallest = k
			}
			count++
		}
		return array[smallest]
	}
	return -1
}
