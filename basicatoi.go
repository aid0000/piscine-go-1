package piscine

func BasicAtoi(s string) int {

	arrayStr := []rune(s)
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		if arrayStr[i] < '0' || arrayStr[i] > '9' {

			return ans
		} else {
			ans *= 10
			ans += int(arrayStr[i]) - '0'
		}

	}
	return ans
}