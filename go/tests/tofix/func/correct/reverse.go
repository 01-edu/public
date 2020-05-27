package correct

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func pushFront(node *NodeAddL, num int) *NodeAddL {
	tmp := &NodeAddL{Num: num}
	if node == nil {
		return tmp
	}
	tmp.Next = node
	return tmp
}

func Reverse(node *NodeAddL) *NodeAddL {
	if node == nil {
		return node
	}
	ans := &NodeAddL{Num: node.Num}
	for tmp := node.Next; tmp != nil; tmp = tmp.Next {
		ans = pushFront(ans, tmp.Num)
	}
	return ans
}
