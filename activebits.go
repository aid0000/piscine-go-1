package piscine

func ActiveBits(n int) uint {
count:=0
	if n<0{
		n=-n
	}
	resto:=0
	divisor:=0
	DivMod(n, 2, &divisor, &resto)
	for divisor>0{	
		DivMod(n, 2, &divisor, &resto)
		n=divisor
		if resto==1{
			count++
		}
	}
	return uint(count)
}