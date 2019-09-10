package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"os/exec"
)

func main(){
	info, err := os.Stdin.Stat()
    	if err != nil {
        	fmt.Println(err.Error())
    	}
	
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println()
		return
	}
	
	reader := bufio.NewReader(os.Stdin)
	var output []rune
	
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}
	
	row, col:= getDimension(string(output))
	fmt.Println(IsRaid1a(row, col))
}

func IsRaid1a (row, col int) string {
	cmd, _ := exec.Command("raid1a", string(row), string(col)).Output()
	return string(cmd)
}


func getDimension(s string) (int, int) {
	col, row:= 0, 0
	for _, res := range s {
		if row == 0 && res != '\n' {
			col++
		}
		if res == '\n' {
			row++
		}
	}
	return row, col
}















 
package piscine

func AtoiBase(s string, base string) int {
	baseF := len(base)
	result := []rune(base)
	pos := 0

	if constrains1(baseF, result) == true {
		for i := 0; i < len(s); i++ {
			pos *= len(base)
			pos += Index1(s, base)
		}
		return pos
	} else {

		return 0
	}
	return 0
}

func constrains1(baseF int, result []rune) bool {

	if baseF == 0 {
		return false
	}
	for a := 0; a < len(result); a++ {
		for x := a + 1; x < len(result); x++ {
			if baseF < 2 || result[a] == result[x] || result[x] == '+' || result[x] == '-' {
				return false
			}
		}
	}
	return true
}

func Index1(s string, toFind string) int {
	arrayStr := []rune(s)
	array := []rune(toFind)
	for j := 0; j < len(s); j++ {
		for x := 0; x < len(toFind); x++ {

			if arrayStr[j] == array[x] { // ex: s[0:3] para dividir a string to 0 atil 3
				return x
			}
		}
	}

	return -1
}
































		







package piscine
prntnbm
import (
	"fmt"
)

func mostrar(n int, table [9]int, tmax [9]int) {
	i := 0
	for i < n {
		fmt.Print(table[i])
		i++
	}
	if table[0] != tmax[0] {
		fmt.Print(", ")
	}
}

func printCombination() {
	table := [9]int{0}
	tmax := [9]int{9}
	for table[0] <= tmax[0] {
		mostrar(1, table, tmax)
		table[0]++
	}
}

func PrintCombN(n int) {
	table := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	tmax := [9]int{}

	if n == 1 {
		printCombination()
	} else {
		i := n - 1
		j := 9
		for i >= 0 {
			tmax[i] = j
			i--
			j--
		}
		i = n - 1
		for table[0] < tmax[0] {
			if table[i] != tmax[i] {
				mostrar(n, table, tmax)
				table[i]++
			}
			if table[i] == tmax[i] {
				if table[i-1] != tmax[i-1] {
					mostrar(n, table, tmax)
					table[i-1]++
					j = i
					for j < n {
						table[j] = table[j-1] + 1
						j++
					}
					i = n - 1
				}
			}
			for table[i] == tmax[i] && table[i-1] == tmax[i-1] && i > 1 {
				i--
			}
		}
		mostrar(n, table, tmax)
	}
	fmt.Println()













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









sort params















package main

import (
	"fmt"
	"os"
)

func main() {
	array := []string(os.Args[1:])

	quicksort(array)
	for i := 0; i < len(array); i++ {

		fmt.Println(array[i])
	}
}




//quicksort
func quicksort(table []string) {
	beg := 0
	end := len(table) - 1

	quickSrot(table, beg, end)
}
func quickSrot(table []string, beg int, end int) {

	if beg < end {
		lockPivo := move(table, beg, end)
		quickSrot(table, beg, lockPivo-1)
		quickSrot(table, lockPivo+1, end)

	}
}

func move(table []string, beg int, end int) int {
	pivo := table[end]
	i := beg - 1

	for j := beg; j < end; j++ {
		if table[j] <= pivo {
			i++
			aux := table[j]
			table[j] = table[i]
			table[i] = aux
		}
	}

	aux := table[end]
	table[end] = table[i+1]
	table[i+1] = aux

	return i + 1
}












	if value < 2 {
		return false
	}
	limit := int(math.Floor(math.Sqrt(float64(value))))
	i := 2
	for i <= limit {
		if value%i == 0 {
			return false
		}
		i++
	}
	return true
}



