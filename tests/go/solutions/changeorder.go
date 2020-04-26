package correct

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

func Changeorder(node *NodeAddL) *NodeAddL {
	if node == nil {
		return node
	}
	ans := &NodeAddL{Num: node.Num}
	for tmp := node; tmp != nil; {
		tmp = tmp.Next
		if tmp == nil {
			break
		}
		tmp = tmp.Next
		if tmp == nil {
			break
		}
		ans = pushBack(ans, tmp.Num)
	}

	if node.Next == nil {
		return ans
	}
	for tmp := node.Next; tmp != nil; {
		ans = pushBack(ans, tmp.Num)
		tmp = tmp.Next
		if tmp == nil {
			break
		}
		tmp = tmp.Next
		if tmp == nil {
			break
		}
	}
	return ans
}
