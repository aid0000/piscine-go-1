package piscine

import (
	"github.com/01-edu/z01"
)
func Printrunes(a string, s int){

	if a!= "" && s==0{
		p:=[]rune(a)
		for _,l:= range p{
			z01.PrintRune(l)

		}
	} else {
		PrintNbr(s)
	}
	z01.PrintRune('\n')
}  