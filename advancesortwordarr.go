package piscine

import (
	//"fmt"
)


func AdvancedSortWordArr(array []string, f func(a, b string) int) {

	quickSrot2(array, 0, len(array)-1)
}



/*	
	for i:=1; i<Lent3(array) ;i++{
		if f(array[i], array[i-1]) > 1 {
			array[i]=array[i-1]                  //anda uma posicao
			array[i-1]=array[i]					//recua uma posicao
			i=0
		}
		
	}*/
/*
var str string
for i := 0; i < Lent3(array); i++ {
	str = str + array[i]
}


arr:=array
quickSrot2(array, 0,f(str, string(len(array)-1)))
fmt.Println(arr) */