package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	size := len(os.Args)
	/*	if size == 1 {
		for { //could use fmt.scanf but not sure if the student can use it
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			fmt.Print(data)
		}
	} else if len(os.Args) > 1 {*/
	for i := 1; i < size; i++ {
		data, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(data))
		fmt.Println()
	}

	//}

}