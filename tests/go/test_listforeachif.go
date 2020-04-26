package main

import (
	"strconv"

	"github.com/01-edu/z01"

	solution "./solutions"
	student "./student"
)

type Node8 = student.NodeL
type List8 = solution.List
type NodeS8 = solution.NodeL
type ListS8 = student.List

// function to apply, in listforeachif
func addOneS(node *NodeS8) {
	data := node.Data.(int)
	data++
	node.Data = interface{}(data)
}

// function to apply, in listforeachif
func addOne(node *Node8) {
	data := node.Data.(int)
	data++
	node.Data = interface{}(data)
}

func subtract1_sol(node *NodeS8) {
	data := node.Data.(int)
	data--
	node.Data = interface{}(data)
}

func subtractOne(node *Node8) {
	data := node.Data.(int)
	data--
	node.Data = interface{}(data)
}

func listToStringStu7(l *ListS8) string {
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

func listPushBackTest8(l1 *ListS8, l2 *List8, data interface{}) {
	n1 := &Node8{Data: data}
	n2 := &NodeS8{Data: data}

	if l1.Head == nil {
		l1.Head = n1
	} else {
		iterator := l1.Head
		for iterator.Next != nil {
			iterator = iterator.Next
		}
		iterator.Next = n1
	}

	if l2.Head == nil {
		l2.Head = n2
	} else {
		iterator1 := l2.Head
		for iterator1.Next != nil {
			iterator1 = iterator1.Next
		}
		iterator1.Next = n2
	}
}

func comparFuncList8(l1 *List8, l2 *ListS8, f func(*Node8) bool, comp func(*Node8)) {
	funcFName := solution.GetName(f)
	funcComp := solution.GetName(comp)
	for l1.Head != nil || l2.Head != nil {
		if (l1.Head == nil && l2.Head != nil) || (l1.Head != nil && l2.Head == nil) {
			z01.Fatalf("\nstudent list:%s\nlist:%s\nfunction f used: %s\nfunction comp: %s\n\nListForEachIf() == %v instead of %v\n\n",
				listToStringStu7(l2), solution.ListToString(l1.Head), funcComp, funcFName, l2.Head, l1.Head)
		}
		if l1.Head.Data != l2.Head.Data {
			z01.Fatalf("\nstudent list:%s\nlist:%s\nfunction f used: %s\nfunction comp: %s\n\nListForEachIf() == %v instead of %v\n\n",
				listToStringStu7(l2), solution.ListToString(l1.Head), funcComp, funcFName, l2.Head.Data, l1.Head.Data)
		}
		l1.Head = l1.Head.Next
		l2.Head = l2.Head.Next
	}
}

// applies a function to an element of the linked solution.ListS
func main() {
	link1 := &ListS8{}
	link2 := &List8{}

	table := []solution.NodeTest{}
	table = append(table,
		solution.NodeTest{
			Data: []interface{}{},
		},
	)

	// just numbers/ints
	for i := 0; i < 3; i++ {
		val := solution.NodeTest{
			Data: solution.ConvertIntToInterface(z01.MultRandInt()),
		}
		table = append(table, val)
	}
	// just strings
	for i := 0; i < 3; i++ {
		val := solution.NodeTest{
			Data: solution.ConvertIntToStringface(z01.MultRandWords()),
		}
		table = append(table, val)
	}
	table = append(table,
		solution.NodeTest{
			Data: []interface{}{"I", 1, "something", 2},
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.Data); i++ {
			listPushBackTest8(link1, link2, arg.Data[i])
		}
		solution.ListForEachIf(link2, addOneS, solution.IsPositive_node)
		student.ListForEachIf(link1, addOne, student.IsPositive_node)
		comparFuncList8(link2, link1, student.IsPositive_node, addOne)

		solution.ListForEachIf(link2, subtract1_sol, solution.IsPositive_node)
		student.ListForEachIf(link1, subtractOne, student.IsPositive_node)
		comparFuncList8(link2, link1, student.IsPositive_node, subtractOne)

		link1 = &ListS8{}
		link2 = &List8{}
	}
}
