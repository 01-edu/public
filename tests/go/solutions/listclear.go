package solutions

//cleans the list
func ListClear(l *List) {

	temp := l.Head
	next := l.Head
	for temp != nil {
		next = temp.Next
		temp = nil
		temp = next
	}

	l.Head = nil
}
