package solutions

type Node struct {
	Data int
}

func CreateElem(n *Node, value int) {
	n.Data = value
}
