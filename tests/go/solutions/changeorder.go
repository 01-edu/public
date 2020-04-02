package solutions

//This solution is the comparing file of the staff
// Because the solution is a function,
//
//1) here the package is solutions
//2) it does not need an empty func main(){}
//3) its location is 1 level below the folder of the nauuo_test.go file

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

func Changeorder(node *NodeAddL) *NodeAddL {
	if node == nil {
		return node
	}
	ans := &NodeAddL{Num: node.Num}
	for tmp := node; tmp != nil; {
		tmp = tmp.Next
		if tmp == nil {
			break
		}
		tmp = tmp.Next
		if tmp == nil {
			break
		}
		ans = pushBack(ans, tmp.Num)
	}

	if node.Next == nil {
		return ans
	}
	for tmp := node.Next; tmp != nil; {
		ans = pushBack(ans, tmp.Num)
		tmp = tmp.Next
		if tmp == nil {
			break
		}
		tmp = tmp.Next
		if tmp == nil {
			break
		}
	}
	return ans
}
