package piscine

import (
	
)
func ConcatParams(args []string) string {
	
	result:= ""
	for i, v:=range args{

			
			result += v 
			if i!=Lent3(args)-1{
				
				result += "\n"
			}
			
		
	}
	return result
}

func Lent3(d []string) int{
	inc:=0
	for _, _= range d{
		inc ++
	}
	 return inc
}