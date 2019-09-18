package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

)

func counter(output []rune) (x, y int) {
	countX := 0
	countY := 0
	flag := true
	for _, s := range output {
		if s == '\n' {
			countY++
			flag = false
		} else {
			if flag == true {
				countX++
			}
		}
	}
	return countX, countY
}

func determineInput(output []rune) {
	x, y := counter(output)
	arrOfRaids := []string{"raid1a", "raid1b", "raid1c", "raid1d", "raid1e"}
	formula1 := ((x * y) + y - 1) - x
	formula2 := (x * y) + y - 2
	count := 0

	sort.Sort(sort.StringSlice(arrOfRaids))
	for i := 0; i < len(arrOfRaids); i++ {
		if count > 0 {
			fmt.Print(" || ")
		}
		if (output[0] == 'o' || arrOfRaids[i] == "raid1a") && count < 1 {
			fmt.Printf("[raid1a] [%d] [%d]", x, y)
			count++
		} else if (output[0] == '/' || arrOfRaids[i] == "raid1b") && count < 2 {
			fmt.Printf("[raid1b] [%d] [%d]", x, y)
			count++
		} else if ((output[0] == 'A' && output[formula1] == 'C' && output[formula2] == 'C') || arrOfRaids[i] == "raid1c") && count < 3 {
			fmt.Printf("[raid1c] [%d] [%d]", x, y)
			count++
		} else if ((output[0] == 'A' && output[formula1] == 'A' && output[formula2] == 'C') || arrOfRaids[i] == "raid1d") && count < 4 {
			fmt.Printf("[raid1d] [%d] [%d]", x, y)
			count++
		} else if ((output[0] == 'A' && output[formula1] == 'C' && output[formula2] == 'A') || arrOfRaids[i] == "raid1e") && count < 5 {
			fmt.Printf("[raid1e] [%d] [%d]", x, y)
			count++
		}
	}
	fmt.Println()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var output []rune
	for {
		input, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		output = append(output, input)
	}

	if output[0] != 'o' && output[0] != '/' && output[0] != 'A' {
		fmt.Println("Not a Raid function")
		return
	}
	determineInput(output)
}



