package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    a := node

    if node.Parent == nil {
        root = rplc
    } else if node == node.Parent.Left {
        a.Parent.Left = rplc
    } else {
        a.Parent.Right = rplc
    }
    a.Parent = node.Parent
    
    return root
}