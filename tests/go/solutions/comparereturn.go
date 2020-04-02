// +build ignore

package solutions

import (
	"reflect"
	"testing"

	"github.com/01-edu/z01"

	student "../student"
)

func CompareReturn(t *testing.T, fn1, fn2, arg1, arg2 interface{}) {
	arar1 := []interface{}{arg1}
	arar2 := []interface{}{arg2}

	out1 := z01.Monitor(fn1, arar1)
	out2 := z01.Monitor(fn2, arar2)

	for i, v := range out1.Results {
		switch str := v.(type) {
		case *solutions.TreeNode:
			CompareNode(t, str, out2.Results[i].(*student.TreeNode))
		default:
			if !reflect.DeepEqual(str, out2.Results[i]) {
				t.Errorf("%s(%s) == %s instead of %s\n",
					z01.NameOfFunc(fn1),
					z01.Format(arg1),
					z01.Format(out1.Results...),
					z01.Format(out2.Results...),
				)
			}
		}
	}
}
