package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {

	replace := node
	
	if root == nil {
		return nil
	}

	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		replace.Parent.Left = rplc
	} else {
		replace.Parent.Right = rplc
	}
	node.Parent=replace.Parent 

	return root
}
