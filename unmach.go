package piscine

func Unmatch(arr []int) int {

	for i := 0; i < Lent4(arr); i++ {
		for j := i + 1; j < Lent4(arr); j++ {
			if arr[i] == arr[j] {
				arr[i] = 0
				arr[j] = 0
				break
			}
		}
	}
	for k := 0; k < Lent4(arr); k++ {
		if arr[k] != 0 {

			return arr[k]
		}
	}
return 0
}
