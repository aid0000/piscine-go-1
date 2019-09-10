package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	var results [1000]rune
	baseF := len(base)
	result := []rune(base)
	if baseF >= 2 && constrains(baseF, result) == true {
		i := 0
		pos := 0
		if baseF == 10 {
			fmt.Print(nbr)
		} else if nbr > 0 {
			for nbr > 0 {
				convert2(pos, i, nbr, baseF, result, results)
			}
			//print the result
			for j := i - 1; j >= 0; j-- {
				z01.PrintRune(results[j])
			}
		} else if nbr < 0 {
			nbr = -nbr
			for nbr > 0 {
				convert2(pos, i, nbr, baseF, result, results)
			}
			//print the result
			fmt.Print("-")
			for j := i - 1; j >= 0; j-- {
				z01.PrintRune(results[j])
			}
		}
	} else {
		results[0] = 'N'
		results[1] = 'V'
		for j := 0; j < 2; j++ {
			z01.PrintRune(results[j])
		}
	}
}

func constrains(baseF int, result []rune) bool {
	if baseF == 0 {
		return false
	}
	for a := 0; a < len(result); a++ {
		for x := a + 1; x < len(result); x++ {
			if baseF < 2 || result[a] == result[x] || result[a] == '+' || result[a] == '-' {
				return false

			}

		}
	}
	return true
}

func convert2(pos, i, nbr, baseF int, result []rune, results [1000]rune) {
	pos = nbr % baseF
	results[i] = result[pos]
	nbr = nbr / baseF
	i++
}