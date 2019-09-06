package piscine

func IterativeFactorial(nb int ) int {

	if nb < 0 || nb > 12 {
		return 0
	} else if nb == 0 {
		return 1
	}
	a := nb - 1
	for i := a; i >= 1; i-- {
		nb *= i
	}
	return nb
}