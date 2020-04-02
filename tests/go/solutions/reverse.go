package solutions

//This solution is the comparing file of the staff
// Because the solution is a function,
//
//1) here the package is solutions
//2) it does not need an empty func main(){}
//3) its location is 1 level below the folder of the nauuo_test.go file

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
