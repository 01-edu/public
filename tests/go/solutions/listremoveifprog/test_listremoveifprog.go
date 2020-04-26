package main

import (
	"strconv"

	"github.com/01-edu/z01"

	solution "../../solutions"
)

type Node10 = NodeL
type List10 = solution.List
type NodeS10 = solution.NodeL
type ListS10 = List

func listToStringStu12(l *ListS10) string {
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

func listPushBackTest10(l1 *ListS10, l2 *List10, data interface{}) {
	n := &Node10{Data: data}
	n1 := &NodeS10{Data: data}
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

func comparFuncList10(l1 *List10, l2 *ListS10, data interface{}) {
	for l1.Head != nil || l2.Head != nil {
		if (l1.Head == nil && l2.Head != nil) || (l1.Head != nil && l2.Head == nil) {
			z01.Fatalf("\ndata used: %v\nstudent list:%s\nlist:%s\n\nListRemoveIf() == %v instead of %v\n\n",
				data, listToStringStu12(l2), solution.ListToString(l1.Head), l2.Head, l1.Head)
		}
		if l1.Head.Data != l2.Head.Data {
			z01.Fatalf("\ndata used: %v\nstudent list:%s\nlist:%s\n\nListRemoveIf() == %v instead of %v\n\n",
				data, listToStringStu12(l2), solution.ListToString(l1.Head), l2.Head.Data, l1.Head.Data)
		}
		l1.Head = l1.Head.Next
		l2.Head = l2.Head.Next
	}
}

// removes all the elements that are equal to a value
func main() {
	link1 := &List10{}
	link2 := &ListS10{}
	table := []solution.NodeTest{}

	table = solution.ElementsToTest(table)

	table = append(table,
		solution.NodeTest{
			Data: []interface{}{"hello", "hello1", "hello2", "hello3"},
		},
	)

	for _, arg := range table {
		for i := 0; i < len(arg.Data); i++ {
			listPushBackTest10(link2, link1, arg.Data[i])
		}
		aux := len(arg.Data) - 1

		index := z01.RandIntBetween(0, aux)
		if link1.Head != nil && link2.Head != nil {
			cho := arg.Data[index]
			ListRemoveIf(link2, cho)
			solution.ListRemoveIf(link1, cho)
			comparFuncList10(link1, link2, cho)
		} else {
			ListRemoveIf(link2, 1)
			solution.ListRemoveIf(link1, 1)
			comparFuncList10(link1, link2, 1)
		}

		link1 = &List10{}
		link2 = &ListS10{}
	}
}
