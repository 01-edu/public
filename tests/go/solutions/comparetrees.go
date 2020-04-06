// +build ignore

package solutions

import (
	"testing"

	student "../student"
)

func CompareTrees(root *solutions.TreeNode, rootS *student.TreeNode, t *testing.T) {
	if root != nil && rootS != nil {
		CompareTrees(root.Left, rootS.Left, t)
		CompareTrees(root.Right, rootS.Right, t)
		if root.Data != rootS.Data {
			t.Fatalf("BTreeInsertData(%v), node == %v instead of %v ",
				root.Data, root.Data, rootS.Data)
		}
	} else if root != nil && rootS == nil {
		t.Fatalf("BTreeInsertData(%v), node == %v instead of %v ", root, root, rootS)
	} else if root == nil && rootS != nil {
		t.Fatalf("BTreeInsertData(%v), node == %v instead of %v ", root, root, rootS)
	}
}
