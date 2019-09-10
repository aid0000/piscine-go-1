package piscine
func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
if root !=nil{
	BTreeApplyPostorder(root.Left, f)
	BTreeApplyPostorder(root.Right, f)
	f(root.Data)
	
}
}


/*Post-order (LRN)


Check if the current node is empty or null.
Traverse the left subtree by recursively calling the post-order function.
Traverse the right subtree by recursively calling the post-order function.
Display the data part of the root (or current node).*/