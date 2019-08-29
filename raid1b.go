package piscine

import (
	"github.com/01-edu/z01"
)



func Raid1b(x, y int) {
	a := "/"
	b := "*"
	c := "\\"
	d := " "

	if y > 1 {
		
		partedecima(a, b, c, x)

		lados(x, y, b, d)
		partedebaixo(b, c, a, x)
	} else {
		partedecima(a, b, c, x)
	}
}

func partedecima(a, b, c string, x int) {
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

func partedebaixo(b, c, a string, x int) {
	if x > 1 {
		for i := x; i <= x; i++ {
			Printrunes2(c,0)
			for f := 1; f <= x-2; f++ {
				Printrunes2(b,0)

			}
		}
		Printrunes2(a,0)
	} else {
		Printrunes2(c,0)
	}
	z01.PrintRune('\n')

}

func lados1(x, y int, b, d string) {
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
