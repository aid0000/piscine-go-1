package piscine

func IsUpper(str string) bool {
h := []rune(str)

for i := 0; i <= lent(h)-1; i++ {
	if (h[i] >= 0) && (h[i] <= 64) || (h[i] >= 91) && (h[i] <= 127){
		return false
	}
		
}
return true
}