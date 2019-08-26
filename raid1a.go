package piscine

import (
	"fmt"
)

func Raid1a(x, y int) {
	a := "o"
	b := "-"
if y>1 {


parteCimaBaixo(a,b, x)

	//-------------------------------------------------

	c := "|"
	d := " "
lados(x,y,c,d)
	parteCimaBaixo(a,b, x)
}else {
	parteCimaBaixo(a,b,x)
}
	
}
func parteCimaBaixo(a,b string, x int ){
	if x>1{
		for i := x; i <= x; i++ {
			fmt.Print(a)
			for f := 1; f <= x-2; f++ {
				fmt.Print(b)
	
			}
		}
		fmt.Print(a)
	} else {
		fmt.Print(a)
	}
	fmt.Print("\n")
}
func lados(x, y int, c,d string ){
	if x>1{
		for l := 0; l < y-2; l++ {
	
			fmt.Print(c)
				for h := 0; h < x-2; h++ {
					fmt.Print(d)
	
			}
			fmt.Print(c)
			fmt.Print("\n")
		}
	}else {
		for l := 0; l < y-2; l++ {
	
			fmt.Print(c)
			fmt.Print("\n")
		}
	
	}
}