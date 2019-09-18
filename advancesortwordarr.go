package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	
	for i:=1; i<Lent3(array) ;i++{
		if f(array[i], array[i-1]) > 0 {
			array[i]=array[i-1]                  //anda uma posicao
			array[i-1]=array[i]					//recua uma posicao
			i=0
		}

	}
}


