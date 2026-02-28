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
func BTreeIsBinary(root *TreeNode) bool {
if root == nil {
	return true
}
if root.Left != nil && root.Right != nil{
	if root.Left.Data > root.Data || root.Right.Data < root.Data{
		return false
	}
	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}
return true
}
func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	fmt.Println(BTreeIsBinary(root))
}
