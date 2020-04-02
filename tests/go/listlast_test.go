package student_test

import (
	"strconv"
	"testing"

	solution "./solutions"
	student "./student"
)

type Node3 = student.NodeL
type List3 = solution.List
type NodeS3 = solution.NodeL
type ListS3 = student.List

func listToStringStu9(l *ListS3) string {
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

//inserts node on two lists
func listPushBackTest3(l *ListS3, l1 *List3, data interface{}) {
	n := &Node3{Data: data}
	n1 := &NodeS3{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
	if l1.Head == nil {
		l1.Head = n1
		l1.Tail = n1
	} else {
		l1.Tail.Next = n1
		l1.Tail = n1
	}
}

//exercise 5
//last element of the solution.ListS
func TestListLast(t *testing.T) {
	link := &List3{}
	link2 := &ListS3{}
	table := []solution.NodeTest{}

	table = solution.ElementsToTest(table)
	table = append(table,
		solution.NodeTest{
			Data: []interface{}{3, 2, 1},
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.Data); i++ {
			listPushBackTest3(link2, link, arg.Data[i])
		}
		aux := solution.ListLast(link)
		aux1 := student.ListLast(link2)

		if aux != aux1 {
			t.Errorf("\nlist:%s\n\nListLast() == %v instead of %v\n\n",
				listToStringStu9(link2), aux1, aux)
		}
		link = &List3{}
		link2 = &ListS3{}
	}
}
