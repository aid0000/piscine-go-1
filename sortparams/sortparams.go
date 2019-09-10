package main

import (
	"fmt"
	"os"
)

func main() {
	array := []string(os.Args[1:])

	quicksort(array)
	for i := 0; i < len(array); i++ {

		fmt.Println(array[i])
	}
}


func quicksort(table []string) {
	beg := 0
	end := len(table) - 1

	Srot2(table, beg, end)
}
func Srot2(table []string, beg int, end int) {

	if beg < end {
		lockPivo := move(table, beg, end)
		Srot2(table, beg, lockPivo-1)
		Srot2(table, lockPivo+1, end)

	}
}

func move(table []string, beg int, end int) int {
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












	

