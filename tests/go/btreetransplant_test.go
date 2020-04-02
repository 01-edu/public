package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
	"fmt"
)

func parentListTransp(root *student.TreeNode) string {
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
	r += parentListTransp(root.Left) + parentListTransp(root.Right)
	return r
}

func FormatTree_transp(root *student.TreeNode) string {
	if root == nil {
		return ""
	}
	res := root.Data + "\n"
	res += formatSubTree_transp(root, "")
	return res
}

func formatSubTree_transp(root *student.TreeNode, prefix string) string {
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
		res += formatSubTree_transp(root.Right, newPrefix)
	}

	if hasLeft {
		if hasRight {
			res += prefix
		}
		res += "└── " + root.Left.Data + "\n"
		res += formatSubTree_transp(root.Left, prefix+"    ")
	}
	return res
}

func errorMessage_transp(t *testing.T, fn interface{}, root, sel, repl *solutions.TreeNode, rootA *solutions.TreeNode, rootAS *student.TreeNode) {
	t.Errorf("%s(\nRoot:\n %s, Selected:\n%s, Replacement:\n%s\n) ==\n%s instead of\n%s\n",
		z01.NameOfFunc(fn),
		solutions.FormatTree(root),
		solutions.FormatTree(sel),
		solutions.FormatTree(repl),
		FormatTree_transp(rootAS),
		solutions.FormatTree(rootA),
	)
}

func CompareTrees_transp(t *testing.T, fn interface{}, root, sel, repl *solutions.TreeNode, rootA *solutions.TreeNode, rootAS *student.TreeNode) {
	solTree := solutions.FormatTree(rootA)
	stuTree := FormatTree_transp(rootAS)

	if solTree != stuTree {
		errorMessage_transp(t, fn, root, sel, repl, rootA, rootAS)
	}
	parentSol := solutions.ParentList(rootA)
	parentStu := parentListTransp(rootAS)

	if parentSol != parentStu {
		fmt.Println("Tree:\n", solTree)
		t.Errorf("Expected\n%s instead of\n%s\n", parentSol, parentStu)
	}
}

func TestBTreeTransplant(t *testing.T) {
	root := &solutions.TreeNode{Data: "04"}
	rootS := &student.TreeNode{Data: "04"}
	rootOr := &solutions.TreeNode{Data: "04"}

	replacement := &solutions.TreeNode{Data: "55"}
	replacementS := &student.TreeNode{Data: "55"}

	ins := []string{"01", "07", "05", "12", "02", "03", "10"}

	rep := []string{"33", "12", "15", "60"}

	for _, v := range ins {
		root = solutions.BTreeInsertData(root, v)
		rootS = student.BTreeInsertData(rootS, v)
		rootOr = solutions.BTreeInsertData(rootOr, v)
	}

	for _, v := range rep {
		replacement = solutions.BTreeInsertData(replacement, v)
		replacementS = student.BTreeInsertData(replacementS, v)
	}

	selected := solutions.BTreeSearchItem(root, "07")
	selectedS := student.BTreeSearchItem(rootS, "07")
	root = solutions.BTreeTransplant(root, selected, replacement)
	rootS = student.BTreeTransplant(rootS, selectedS, replacementS)
	fn := interface{}(solutions.BTreeTransplant)

	CompareTrees_transp(t, fn, rootOr, selected, replacement, root, rootS)
}
