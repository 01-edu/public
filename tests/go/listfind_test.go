package student_test

import (
	"testing"

	solution "./solutions"
	student "./student"
)

type Node9 = student.NodeL
type List9 = solution.List
type NodeS9 = solution.NodeL
type ListS9 = student.List

func listPushBackTest9(l *ListS9, l1 *List9, data interface{}) {
	n := &Node9{Data: data}
	n1 := &NodeS9{Data: data}
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

// exercise 11
func TestListFind(t *testing.T) {
	link := &List9{}
	link2 := &ListS9{}

	table := []solution.NodeTest{}
	table = solution.ElementsToTest(table)
	table = append(table,
		solution.NodeTest{
			Data: []interface{}{"hello", "hello1", "hello2", "hello3"},
		},
	)

	for _, arg := range table {
		for i := 0; i < len(arg.Data); i++ {
			listPushBackTest9(link2, link, arg.Data[i])
		}
		if len(arg.Data) != 0 {
			aux := student.ListFind(link2, arg.Data[(len(arg.Data)-1)/2], student.CompStr)
			aux1 := solution.ListFind(link, arg.Data[(len(arg.Data)-1)/2], solution.CompStr)

			if aux != nil || aux1 != nil {
				if *aux != *aux1 {
					t.Errorf("ListFind(ref: %s) == %s instead of %s\n", arg.Data[(len(arg.Data)-1)/2], *aux, *aux1)
				}
			}
		}
		link = &List9{}
		link2 = &ListS9{}
	}

	for i := 0; i < len(table[0].Data); i++ {
		listPushBackTest9(link2, link, table[0].Data[i])
	}

	aux := student.ListFind(link2, "lksdf", student.CompStr)
	aux1 := solution.ListFind(link, "lksdf", solution.CompStr)
	if aux != nil && aux1 != nil {
		if *aux != *aux1 {
			t.Errorf("ListFind(ref: lksdf) == %s instead of %s\n", *aux, *aux1)
		}
	}

}
