package palindrome

import "math"

func IsPalindrome(x int) bool {
	newNum := 0
	tempX := x
	count := 0

	for tempX > 0 {
		count++

		tempX = tempX / 10
	}

	tempX = x

	for tempX > 0 {
		digit := tempX % 10
		newNum += digit * int(math.Pow(10, float64(count-1)))

		count--
		tempX = tempX / 10
	}

	return x == newNum
}
