package solutions

type Nodelist struct {
	Data int
	Next *Nodelist
}

func SortList(l *Nodelist, cmp func(a, b int) bool) *Nodelist {

	head := l
	if head == nil {
		return nil
	}
	head.Next = SortList(head.Next, cmp)

	if head.Next != nil && cmp(head.Data, head.Next.Data) {
		head = moveValue(head, cmp)
	}
	return head
}

func moveValue(l *Nodelist, cmp func(a, b int) bool) *Nodelist {
	p := l
	n := l.Next
	ret := n

	for n != nil && cmp(l.Data, n.Data) {
		p = n
		n = n.Next
	}
	p.Next = l
	l.Next = n
	return ret
}
