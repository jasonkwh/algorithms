package romantoint

func RomanToInt(s string) int {
	output := 0
	skipNext := false

	for i := 0; i < len(s); i++ {
		if skipNext {
			skipNext = false
			continue
		}
		if string(s[i]) == "M" {
			output += 1000
		}
		if string(s[i]) == "D" {
			output += 500
		}
		if string(s[i]) == "C" {
			if i != (len(s) - 1) {
				if string(s[i+1]) == "M" {
					output += 900
					skipNext = true
					continue
				}
				if string(s[i+1]) == "D" {
					output += 400
					skipNext = true
					continue
				}
			}
			output += 100
		}
		if string(s[i]) == "L" {
			output += 50
		}
		if string(s[i]) == "X" {
			if i != (len(s) - 1) {
				if string(s[i+1]) == "C" {
					output += 90
					skipNext = true
					continue
				}
				if string(s[i+1]) == "L" {
					output += 40
					skipNext = true
					continue
				}
			}
			output += 10
		}
		if string(s[i]) == "V" {
			output += 5
		}
		if string(s[i]) == "I" {
			if i != (len(s) - 1) {
				if string(s[i+1]) == "X" {
					output += 9
					break
				}
				if string(s[i+1]) == "V" {
					output += 4
					break
				}
			}
			output += 1
		}
	}

	return output
}
