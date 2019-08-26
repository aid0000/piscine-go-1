package piscine

import (
	"fmt"
)

func Raid1a(x, y int) {

	a := "o"
	b := "-"

	for i := x; i <= x; i++ {
		fmt.Print(a)
		for f := +3; f <= x; f++ {
			fmt.Print(b)

		}
	}
	fmt.Print(a)
	fmt.Print("\n")

	c := "|"
	d := " "

	for g := y; g <= y; g++ {
		fmt.Print(c)
		for g := +4; g <= y; g++ {
			fmt.Print("\n", c)

			for h := +3; h <= x; h++ {
				fmt.Print(d)

			}
		}
	}

	fmt.Print("\n")

	for i := x; i <= x; i++ {
		fmt.Print(a)
		for f := +3; f <= x; f++ {
			fmt.Print(b)

		}
	}
	fmt.Print(a)
	fmt.Print("\n")

}