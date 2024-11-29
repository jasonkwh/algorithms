package others

import "sort"

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)
	min := 0

	for _, coin := range coins {
		if coin > min+1 {
			break
		}
		min += coin
	}
	return min + 1
}
