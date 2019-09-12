package piscine

import(
	
	"fmt"
	"strings"
)

func Reverse(s string) string {
    var reverse string
    for i := lent2(s)-1; i >= 0; i-- {
        reverse += string(s[i])
    }
    return reverse 
}

func PrintNbrBase(nbr int, str string)(){
	indx := 0
	for _,res:= range str {
		if string(res) == "-" || string(res) == "+" || strings.Count(str, string(res)) > 1 {
			indx = 1
			break
		}
	}
	if indx == 1 || lent2(str) < 2{
		fmt.Print("NV")	
	}else if 2147483647  <nbr || -2147483648 > nbr{
		fmt.Print(int64(nbr))
	}else{
		if nbr < 0 {
			fmt.Print("-")
			nbr *= -1	
		}
		i:=0
		nan:=""
		for nbr >= lent2(str) {
			if nbr >= lent2(str) {
				nan +=string(str[nbr % lent2(str)])
				nbr = nbr/lent2(str)
				i++
			}
		}
		nan +=string(str[nbr])
		fmt.Print(Reverse(nan))
	}
}