package student_test

import (
	"strconv"
	"testing"

	solution "./solutions"
	student "./student"
)

type Node6 = student.NodeL
type List6 = solution.List
type NodeS6 = solution.NodeL
type ListS6 = student.List

func listToStringStu13(l *ListS6) string {
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

func comparFuncList6(l *List6, l1 *ListS6, t *testing.T) {
	for l.Head != nil || l1.Head != nil {
		if (l.Head == nil && l1.Head != nil) || (l.Head != nil && l1.Head == nil) {
			t.Errorf("\nstudent list:%s\nlist:%s\n\nListReverse() == %v instead of %v\n\n",
				listToStringStu13(l1), solution.ListToString(l.Head), l1.Head, l.Head)
			return
		} else if l.Head.Data != l1.Head.Data {
			t.Errorf("\nstudent list:%s\nlist:%s\n\nListReverse() == %v instead of %v\n\n",
				listToStringStu13(l1), solution.ListToString(l.Head), l1.Head.Data, l.Head.Data)
			return
		}
		l.Head = l.Head.Next
		l1.Head = l1.Head.Next
	}
}
func listPushBackTest6(l *ListS6, l1 *List6, data interface{}) {
	n := &Node6{Data: data}
	n1 := &NodeS6{Data: data}
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

// exercise 8
func TestListReverse(t *testing.T) {
	link := &List6{}
	link2 := &ListS6{}
	table := []solution.NodeTest{}
	table = solution.ElementsToTest(table)
	table = append(table,
		solution.NodeTest{
			Data: []interface{}{"I", 1, "something", 2},
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.Data); i++ {
			listPushBackTest6(link2, link, arg.Data[i])
		}
		student.ListReverse(link2)
		solution.ListReverse(link)
		comparFuncList6(link, link2, t)
		link = &List6{}
		link2 = &ListS6{}
	}
}
