package correct

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
