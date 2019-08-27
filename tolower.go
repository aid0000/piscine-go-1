package piscine

func ToLower(s string) string {
	h := []rune(s)
	result := ""
	for i := 0; i <= lent(h)-1; i++ {

		if (h[i] >= 'A') && (h[i] <= 'Z') {

			h[i] = h[i] + 32

		}
		result += string(h[i])

	}
	return result
}
