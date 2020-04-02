package main

//This solution is the placeholder of the student solution
// for an exercise in the exam asking for a Function
//Remember the disclaimer!!!!
//1) here the package is main
//2) It does need an empty func main(){}

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func pushBack(node *NodeAddL, num int) *NodeAddL {
	nw := &NodeAddL{Num: num}
	if node == nil {
		return nw
	}
	for tmp := node; tmp != nil; tmp = tmp.Next {
		if tmp.Next == nil {
			tmp.Next = nw
			return node
		}
	}
	return node
}

func Sortll(node *NodeAddL) *NodeAddL {
	if node == nil {
		return node
	}

	for first := node; first != nil; first = first.Next {
		for second := first.Next; second != nil; second = second.Next {
			if first.Num < second.Num {
				tmp := first.Num
				first.Num = second.Num
				second.Num = tmp
			}
		}
	}
	return node
}
func main() {

}
