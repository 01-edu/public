package main

import (
	"strconv"
	"testing"

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

func compareNodes(t *testing.T, stuResult *stuNode, solResult *solNode, num1, num2 int) {
	if stuResult == nil && solResult == nil {

	} else if stuResult != nil && solResult == nil {
		stuNum := stuNodeString(stuResult)
		t.Errorf("\nAddLinkedNumbers(%v, %v) == %v instead of %v\n\n",
			num1, num2, stuNum, "")
	} else if stuResult == nil && solResult != nil {
		solNum := solNodeString(solResult)
		t.Errorf("\nAddLinkedNumbers(%v, %v) == %v instead of %v\n\n",
			num1, num2, "", solNum)
	} else {
		stuNum := stuNodeString(stuResult)
		solNum := solNodeString(solResult)
		if stuNum != solNum {
			t.Errorf("\nAddLinkedNumbers(%v, %v) == %v instead of %v\n\n",
				num1, num2, stuNum, solNum)
		}
	}
}

func TestAddLinkedNumbers(t *testing.T) {
	type node struct {
		num1 int
		num2 int
	}

	table := []node{}

	table = append(table,
		node{315, 592},
	)

	for i := 0; i < 15; i++ {
		value := node{
			num1: z01.RandIntBetween(0, 1000000000),
			num2: z01.RandIntBetween(0, 1000000000),
		}

		table = append(table, value)
	}

	for _, arg := range table {
		stuResult := AddLinkedNumbers(stuNumToList(arg.num1), stuNumToList(arg.num2))
		solResult := solutions.AddLinkedNumbers(solNumToList(arg.num1), solNumToList(arg.num2))

		compareNodes(t, stuResult, solResult, arg.num1, arg.num2)
	}
}
