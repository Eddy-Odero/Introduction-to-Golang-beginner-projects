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
func BTreeMin(root *TreeNode) *TreeNode {
if root == nil || root.Left == nil {
	return root
}
return BTreeMin(root.Left)
}
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		

		
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left == nil {
			root.Right.Parent = root.Parent
			return root.Right
		}
		if root.Right == nil {
			root.Left.Parent = root.Parent
			return root.Left
		}

		successor := BTreeMin(root.Right)
		root.Data = successor.Data
		root.Right = BTreeDeleteNode(root.Right, successor)
	}

	return root
}
func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	node := BTreeSearchItem(root, "4")
	fmt.Println("Before delete:")
	BTreeApplyInorder(root, fmt.Println)
	root = BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	BTreeApplyInorder(root, fmt.Println)
}
