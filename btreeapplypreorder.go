package piscine 
func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {

if root!=nil{
	f(root.Data)
	BTreeApplyPreorder(root.Left, f)
	BTreeApplyPreorder(root.Right, f)
}

}
/*Pre-order (NLR)


Check if the current node is empty or null.
Display the data part of the root (or current node).
Traverse the left subtree by recursively calling the pre-order function.
Traverse the right subtree by recursively calling the pre-order function.*/