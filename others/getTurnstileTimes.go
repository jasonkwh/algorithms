package others

func GetTimes(time []int32, direction []int32) []int32 {
	result := make([]int32, len(time))

	endTime := int32(0)
	for _, t := range time {
		if t > endTime {
			endTime = t
		}
	}
	if endTime == 0 {
		return result
	}

	in := make([]int, 0)
	out := make([]int, 0)
	currentIndex := 0
	turnstile := 0

	for i := int32(0); i <= endTime; i++ {
		for currentIndex < len(time) && time[currentIndex] == i {
			if direction[currentIndex] == 0 {
				in = append(in, currentIndex)
			}
			if direction[currentIndex] == 1 {
				out = append(out, currentIndex)
			}
			currentIndex++
		}

		if len(in)+len(out) == 0 {
			turnstile = 0
			continue
		}
		if (turnstile == 1 || len(out) == 0) && len(in) > 0 {
			result[in[0]] = i
			in = in[1:]
			turnstile = 1
			continue
		}
		if len(out) > 0 {
			result[out[0]] = i
			out = out[1:]
			turnstile = 2
		}
	}
	for len(in) > 0 {
		result[in[0]] = endTime + 1
		in = in[1:]
		endTime++
	}
	return result
}
