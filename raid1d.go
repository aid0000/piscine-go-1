package piscine

import (
	"fmt"
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
			fmt.Print(a)
			for f := 1; f <= x-2; f++ {
				fmt.Print(b)

			}
		}
		fmt.Print(c)
	} else {
		fmt.Print(a)
	}
	fmt.Print("\n")
}

func lados3(x, y int, b, d string) {
	if x > 1 {
		for l := 0; l < y-2; l++ {

			fmt.Print(b)
			for h := 0; h < x-2; h++ {
				fmt.Print(d)

			}
			fmt.Print(b)
			fmt.Print("\n")
		}
	} else {
		for l := 0; l < y-2; l++ {

			fmt.Print(b)
			fmt.Print("\n")
		}

	}
}
