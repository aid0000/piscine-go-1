package piscine

func Split(str, charset string) []string {
	resp := []string{}
	new := ""
		for i := 0; i < len(str); i++ {
			if Charset(str, charset, i) {
				if new != "" {
					resp = append(resp, new)
					new = ""
					i = i + len(charset) - 1
				}
			} else {
				new = new + string(str[i])
			}
		}
		if new != "" {
			resp = append(resp, new)
		}
		return resp
}


func Charset(str, charset string, i int) bool {

	j := 0
		for j < len(charset) && i < len(str) {
			if str[i] != charset[j] {
				return false
			}
			i++
			j++
		}
		return true
}

/*1º--> saber e edentificar o charset. fazer funcao para quando aparecer charset na str retorna true
2º-->criar []string vazio para retornar
3º-->criar variavel vazia par poder substituir por charset no arr de str

*/
 
