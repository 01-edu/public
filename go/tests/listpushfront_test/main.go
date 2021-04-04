package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/public/go/tests/func/correct"

	"github.com/01-edu/public/go/tests/lib"
)

type (
	ListSa = correct.List
	Lista  = student.List
)

func listToStringStu11(l *Lista) string {
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

// makes the test, compares 2 lists one from the solutions and the other from the student
func comparFuncList1(l *Lista, l1 *ListSa, data []interface{}) {
	for l.Head != nil || l1.Head != nil {
		if (l.Head == nil && l1.Head != nil) || (l.Head != nil && l1.Head == nil) {
			lib.Fatalf("\ndata used: %v\nstudent list:%s\nlist:%s\n\nListPushFront()== %v instead of %v\n\n",
				data, listToStringStu11(l), correct.ListToString(l1.Head), l.Head, l1.Head)
		}
		if l.Head.Data != l1.Head.Data {
			lib.Fatalf("\ndata used: %v\nstudent list:%s\nlist:%s\n\nListPushFront()== %v instead of %v\n\n",
				data, listToStringStu11(l), correct.ListToString(l1.Head), l.Head, l1.Head)
		}
		l1.Head = l1.Head.Next
		l.Head = l.Head.Next
	}
}

// to insert a value in the first position of the list
func main() {
	link1 := &Lista{}
	link2 := &ListSa{}
	table := []correct.NodeTest{}
	table = correct.ElementsToTest(table)
	table = append(table,
		correct.NodeTest{
			Data: []interface{}{"Hello", "man", "how are you"},
		},
	)
	for _, arg := range table {
		for _, item := range arg.Data {
			student.ListPushFront(link1, item)
			correct.ListPushFront(link2, item)
		}
		comparFuncList1(link1, link2, arg.Data)
		link1 = &Lista{}
		link2 = &ListSa{}
	}
}
