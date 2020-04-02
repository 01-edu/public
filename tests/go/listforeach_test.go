package student_test

import (
	"strconv"
	"testing"

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

func listPushBackTest7(l *ListS7, l1 *List7, data interface{}) {
	n := &Node7{Data: data}
	n1 := &NodeS7{Data: data}
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

func comparFuncList7(l *List7, l1 *ListS7, t *testing.T, f func(*Node7)) {
	funcName := solution.GetName(f)
	for l.Head != nil || l1.Head != nil {
		if (l.Head == nil && l1.Head != nil) || (l.Head != nil && l1.Head == nil) {
			t.Errorf("\nstudent list: %s\nlist: %s\nfunction used: %s\n\nListForEach() == %v instead of %v\n\n",
				listToStringStu8(l1), solution.ListToString(l.Head), funcName, l1.Head, l.Head)
			return
		} else if l.Head.Data != l1.Head.Data {
			t.Errorf("\nstudent list: %s\nlist: %s\nfunction used: %s\n\nListForEach() == %v instead of %v\n\n",
				listToStringStu8(l1), solution.ListToString(l.Head), funcName, l1.Head.Data, l.Head.Data)
			return
		}
		l.Head = l.Head.Next
		l1.Head = l1.Head.Next
	}
}

//exercise 9
//applies a function to the solution.ListS
func TestListForEach(t *testing.T) {
	link := &List7{}
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
			listPushBackTest7(link2, link, arg.Data[i])
		}
		student.ListForEach(link2, student.Add2_node)
		solution.ListForEach(link, solution.Add2_node)

		comparFuncList7(link, link2, t, student.Add2_node)

		student.ListForEach(link2, student.Subtract3_node)
		solution.ListForEach(link, solution.Subtract3_node)

		comparFuncList7(link, link2, t, student.Subtract3_node)

		link = &List7{}
		link2 = &ListS7{}
	}
}
