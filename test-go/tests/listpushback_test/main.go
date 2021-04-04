package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/public/test-go/lib"
	"github.com/01-edu/public/test-go/solutions"
)

type (
	ListS = solutions.List
	List  = student.List
)

func listToStringStu10(l *List) string {
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
func comparFuncList(l *ListS, l1 *List, data []interface{}) {
	for l.Head != nil || l1.Head != nil {
		if (l.Head == nil && l1.Head != nil) || (l.Head != nil && l1.Head == nil) {
			lib.Fatalf("\ndata used: %v\nstudent list:%s\nlist:%s\n\nListPushBack()== %v instead of %v\n\n",
				data, listToStringStu10(l1), solutions.ListToString(l.Head), l1.Head, l.Head)
		}
		if l.Head.Data != l1.Head.Data {
			lib.Fatalf("\ndata used: %v\nstudent list:%s\nlist:%s\n\nListPushBack()== %v instead of %v\n\n",
				data, listToStringStu10(l1), solutions.ListToString(l.Head), l1.Head.Data, l.Head.Data)
		}
		l.Head = l.Head.Next
		l1.Head = l1.Head.Next
	}
}

func main() {
	link1 := &ListS{}
	link2 := &List{}

	table := []solutions.NodeTest{}
	table = solutions.ElementsToTest(table)
	table = append(table,
		solutions.NodeTest{
			Data: []interface{}{"Hello", "man", "how are you"},
		},
	)
	for _, arg := range table {
		for _, item := range arg.Data {
			student.ListPushBack(link2, item)
			solutions.ListPushBack(link1, item)
		}
		comparFuncList(link1, link2, arg.Data)
	}
}
