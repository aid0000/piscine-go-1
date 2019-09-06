package piscine

func ListRemoveIf(l *List, data_ref interface{}) {

	temp := l.Head
	prev := l.Head
	 								// If head node itself holds the key to be deleted 
	if temp != nil && temp.Data == data_ref {
		l.Head = temp.Next  		// Changed head 
		temp = l.Head
	}
    								// Search for the key to be deleted, keep track of the 
    								// previous node as we need to change 'prev.next' 
	for temp != nil {
		for temp != nil && temp.Data != data_ref {
			prev = temp
			temp = temp.Next
		}
		if temp == nil {			// If key was not present in linked list 
			return
		}
		prev.Next = temp.Next		// Unlink the node from linked list
		temp = prev.Next
	}
}
