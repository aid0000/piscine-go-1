package piscine

func Atoi(s string) int {

	ar := []rune(s)
	n := len(s)
	res := 0
	neg := 1
	for i := 0; i < n; i++ {
		if ar[i] == ' ' {
			return 0
		}
		if ar[i] == '-' && i != 0 {
			return 0
		}

		if ar[i] < '0' && ar[i] != '-' && ar[i] != '+' || ar[i] > '9' && ar[i] != '-' && ar[i] != '+' {
			return 0
		} else {
			if ar[i] == '+' || ar[i] == '-' {
				if ar[i] == '-' {
					neg = -1
				}
				if len(ar) > 1 {
					i++

					if ar[i] == '+' || ar[i] == '-' {
						return 0
					}
				}
			}
			res *= 10
			res += int(ar[i]) - '0'
			if i == n-1 {
				res *= neg

			}
		}
	}
	return res
}