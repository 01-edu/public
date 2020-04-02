package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func parentListDelete(root *student.TreeNode) string {
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
	r += parentListDelete(root.Left) + parentListDelete(root.Right)
	return r
}

func FormatTree_delete(root *student.TreeNode) string {
	if root == nil {
		return ""
	}
	res := root.Data + "\n"
	res += formatSubTree_delete(root, "")
	return res
}

func formatSubTree_delete(root *student.TreeNode, prefix string) string {
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
		res += formatSubTree_delete(root.Right, newPrefix)
	}

	if hasLeft {
		if hasRight {
			res += prefix
		}
		res += "└── " + root.Left.Data + "\n"
		res += formatSubTree_delete(root.Left, prefix+"    ")
	}
	return res
}

func errorMessage_delete(t *testing.T, fn interface{}, deleted string, rootOr, root *solutions.TreeNode, rootS *student.TreeNode) {
	t.Errorf("%s(\n%s, %s\n) ==\n%s instead of\n%s\n",
		z01.NameOfFunc(fn),
		solutions.FormatTree(rootOr),
		deleted,
		FormatTree_delete(rootS),
		solutions.FormatTree(root),
	)
}

func CompareTrees_delete(t *testing.T, fn interface{}, deleted string, rootOr, root *solutions.TreeNode, rootS *student.TreeNode) {
	sel := student.BTreeSearchItem(rootS, deleted)

	if !student.BTreeIsBinary(rootS) || sel != nil {
		errorMessage_delete(t, fn, deleted, rootOr, root, rootS)
	}

}

func TestBTreeDeleteNode(t *testing.T) {
	root := &solutions.TreeNode{Data: "04"}
	rootS := &student.TreeNode{Data: "04"}
	rootOr := &solutions.TreeNode{Data: "04"}

	ins := []string{"01", "07", "05", "12", "02", "03", "10"}

	for _, v := range ins {
		root = solutions.BTreeInsertData(root, v)
		rootS = student.BTreeInsertData(rootS, v)
		rootOr = solutions.BTreeInsertData(rootOr, v)
	}

	selected := solutions.BTreeSearchItem(root, "04")
	selectedS := student.BTreeSearchItem(rootS, "04")

	root = solutions.BTreeDeleteNode(root, selected)
	rootS = student.BTreeDeleteNode(rootS, selectedS)
	fn := interface{}(solutions.BTreeDeleteNode)
	CompareTrees_delete(t, fn, selected.Data, rootOr, root, rootS)
}
