package main

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	count := 0
	iterator := l.Head
	for iterator != nil {
		count++
		iterator = iterator.Next
	}
	return count
}

func main() {
}
