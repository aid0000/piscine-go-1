package piscine

func ToUpper(s string) string {
	h := []rune(s)
	result := ""
	for i := 0; i <= lent(h)-1; i++ {

		if (h[i] >= 'a') && (h[i] <= 'z') {

			h[i] = h[i] - 32

		}
		result += string(h[i])

	}
	return result
}
