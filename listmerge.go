package piscine

func ListMerge(l1 *List, l2 *List) {

	novohead := l1.Head
	novohead2 := l2.Head
	for novohead != nil {			//percorre a list ate ao fim quando for nill

		if novohead.Next == nil {	//se o proximo passo for nill	
			novohead.Next = novohead2	//esse passo torna-se no primeiro elemento da lista 2
			return
		}
		novohead = novohead.Next
	}
}
