package main

import (
	"io"
	// "os"
	"fmt"
    "math/rand"
	"testing"
    "strconv"

	// "github.com/01-edu/z01"
	solutions "../../solutions"
)

type stuNode = TNode
type solNode = solutions.TNode

func solInsert(N *solNode, newVal int) {
    if N == nil {
        return
    } else if newVal <= N.Val {
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
    } else if newVal <= N.Val {
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

func IsIdentical(root1 *solNode, root2 *stuNode) int {
    if root1 == nil && root2 == nil {
        return 1
    } else if root1 == nil && root2 != nil {
        return 0
    } else if root2 == nil && root1 != nil {
        return 0
    } else {
        if root1.Val == root2.Val && IsIdentical(root1.Left, root2.Left) == 1 && IsIdentical(root1.Right, root2.Right) == 1 {
            return 1
        } else {
            return 0
        }
    }
    return 1
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
    if (root == nil) {
        return ""
    }
    ans := strconv.Itoa(root.Val)
    if (root.Left == nil && root.Right == nil) {
        return ans
    }
    if (root.Left != nil) {
       ans += " " + returnStuTree(root.Left)
    }
    if (root.Right != nil) {
       ans += " " + returnStuTree(root.Right)
    }
    return ans
}

func returnSolTree(root *solNode) string{
    if (root == nil) {
        return ""
    }
    ans := strconv.Itoa(root.Val)
    if (root.Left == nil && root.Right == nil) {
        return ans
    }
    if (root.Left != nil) {
       ans += " " + returnSolTree(root.Left)
    }
    if (root.Right != nil) {
       ans += " " + returnSolTree(root.Right)
    }
    return ans
}

func TestInvertTree(t *testing.T) {
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
	    solutions.InvertTree(tree)
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

	    ret := IsIdentical(tree, TestTree)
	    if ret != 1 {
            tree1 := returnSolTree(temp)
            tree2 := returnStuTree(tmp)
            t.Errorf("\n\"%v\" instead of \"%v\"\n\n", tree1, tree2)
	        // t.Errorf("\nError\n\n")
	    }
	}
}
