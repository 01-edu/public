package solutions

//gives the length of the List
func ListSize(l *List) int {
	count := 0
	iterator := l.Head
	for iterator != nil {
		count++
		iterator = iterator.Next
	}
	return count
}
