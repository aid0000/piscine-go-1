package piscine
func Capitalize(s string) string {

	h := []rune(s)
	result := ""
	for i := 0; i <= lent(h)-1; i++ {
		if (h[i] >= 0) && (h[i] <= 47) {

		
		if (h[i+1] >= 'a') && (h[i+1] <= 'z') {

			h[i+1] = h[i+1] - 32


		}
	}
		result += string(h[i])
	
	}
	return result

}