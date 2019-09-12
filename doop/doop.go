package main

import (
	"fmt"
	"os"
	"strconv"
	piscine".."
)

func isSignal(str string, arr[]string) bool{
	for _,v := range arr{
		if str == v{
			return true
		}
	}
	return false
}

func main(){
	signal := []string{"+","*","-","/","%"}
	args := os.Args[1:]
	if piscine.Lent3(args) != 3 {

	}else{
		if isSignal(args[1], signal){
			nb1, err := strconv.Atoi(args[0])
			nb2, err2 := strconv.Atoi(args[2])
			if err == nil && err2 == nil{
				switch args[1] {
				case "+":
					fmt.Println(nb1 + nb2)
				case "-":
					fmt.Println(nb1 - nb2)
				case "/":
					if nb2 == 0 {
						fmt.Println("No division by 0")
					}else{
						fmt.Println(nb1 / nb2)
					}
				case "%":
					if nb2 == 0 {
						fmt.Println("No Modulo by 0")
					}else{
						fmt.Println(nb1 % nb2)
					}
				case "*":
					fmt.Println(nb1 * nb2)

				}
			}else{
				fmt.Println("1")
			}
		}else{
			fmt.Println("0")
		}
	}
	
}