package piscine
func MakeRange(min, max int) []int {

	if min>=max{
		return nil
	}
	resultado:= make([]int, max- min)
	for i := range resultado{
		resultado[i] = min + i
	}
	return resultado
}
