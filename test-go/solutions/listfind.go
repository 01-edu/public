package solutions

// finds the element and returns the first data from the node that is a string
func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	iterator := l.Head
	for iterator != nil {
		if comp(iterator.Data, ref) {
			return &iterator.Data
		}

		iterator = iterator.Next
	}
	return nil
}

// defines for two elements the equality criteria
func CompStr(a, b interface{}) bool {
	return a == b
}
