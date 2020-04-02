// +build ignore

package solutions

import (
	"testing"

	student "../student"
)

func CompareNode(t *testing.T, a *solutions.TreeNode, b *student.TreeNode) {
	if a != nil && b != nil {
		if a.Data != b.Data {
			t.Errorf("expected %s instead of %s\n",
				a.Data,
				b.Data,
			)
			CompareNode(t, a.Parent, b.Parent)
			CompareNode(t, a.Left, b.Left)
			CompareNode(t, a.Right, b.Right)
		}
	} else if a != nil && b == nil {
		t.Errorf("expected %s instead of %v\n",
			a.Data,
			b,
		)
	} else if a == nil && b != nil {
		t.Errorf("expected %v instead of %v\n",
			a,
			b.Data,
		)
	}
}
