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
func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil || root.Right == nil{
		return root
	}
return BTreeMax(root.Right)

}
func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	max :=BTreeMax(root)
	fmt.Println(max.Data)
}
