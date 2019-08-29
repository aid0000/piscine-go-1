package piscine

import (
	
)
func ConcatParams(args []string) string {
	
	result:= ""
	for i, v:=range args{

			
			result += v 
			if i!=lent3(args)-1{
				
				result += "\n"
			}
			
		
	}
	return result
}

func lent3(d []string) int{
	inc:=0
	for _, _= range d{
		inc ++
	}
	 return inc
}