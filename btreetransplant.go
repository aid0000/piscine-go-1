package piscine

func BTreeTransplant(root, node, repla *TreeNode) *TreeNode {

	replace := node
	
	if root == nil {
		return nil
	}

	if root.Parent == nil {
		root = repla
	} else if node == node.Parent.Left {
		replace.Parent.Left = repla
	} else {
		replace.Parent.Right = repla
	}
	replace.Parent = node.Parent

	return root
}
