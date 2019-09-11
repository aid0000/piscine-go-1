package piscine

import (
	"github.com/01-edu/z01"
)

func BTreeApplyByLevel(root *TreeNode, fn interface{})  {
	
	if root != nil {
		ar:=[]interface{}{root.Data}
		z01.Call(fn, ar)
		BTreeApplyByLevel(root.Left, fn)         //f is the in orderfunc
		BTreeApplyByLevel(root.Right, fn)
		
	}
}