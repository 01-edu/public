package student_test

import (
	"strconv"
	"testing"

	solution "./solutions"
	student "./student"
	"github.com/01-edu/z01"
)

type NodeI12 = student.NodeI
type NodeIS12 = solution.NodeI

func printListStudent(n *NodeI12) string {
	var res string
	it := n
	for it != nil {
		res += strconv.Itoa(it.Data) + "-> "
		it = it.Next
	}
	res += "<nil>"
	return res
}

func nodePushBackListInt12(l *NodeI12, l1 *NodeIS12, data int) {
	n := &NodeI12{Data: data}
	n1 := &NodeIS12{Data: data}

	if l == nil {
		l = n
	} else {
		iterator := l
		for iterator.Next != nil {
			iterator = iterator.Next
		}
		iterator.Next = n
	}

	if l1 == nil {
		l1 = n1
	} else {
		iterator1 := l1
		for iterator1.Next != nil {
			iterator1 = iterator1.Next
		}
		iterator1.Next = n1
	}
}

func comparFuncNodeInt12(l *NodeI12, l1 *NodeIS12, t *testing.T) {
	for l != nil || l1 != nil {
		if (l == nil && l1 != nil) || (l != nil && l1 == nil) {
			t.Errorf("\nstudent list:%s\nlist:%s\n\nListSort() == %v instead of %v\n\n",
				printListStudent(l), solution.PrintList(l1), l, l1)
			return
		} else if l.Data != l1.Data {
			t.Errorf("\nstudent list:%s\nlist:%s\n\nListSort() == %v instead of %v\n\n",
				printListStudent(l), solution.PrintList(l1), l.Data, l1.Data)
			return
		}
		l = l.Next
		l1 = l1.Next
	}
}

//exercise 15
func TestListSort(t *testing.T) {
	var link *NodeI12
	var link2 *NodeIS12

	type nodeTest struct {
		data []int
	}
	table := []nodeTest{}

	table = append(table,
		nodeTest{
			data: []int{},
		})
	//just numbers/ints
	for i := 0; i < 2; i++ {
		val := nodeTest{
			data: z01.MultRandInt(),
		}
		table = append(table, val)
	}

	table = append(table,
		nodeTest{
			data: []int{5, 4, 3, 2, 1},
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.data); i++ {
			nodePushBackListInt12(link, link2, arg.data[i])
		}
		aux := solution.ListSort(link2)
		aux2 := student.ListSort(link)

		comparFuncNodeInt12(aux2, aux, t)

		link = &NodeI12{}
		link2 = &NodeIS12{}
	}
}
