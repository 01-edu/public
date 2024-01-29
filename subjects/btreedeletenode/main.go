package main

import (
	"fmt"
	"piscine"
)

func main() {
	root := &piscine.TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	node := BTreeSearchItem(root, "4")
	fmt.Println("Before delete:")
	BTreeApplyInorder(root, fmt.Println)
	root = piscine.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	BTreeApplyInorder(root, fmt.Println)
}

func BTreeInsertData(root *piscine.TreeNode, data string) *piscine.TreeNode {
	if root == nil {
		return &piscine.TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}

	return root
}

func BTreeSearchItem(root *piscine.TreeNode, elem string) *piscine.TreeNode {
	if root == nil {
		return nil
	}

	if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	} else if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	} else if elem == root.Data {
		return root
	} else {
		return root
	}
}

func BTreeApplyInorder(root *piscine.TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}
