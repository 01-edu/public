package solutions

//This solution is the comparing file of the staff
// Because the solution is a function,
//
//1) here the package is solutions
//2) it does not need an empty func main(){}
//3) its location is 1 level below the folder of the nauuo_test.go file

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
