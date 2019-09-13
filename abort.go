package piscine

import (
	
)
func Abort(a, b, c, d, e int) int {
	arg := []int{a, b, c, d, e}
	SortIntegerTable(arg)			//organiza os args ordem+
	return arg[2]                    //porque e o arg do meo
}
