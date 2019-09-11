package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		n := nb + 1
		for IsPrime(n) == false {
			n++
		}
		return n
	}
}
/*
 
package piscine

import "math"

func FindNextPrime(nb int)int{
	nextprime:=nb-1
	i:=nb+1
	for i>nb{
		nextprime++
		if premier(nextprime){
			return nextprime
		}
		i++	
	}
	return nextprime
}


func premier(nb int) bool{
	decision:=true
	if nb<=1{
		return false
	}
	for i:=2;i<int(math.Round(math.Sqrt(float64(nb))))+1;i++{
		if nb%i==0{
			decision=false
		}
	}
	return decision
}*/
