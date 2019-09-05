package main

/*import (
	"os"
	"fmt"
	"bufio"
	"io"
	"os/exec"
)
*/
func main(){
	
}


//func israid1()
//{/
	




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
