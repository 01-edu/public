package solutions

// returs the node in a given position
func ListAt(l *NodeL, nbr int) *NodeL {
	index := 0
	count := 0
	head := l

	for head != nil {
		index++
		head = head.Next
	}
	if nbr <= index {
		for l != nil {
			if count == nbr {
				return l
			}
			count++
			l = l.Next
		}
	}
	return nil
}
