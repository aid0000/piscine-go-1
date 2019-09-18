package piscine

func SortWordArr(array []string) {

	quickSrot2(array, 0, len(array)-1)

}


func quickSrot2(table []string, beg int, end int) {

	if beg < end {
		lockPivo := mudaVariavel2(table, beg, end)
		quickSrot2(table, beg, lockPivo-1)
		quickSrot2(table, lockPivo+1, end)

	}
}

func mudaVariavel2(table []string, beg int, end int) int {
	pivo := table[end]
	i := beg - 1

	for j := beg; j < end; j++ {
		if table[j] <= pivo {
			i++
			aux := table[j]
			table[j] = table[i]
			table[i] = aux
		}
	}

	aux := table[end]
	table[end] = table[i+1]
	table[i+1] = aux

	return i + 1
}



/*
	var str string
	for i := 0; i < Lent3(array); i++ {
		str = str + array[i]

	}
	arr := []rune(str)*/
	//	for p := 0; p < len(arr); p++ {

	//	array[p] = string(arr[p])

	//}