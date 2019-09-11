

 
package piscine


func FindNextPrime(nb int)int{
	nextprime:=nb-1
	i:=nb+1
	for i>nb{
		nextprime++
		if IsPrime(nextprime){
			return nextprime
		}
		i++	
	}
	return nextprime
}



