package piscine
func Join(strs []string, sep string) string {
	strV := ""

	for i := 0; i < Lent3(strs); i++ {
		strV += strs[i]                     
		if i < Lent3(strs)-1 {
			strV += sep
		}
	}
	return strV
}


