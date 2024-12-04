package reverseinteger

func Reverse(x int) int {
	reversed := 0
	negative := false

	if x < 0 {
		negative = true
		x = x * -1
	}

	for x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}

	if negative {
		reversed = reversed * -1
	}

	temp := int32(reversed)

	if int(temp) != reversed {
		return 0
	}

	return reversed
}
