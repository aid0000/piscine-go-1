package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		a := FindNextPrime(nb+1)
		return a
	}
}