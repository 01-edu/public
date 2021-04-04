package main

import (
	"fmt"

	student "student"

	"github.com/01-edu/public/test-go/lib"
	"github.com/01-edu/public/test-go/tests/correct"
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

func errorMessage_transp(fn interface{}, root, sel, repl *correct.TreeNode, rootA *correct.TreeNode, rootAS *student.TreeNode) {
	lib.Fatalf("%s(\nRoot:\n %s, Selected:\n%s, Replacement:\n%s\n) ==\n%s instead of\n%s\n",
		"BTreeTransplant",
		correct.FormatTree(root),
		correct.FormatTree(sel),
		correct.FormatTree(repl),
		FormatTree_transp(rootAS),
		correct.FormatTree(rootA),
	)
}

func CompareTrees_transp(fn interface{}, root, sel, repl *correct.TreeNode, rootA *correct.TreeNode, rootAS *student.TreeNode) {
	solTree := correct.FormatTree(rootA)
	stuTree := FormatTree_transp(rootAS)

	if solTree != stuTree {
		errorMessage_transp(fn, root, sel, repl, rootA, rootAS)
	}
	parentSol := correct.ParentList(rootA)
	parentStu := parentListTransp(rootAS)

	if parentSol != parentStu {
		fmt.Println("Tree:\n", solTree)
		lib.Fatalf("Expected\n%s instead of\n%s\n", parentSol, parentStu)
	}
}

func main() {
	root := &correct.TreeNode{Data: "04"}
	rootS := &student.TreeNode{Data: "04"}
	rootOr := &correct.TreeNode{Data: "04"}

	replacement := &correct.TreeNode{Data: "55"}
	replacementS := &student.TreeNode{Data: "55"}

	ins := []string{"01", "07", "05", "12", "02", "03", "10"}

	rep := []string{"33", "12", "15", "60"}

	for _, v := range ins {
		root = correct.BTreeInsertData(root, v)
		rootS = student.BTreeInsertData(rootS, v)
		rootOr = correct.BTreeInsertData(rootOr, v)
	}

	for _, v := range rep {
		replacement = correct.BTreeInsertData(replacement, v)
		replacementS = student.BTreeInsertData(replacementS, v)
	}

	selected := correct.BTreeSearchItem(root, "07")
	selectedS := student.BTreeSearchItem(rootS, "07")
	root = correct.BTreeTransplant(root, selected, replacement)
	rootS = student.BTreeTransplant(rootS, selectedS, replacementS)
	fn := interface{}(correct.BTreeTransplant)

	CompareTrees_transp(fn, rootOr, selected, replacement, root, rootS)
}
