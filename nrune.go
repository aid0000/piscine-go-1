package piscine

func NRune(s string, n int) rune {

			ar:=[]rune(s)
			
			for i:=0;i<=n;i++{
				if i==n{
					return (ar[i-1])
				}
			}
		return 0

}