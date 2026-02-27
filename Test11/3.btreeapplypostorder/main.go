package main

import (
	"fmt"
	
)
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                 string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
if root == nil {
	return &TreeNode{Data:data}
}
if data < root.Data{
	leftchild := BTreeInsertData(root.Left, data)
	root.Left = leftchild
	leftchild.Parent = root
}else if data > root.Data{
	rightchild := BTreeInsertData(root.Right,data)
	root.Right = rightchild
	rightchild.Parent = root
}
return root
}
func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
if root == nil {
	return 
}
BTreeApplyPostorder(root.Left,f)
BTreeApplyPostorder(root.Right,f)
f(root.Data)
}
func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	BTreeApplyPostorder(root, fmt.Println)

}
