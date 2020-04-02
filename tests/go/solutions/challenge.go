package solutions

import (
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

func FormatTree(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := root.Data + "\n"
	res += formatSubTree(root, "")
	return res
}

func formatSubTree(root *TreeNode, prefix string) string {
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
		res += formatSubTree(root.Right, newPrefix)
	}

	if hasLeft {
		if hasRight {
			res += prefix
		}
		res += "└── " + root.Left.Data + "\n"
		res += formatSubTree(root.Left, prefix+"    ")
	}
	return res
}

func ParentList(root *TreeNode) string {
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
	r += ParentList(root.Left) + ParentList(root.Right)
	return r
}

func ChallengeTree(t *testing.T,
	fn1, fn2 interface{},
	arg1 *TreeNode, arg2 interface{},
	args ...interface{}) {

	args1 := []interface{}{arg1}
	args2 := []interface{}{arg2}

	if args != nil {
		for _, v := range args {
			args1 = append(args1, v)
			args2 = append(args2, v)
		}
	}
	st1 := z01.Monitor(fn1, args1)
	st2 := z01.Monitor(fn2, args2)

	if st1.Stdout != st2.Stdout {
		t.Errorf("%s(\n%s)\n prints %s instead of %s\n",
			z01.NameOfFunc(fn2),
			FormatTree(arg1),
			z01.Format(st2.Stdout),
			z01.Format(st1.Stdout),
		)
	}

}

func Challenge(t *testing.T, fn1, fn2 interface{}, arg1, arg2 interface{}, args ...interface{}) {
	args1 := []interface{}{arg1}
	args2 := []interface{}{arg2}

	if args != nil {
		for _, v := range args {
			args1 = append(args1, v)
			args2 = append(args2, v)
		}
	}
	st1 := z01.Monitor(fn1, args1)
	st2 := z01.Monitor(fn2, args2)

	if st1.Stdout != st2.Stdout {
		t.Errorf("%s(%s) prints %s instead of %s\n",
			z01.NameOfFunc(fn2),
			z01.Format(arg2),
			z01.Format(st2.Stdout),
			z01.Format(st1.Stdout),
		)
	}

}

func PrintList(n *NodeI) string {
	var res string
	it := n
	for it != nil {
		res += strconv.Itoa(it.Data) + "-> "
		it = it.Next
	}
	res += "<nil>"
	return res
}

func ListToString(n *NodeL) string {
	var res string
	it := n
	for it != nil {
		switch it.Data.(type) {
		case int:
			res += strconv.Itoa(it.Data.(int)) + "-> "
		case string:
			res += it.Data.(string) + "-> "
		}
		it = it.Next
	}
	res += "<nil>"
	return res
}

func ConvertIntToInterface(t []int) []interface{} {
	RandLen := z01.RandIntBetween(0, len(t))
	s := make([]interface{}, RandLen)
	for j := 0; j < RandLen; j++ {
		for i := 0; i < z01.RandIntBetween(1, len(t)); i++ {
			s[j] = t[i]
		}
	}
	return s
}

func ConvertIntToStringface(t []string) []interface{} {
	RandLen := z01.RandIntBetween(0, len(t))
	s := make([]interface{}, RandLen)
	for j := 0; j < RandLen; j++ {
		for i := 0; i < z01.RandIntBetween(1, len(t)); i++ {
			s[j] = t[i]
		}
	}
	return s
}

type NodeTest struct {
	Data []interface{}
}

func ElementsToTest(table []NodeTest) []NodeTest {
	table = append(table,
		NodeTest{
			Data: []interface{}{},
		},
	)
	for i := 0; i < 3; i++ {
		val := NodeTest{
			Data: ConvertIntToInterface(z01.MultRandInt()),
		}
		table = append(table, val)
	}
	for i := 0; i < 3; i++ {
		val := NodeTest{
			Data: ConvertIntToStringface(z01.MultRandWords()),
		}
		table = append(table, val)
	}
	return table
}

// GetName gets the function name
func GetName(f interface{}) string {
	pathFuncUsed := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), "/")
	return pathFuncUsed[len(pathFuncUsed)-1]
}
