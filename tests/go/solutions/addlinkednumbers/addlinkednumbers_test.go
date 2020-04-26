package main

import (
	"strconv"

	"github.com/01-edu/z01"

	solutions "../../solutions"
)

type stuNode = NodeAddL
type solNode = solutions.NodeAddL

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

func stuNodeString(node *stuNode) string {
	var result string
	for tmp := node; tmp != nil; tmp = tmp.Next {
		result += strconv.Itoa(tmp.Num)
		if tmp.Next != nil {
			result += "->"
		}
	}
	return result
}

func solNodeString(node *solNode) string {
	var result string
	for tmp := node; tmp != nil; tmp = tmp.Next {
		result += strconv.Itoa(tmp.Num)
		if tmp.Next != nil {
			result += "->"
		}
	}
	return result
}

func compareNodes(stuResult *stuNode, solResult *solNode, num1, num2 int) {
	if stuResult == nil && solResult == nil {
		return
	}
	if stuResult != nil && solResult == nil {
		stuNum := stuNodeString(stuResult)
		z01.Fatalf("\nAddLinkedNumbers(%v, %v) == %v instead of %v\n\n",
			num1, num2, stuNum, "")
	}
	if stuResult == nil && solResult != nil {
		solNum := solNodeString(solResult)
		z01.Fatalf("\nAddLinkedNumbers(%v, %v) == %v instead of %v\n\n",
			num1, num2, "", solNum)
	}
	stuNum := stuNodeString(stuResult)
	solNum := solNodeString(solResult)
	if stuNum != solNum {
		z01.Fatalf("\nAddLinkedNumbers(%v, %v) == %v instead of %v\n\n",
			num1, num2, stuNum, solNum)
	}
}

func main() {
	args := [][2]int{{315, 592}}

	for i := 0; i < 15; i++ {
		args = append(args, [2]int{z01.RandPosZ(), z01.RandPosZ()})
	}

	for _, arg := range args {
		stuResult := AddLinkedNumbers(stuNumToList(arg[0]), stuNumToList(arg[1]))
		solResult := solutions.AddLinkedNumbers(solNumToList(arg[0]), solNumToList(arg[1]))

		compareNodes(stuResult, solResult, arg[0], arg[1])
	}
}
