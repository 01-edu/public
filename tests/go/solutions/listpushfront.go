package solutions

//inserts node on the first position of the list
func ListPushFront(l *List, data interface{}) {

	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		return
	}

	n.Next = l.Head
	l.Head = n
}
