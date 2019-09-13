package piscine

func Max(arr []int) int {
	result := arr[0]
	for i := 0; i < Lent4(arr); i++ {

		if result < arr[i] {
			result = arr[i]
		}

	}
	return result
}
