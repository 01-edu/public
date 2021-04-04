package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
	"github.com/01-edu/public/test-go/solutions"
)

type (
	Node9  = student.NodeL
	List9  = solutions.List
	NodeS9 = solutions.NodeL
	ListS9 = student.List
)

func listPushBackTest9(l1 *ListS9, l2 *List9, data interface{}) {
	n := &Node9{Data: data}
	n1 := &NodeS9{Data: data}
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

func main() {
	link1 := &List9{}
	link2 := &ListS9{}

	table := []solutions.NodeTest{}
	table = solutions.ElementsToTest(table)
	table = append(table,
		solutions.NodeTest{
			Data: []interface{}{"hello", "hello1", "hello2", "hello3"},
		},
	)

	for _, arg := range table {
		for i := 0; i < len(arg.Data); i++ {
			listPushBackTest9(link2, link1, arg.Data[i])
		}
		if len(arg.Data) != 0 {
			aux1 := student.ListFind(link2, arg.Data[(len(arg.Data)-1)/2], student.CompStr)
			aux2 := solutions.ListFind(link1, arg.Data[(len(arg.Data)-1)/2], solutions.CompStr)

			if aux1 != nil || aux2 != nil {
				if *aux1 != *aux2 {
					lib.Fatalf("ListFind(ref: %s) == %s instead of %s\n", arg.Data[(len(arg.Data)-1)/2], *aux1, *aux2)
				}
			}
		}
		link1 = &List9{}
		link2 = &ListS9{}
	}

	for i := 0; i < len(table[0].Data); i++ {
		listPushBackTest9(link2, link1, table[0].Data[i])
	}

	aux1 := student.ListFind(link2, "lksdf", student.CompStr)
	aux2 := solutions.ListFind(link1, "lksdf", solutions.CompStr)
	if aux1 != nil && aux2 != nil {
		if *aux1 != *aux2 {
			lib.Fatalf("ListFind(ref: lksdf) == %s instead of %s\n", *aux1, *aux2)
		}
	}
}
