package student_test

import (
	"strconv"
	"testing"

	solution "./solutions"
	student "./student"
	"github.com/01-edu/z01"
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

func listPushBackTest8(l *ListS8, l1 *List8, data interface{}) {

	n := &Node8{Data: data}
	n1 := &NodeS8{Data: data}

	if l.Head == nil {
		l.Head = n
	} else {
		iterator := l.Head
		for iterator.Next != nil {
			iterator = iterator.Next
		}
		iterator.Next = n
	}

	if l1.Head == nil {
		l1.Head = n1
	} else {
		iterator1 := l1.Head
		for iterator1.Next != nil {
			iterator1 = iterator1.Next
		}
		iterator1.Next = n1
	}
}

func comparFuncList8(l *List8, l1 *ListS8, t *testing.T, f func(*Node8) bool, comp func(*Node8)) {
	funcFName := solution.GetName(f)
	funcComp := solution.GetName(comp)
	for l.Head != nil || l1.Head != nil {
		if (l.Head == nil && l1.Head != nil) || (l.Head != nil && l1.Head == nil) {
			t.Errorf("\nstudent list:%s\nlist:%s\nfunction f used: %s\nfunction comp: %s\n\nListForEachIf() == %v instead of %v\n\n",
				listToStringStu7(l1), solution.ListToString(l.Head), funcComp, funcFName, l1.Head, l.Head)
			return
		} else if l.Head.Data != l1.Head.Data {
			t.Errorf("\nstudent list:%s\nlist:%s\nfunction f used: %s\nfunction comp: %s\n\nListForEachIf() == %v instead of %v\n\n",
				listToStringStu7(l1), solution.ListToString(l.Head), funcComp, funcFName, l1.Head.Data, l.Head.Data)
			return
		}
		l.Head = l.Head.Next
		l1.Head = l1.Head.Next
	}
}

// exercise 10
// applies a function to an element of the linked solution.ListS
func TestListForEachIf(t *testing.T) {
	link := &ListS8{}
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
			listPushBackTest8(link, link2, arg.Data[i])
		}
		solution.ListForEachIf(link2, addOneS, solution.IsPositive_node)
		student.ListForEachIf(link, addOne, student.IsPositive_node)
		comparFuncList8(link2, link, t, student.IsPositive_node, addOne)

		solution.ListForEachIf(link2, subtract1_sol, solution.IsPositive_node)
		student.ListForEachIf(link, subtractOne, student.IsPositive_node)
		comparFuncList8(link2, link, t, student.IsPositive_node, subtractOne)

		link = &ListS8{}
		link2 = &List8{}
	}
}
