package student_test

import (
	"strconv"
	"testing"

	"github.com/01-edu/z01"

	solution "./solutions"
	student "./student"
)

type Node11 = student.NodeL
type List11 = solution.List
type NodeS11 = solution.NodeL
type ListS11 = student.List

func listToStringStu(l *ListS11) string {
	var res string
	it := l.Head
	for it != nil {
		switch it.Data.(type) {
		case int:
			res += strconv.Itoa(it.Data.(int)) + "->"
		case string:
			res += it.Data.(string) + "->"
		}
		it = it.Next
	}
	res += "<nil>"
	return res
}

func listPushBackTest11(l *ListS11, l1 *List11, data interface{}) {
	n := &Node11{Data: data}
	n1 := &NodeS11{Data: data}

	if l.Head == nil {
		l.Head = n
	} else {
		iterator := l.Head
		for iterator.Next != nil {
			iterator = iterator.Next
		}
		iterator.Next = n
	}
	l.Tail = n

	if l1.Head == nil {
		l1.Head = n1
	} else {
		iterator1 := l1.Head
		for iterator1.Next != nil {
			iterator1 = iterator1.Next
		}
		iterator1.Next = n1
	}
	l1.Tail = n1
}

func comparFuncList11(l *List11, l1 *ListS11, t *testing.T) {
	for l.Head != nil || l1.Head != nil {
		if (l.Head == nil && l1.Head != nil) || (l.Head != nil && l1.Head == nil) {
			t.Errorf("\nstudent list:%s\nlist:%s\n\nListMerge() == %v instead of %v\n\n",
				listToStringStu(l1), solution.ListToString(l.Head), l1.Head, l.Head)
			return
		} else if l.Head.Data != l1.Head.Data {
			t.Errorf("\nstudent list:%s\nlist:%s\n\nListMerge() == %v instead of %v\n\n",
				listToStringStu(l1), solution.ListToString(l.Head), l1.Head.Data, l.Head.Data)
			return
		}
		l.Head = l.Head.Next
		l1.Head = l1.Head.Next
	}
}

//exercise 14
func TestListMerge(t *testing.T) {
	link := &List11{}
	linkTest := &List11{}
	link2 := &ListS11{}
	link2Test := &ListS11{}

	type nodeTest struct {
		data  []interface{}
		data2 []interface{}
	}
	table := []nodeTest{}
	// empty list
	table = append(table,
		nodeTest{
			data:  []interface{}{},
			data2: []interface{}{},
		})
	table = append(table,
		nodeTest{
			data:  solution.ConvertIntToInterface(z01.MultRandInt()),
			data2: []interface{}{},
		})
	// jut ints
	for i := 0; i < 3; i++ {
		val := nodeTest{
			data:  solution.ConvertIntToInterface(z01.MultRandInt()),
			data2: solution.ConvertIntToInterface(z01.MultRandInt()),
		}
		table = append(table, val)
	}
	// just strings
	for i := 0; i < 2; i++ {
		val := nodeTest{
			data:  solution.ConvertIntToStringface(z01.MultRandWords()),
			data2: solution.ConvertIntToStringface(z01.MultRandWords()),
		}
		table = append(table, val)
	}
	table = append(table,
		nodeTest{
			data:  []interface{}{},
			data2: []interface{}{"a", 1},
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.data); i++ {
			listPushBackTest11(link2, link, arg.data[i])
		}
		for i := 0; i < len(arg.data2); i++ {
			listPushBackTest11(link2Test, linkTest, arg.data2[i])
		}

		solution.ListMerge(link, linkTest)

		student.ListMerge(link2, link2Test)

		comparFuncList11(link, link2, t)

		link = &List11{}
		linkTest = &List11{}
		link2 = &ListS11{}
		link2Test = &ListS11{}
	}
}
