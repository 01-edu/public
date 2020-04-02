package main

type TreeNodeL struct {
    Left    *TreeNodeL
    Val     int
    Right   *TreeNodeL
}

func IsSameTree(p *TreeNodeL, q *TreeNodeL) bool {
    if (p == nil && q == nil) {
        return true
    }
    if (checkIfEq(p, q) == true) {
        return true
    }
    return false
}

func checkIfEq(t1 *TreeNodeL, t2 *TreeNodeL) bool {
    if (t1 == nil && t2 == nil) {
        return true
    }
    if (t1 == nil || t2 == nil) {
        return false;
    }
    return (t1.Val == t2.Val && checkIfEq(t1.Right, t2.Right) && checkIfEq(t1.Left, t2.Left))
}

func main() {
    
}