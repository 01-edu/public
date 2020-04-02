package solutions

func ListSort(l *NodeI) *NodeI {

	Head := l
	if Head == nil {
		return nil
	}
	Head.Next = ListSort(Head.Next)

	if Head.Next != nil && Head.Data > Head.Next.Data {
		Head = move(Head)
	}
	return Head
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
