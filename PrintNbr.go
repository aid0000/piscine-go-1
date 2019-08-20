package piscine

import (
		"github.com/01-edu/z01"
		
)

func PrintNbr(n int) {
	//var a [2]string
	//a :=[3] string
	a:= string(n)
	ar := []rune(a)
	for i := 0; i < len(ar); i++ {
		
		z01.PrintRune(ar[i])
	}
	z01.PrintRune('\n')
}
