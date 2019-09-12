package main

import (
	"fmt"
	"io/ioutil"
	"os"
	piscine".."
)

func main() {
	size := piscine.Lent3(os.Args)
	
	for i := 1; i < size; i++ {
		data, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(data))
		fmt.Println()
	}

	

}