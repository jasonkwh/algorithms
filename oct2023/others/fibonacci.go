package others

func GetNthFib(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		// fibonacci formula
		// f(n) == f(n-1)+f(n-2)
		return GetNthFib(n-1) + GetNthFib(n-2)
	}
}
