package piscine

import (
	"fmt"
)
func PrintNbrBase(nbr int, base string) () {



var binary []int
   i:=0
	for nbr > 0 {
		 binary[i]= nbr% 2
		 nbr = nbr/2
		 i++
	}
	for j:=i-1;j>=0;j++{
		fmt.Println(i)
	}
}