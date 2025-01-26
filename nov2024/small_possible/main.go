package smallpossible

func Solution(A []int) int {
	smallestNum := 0
	largestNum := 0
	negatives := make([]bool, 0)

	for _, num := range A {
		if smallestNum == 0 || num < smallestNum {
			smallestNum = num
		}

		if largestNum == 0 || num > largestNum {
			largestNum = num
		}

		if num < 0 {
			negatives = append(negatives, true)
		} else {
			negatives = append(negatives, false)
		}
	}

	for _, negative := range negatives {
		if !negative {
			break
		} else {
			return 1
		}
	}

	diffarr := make([]int, 0)

	for i := smallestNum; i < largestNum; i++ {
		sameVal := false

		for _, num := range A {
			if num == i && !sameVal {
				sameVal = true

				break
			}
		}

		if !sameVal {
			diffarr = append(diffarr, i)
		}
	}

	if len(diffarr) == 0 {
		return largestNum + 1
	}

	smallestNum = 0

	for _, num := range diffarr {
		if smallestNum == 0 || num < smallestNum {
			smallestNum = num
		}
	}

	return smallestNum
}
