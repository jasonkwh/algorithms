package lengthoflongestsubstring

// not working for case "aab"
func LengthOfLongestSubstring(s string) int {
	length := 0

	for i := 0; i < len(s); i++ {
		if i+1 == len(s) {
			break
		}

		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i > length {
					length = j - i
				}

				break
			}
		}
	}

	return length
}
