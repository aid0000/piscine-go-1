package piscine

import (
	"strings"
)
func AtoiBase(s string, base string) int {


	indx:=0
	for _,res:= range base {
		if string(res) == "-" || string(res) == "+" || strings.Count(base, string(res)) > 1 { //nao pode ser menos de 2 numeros a base e nao aceita +-
			indx = 1
			break
		}

	}
	if indx == 1 || lent2(base) < 2{ //se for menos de 2 a base, retorna zero
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
	