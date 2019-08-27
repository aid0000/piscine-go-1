package piscine

func Index(s string, toFind string) int {
	

j:=[]rune(s)
l:=[]rune(toFind)
n:=lent(j)
k:=lent(l)

		for i:=0; i<= n-k; i++ {
				if toFind==s[i:i+k]{
					return (i)
				}
		}
		return -1
}











