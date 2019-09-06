package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {

	novohead := l.Head
	//o caminho ta guardado dentro da Data,
	//tenho de igualar a data ao ref com a função compr o a=b logo o a->Data e o b->ref
	//indicar o caminho para a data (l.Head.Data) l.head= novo head
	//uso o simbolo & para buscar o caminho
	for novohead != nil {
		if comp(novohead.Data, ref) {

			return &novohead.Data
		}
		//usado para imprimir o que ta na posicao
		novohead = novohead.Next
	}

	return nil
}
