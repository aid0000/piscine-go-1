package piscine

import (
	"fmt"
)

func Raid1a(x, y int) {

	a := "o"
	b := "-"

	for i := x; i <= x; i++ {
		fmt.Print(a)
		for f := 1; f <= x-2; f++ {
			fmt.Print(b)

		}
	}
	fmt.Print(a)

	fmt.Print("\n")

	//-------------------------------------------------

	c := "|"
	d := " "

	for l := 1 + 2; l <= y; l++ {

		fmt.Print(c)
		for h := y; h <= x; h++ {
			fmt.Print(d)

		}
		fmt.Print("\n")
	}
	for i := x; i <= x; i++ {
		fmt.Print(a)
		for f := 1; f <= x-2; f++ {
			fmt.Print(b)

		}
	}
	fmt.Print(a)

	fmt.Print("\n")
}
