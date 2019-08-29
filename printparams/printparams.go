package main

import (
   piscine ".."
	"os"
	
)

func main() {
	
	for i, v:=range os.Args{
			if i!=0{

				piscine.Printrunes(v, 0)
			}
	}
}
