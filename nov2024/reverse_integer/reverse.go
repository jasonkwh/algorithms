package reverseinteger

import (
	"fmt"
	"math"
	"strconv"
)

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

	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0 // Return 0 for overflow
	}

	return reversed
}

func Reverse2(x int) int {
	negative := false

	if x < 0 {
		negative = true
		x = x * -1
	}

	s := strconv.Itoa(x)
	ts := ""
	for i := len(s) - 1; i >= 0; i-- {
		ts = fmt.Sprintf("%s%s", ts, string(s[i]))
	}

	if negative {
		ts = fmt.Sprintf("-%s", ts)
	}

	reversed, err := strconv.Atoi(ts)
	if err != nil {
		return 0
	}

	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0 // Return 0 for overflow
	}

	return reversed
}
