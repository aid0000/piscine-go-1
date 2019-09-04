package piscine



func Map(f func(int) bool, arr []int) []bool {
	r:=[]bool{}
	for _, s := range arr {
		
		r=append(r,f(s))
	}
	
	return r
}