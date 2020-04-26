package main

import (
	"strconv"

	"github.com/01-edu/z01"

	solution "./solutions"
	student "./student"
)

type Node7 = student.NodeL
type List7 = solution.List
type NodeS7 = solution.NodeL
type ListS7 = student.List

func listToStringStu8(l *ListS7) string {
	var res string
	it := l.Head
	for it != nil {
		switch it.Data.(type) {
		case int:
			res += strconv.Itoa(it.Data.(int)) + "-> "
		case string:
			res += it.Data.(string) + "-> "
		}
		it = it.Next
	}
	res += "<nil>"
	return res
}

func listPushBackTest7(l1 *ListS7, l2 *List7, data interface{}) {
	n := &Node7{Data: data}
	n1 := &NodeS7{Data: data}
	if l1.Head == nil {
		l1.Head = n
	} else {
		iterator := l1.Head
		for iterator.Next != nil {
			iterator = iterator.Next
		}
		iterator.Next = n
	}
	if l2.Head == nil {
		l2.Head = n1
	} else {
		iterator1 := l2.Head
		for iterator1.Next != nil {
			iterator1 = iterator1.Next
		}
		iterator1.Next = n1
	}
}

func comparFuncList7(l1 *List7, l2 *ListS7, f func(*Node7)) {
	funcName := solution.GetName(f)
	for l1.Head != nil || l2.Head != nil {
		if (l1.Head == nil && l2.Head != nil) || (l1.Head != nil && l2.Head == nil) {
			z01.Fatalf("\nstudent list: %s\nlist: %s\nfunction used: %s\n\nListForEach() == %v instead of %v\n\n",
				listToStringStu8(l2), solution.ListToString(l1.Head), funcName, l2.Head, l1.Head)
		}
		if l1.Head.Data != l2.Head.Data {
			z01.Fatalf("\nstudent list: %s\nlist: %s\nfunction used: %s\n\nListForEach() == %v instead of %v\n\n",
				listToStringStu8(l2), solution.ListToString(l1.Head), funcName, l2.Head.Data, l1.Head.Data)
		}
		l1.Head = l1.Head.Next
		l2.Head = l2.Head.Next
	}
}

// applies a function to the solution.ListS
func main() {
	link1 := &List7{}
	link2 := &ListS7{}
	table := []solution.NodeTest{}
	table = solution.ElementsToTest(table)
	table = append(table,
		solution.NodeTest{
			Data: []interface{}{"I", 1, "something", 2},
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.Data); i++ {
			listPushBackTest7(link2, link1, arg.Data[i])
		}
		student.ListForEach(link2, student.Add2_node)
		solution.ListForEach(link1, solution.Add2_node)

		comparFuncList7(link1, link2, student.Add2_node)

		student.ListForEach(link2, student.Subtract3_node)
		solution.ListForEach(link1, solution.Subtract3_node)

		comparFuncList7(link1, link2, student.Subtract3_node)

		link1 = &List7{}
		link2 = &ListS7{}
	}
}
