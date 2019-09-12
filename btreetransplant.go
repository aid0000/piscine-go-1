package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    a := node

    if node.Parent == nil {
        root = rplc
    } else if node == node.Parent.Right {
        a.Parent.Right = rplc
    } else {
        a.Parent.Left = rplc
    }
    a.Parent = node.Parent
    
    return root
}