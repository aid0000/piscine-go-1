package piscine

import (
	"fmt"
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

func partedebaixo3(b, c, a string, x int) {
	if x > 1 {
		for i := x; i <= x; i++ {
			fmt.Print(c)
			for f := 1; f <= x-2; f++ {
				fmt.Print(b)

			}
		}
		fmt.Print(a)
	} else {
		fmt.Print(c)
	}
	fmt.Print("\n")

}

func lados4(x, y int, b, d string) {
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
