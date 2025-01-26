package maxarea

func MaxArea(height []int) int {
	if len(height) == 2 {
		mh := findMinHeight(height[0], height[1])

		return mh
	}

	area1 := 0

	for i := 0; i < len(height); i++ {
		for j := len(height) - 1; j >= 0; j-- {
			minHeight := findMinHeight(height[i], height[j])

			tempArea := minHeight * (j - i)
			if tempArea > area1 {
				area1 = tempArea
			}
		}
	}

	return area1
}

func findMinHeight(heightA, heightB int) int {
	minHeight := 0

	if heightA <= heightB {
		minHeight = heightA
	}

	if heightB < heightA {
		minHeight = heightB
	}

	return minHeight
}
