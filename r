package piscine

func IterativePower(nb,power int) int {

	
    if power== 0 { 
        return 1 
    }else if power % 2 == 0 {
        return IterativePower(nb, power / 2) * IterativePower(nb, power / 2)
    }else{
        return nb * IterativePower(nb, power / 2) * IterativePower(nb, power / 2) 
} 
	}

