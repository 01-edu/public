package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
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

func errorMessage_insert(t *testing.T, fn interface{}, inserted string, root *solutions.TreeNode, rootS *student.TreeNode) {
	t.Errorf("%s(\n%s, %s\n) ==\n%s instead of\n%s\n",
		z01.NameOfFunc(fn),
		solutions.FormatTree(root),
		inserted,
		FormatTree_insert(rootS),
		solutions.FormatTree(root),
	)
}

func CompareTrees_insert(t *testing.T, fn interface{}, inserted string, root *solutions.TreeNode, rootS *student.TreeNode) {
	solTree := solutions.FormatTree(root)
	stuTree := FormatTree_insert(rootS)

	if solTree != stuTree {
		errorMessage_insert(t, fn, inserted, root, rootS)
	}
}

func TestBTreeInsertData(t *testing.T) {
	root := &solutions.TreeNode{Data: "08"}
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
	fn := interface{}(solutions.BTreeInsertData)
	for _, arg := range pos {
		root = solutions.BTreeInsertData(root, arg)
		rootS = student.BTreeInsertData(rootS, arg)
		CompareTrees_insert(t, fn, arg, root, rootS)
	}

}
