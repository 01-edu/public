package main

import (
	"reflect"

	solutions "./solutions"
	student "./student"

	"github.com/01-edu/z01"
)

func errorMessage_max(fn interface{}, root, a *solutions.TreeNode, b *student.TreeNode) {
	z01.Fatalf("%s(\n%s) == %s instead of %s\n",
		"BTreeMax",
		solutions.FormatTree(root),
		b.Data,
		a.Data,
	)
}

func CompareNode_max(fn interface{}, arg1, a *solutions.TreeNode, b *student.TreeNode) {
	if a == nil || b == nil {
		z01.Fatalf("Expected %v instead of %v\n", a, b)
		return
	}

	if a.Data != b.Data {
		errorMessage_max(fn, arg1, a, b)
	}

	if a.Parent != nil && b.Parent != nil {
		if a.Parent.Data != b.Parent.Data {
			errorMessage_max(fn, arg1, a, b)
			z01.Fatalf("Expected parent value %v instead of %v\n", a.Parent.Data, b.Parent.Data)
		}
	} else if (a.Parent == nil && b.Parent != nil) || (a.Parent != nil && b.Parent == nil) {
		z01.Fatalf("Expected parent value %v instead of %v\n", a.Parent, b.Parent)
	}

	if a.Right != nil && b.Right != nil {
		if a.Right.Data != b.Right.Data {
			errorMessage_max(fn, arg1, a, b)
			z01.Fatalf("Expected right child value %v instead of %v\n", a.Right.Data, b.Right.Data)
		}
	} else if (a.Right == nil && b.Right != nil) || (a.Right != nil && b.Right == nil) {
		z01.Fatalf("Expected right child value %v instead of %v\n", a.Right, b.Right)
	}

	if a.Left != nil && b.Left != nil {
		if a.Left.Data != b.Left.Data {
			errorMessage_max(fn, arg1, a, b)
			z01.Fatalf("Expected left child value %v instead of %v\n", a.Left, b.Left)
		}
	} else if (a.Left == nil && b.Left != nil) || (a.Left != nil && b.Left == nil) {
		z01.Fatalf("Expected left child value %v instead of %v\n", a, b)
	}
}

func CompareReturn_max(fn1, fn2 interface{}, arg1 *solutions.TreeNode, arg2 interface{}) {
	arar1 := []interface{}{arg1}
	arar2 := []interface{}{arg2}

	out1 := z01.Monitor(fn1, arar1)
	out2 := z01.Monitor(fn2, arar2)

	for i, v := range out1.Results {
		switch str := v.(type) {
		case *solutions.TreeNode:
			CompareNode_max(fn1, arg1, str, out2.Results[i].(*student.TreeNode))
		default:
			if !reflect.DeepEqual(str, out2.Results[i]) {
				z01.Fatalf("%s(\n%s) == %s instead of\n %s\n",
					"BTreeMax",
					solutions.FormatTree(arg1),
					z01.Format(out2.Results...),
					z01.Format(out1.Results...),
				)
			}
		}
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

	CompareReturn_max(solutions.BTreeMax, student.BTreeMax, root, rootS)
}
