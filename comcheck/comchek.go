package main

import (
	piscine ".."
	"fmt"
	"os"
)

func main() {
	args := []string(os.Args)
	str := []string{"01", "galaxy", "galaxy 01"}
	var result string

	for k := 0; k < piscine.Lent3(str); k++ {
		for i := 1; i < piscine.Lent3(args); i++ {
			result = args[i]
			for _, s := range str {
				if result == s {
					fmt.Println("Alert!!!")
					return

				}
			}

		}
	}
}
