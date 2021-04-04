package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/public/test-go/lib"
	"github.com/01-edu/public/test-go/tests/correct"
)

type (
	Node8  = student.NodeL
	List8  = correct.List
	NodeS8 = correct.NodeL
	ListS8 = student.List
)

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
	funcFName := correct.GetName(f)
	funcComp := correct.GetName(comp)
	for l1.Head != nil || l2.Head != nil {
		if (l1.Head == nil && l2.Head != nil) || (l1.Head != nil && l2.Head == nil) {
			lib.Fatalf("\nstudent list:%s\nlist:%s\nfunction f used: %s\nfunction comp: %s\n\nListForEachIf() == %v instead of %v\n\n",
				listToStringStu7(l2), correct.ListToString(l1.Head), funcComp, funcFName, l2.Head, l1.Head)
		}
		if l1.Head.Data != l2.Head.Data {
			lib.Fatalf("\nstudent list:%s\nlist:%s\nfunction f used: %s\nfunction comp: %s\n\nListForEachIf() == %v instead of %v\n\n",
				listToStringStu7(l2), correct.ListToString(l1.Head), funcComp, funcFName, l2.Head.Data, l1.Head.Data)
		}
		l1.Head = l1.Head.Next
		l2.Head = l2.Head.Next
	}
}

// applies a function to an element of the linked correct.ListS
func main() {
	link1 := &ListS8{}
	link2 := &List8{}

	table := []correct.NodeTest{}
	table = append(table,
		correct.NodeTest{
			Data: []interface{}{},
		},
	)

	// just numbers/ints
	for i := 0; i < 3; i++ {
		val := correct.NodeTest{
			Data: correct.ConvertIntToInterface(lib.MultRandInt()),
		}
		table = append(table, val)
	}
	// just strings
	for i := 0; i < 3; i++ {
		val := correct.NodeTest{
			Data: correct.ConvertIntToStringface(lib.MultRandWords()),
		}
		table = append(table, val)
	}
	table = append(table,
		correct.NodeTest{
			Data: []interface{}{"I", 1, "something", 2},
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.Data); i++ {
			listPushBackTest8(link1, link2, arg.Data[i])
		}
		correct.ListForEachIf(link2, addOneS, correct.IsPositive_node)
		student.ListForEachIf(link1, addOne, student.IsPositive_node)
		comparFuncList8(link2, link1, student.IsPositive_node, addOne)

		correct.ListForEachIf(link2, subtract1_sol, correct.IsPositive_node)
		student.ListForEachIf(link1, subtractOne, student.IsPositive_node)
		comparFuncList8(link2, link1, student.IsPositive_node, subtractOne)

		link1 = &ListS8{}
		link2 = &List8{}
	}
}
