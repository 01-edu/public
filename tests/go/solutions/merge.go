package solutions

type TreeNodeM struct {
	Left  *TreeNodeM
	Val   int
	Right *TreeNodeM
}

func MergeTrees(t1 *TreeNodeM, t2 *TreeNodeM) *TreeNodeM {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val = t1.Val + t2.Val
	t1.Left = MergeTrees(t1.Left, t2.Left)
	t1.Right = MergeTrees(t1.Right, t2.Right)
	return t1
}
