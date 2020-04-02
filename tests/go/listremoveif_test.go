package student_test

import (
	"strconv"
	"testing"

	solution "./solutions"
	student "./student"
	"github.com/01-edu/z01"
)

type Node10 = student.NodeL
type List10 = solution.List
type NodeS10 = solution.NodeL
type ListS10 = student.List

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

func comparFuncList10(l *List10, l1 *ListS10, t *testing.T, data interface{}) {
	for l.Head != nil || l1.Head != nil {
		if (l.Head == nil && l1.Head != nil) || (l.Head != nil && l1.Head == nil) {
			t.Errorf("\ndata used: %v\nstudent list:%s\nlist:%s\n\nListRemoveIf() == %v instead of %v\n\n",
				data, listToStringStu12(l1), solution.ListToString(l.Head), l1.Head, l.Head)
			return
		} else if l.Head.Data != l1.Head.Data {
			t.Errorf("\ndata used: %v\nstudent list:%s\nlist:%s\n\nListRemoveIf() == %v instead of %v\n\n",
				data, listToStringStu12(l1), solution.ListToString(l.Head), l1.Head.Data, l.Head.Data)
			return
		}
		l.Head = l.Head.Next
		l1.Head = l1.Head.Next
	}
}

//exercise 13
//removes all the elements that are equal to a value
func TestListRemoveIf(t *testing.T) {
	link := &List10{}
	link2 := &ListS10{}
	var index int
	table := []solution.NodeTest{}

	table = solution.ElementsToTest(table)

	table = append(table,
		solution.NodeTest{
			Data: []interface{}{"hello", "hello1", "hello2", "hello3"},
		},
	)

	for _, arg := range table {
		for i := 0; i < len(arg.Data); i++ {
			listPushBackTest10(link2, link, arg.Data[i])
		}
		aux := len(arg.Data) - 1

		index = z01.RandIntBetween(0, aux)
		if link.Head != nil && link2.Head != nil {
			cho := arg.Data[index]
			student.ListRemoveIf(link2, cho)
			solution.ListRemoveIf(link, cho)
			comparFuncList10(link, link2, t, cho)
		} else {
			student.ListRemoveIf(link2, 1)
			solution.ListRemoveIf(link, 1)
			comparFuncList10(link, link2, t, 1)
		}

		link = &List10{}
		link2 = &ListS10{}
	}
}
