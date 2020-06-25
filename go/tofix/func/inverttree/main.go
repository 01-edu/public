package correct

import (
	"fmt"
	"io"
	"math/rand"
	"strconv"
)

type TNode struct {
	Val   int
	Left  *TNode
	Right *TNode
}

func Invert(root *TNode) {
	if root != nil {
		temp := root.Left
		root.Left = root.Right
		root.Right = temp

		Invert(root.Left)
		Invert(root.Right)
	}
}

func InvertTree(root *TNode) *TNode {
	Invert(root)
	return root
}

type stuNode = TNode
type solNode = correct.TNode

func solInsert(N *solNode, newVal int) {
	if N == nil {
		return
	}
	if newVal <= N.Val {
		if N.Left == nil {
			N.Left = &solNode{Val: newVal, Left: nil, Right: nil}
		} else {
			solInsert(N.Left, newVal)
		}
	} else {
		if N.Right == nil {
			N.Right = &solNode{Val: newVal, Left: nil, Right: nil}
		} else {
			solInsert(N.Right, newVal)
		}
	}
}

func stuInsert(N *stuNode, newVal int) {
	if N == nil {
		return
	}
	if newVal <= N.Val {
		if N.Left == nil {
			N.Left = &stuNode{Val: newVal, Left: nil, Right: nil}
		} else {
			stuInsert(N.Left, newVal)
		}
	} else {
		if N.Right == nil {
			N.Right = &stuNode{Val: newVal, Left: nil, Right: nil}
		} else {
			stuInsert(N.Right, newVal)
		}
	}
}

func IsIdentical(root1 *solNode, root2 *stuNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil && root2 != nil {
		return false
	}
	if root2 == nil && root1 != nil {
		return false
	}
	return root1.Val == root2.Val && IsIdentical(root1.Left, root2.Left) && IsIdentical(root1.Right, root2.Right)
}

func stuPrint(w io.Writer, node *stuNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Val)
	stuPrint(w, node.Left, ns+2, 'L')
	stuPrint(w, node.Right, ns+2, 'R')
}

func solPrint(w io.Writer, node *solNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Val)
	solPrint(w, node.Left, ns+2, 'L')
	solPrint(w, node.Right, ns+2, 'R')
}

func returnStuTree(root *stuNode) string {
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

func returnSolTree(root *solNode) string {
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

func main() {
	root, val1, val2, val3, val4 := 0, 0, 0, 0, 0

	root = rand.Intn(30)
	tree := &solNode{Val: root, Left: nil, Right: nil}
	TestTree := &stuNode{Val: root, Left: nil, Right: nil}
	for i := 0; i < 15; i++ {
		tree = &solNode{Val: root, Left: nil, Right: nil}
		temp := tree
		val1, val2, val3, val4 = rand.Intn(30), rand.Intn(30), rand.Intn(30), rand.Intn(30)
		solInsert(tree, val1)
		solInsert(tree, val2)
		solInsert(tree, val3)
		solInsert(tree, val4)
		// solPrint(os.Stdout, tree, 0, 'M')
		correct.InvertTree(tree)
		// solPrint(os.Stdout, tree, 0, 'M')

		TestTree = &stuNode{Val: root, Left: nil, Right: nil}
		tmp := TestTree
		stuInsert(TestTree, val1)
		stuInsert(TestTree, val2)
		stuInsert(TestTree, val3)
		stuInsert(TestTree, val4)
		// stuPrint(os.Stdout, TestTree, 0, 'M')
		InvertTree(TestTree)
		// stuPrint(os.Stdout, TestTree, 0, 'M')

		if !IsIdentical(tree, TestTree) {
			tree1 := returnSolTree(temp)
			tree2 := returnStuTree(tmp)
			lib.Fatalf("\n\"%v\" instead of \"%v\"\n\n", tree1, tree2)
			// lib.Fatalf("\nError\n\n")
		}
	}
}
