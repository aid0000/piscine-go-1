package piscine
func CountIf(f func(string) bool, tab []string) int {
numberelements:=0
	for _, s := range tab{
		if f(s) == true {
			numberelements++
		}
	}
	return numberelements
}
