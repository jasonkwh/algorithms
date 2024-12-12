package regexmatch

// passed 267/354 test cases
func IsMatch(s string, p string) bool {
	for i := 0; i < len(s); i++ {
		if i == len(p)-1 {
			if string(p[i]) == "*" {
				return true
			}

			if i != len(s)-1 {
				return false
			}
		}

		if string(p[i]) == "." {
			continue
		}

		if s[i] == p[i] {
			continue
		}

		return false
	}

	return true
}
