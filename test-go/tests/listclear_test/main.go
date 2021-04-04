package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/public/test-go/lib"
	"github.com/01-edu/public/test-go/tests/correct"
)

type (
	Node4  = student.NodeL
	List4  = correct.List
	NodeS4 = correct.NodeL
	ListS4 = student.List
)

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

func listPushBackTest4(l1 *ListS4, l2 *List4, data interface{}) {
	n := &Node4{Data: data}
	n1 := &NodeS4{Data: data}
	if l1.Head == nil {
		l1.Head = n
	} else {
		iterator := l1.Head
		for iterator.Next != nil {
			iterator = iterator.Next
		}
		iterator.Next = n
	}
	if l2.Head == nil {
		l2.Head = n1
	} else {
		iterator1 := l2.Head
		for iterator1.Next != nil {
			iterator1 = iterator1.Next
		}
		iterator1.Next = n1
	}
}

// simply cleans the linked correct.ListS
func main() {
	link1 := &List4{}
	link2 := &ListS4{}

	table := []correct.NodeTest{}

	table = correct.ElementsToTest(table)

	table = append(table,
		correct.NodeTest{
			Data: []interface{}{"I", 1, "something", 2},
		},
	)

	for _, arg := range table {
		for i := 0; i < len(arg.Data); i++ {
			listPushBackTest4(link2, link1, arg.Data[i])
		}
		correct.ListClear(link1)
		student.ListClear(link2)

		if link2.Head != nil {
			lib.Fatalf("\nstudent list:%s\nlist:%s\n\nListClear() == %v instead of %v\n\n",
				listToStringStu5(link2), correct.ListToString(link1.Head), link2.Head, link1.Head)
		}
	}
}
