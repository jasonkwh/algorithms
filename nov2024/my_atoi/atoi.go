package myatoi

import (
	"math"
)

func MyAtoi(s string) int {
	negative := false
	output := 0

	for i := len(s) - 1; i >= 0; i-- {
		switch string(s[i]) {
		case "-":
			negative = true
			continue
		case "+":
			continue
		case "1":
			output += outputNum(1, len(s), i)
			continue
		case "2":
			output += outputNum(2, len(s), i)
			continue
		case "3":
			output += outputNum(3, len(s), i)
			continue
		case "4":
			output += outputNum(4, len(s), i)
			continue
		case "5":
			output += outputNum(5, len(s), i)
			continue
		case "6":
			output += outputNum(6, len(s), i)
			continue
		case "7":
			output += outputNum(7, len(s), i)
			continue
		case "8":
			output += outputNum(8, len(s), i)
			continue
		case "9":
			output += outputNum(9, len(s), i)
			continue
		case "0":
			continue
		default:
		}

		break
	}

	if negative {
		output = output * -1
	}

	return output
}

func outputNum(num, slen, i int) int {
	return num * int(math.Pow(10, float64(slen-1-i)))
}
