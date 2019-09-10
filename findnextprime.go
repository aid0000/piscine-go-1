package piscine

func FindNextPrime(nb int) int {
	if nb>2147483647{
		return 0
	}
	for !IsPrime(nb) {
		nb++
		} 
		return nb
	
}

/*int8   : -128 to 127 
int16  : -32768 to 32767 
						
int32  : -2147483648 to 2147483647 
int64  : -9223372036854775808 to 9223372036854775807*/