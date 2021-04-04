package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
	"github.com/01-edu/public/test-go/solutions"
)

type (
	Node2  = student.NodeL
	List2  = solutions.List
	NodeS2 = solutions.NodeL
	ListS2 = student.List
)

func listPushBackTest2(l *ListS2, l1 *List2, data interface{}) {
	n := &Node2{Data: data}
	n1 := &NodeS2{Data: data}
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
	link := &List2{}
	link2 := &ListS2{}
	table := []solutions.NodeTest{}
	table = solutions.ElementsToTest(table)
	table = append(table,
		solutions.NodeTest{
			Data: []interface{}{"Hello", "man", "how are you"},
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.Data); i++ {
			listPushBackTest2(link2, link, arg.Data[i])
		}
		aux := solutions.ListSize(link)
		aux2 := student.ListSize(link2)
		if aux != aux2 {
			lib.Fatalf("ListSize(%v) == %d instead of %d\n", solutions.ListToString(link.Head), aux2, aux)
		}
		link = &List2{}
		link2 = &ListS2{}
	}
}
