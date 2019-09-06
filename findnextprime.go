package piscine

func FindNextPrime(nb int) int {
	if isPrime2(nb) {
		return nb
	} else {
		return FindNextPrime(nb+1) 
	}
}
func isPrime2(value int) bool {

    if value<=1{
        return false
    }
      for i:=2; i<value; i++{
         if value % i == 0{
          return false
        }
    }
    return true
}



func IsPrime2(value int) bool {
	if value < 2 {
		return false
	}
	b:=string(value)
	a:=[]rune(b)
	for p,_:=range a{

		i := 2
		for i <= p {
			if value%i == 0 {
				return false
			}
			i++
		}
	}
	return true
}