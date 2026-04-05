package maain

import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		left := BTreeInsertData(root.Left, data)
		root.Left = left
		left.Parent = root
	} else if data > root.Data {
		right := BTreeInsertData(root.Right, data)
		root.Right = right
		right.Parent = root
	}
	return root
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem == root.Data {
		return root
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
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
		if root.Left != nil {
			root.Left.Parent = root
		}
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
		if root.Right != nil {
			root.Right.Parent = root
		}
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left == nil {
			if root.Right != nil {
				root.Right.Parent = root.Parent
			}
			return root.Right
		}
		if root.Right == nil {
			if root.Left != nil {
				root.Left.Parent = root.Parent
			}
			return root.Left
		}
		successor := BTreeMin(root.Right)
		root.Data = successor.Data
		root.Right = BTreeDeleteNode(root.Right, successor)
		if root.Right != nil {
			root.Right.Parent = root
		}
	}
	return root
}

func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, "", "")
}

func isBST(node *TreeNode, min, max string) bool {
	if node == nil {
		return true
	}
	if (min != "" && node.Data <= min) || (max != "" && node.Data >= max) {
		return false
	}
	return isBST(node.Left, min, node.Data) && isBST(node.Right, node.Data, max)
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
