package piscine

func IsPrintable(str string) bool {
h := []rune(str)

for i := 0; i <= lent(h)-1; i++ {
	if (h[i] >= 0) && (h[i] <= 31) {
		return false
	}
		
}
return true
}