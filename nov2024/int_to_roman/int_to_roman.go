package inttoroman

import "fmt"

func IntToRoman(num int) string {
	ms := num / 1000
	num = num % 1000
	cs := num / 100
	num = num % 100
	tens := num / 10
	num = num % 10

	return fmt.Sprintf("%s%s%s%s",
		mToRoman(ms),
		cToRoman(cs),
		tenToRoman(tens),
		numToRoman(num))
}

func mToRoman(ms int) string {
	if ms == 0 {
		return ""
	}

	output := "M"

	for i := 1; i < ms; i++ {
		output = fmt.Sprintf("M%s", output)
	}

	return output
}

func cToRoman(cs int) string {
	if cs == 0 {
		return ""
	}

	output := ""

	if cs < 5 {
		if cs == 4 {
			return "CD"
		}

		output = "C"

		for i := 1; i < cs; i++ {
			output = fmt.Sprintf("C%s", output)
		}
	}

	if cs >= 5 {
		if cs == 9 {
			return "CM"
		}

		output = "D"
		for i := 5; i < cs; i++ {
			output = fmt.Sprintf("%sC", output)
		}
	}

	return output
}

func tenToRoman(ten int) string {
	if ten == 0 {
		return ""
	}

	output := ""

	if ten < 5 {
		if ten == 4 {
			return "XL"
		}

		output = "X"

		for i := 1; i < ten; i++ {
			output = fmt.Sprintf("X%s", output)
		}
	}

	if ten >= 5 {
		if ten == 9 {
			return "XC"
		}
		output = "L"

		for i := 5; i < ten; i++ {
			output = fmt.Sprintf("%sX", output)
		}
	}

	return output
}

func numToRoman(num int) string {
	if num == 0 {
		return ""
	}

	output := ""

	if num < 5 {
		if num == 4 {
			return "IV"
		}

		output = "I"

		for i := 1; i < num; i++ {
			output = fmt.Sprintf("I%s", output)
		}
	}

	if num >= 5 {
		if num == 9 {
			return "IX"
		}
		output = "V"

		for i := 5; i < num; i++ {
			output = fmt.Sprintf("%sI", output)
		}
	}

	return output
}
