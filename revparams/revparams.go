package main

import (
    piscine ".."
	"os"
	
)

func main() {
	
	for i:=lent(os.Args)-1; i>=1; i--{
		
		piscine.Printrunes(os.Args[i], 0)
	}
}
func lent(d []string) int{
	inc:=0
	for _, _= range d{
		inc ++
	}
	 return inc
}

