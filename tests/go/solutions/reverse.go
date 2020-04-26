package solutions

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
