package piscine

import (
	"github.com/01-edu/z01"
	
)

func Raid1e(x, y int) {
	a := "A"
	b := "B"
	c := "C"
	d := " "

	if y > 1 {
		
		partedecima3(a, b, c, x)

		lados4(x, y, b, d)
		partedebaixo3(b, c, a, x)
	} else {
		partedecima3(a, b, c, x)
	}
}

func partedecima3(a, b, c string, x int) {
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

func partedebaixo3(b, c, a string, x int) {
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

func lados4(x, y int, b, d string) {
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
