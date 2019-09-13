package piscine

func Rot14(str string) string {
	arrayRune := []rune(str)
	var result string

	for i := 0; i < lent(arrayRune); i++ {				 	//quando é treze e +13 de A a M
															// e -13 de n a z
		if arrayRune[i] >= 'a' && arrayRune[i] <= 'z' { 	//como é base 14 subtrai 12 quando é
															//maior que m e soma 14 quando menor

			if arrayRune[i] >= 'm' {
				arrayRune[i] -= 12
			} else {
				arrayRune[i] += 14
			}
		} else if arrayRune[i] >= 'A' && arrayRune[i] <= 'Z' {
			if arrayRune[i] >= 'M' {
				arrayRune[i] -= 12
			} else {
				arrayRune[i] += 14
			}
		}
		result += string(arrayRune[i])
	}
	return result
}
