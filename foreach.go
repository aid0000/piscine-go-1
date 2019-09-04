package piscine

import (
	"github.com/01-edu/z01"
)

func ForEach(f func(int), arr []int) {
	for _, s := range arr {
		
		f(s)
	}
	z01.PrintRune('\n')
}