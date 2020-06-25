package correct

import (
	"math/rand"
	"strconv"
)

type TreeNodeL struct {
	Left  *TreeNodeL
	Val   int
	Right *TreeNodeL
}

func IsSameTree(p *TreeNodeL, q *TreeNodeL) bool {
	if p == nil && q == nil {
		return true
	}
	return checkIfEq(p, q)
}

func checkIfEq(t1 *TreeNodeL, t2 *TreeNodeL) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && checkIfEq(t1.Right, t2.Right) && checkIfEq(t1.Left, t2.Left)
}

type stuTreeNode = TreeNodeL
type solTreeNode = correct.TreeNodeL

func insertStu(t *stuTreeNode, v int) *stuTreeNode {
	if t == nil {
		return &stuTreeNode{Left: nil, Val: v, Right: nil}
	}
	if v < t.Val {
		t.Left = insertStu(t.Left, v)
		return t
	}
	t.Right = insertStu(t.Right, v)
	return t
}

func insertSol(t *solTreeNode, v int) *solTreeNode {
	if t == nil {
		return &solTreeNode{Left: nil, Val: v, Right: nil}
	}
	if v < t.Val {
		t.Left = insertSol(t.Left, v)
		return t
	}
	t.Right = insertSol(t.Right, v)
	return t
}

func New(n, k int) (*solTreeNode, *stuTreeNode, *solTreeNode) {
	var cop *solTreeNode
	var stu *stuTreeNode
	var sol *solTreeNode
	for _, v := range rand.Perm(n) {
		cop = insertSol(cop, (1+v)*k)
		stu = insertStu(stu, (1+v)*k)
		sol = insertSol(sol, (1+v)*k)
	}
	return cop, stu, sol
}

func returnSolTree(root *solTreeNode) string {
	if root == nil {
		return ""
	}
	ans := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return ans
	}
	if root.Left != nil {
		ans += " " + returnSolTree(root.Left)
	}
	if root.Right != nil {
		ans += " " + returnSolTree(root.Right)
	}
	return ans
}

func returnStuTree(root *stuTreeNode) string {
	if root == nil {
		return ""
	}
	ans := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return ans
	}
	if root.Left != nil {
		ans += " " + returnStuTree(root.Left)
	}
	if root.Right != nil {
		ans += " " + returnStuTree(root.Right)
	}
	return ans
}

func compareResults(stuResult, solResult bool, solTree1, solTree2 *solTreeNode) {
	if stuResult != solResult {
		tree1 := returnSolTree(solTree1)
		tree2 := returnSolTree(solTree2)
		lib.Fatalf("\nIsSameTree(\"%v\", \"%v\") == \"%v\" instead of \"%v\"\n\n", tree1, tree2, stuResult, solResult)
	}
}

func main() {
	type node struct {
		n int
		k int
	}

	table := []node{}
	for i := 0; i < 15; i++ {
		value := node{lib.RandIntBetween(10, 15), lib.RandIntBetween(1, 10)}
		table = append(table, value)
	}

	// Check for different trees
	for _, arg := range table {
		cop1, stuTree1, solTree1 := New(arg.n, arg.k)
		cop2, stuTree2, solTree2 := New(arg.n, arg.k)
		stuResult := IsSameTree(stuTree1, stuTree2)
		solResult := correct.IsSameTree(solTree1, solTree2)

		compareResults(stuResult, solResult, cop1, cop2)
	}

	// Check for same trees
	for _, arg := range table {
		cop1, stuTree1, solTree1 := New(arg.n, arg.k)
		stuResult := IsSameTree(stuTree1, stuTree1)
		solResult := correct.IsSameTree(solTree1, solTree1)

		compareResults(stuResult, solResult, cop1, cop1)
	}
}
