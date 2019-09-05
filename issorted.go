package piscine

func IsSorted(f func(int, int) int, arr []int) bool {

	cesce := true
	decresce := true

	for i := 1; i < Lent4(arr); i++ {
		if !(f(arr[i-1], arr[i]) >= 0) {
			cesce = false
		}
	}

	for i := 1; i < Lent4(arr); i++ {
		if !(f(arr[i-1], arr[i]) <= 0) {
			decresce = false
		}
	}
	return cesce || decresce

}
func F(a,b int) int{
	if a>b{
		return 1
	}else if a==b{
		return 0
	}else {
		return -1
	}
}
func Lent4(d []int) int{
	inc:=0
	for _, _= range d{
		inc ++
	}
	 return inc
}