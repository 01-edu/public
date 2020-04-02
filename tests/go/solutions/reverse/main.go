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

func pushFront(node *NodeAddL, num int) *NodeAddL {
	tmp := &NodeAddL{Num: num}
	if node == nil {
		return tmp
	}
	tmp.Next = node
	return tmp
}

func Reverse(node *NodeAddL) *NodeAddL {
	if node == nil {
		return node
	}
	ans := &NodeAddL{Num: node.Num}
	for tmp := node.Next; tmp != nil; tmp = tmp.Next {
		ans = pushFront(ans, tmp.Num)
	}
	return ans
}

func main() {

}
