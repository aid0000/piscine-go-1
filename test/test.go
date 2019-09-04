package main

import (
	"fmt"
	piscine ".."
)

func main() {
	str := "Hello\thow\nare you?"
	fmt.Println(piscine.SplitWhiteSpaces(str))
}