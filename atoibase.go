package piscine

import (
	"strings"
)
func AtoiBase(s string, base string) int {


	indx:=0
	for _,res:= range base {
		if string(res) == "-" || string(res) == "+" {
			indx = 1
			break
		}

	}
	if indx == 1 || lent2(base) < 2{
		return 0	
	}else{
		result := 0
		for i,res:= range s {
			ind := strings.Index(base,string(res))
			result += ind* RecursivePower(lent2(base),lent2(s)-1 - i)
		}
		return result
	}
}
	