package piscine
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {

	if root ==nil || root.Data == elem{
		return root
	}
	if root.Data <elem{
		return BTreeSearchItem(root.Right, elem)

	}

return BTreeSearchItem(root.Left, elem)

}
