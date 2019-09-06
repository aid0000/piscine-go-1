
package piscine

func Sqrt(nb int) int {
	a := 1

	for nb != IterativePower(a, 2) {
		a++
		if a > nb {
			return 0
		}
	}
	if nb == IterativePower(a, 2) {
		return a
	}
	return 0
}