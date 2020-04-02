package student_test

import (
	"strconv"
	"testing"

	solution "./solutions"
	student "./student"
)

type Node4 = student.NodeL
type List4 = solution.List
type NodeS4 = solution.NodeL
type ListS4 = student.List

func listToStringStu5(l *ListS4) string {
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

func listPushBackTest4(l *ListS4, l1 *List4, data interface{}) {
	n := &Node4{Data: data}
	n1 := &NodeS4{Data: data}
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

//exercise 6
//simply cleans the linked solution.ListS
func TestListClear(t *testing.T) {
	link := &List4{}
	link2 := &ListS4{}

	table := []solution.NodeTest{}

	table = solution.ElementsToTest(table)

	table = append(table,
		solution.NodeTest{
			Data: []interface{}{"I", 1, "something", 2},
		},
	)

	for _, arg := range table {

		for i := 0; i < len(arg.Data); i++ {
			listPushBackTest4(link2, link, arg.Data[i])
		}
		solution.ListClear(link)
		student.ListClear(link2)

		if link2.Head != nil {
			t.Errorf("\nstudent list:%s\nlist:%s\n\nListClear() == %v instead of %v\n\n",
				listToStringStu5(link2), solution.ListToString(link.Head), link2.Head, link.Head)
		}
	}
}
