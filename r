package piscine

func Unmatch(arr []int) int {
	result:=4
	for i:=range arr{
		if !ispair(arr[i],i, arr){
			result=arr[i]
			return result		
		}
	}
	return result
}

func ispair(elt,index int, arr[]int) bool{
	for i:=range arr{
		if i!=index{
			if arr[i]==elt{
				return true			
			}
		}
	}
	return false	
}    










package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return NbrBase(AtoiBase(nbr, baseFrom), baseTo)
}

func NbrBase(nbr int, str string) string {
	indx := 0
	for _, res := range str {
		if string(res) == "-" || string(res) == "+" {
			indx = 1
			break
		}
	}
	if indx == 1 || len(str) < 2 {
		return "0"
	} else if 2147483647 < nbr || -2147483648 > nbr {
		return string(nbr)
	} else {
		neg := ""
		if nbr < 0 {
			neg = "-"
			nbr *= -1
		}
		i := 0
		smt := ""
		for nbr >= lent2(str) {
			if nbr >= lent2(str) {
				smt += string(str[nbr%lent2(str)])
				nbr = nbr / lent2(str)
				i++
			}
		}
		smt += string(str[nbr])
		return neg + Reverse(smt)
	}
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






----------------------------------------------------------------------------------------

import (
	"math"
	"strings"
)

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return  NbrBase(AtoiBase(nbr, baseFrom), baseTo)
}

func NbrBase(nbr int, str string) string {
	indx := 0
	for _,res:= range str {
		if string(res) == "-" || string(res) == "+" || strings.Count(str, string(res)) > 1 {
			indx = 1
			break
		}
	}
	if indx == 1 || len(str) < 2{
		return "0"
	}else if 2147483647  <nbr || -2147483648 > nbr{
		return string(nbr)
	}else{
		neg := ""
		if nbr < 0 {
			neg = "-"
			nbr *= -1	
		}
		i:=0
		nan:= ""
		for nbr >= len(str) {
			if nbr >= len(str) {
				nan +=string(str[nbr % len(str)])
				nbr = nbr/len(str)
				i++
			}
		}
		nan +=string(str[nbr])
		return neg+Reverse(nan)
	}
}



rrrrrrrrrrrrrrrrrrrrr

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
		
		fmt.Println(output)
	}
	row, col:= getsize(string(output))
	fmt.Println(IsRaid1a(row, col))
}

func IsRaid1a (row, col int) string {
	cmd, _ := exec.Command("raid1a", string(row), string(col)).Output()
	return string(cmd)
}


func getsize(s string) (int, int) {
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








ztail------------------------------------------------------------
package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
	"strconv"
)

func numberOfBytes(args []string) (int, []string) {
	n := len(args)
	nbytes := 0
	var files []string
	for i, v := range args {
		var err error
		_, err = strconv.Atoi(v)
		if v == "-c" {
			if i >= n-1 {
				fmt.Printf("tail: option requires an argument -- 'c'\nTry 'tail --help' for more information.")
				os.Exit(1)
			}
			arg := args[i+1]

			nbytes, err = strconv.Atoi(arg)

			if err != nil {
				fmt.Printf("tail: invalid number of bytes: %s\n", arg)
				os.Exit(1)
			}
			continue
		}

		if err != nil {
			files = append(files, v)
		}

	}
	return nbytes, files
}

func fileSize(fi *os.File) int64 {
	fil, err := fi.Stat()

	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return fil.Size()
}

func main() {

	n := len(os.Args)
	if n < 4 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	nbytes, files := numberOfBytes(os.Args[1:])

	printName := len(files) > 1

	//open files for reading only
	for j, f := range files {
		fi, err := os.Open(f)
		if err != nil {
			fmt.Printf("tail: cannot open '%s' for reading: No such file or directory\n", f)
			os.Exit(1)
		}
		if printName {
			fmt.Printf("==> %s <==\n", f)
			
		}
		read := make([]byte, int(nbytes))
		_, er := fi.ReadAt(read, fileSize(fi)-int64(nbytes))
		if er != nil {
			fmt.Println(er.Error())
		}

		for _, c := range read {
			z01.PrintRune(rune(c))
		}

		if j < len(files)-1 {
			z01.PrintRune('\n')
		}

		fi.Close()
	}
}