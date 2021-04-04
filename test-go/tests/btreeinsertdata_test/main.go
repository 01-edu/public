package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
	"github.com/01-edu/public/test-go/tests/correct"
)

func parentListInsert(root *student.TreeNode) string {
	if root == nil {
		return ""
	}

	var parent string

	if root.Parent == nil {
		parent = "nil"
	} else {
		parent = root.Parent.Data
	}

	r := "Node: " + root.Data + " Parent: " + parent + "\n"
	r += parentListInsert(root.Left) + parentListInsert(root.Right)
	return r
}

func FormatTree_insert(root *student.TreeNode) string {
	if root == nil {
		return ""
	}
	res := root.Data + "\n"
	res += formatSubTree_insert(root, "")
	return res
}

func formatSubTree_insert(root *student.TreeNode, prefix string) string {
	if root == nil {
		return ""
	}

	var res string

	hasLeft := root.Left != nil
	hasRight := root.Right != nil

	if !hasLeft && !hasRight {
		return res
	}

	res += prefix
	if hasLeft && hasRight {
		res += "├── "
	}

	if !hasLeft && hasRight {
		res += "└── "
	}

	if hasRight {
		printStrand := (hasLeft && hasRight && (root.Right.Right != nil || root.Right.Left != nil))
		newPrefix := prefix
		if printStrand {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		res += root.Right.Data + "\n"
		res += formatSubTree_insert(root.Right, newPrefix)
	}

	if hasLeft {
		if hasRight {
			res += prefix
		}
		res += "└── " + root.Left.Data + "\n"
		res += formatSubTree_insert(root.Left, prefix+"    ")
	}
	return res
}

func errorMessage_insert(fn interface{}, inserted string, root *correct.TreeNode, rootS *student.TreeNode) {
	lib.Fatalf("%s(\n%s, %s\n) ==\n%s instead of\n%s\n",
		"BTreeInsertData",
		correct.FormatTree(root),
		inserted,
		FormatTree_insert(rootS),
		correct.FormatTree(root),
	)
}

func CompareTrees_insert(fn interface{}, inserted string, root *correct.TreeNode, rootS *student.TreeNode) {
	solTree := correct.FormatTree(root)
	stuTree := FormatTree_insert(rootS)

	if solTree != stuTree {
		errorMessage_insert(fn, inserted, root, rootS)
	}
}

func main() {
	root := &correct.TreeNode{Data: "08"}
	rootS := &student.TreeNode{Data: "08"}

	var pos []string

	pos = append(pos,
		"x",
		"z",
		"y",
		"t",
		"r",
		"q",
		"01",
		"b",
		"c",
		"a",
		"d",
	)
	fn := interface{}(correct.BTreeInsertData)
	for _, arg := range pos {
		root = correct.BTreeInsertData(root, arg)
		rootS = student.BTreeInsertData(rootS, arg)
		CompareTrees_insert(fn, arg, root, rootS)
	}
}
