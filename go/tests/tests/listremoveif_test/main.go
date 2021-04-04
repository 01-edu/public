package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/public/go/tests/func/correct"

	"github.com/01-edu/public/go/tests/lib"
)

type (
	Node10  = student.NodeL
	List10  = correct.List
	NodeS10 = correct.NodeL
	ListS10 = student.List
)

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

func listPushBackTest10(l *ListS10, l1 *List10, data interface{}) {
	n := &Node10{Data: data}
	n1 := &NodeS10{Data: data}
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

func comparFuncList10(l *List10, l1 *ListS10, data interface{}) {
	for l.Head != nil || l1.Head != nil {
		if (l.Head == nil && l1.Head != nil) || (l.Head != nil && l1.Head == nil) {
			lib.Fatalf("\ndata used: %v\nstudent list:%s\nlist:%s\n\nListRemoveIf() == %v instead of %v\n\n",
				data, listToStringStu12(l1), correct.ListToString(l.Head), l1.Head, l.Head)
		}
		if l.Head.Data != l1.Head.Data {
			lib.Fatalf("\ndata used: %v\nstudent list:%s\nlist:%s\n\nListRemoveIf() == %v instead of %v\n\n",
				data, listToStringStu12(l1), correct.ListToString(l.Head), l1.Head.Data, l.Head.Data)
		}
		l.Head = l.Head.Next
		l1.Head = l1.Head.Next
	}
}

// removes all the elements that are equal to a value
func main() {
	link1 := &List10{}
	link2 := &ListS10{}
	var index int
	table := []correct.NodeTest{}

	table = correct.ElementsToTest(table)

	table = append(table,
		correct.NodeTest{
			Data: []interface{}{"hello", "hello1", "hello2", "hello3"},
		},
	)

	for _, arg := range table {
		for i := 0; i < len(arg.Data); i++ {
			listPushBackTest10(link2, link1, arg.Data[i])
		}
		aux := len(arg.Data) - 1

		index = lib.RandIntBetween(0, aux)
		if link1.Head != nil && link2.Head != nil {
			cho := arg.Data[index]
			student.ListRemoveIf(link2, cho)
			correct.ListRemoveIf(link1, cho)
			comparFuncList10(link1, link2, cho)
		} else {
			student.ListRemoveIf(link2, 1)
			correct.ListRemoveIf(link1, 1)
			comparFuncList10(link1, link2, 1)
		}

		link1 = &List10{}
		link2 = &ListS10{}
	}
}
