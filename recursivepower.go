package piscine

func RecursivePower(nb,power int) int {
if power<0{
	return 0
}
	if power== 0 { 
        return 1 
    }else if power % 2 == 0 {
        return RecursivePower(nb, power / 2) * RecursivePower(nb, power / 2)
    }else{
        return nb * RecursivePower(nb, power / 2) * RecursivePower(nb, power / 2) 
} 
	}