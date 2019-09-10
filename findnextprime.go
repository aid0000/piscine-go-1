package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		return FindNextPrime(nb+1) 
	}
}

