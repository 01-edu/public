package correct

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func pushBack(node *NodeAddL, num int) *NodeAddL {
	nw := &NodeAddL{Num: num}
	if node == nil {
		return nw
	}
	for tmp := node; tmp != nil; tmp = tmp.Next {
		if tmp.Next == nil {
			tmp.Next = nw
			return node
		}
	}
	return node
}

func Sortll(node *NodeAddL) *NodeAddL {
	if node == nil {
		return node
	}

	for first := node; first != nil; first = first.Next {
		for second := first.Next; second != nil; second = second.Next {
			if first.Num < second.Num {
				tmp := first.Num
				first.Num = second.Num
				second.Num = tmp
			}
		}
	}
	return node
}
