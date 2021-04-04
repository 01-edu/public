package main

import (
	"reflect"

	student "student"

	"github.com/01-edu/public/test-go/lib"
	"github.com/01-edu/public/test-go/solutions"
)

func BTreeMinStu(root *student.TreeNode) *student.TreeNode {
	if root == nil || root.Left == nil {
		return root
	}
	return BTreeMinStu(root.Left)
}

func errorMessage_isbin(fn interface{}, root, a *solutions.TreeNode, b *student.TreeNode) {
	lib.Fatalf("%s(\n%s\n) == %s instead of %s\n",
		"BTreeIsBinary",
		solutions.FormatTree(root),
		b.Data,
		a.Data,
	)
}

func CompareNode_isbin(fn interface{}, arg1, a *solutions.TreeNode, b *student.TreeNode) {
	if a == nil || b == nil {
		lib.Fatalf("Expected %v instead of %v\n", a, b)
	}
	if a.Data != b.Data {
		errorMessage_isbin(fn, arg1, a, b)
	}
	if a.Parent != nil && b.Parent != nil && a.Parent.Data != b.Parent.Data {
		errorMessage_isbin(fn, arg1, a, b)
		lib.Fatalf("Expected parent value %v instead of %v\n", a.Parent.Data, b.Parent.Data)
	}
	if (a.Parent == nil && b.Parent != nil) || (a.Parent != nil && b.Parent == nil) {
		lib.Fatalf("Expected parent value %v instead of %v\n", a.Parent, b.Parent)
	}
	if a.Right != nil && b.Right != nil && a.Right.Data != b.Right.Data {
		errorMessage_isbin(fn, arg1, a, b)
		lib.Fatalf("Expected right child value %v instead of %v\n", a.Right.Data, b.Right.Data)
	}
	if (a.Right == nil && b.Right != nil) || (a.Right != nil && b.Right == nil) {
		lib.Fatalf("Expected right child value %v instead of %v\n", a.Right, b.Right)
	}
	if a.Left != nil && b.Left != nil && a.Left.Data != b.Left.Data {
		errorMessage_isbin(fn, arg1, a, b)
		lib.Fatalf("Expected left child value %v instead of %v\n", a.Left, b.Left)
	}
	if (a.Left == nil && b.Left != nil) || (a.Left != nil && b.Left == nil) {
		lib.Fatalf("Expected left child value %v instead of %v\n", a, b)
	}
}

func CompareReturn_isbin(fn1, fn2 interface{}, arg1 *solutions.TreeNode, arg2 interface{}) {
	arar1 := []interface{}{arg1}
	arar2 := []interface{}{arg2}

	out1 := lib.Monitor(fn1, arar1)
	out2 := lib.Monitor(fn2, arar2)

	for i, v := range out1.Results {
		switch str := v.(type) {
		case *solutions.TreeNode:
			CompareNode_isbin(fn1, arg1, str, out2.Results[i].(*student.TreeNode))
		default:
			if !reflect.DeepEqual(str, out2.Results[i]) {
				lib.Fatalf("%s(\n%s) == %s instead of %s\n",
					"BTreeIsBinary",
					solutions.FormatTree(arg1),
					lib.Format(out2.Results...),
					lib.Format(out1.Results...),
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

	CompareReturn_isbin(solutions.BTreeIsBinary, student.BTreeIsBinary, root, rootS)

	rootNB := &solutions.TreeNode{Data: "04"}
	rootNB_stu := &student.TreeNode{Data: "04"}
	// Test a non-binarysearch tree
	for _, v := range ins {
		rootNB = solutions.BTreeInsertData(rootNB, v)
		rootNB_stu = student.BTreeInsertData(rootNB_stu, v)
	}

	min := solutions.BTreeMin(rootNB)
	minStu := BTreeMinStu(rootNB_stu)

	min.Left = &solutions.TreeNode{Data: "123"}
	minStu.Left = &student.TreeNode{Data: "123"}

	CompareReturn_isbin(solutions.BTreeIsBinary, student.BTreeIsBinary, rootNB, rootNB_stu)
}
