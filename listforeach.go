package piscine
func ListForEach(l *List, f func(*NodeL)) {

	
	novohead:=l.Head
	
	for novohead != nil {
		f(novohead)
		novohead=novohead.Next
	}	
}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
