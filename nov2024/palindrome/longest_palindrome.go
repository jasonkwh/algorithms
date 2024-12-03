package palindrome

// not working for case "ccc"
func LongestPalindrome(s string) string {
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if i == j {
				continue
			}

			if s[i] == s[j] {
				return s[i : j+1]
			}
		}
	}

	if len(s) >= 1 {
		return string(s[0])
	}

	return ""
}
