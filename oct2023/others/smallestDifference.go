package others

import "math"

func SmallestDifference(array1, array2 []int) []int {
	small1, small2, diff := 0, 0, 0.0

	for _, v1 := range array1 {
		for _, v2 := range array2 {
			tempdiff := math.Abs(float64(v1 - v2))

			if small1 == 0 || small2 == 0 || tempdiff < diff {
				small1 = v1
				small2 = v2
				diff = tempdiff
				continue
			}
		}
	}

	return []int{small1, small2}
}
