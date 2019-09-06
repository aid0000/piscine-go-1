package piscine
func ListAt(l *NodeL, pos int) *NodeL{
	
	iterator:=l
		inc:=0
	for iterator!= nil{
		if pos==inc{
		return iterator	
		}
		inc ++
		iterator= iterator.Next
	}
	return nil
}
