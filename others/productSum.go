package others

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	output := 0
	for i := 0; i < len(array); i++ {
		output += productSum(array[i], 1)
	}
	return output
}

func productSum(item interface{}, depth int) int {
	switch v := item.(type) {
	case int:
		return v
	case SpecialArray:
		depth++
		output := 0
		for i := 0; i < len(v); i++ {
			output += productSum(v[i], depth)
		}
		return depth * output
	default:
		return 0
	}
}
