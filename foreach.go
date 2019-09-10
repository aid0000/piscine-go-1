package piscine

func ForEach(f func(int), arr []int) {
	for _, s := range arr {
		f(s)
	}
}
