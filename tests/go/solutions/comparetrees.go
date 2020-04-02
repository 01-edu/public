// +build ignore

package solutions

import (
	student "../student"
	"testing"
)

func CompareTrees(root *solutions.TreeNode, rootS *student.TreeNode, t *testing.T) {
	if root != nil && rootS != nil {
		CompareTrees(root.Left, rootS.Left, t)
		CompareTrees(root.Right, rootS.Right, t)
		if root.Data != rootS.Data {
			t.Errorf("BTreeInsertData(%v), node == %v instead of %v ",
				root.Data, root.Data, rootS.Data)
		}
	} else if root != nil && rootS == nil {
		t.Errorf("BTreeInsertData(%v), node == %v instead of %v ", root, root, rootS)
	} else if root == nil && rootS != nil {
		t.Errorf("BTreeInsertData(%v), node == %v instead of %v ", root, root, rootS)
	}
}
