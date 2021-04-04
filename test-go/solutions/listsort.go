package solutions

func ListSort(l *NodeI) *NodeI {
	head := l
	if head == nil {
		return nil
	}
	head.Next = ListSort(head.Next)

	if head.Next != nil && head.Data > head.Next.Data {
		head = move(head)
	}
	return head
}

func move(l *NodeI) *NodeI {
	p := l
	n := l.Next
	ret := n

	for n != nil && l.Data > n.Data {
		p = n
		n = n.Next
	}
	p.Next = l
	l.Next = n
	return ret
}
