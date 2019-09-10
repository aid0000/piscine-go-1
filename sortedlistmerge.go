package piscine

func SortedListMerge(l1 *NodeI, l2 *NodeI) *NodeI {

	l1 = ListSort(l1)
	l2 = ListSort(l2)

	if l1 == nil {    
		return l2
	}								//se tiverem vazias retorna a que tem algo
	if l2 == nil {
		return l1
	}
	if l1.Data <= l2.Data {
		l1.Next = SortedListMerge(l1.Next, l2)		//Escolha l1 maior igual a l2 e retorna l1 se sim cao contrario retorna o l2
		return l1
	} else {
		l2.Next = SortedListMerge(l1, l2.Next)
		return l2
	}
}