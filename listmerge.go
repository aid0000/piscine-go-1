package piscine

func ListMerge(l1 *List, l2 *List) {
	if l2 == nil || l1 == nil {
		return
	}

	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Head
		return
	}

	l1.Tail.Next = l2.Head
}
