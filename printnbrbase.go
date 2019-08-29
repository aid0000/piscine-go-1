package piscine

import (
	
	"fmt"
)
func PrintNbrBase(nbr int, base string) () {



var binary [100]int
   i:=0
   
   
   if lent2(base)==10{
	
	   fmt.Println(nbr)
	   
   }else {
	for nbr > 0 {
		 binary[i]= nbr% lent2(base)
		 nbr = nbr/2
		 i++
	}
	for j:=i-1;j>=0;j--{
		
		fmt.Print(binary[j])
	}
   }

}
func lent2(d string) int{
	inc:=0
	for _, _= range d{
		inc ++
	}
	 return inc
}
