package piscine

func SortIntegerTable(table []int) {
	beg := 0
	end := len(table) - 1

	quickSrot(table, beg, end)
}
func quickSrot(table []int, beg int, end int) {

	if beg < end {
		lockPivo := mudaVariavel(table, beg, end)
		quickSrot(table, beg, lockPivo-1)
		quickSrot(table, lockPivo+1, end)

	}
}

func mudaVariavel(table []int, beg int, end int) int {
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