package main

import (
	piscine".."
	"fmt"
)


func main(){
	
	i:=0
     base:= 2
	
	
	for i=0; i<=25; i++{

		fmt.Printf("The decimal number %d in base %d is %d  \n", i,base, DecToBase( base, i) )
		
		
		
	}
}

/*
     n = number in base 10 = 4
     base = the base that we want to convert to = 2
     4 /2  =  2  Remainder= 0
     2 / 2 = 1   Remainder= 0
     1 / 2 = 0   Remainder= 1
     
     4 in binary is 1 * 100 + 0 *10 + 0 *1 = 100
     
     4 in binary is 100
*/


func DecToBase(base, n int) int{

	num:=0// the number that we want to return in the new base
	i:= 0 // the iterative variable used for the place 
	quotient:= n // the quotient of the number in base 10 
	remainder := 0 // the remainder of the number in base 10 
	
	for quotient != 0{
		
		remainder = quotient%base
		quotient = quotient / base
		
		num = piscine.IterativePower(10,i)*remainder + num
		i++
	}
	
	return num
}

