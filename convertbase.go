package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return NbrBase(AtoiBase(nbr, baseFrom), baseTo)
}

func NbrBase(nbr int, str string) string {
	indx := 0
	for _, res := range str {
		if string(res) == "-" || string(res) == "+" {
			indx = 1
			break
		}
	}
	if indx == 1 || len(str) < 2 {
		return "0"
	} else if 2147483647 < nbr || -2147483648 > nbr {
		return string(nbr)
	} else {
		
		
		i := 0
		smt := ""
		for nbr >= lent2(str) {
			if nbr >= lent2(str) {
				smt += string(str[nbr%lent2(str)])
				nbr = nbr / lent2(str)
				i++
			}
		}
		smt += string(str[nbr])
		return Reverse(smt)
	}
}
