package piscine
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {

	if root != nil {
		BTreeApplyInorder(root.Left, f)         //f is the in orderfunc
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
}

/*In-order (LNR)


Check if the current node is empty or nil.
Traverse the left subtree by recursively calling the in-order function.
Display the data part of the root (or current node).
Traverse the right subtree by recursively calling the in-order function.*/