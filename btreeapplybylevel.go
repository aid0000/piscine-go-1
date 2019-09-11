package piscine

import (
	"github.com/01-edu/z01"
	
	
)

func BTreeApplyByLevel(root *TreeNode, fn interface{})  {
	novo:=BTreeLevelCount(root)
	for d:=1; d<=novo;d++{
	
		PrintNodesLevel(root, d, fn)
	}
}
 func PrintNodesLevel(root *TreeNode, level int, fn interface{} ){
if root!=nil{
	if level==1{
		ar:=[]interface{}{root.Data}
		z01.Call(fn, ar)
		
		
	}else if level > 1{
		PrintNodesLevel(root.Left, level-1, fn)
		PrintNodesLevel(root.Right, level-1, fn)
	   }
	}
 }




 /*
func BTreeApplyByLevel(root *TreeNode, fn interface{})  {
	
	if root != nil {
		BTreeApplyByLevel(root.Left, fn)         //f is the in orderfunc
		ar:=[]interface{}{root.Data}
		z01.Call(fn, ar)
		BTreeApplyByLevel(root.Right, fn)
		
	}
}*/