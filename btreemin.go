package piscine
func BTreeMin(root *TreeNode) *TreeNode {
	
	if root == nil || root.Left == nil {
		return root
	}

	return BTreeMax(root.Left)
}
