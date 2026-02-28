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
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
if root == nil {
	return 
}
BTreeApplyInorder(root.Left,f)
f(root.Data)
BTreeApplyInorder(root.Right,f)
}
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
if root == nil {
	return nil
}
if elem == root.Data{
return root
}else if elem < root.Data{
	return BTreeInsertData(root.Left, elem)
}else{
	return BTreeInsertData(root.Right, elem)
}

}
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
if root == nil {
		return nil
	}
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}
	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root
}
func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	node := BTreeSearchItem(root, "1")
	rplc := &TreeNode{Data: "3"}
	root = BTreeTransplant(root, node, rplc)
	BTreeApplyInorder(root, fmt.Println)
}
