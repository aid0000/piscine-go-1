package piscine

import "fmt"

const N = 8

var position = [N]int{}

func talivre(queennumber, rowposition int) bool {
	for i := 0; i < queennumber; i++ { //verifica todas as rainhas e as suas posicoes para entao avancar
		other_row_pos := position[i] //vai buscar a posicao de outra rainha

		if other_row_pos == rowposition || //mesma linha
			other_row_pos == rowposition-(queennumber-i) || //diagonais
			other_row_pos == rowposition+(queennumber-i) {
			return false
		}
	}
	return true
}

func resolverpuzzle(k int) {

	if k == N {
		for i := 0; i < N; i++ { //gera as possibilidades
			fmt.Print(position[i] + 1)
		}
		fmt.Print("\n")
	} else {
		for i := 0; i < N; i++ {
			if talivre(k, i) { //verifica se ta livre antes de colocar a rainha na posicao k
				position[k] = i
				//coloca outra rainha
				resolverpuzzle(k + 1)
			}
		}
	}
}
func EightQueens() {
	resolverpuzzle(0)
}
