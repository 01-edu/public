package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
)

type stuNode = NodeAddL
type solNode = correct.NodeAddL

func stuPushFront(node *stuNode, num int) *stuNode {
	tmp := &stuNode{Num: num}
	tmp.Next = node
	return tmp
}

func stuNumToList(num int) *stuNode {
	var res *stuNode
	for num > 0 {
		res = stuPushFront(res, num%10)
		num /= 10
	}
	return res
}

func stuListToNum(node *stuNode) int {
	var n int

	for tmp := node; tmp != nil; tmp = tmp.Next {
		n = n*10 + tmp.Num
	}
	return n
}

func solPushFront(node *solNode, num int) *solNode {
	tmp := &solNode{Num: num}
	tmp.Next = node
	return tmp
}

func solNumToList(num int) *solNode {
	var res *solNode
	for num > 0 {
		res = solPushFront(res, num%10)
		num /= 10
	}
	return res
}

func solListToNum(node *solNode) int {
	var n int

	for tmp := node; tmp != nil; tmp = tmp.Next {
		n = n*10 + tmp.Num
	}
	return n
}

func compareNodes(stuResult *stuNode, solResult *solNode, num1 int) {
	if stuResult != nil && solResult != nil {
		stuNum := stuListToNum(stuResult)
		solNum := solListToNum(solResult)
		if stuNum != solNum {
			z01.Fatalf("\nReverse(%d) == %v instead of %v\n\n",
				num1, stuNum, solNum)
		}
	} else if stuResult != nil && solResult == nil {
		stuNum := stuListToNum(stuResult)
		z01.Fatalf("\nReverse(%d) == %v instead of %v\n\n",
			num1, stuNum, "")
	} else if stuResult == nil && solResult != nil {
		solNum := solListToNum(solResult)
		z01.Fatalf("\nReverse(%d) == %v instead of %v\n\n",
			num1, "", solNum)
	}
}

func main() {
	table := []int{123456543}

	table = append(table, z01.MultRandIntBetween(0, 1000000000)...)
	for _, arg := range table {
		stuResult := Reverse(stuNumToList(arg))
		solResult := correct.Reverse(solNumToList(arg))

		compareNodes(stuResult, solResult, arg)
	}
}
