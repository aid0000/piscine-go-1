package piscine

func ListClear(l *List) {
	a := l.Head
	b := l.Head
	for a != nil {
		b = a.Next
		a = b
	}
	l.Head = nil
}