package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/public/go/tests/func/correct"

	"github.com/01-edu/public/go/tests/lib"
)

type (
	Node6  = student.NodeL
	List6  = correct.List
	NodeS6 = correct.NodeL
	ListS6 = student.List
)

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

func comparFuncList6(l *List6, l1 *ListS6) {
	for l.Head != nil || l1.Head != nil {
		if (l.Head == nil && l1.Head != nil) || (l.Head != nil && l1.Head == nil) {
			lib.Fatalf("\nstudent list:%s\nlist:%s\n\nListReverse() == %v instead of %v\n\n",
				listToStringStu13(l1), correct.ListToString(l.Head), l1.Head, l.Head)
		}
		if l.Head.Data != l1.Head.Data {
			lib.Fatalf("\nstudent list:%s\nlist:%s\n\nListReverse() == %v instead of %v\n\n",
				listToStringStu13(l1), correct.ListToString(l.Head), l1.Head.Data, l.Head.Data)
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

func main() {
	link1 := &List6{}
	link2 := &ListS6{}
	table := []correct.NodeTest{{
		Data: []interface{}{"I", 1, "something", 2},
	}}
	table = correct.ElementsToTest(table)
	for _, arg := range table {
		for _, item := range arg.Data {
			listPushBackTest6(link2, link1, item)
		}
		student.ListReverse(link2)
		correct.ListReverse(link1)
		comparFuncList6(link1, link2)
		link1 = &List6{}
		link2 = &ListS6{}
	}
}
