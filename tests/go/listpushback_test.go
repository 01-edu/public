package student_test

import (
	"strconv"
	"testing"

	solution "./solutions"
	student "./student"
)

type ListS = solution.List
type List = student.List

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
func comparFuncList(l *ListS, l1 *List, t *testing.T, data []interface{}) {
	for l.Head != nil || l1.Head != nil {
		if (l.Head == nil && l1.Head != nil) || (l.Head != nil && l1.Head == nil) {
			t.Errorf("\ndata used: %v\nstudent list:%s\nlist:%s\n\nListPushBack()== %v instead of %v\n\n",
				data, listToStringStu10(l1), solution.ListToString(l.Head), l1.Head, l.Head)
			return
		} else if l.Head.Data != l1.Head.Data {
			t.Errorf("\ndata used: %v\nstudent list:%s\nlist:%s\n\nListPushBack()== %v instead of %v\n\n",
				data, listToStringStu10(l1), solution.ListToString(l.Head), l1.Head.Data, l.Head.Data)
			return
		}
		l.Head = l.Head.Next
		l1.Head = l1.Head.Next
	}
}

// exercise 2
func TestListPushBack(t *testing.T) {
	link := &ListS{}
	link2 := &List{}

	table := []solution.NodeTest{}
	table = solution.ElementsToTest(table)
	table = append(table,
		solution.NodeTest{
			Data: []interface{}{"Hello", "man", "how are you"},
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.Data); i++ {
			student.ListPushBack(link2, arg.Data[i])
			solution.ListPushBack(link, arg.Data[i])
		}
		comparFuncList(link, link2, t, arg.Data)
	}
}
