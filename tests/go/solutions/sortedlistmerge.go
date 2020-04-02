package solutions

func SortedListMerge(l1 *NodeI, l2 *NodeI) *NodeI {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Data <= l2.Data {
		l1.Next = SortedListMerge(l1.Next, l2)
		return l1
	}
	l2.Next = SortedListMerge(l1, l2.Next)
	return l2

}
