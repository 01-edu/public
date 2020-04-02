package student_test

import (
	"reflect"
	"testing"

	solutions "./solutions"
	student "./student"

	"github.com/01-edu/z01"
)

func errorMessage_min(t *testing.T, fn interface{}, root, a *solutions.TreeNode, b *student.TreeNode) {
	t.Errorf("%s(\n%s) == %s instead of %s\n",
		z01.NameOfFunc(fn),
		solutions.FormatTree(root),
		b.Data,
		a.Data,
	)
}

func CompareNode_min(t *testing.T, fn interface{}, arg1, a *solutions.TreeNode, b *student.TreeNode) {
	if a == nil || b == nil {
		t.Errorf("Expected %v instead of %v\n", a, b)
		return
	}

	if a.Data != b.Data {
		errorMessage_min(t, fn, arg1, a, b)
	}

	if a.Parent != nil && b.Parent != nil {
		if a.Parent.Data != b.Parent.Data {
			errorMessage_min(t, fn, arg1, a, b)
			t.Errorf("Expected parent value %v instead of %v\n", a.Parent.Data, b.Parent.Data)
		}
	} else if (a.Parent == nil && b.Parent != nil) || (a.Parent != nil && b.Parent == nil) {
		t.Errorf("Expected parent value %v instead of %v\n", a.Parent, b.Parent)
	}

	if a.Right != nil && b.Right != nil {
		if a.Right.Data != b.Right.Data {
			errorMessage_min(t, fn, arg1, a, b)
			t.Errorf("Expected right child value %v instead of %v\n", a.Right.Data, b.Right.Data)
		}
	} else if (a.Right == nil && b.Right != nil) || (a.Right != nil && b.Right == nil) {
		t.Errorf("Expected right child value %v instead of %v\n", a.Right, b.Right)
	}

	if a.Left != nil && b.Left != nil {
		if a.Left.Data != b.Left.Data {
			errorMessage_min(t, fn, arg1, a, b)
			t.Errorf("Expected left child value %v instead of %v\n", a.Left, b.Left)
		}
	} else if (a.Left == nil && b.Left != nil) || (a.Left != nil && b.Left == nil) {
		t.Errorf("Expected left child value %v instead of %v\n", a, b)
	}
}

func CompareReturn_min(t *testing.T, fn1, fn2, arg1, arg2 interface{}) {
	arar1 := []interface{}{arg1}
	arar2 := []interface{}{arg2}

	out1 := z01.Monitor(fn1, arar1)
	out2 := z01.Monitor(fn2, arar2)

	for i, v := range out1.Results {
		switch str := v.(type) {
		case *solutions.TreeNode:
			CompareNode_min(t, fn1, arg1.(*solutions.TreeNode), str, out2.Results[i].(*student.TreeNode))
		default:
			if !reflect.DeepEqual(str, out2.Results[i]) {
				t.Errorf("%s(%s) == %s instead of %s\n",
					z01.NameOfFunc(fn1),
					z01.Format(arg1),
					z01.Format(out2.Results...),
					z01.Format(out1.Results...),
				)
			}
		}
	}
}

func TestBTreeMin(t *testing.T) {
	root := &solutions.TreeNode{Data: "04"}
	rootS := &student.TreeNode{Data: "04"}

	ins := []string{"03", "02", "01", "07", "05", "12", "10"}

	for _, v := range ins {
		root = solutions.BTreeInsertData(root, v)
		rootS = student.BTreeInsertData(rootS, v)
	}

	CompareReturn_min(t, solutions.BTreeMin, student.BTreeMin, root, rootS)
}
