package piscine

import "fmt"

const N = 8

var position = [N]int{}

func talivre(queen_number, row_position int) bool {
	for i := 0; i < queen_number; i++ {
		other_row_pos := position[i]

		if other_row_pos == row_position || other_row_pos == row_position-(queen_number-i) || other_row_pos == row_position+(queen_number-i) {
			return false
		}
	}
	return true
}

func resolverpuzzle(k int) {
	if k == N {
		for i := 0; i < N; i++ {
			fmt.Print(position[i] + 1)
		}
		fmt.Print("\n")
	} else {
		for i := 0; i < N; i++ {
			if talivre(k, i) {
				position[k] = i
				resolverpuzzle(k + 1)
			}
		}
	}
}

func EightQueens() {
	resolverpuzzle(0)
}