package solutions

// gives the last element of the list
func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	for l.Head != nil {
		if l.Head.Next == nil {
			return l.Head.Data
		}
		l.Head = l.Head.Next
	}
	return l.Head.Data
}
