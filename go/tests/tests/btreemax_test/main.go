package main

import (
	"reflect"

	student "student"

	"github.com/01-edu/public/go/tests/func/correct"

	"github.com/01-edu/public/go/tests/lib"
)

func errorMessage_max(fn interface{}, root, a *correct.TreeNode, b *student.TreeNode) {
	lib.Fatalf("%s(\n%s) == %s instead of %s\n",
		"BTreeMax",
		correct.FormatTree(root),
		b.Data,
		a.Data,
	)
}

func CompareNode_max(fn interface{}, arg1, a *correct.TreeNode, b *student.TreeNode) {
	if a == nil || b == nil {
		lib.Fatalf("Expected %v instead of %v\n", a, b)
		return
	}

	if a.Data != b.Data {
		errorMessage_max(fn, arg1, a, b)
	}

	if a.Parent != nil && b.Parent != nil {
		if a.Parent.Data != b.Parent.Data {
			errorMessage_max(fn, arg1, a, b)
			lib.Fatalf("Expected parent value %v instead of %v\n", a.Parent.Data, b.Parent.Data)
		}
	} else if (a.Parent == nil && b.Parent != nil) || (a.Parent != nil && b.Parent == nil) {
		lib.Fatalf("Expected parent value %v instead of %v\n", a.Parent, b.Parent)
	}

	if a.Right != nil && b.Right != nil {
		if a.Right.Data != b.Right.Data {
			errorMessage_max(fn, arg1, a, b)
			lib.Fatalf("Expected right child value %v instead of %v\n", a.Right.Data, b.Right.Data)
		}
	} else if (a.Right == nil && b.Right != nil) || (a.Right != nil && b.Right == nil) {
		lib.Fatalf("Expected right child value %v instead of %v\n", a.Right, b.Right)
	}

	if a.Left != nil && b.Left != nil {
		if a.Left.Data != b.Left.Data {
			errorMessage_max(fn, arg1, a, b)
			lib.Fatalf("Expected left child value %v instead of %v\n", a.Left, b.Left)
		}
	} else if (a.Left == nil && b.Left != nil) || (a.Left != nil && b.Left == nil) {
		lib.Fatalf("Expected left child value %v instead of %v\n", a, b)
	}
}

func CompareReturn_max(fn1, fn2 interface{}, arg1 *correct.TreeNode, arg2 interface{}) {
	arar1 := []interface{}{arg1}
	arar2 := []interface{}{arg2}

	out1 := lib.Monitor(fn1, arar1)
	out2 := lib.Monitor(fn2, arar2)

	for i, v := range out1.Results {
		switch str := v.(type) {
		case *correct.TreeNode:
			CompareNode_max(fn1, arg1, str, out2.Results[i].(*student.TreeNode))
		default:
			if !reflect.DeepEqual(str, out2.Results[i]) {
				lib.Fatalf("%s(\n%s) == %s instead of\n %s\n",
					"BTreeMax",
					correct.FormatTree(arg1),
					lib.Format(out2.Results...),
					lib.Format(out1.Results...),
				)
			}
		}
	}
}

func main() {
	root := &correct.TreeNode{Data: "04"}
	rootS := &student.TreeNode{Data: "04"}

	ins := []string{"01", "07", "05", "12", "02", "03", "10"}

	for _, v := range ins {
		root = correct.BTreeInsertData(root, v)
		rootS = student.BTreeInsertData(rootS, v)
	}

	CompareReturn_max(correct.BTreeMax, student.BTreeMax, root, rootS)
}
