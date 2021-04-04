package main

import (
	"fmt"

	student "student"

	"github.com/01-edu/public/test-go/lib"
	"github.com/01-edu/public/test-go/solutions"
)

func errorMessage_search(fn interface{}, root, a *solutions.TreeNode, b *student.TreeNode,
	seaVal string) {
	lib.Fatalf("%s(\n%s\n, %s) == %s instead of %s\n",
		"BTreeSearchItem",
		solutions.FormatTree(root),
		seaVal,
		b.Data,
		a.Data,
	)
}

func CompareNode_search(fn interface{}, arg1, a *solutions.TreeNode, b *student.TreeNode,
	seaVal string) {
	if a == nil && b == nil {
		return
	}

	if (a == nil && b != nil) || (b == nil && a != nil) {
		lib.Fatalf("Expected %v instead of %v\n", a, b)
		return
	}

	if a.Data != b.Data {
		errorMessage_search(fn, arg1, a, b, seaVal)
	}

	if a.Parent != nil && b.Parent != nil {
		if a.Parent.Data != b.Parent.Data {
			errorMessage_search(fn, arg1, a, b, seaVal)
			lib.Fatalf("Expected parent value %v instead of %v\n", a.Parent.Data, b.Parent.Data)
			fmt.Println("Parent.Data", a.Parent.Data, b.Parent.Data)
		}
	} else if (a.Parent == nil && b.Parent != nil) || (a.Parent != nil && b.Parent == nil) {
		lib.Fatalf("Expected parent value %v instead of %v\n", a.Parent, b.Parent)
	}

	if a.Right != nil && b.Right != nil {
		if a.Right.Data != b.Right.Data {
			errorMessage_search(fn, arg1, a, b, seaVal)
			lib.Fatalf("Expected right child value %v instead of %v\n", a.Right.Data, b.Right.Data)
			fmt.Println("Right.Data", a.Right.Data, b.Right.Data)
		}
	} else if (a.Right == nil && b.Right != nil) || (a.Right != nil && b.Right == nil) {
		lib.Fatalf("Expected right child value %v instead of %v\n", a.Right, b.Right)
	}

	if a.Left != nil && b.Left != nil {
		if a.Left.Data != b.Left.Data {
			errorMessage_search(fn, arg1, a, b, seaVal)
			lib.Fatalf("Expected left child value %v instead of %v\n", a.Left, b.Left)
			fmt.Println("Left.Data", a.Left.Data, b.Left.Data)
		}
	} else if (a.Left == nil && b.Left != nil) || (a.Left != nil && b.Left == nil) {
		lib.Fatalf("Expected left child value %v instead of %v\n", a, b)
	}
}

func main() {
	root := &solutions.TreeNode{Data: "04"}
	rootS := &student.TreeNode{Data: "04"}

	ins := []string{"01", "07", "05", "12", "02", "03", "10"}

	for _, v := range ins {
		root = solutions.BTreeInsertData(root, v)
		rootS = student.BTreeInsertData(rootS, v)
	}
	fn := interface{}(solutions.BTreeSearchItem)

	ins = append(ins, "322")
	for _, v := range ins {
		selectedSol := solutions.BTreeSearchItem(root, v)
		selectedStu := student.BTreeSearchItem(rootS, v)
		CompareNode_search(fn, root, selectedSol, selectedStu, v)
	}
}
