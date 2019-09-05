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


























































































package main

import (
	"os"
	"fmt"
	// "io"
	"log"
	"strings"
	// "strconv"
	// "os/exec"
	// "io/ioutil"
	"bufio"
	// "bytes"
)
const FICHIER = "log.txt"

func main(){

	if len(os.Args)==1{
		file, err := os.Open(FICHIER)
		file2, _ := os.Open(FICHIER)
		file3, _ := os.Open(FICHIER)

	    if err != nil {
	        log.Fatal(err)
	    }
	    defer file.Close()

	    scanner := bufio.NewScanner(file)
	    scanner2 := bufio.NewScanner(file2)
	    scanner3 := bufio.NewScanner(file3)
	    compteur:=0

	    for scanner.Scan() {
	    	compteur++
	    }

	    i:=0
	    lastElmt := ""

	    for scanner2.Scan() {
	    	i++
	    	if compteur == i{
	    		lastElmt = fmt.Sprintf(scanner.Text())
	    	}
	    }

	    splitLastElmt := strings.Index(lastElmt," ")
	    lastElmt2 := []rune(lastElmt)
	    lastElmt = string(lastElmt2[splitLastElmt+1:])

	    resultat := []string{}

	    for scanner3.Scan() {
	    	if strings.Index(scanner3.Text(),lastElmt) != -1{
	    		resultat = append(resultat,scanner3.Text())
	    	}
	    }

	    fmt.Println(len(resultat))

	    for i,v := range resultat{
	    	if i<len(resultat)-1{
	    		fmt.Print(v," || ")
	    	}else{
	    		fmt.Println(v)
	    	}
	    }
	    if err := scanner.Err(); err != nil {
	        log.Fatal(err)
	    }
		if err != nil{
			return
		}

	}else{
		fmt.Println("Not a Raid function")
	}
}

// --- FAIL: TestRaid3 (2.40s)
//     raid3_test.go:82: ./student/raid3 raid1a/raid1a  prints "" instead of "[raid1a] [37] [19] || [raid1b] [37] [19] || [raid1c] [37] [19] || [raid1d] [37] [19] || [raid1e] [37] [19]\n"
//     raid3_test.go:82: ./student/raid3 raid1b/raid1b  prints "" instead of "[raid1a] [5] [26] || [raid1b] [5] [26] || [raid1c] [5] [26] || [raid1d] [5] [26] || [raid1e] [5] [26]\n"
//     raid3_test.go:82: ./student/raid3 raid1c/raid1c  prints "" instead of "[raid1a] [4] [8] || [raid1b] [4] [8] || [raid1c] [4] [8] || [raid1d] [4] [8] || [raid1e] [4] [8]\n"
//     raid3_test.go:82: ./student/raid3 raid1d/raid1d  prints "" instead of "[raid1a] [49] [45] || [raid1b] [49] [45] || [raid1c] [49] [45] || [raid1d] [49] [45] || [raid1e] [49] [45]\n"
//     raid3_test.go:82: ./student/raid3 raid1e/raid1e  prints "" instead of "[raid1a] [10] [33] || [raid1b] [10] [33] || [raid1c] [10] [33] || [raid1d] [10] [33] || [raid1e] [10] [33]\n"
// FAIL
// exit status 1
// FAIL	_/jail	2.403s

// --- FAIL: TestRaid3 (2.29s)
//     raid3_test.go:82: ./student/raid3 raid1a/raid1a  prints "[/jail/student/raid3/raid3]\n" 
//     instead of "[raid1a] [39] [33] || [raid1b] [39] [33] || [raid1c] [39] [33] || [raid1d] [39] [33] || [raid1e] [39] [33]\n"
//     raid3_test.go:82: ./student/raid3 raid1b/raid1b  prints "[/jail/student/raid3/raid3]\n" 
//     instead of "[raid1a] [35] [20] || [raid1b] [35] [20] || [raid1c] [35] [20] || [raid1d] [35] [20] || [raid1e] [35] [20]\n"
//     raid3_test.go:82: ./student/raid3 raid1c/raid1c  prints "[/jail/student/raid3/raid3]\n" 
//     instead of "[raid1a] [27] [4] || [raid1b] [27] [4] || [raid1c] [27] [4] || [raid1d] [27] [4] || [raid1e] [27] [4]\n"
//     raid3_test.go:82: ./student/raid3 raid1d/raid1d  prints "[/jail/student/raid3/raid3]\n" 
//     instead of "[raid1a] [13] [27] || [raid1b] [13] [27] || [raid1c] [13] [27] || [raid1d] [13] [27] || [raid1e] [13] [27]\n"
//     raid3_test.go:82: ./student/raid3 raid1e/raid1e  prints "[/jail/student/raid3/raid3]\n" 
//     instead of "[raid1a] [45] [16] || [raid1b] [45] [16] || [raid1c] [45] [16] || [raid1d] [45] [16] || [raid1e] [45] [16]\n"
// FAIL
// exit status 1
// FAIL	_/jail	2.290s











vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv


//gives the last element of the list
func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	for l.Head != nil {
		if l.Head.Next == nil {
			return l.Head.Data
		}
		l.Head = l.Head.Next
	}
	return l.Head.Data
}

vvvvvvvvvvvvvvvvvvvvvvvvvv
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



aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
package piscine

func IsSorted(f func(int, int) int, arr []int) bool {

	cesce := true
	decresce := true

	for i := 1; i < Lent4(arr); i++ {
		if !(f(arr[i-1], arr[i]) >= 0) {
			cesce = false
		}
	}

	for i := 1; i < Lent4(arr); i++ {
		if !(f(arr[i-1], arr[i]) <= 0) {
			decresce = false
		}
	}

	return cesce || decresce

}
func F(a,b int) int{
	if a>b{
		return 1
	}else if a==b{
		return 0
	}else {
		return -1
	}
}
func Lent4(d []int) int{
	inc:=0
	for _, _= range d{
		inc ++
	}
	 return inc
}