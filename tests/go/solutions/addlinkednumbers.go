package solutions

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

func AddLinkedNumbers(num1, num2 *NodeAddL) *NodeAddL {
	var n1, n2, r int
	var result *NodeAddL

	for tmp := num1; tmp != nil; tmp = tmp.Next {
		n1 = n1*10 + tmp.Num
	}

	for tmp := num2; tmp != nil; tmp = tmp.Next {
		n2 = n2*10 + tmp.Num
	}

	r = n1 + n2
	for r > 0 {
		mod := r % 10
		r /= 10
		result = pushFront(result, mod)
	}
	return result
}
