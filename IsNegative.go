package main

import (
        "github.com/01-edu/z01"
)

func main() {
        var a int =-9
        if a<=0 {
                z01.PrintRune('T')
        }else{
                z01.PrintRune('F')
        }
        z01.PrintRune('\n')
}