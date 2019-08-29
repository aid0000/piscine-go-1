package piscine

import (
	"github.com/01-edu/z01"
	
)

func Raid1d(x, y int) {
	a := "A"
	b := "B"
	c := "C"
	d := " "
	if y > 1 {
		
		partedecima2(a, b, c, x)

		lados3(x, y, b, d)
		
		partedecima2(a, b, c, x)
	} else {
		partedecima2(a, b, c, x)
	}
}

func partedecima2(a, b, c string, x int) {
	if x > 1 {
		for i := x; i <= x; i++ {
			Printrunes2(a,0)
			for f := 1; f <= x-2; f++ {
				Printrunes2(b,0)

			}
		}
		Printrunes2(c,0)
	} else {
		Printrunes2(a,0)
	}
	z01.PrintRune('\n')
}

func lados3(x, y int, b, d string) {
	if x > 1 {
		for l := 0; l < y-2; l++ {

			Printrunes2(b,0)
			for h := 0; h < x-2; h++ {
				Printrunes2(d,0)

			}
			Printrunes2(b,0)
			z01.PrintRune('\n')
		}
	} else {
		for l := 0; l < y-2; l++ {

			Printrunes2(b,0)
			z01.PrintRune('\n')
		}

	}
}
