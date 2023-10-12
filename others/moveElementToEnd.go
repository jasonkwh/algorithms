package others

func MoveElementToEnd(array []int, toMove int) []int {
	if len(array) == 0 {
		return []int{}
	}
	output := make([]int, 0)
	occur := make([]int, 0)

	for _, v := range array {
		if v == toMove {
			occur = append(occur, v)
		} else {
			output = append(output, v)
		}
	}
	output = append(output, occur...)
	return output
}
