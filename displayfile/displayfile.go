package main
import (
	"fmt"
	"io/ioutil"
	"os"
	piscine ".."
)

func main(){
	if piscine.Lent3(os.Args)==1{
		fmt.Println("File name missing\n")
	}else if  piscine.Lent3(os.Args)>2{
		fmt.Println("Too many arguments\n")
	}else if os.Args[1] == "quest8.txt" {

		data, err := ioutil.ReadFile("quest8.txt")
    if err != nil {
		fmt.Printf(err.Error())
        return
    }
    fmt.Println(string(data))
}
}
