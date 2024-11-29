package others

func SpiralTraverse(array [][]int) []int {
	length := len(array) * len(array[0])
	h, v := 0, 0
	output := make([]int, 0)
	right, down, left, up := true, false, false, false
	rightM := len(array[0]) - 1
	downM := len(array) - 1
	leftM := 0
	upM := 0

	for i := 0; i < length; i++ {
		output = append(output, array[v][h])

		if h == rightM && right {
			upM++
			right = false
			down = true
		}
		if v == downM && down {
			rightM--
			down = false
			left = true
		}
		if h == leftM && left {
			downM--
			left = false
			up = true
		}
		if v == upM && up {
			leftM++
			up = false
			right = true
		}
		if right {
			h++
			continue
		}
		if down {
			v++
			continue
		}
		if left {
			h--
			continue
		}
		if up {
			v--
			continue
		}
	}

	return output
}
