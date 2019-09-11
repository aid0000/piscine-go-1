package piscine

import (
	"math"
)

func FindNextPrime(nb int) int {
ns:=0
	for i:=nb;i<2*nb; i++{

		if IsPrime2(i) {
			ns=i
			break
		}
			}
		
	return ns
}




func IsPrime2(value int) bool {
	if value < 2 {
		return false
	}
	limit := int(math.Floor(math.Sqrt(float64(value))))
	i := 2
	for i <= limit {
		if value%i == 0 {
			return false
		}
		i++
	}
	return true
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
