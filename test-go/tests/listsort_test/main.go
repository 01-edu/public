package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/public/test-go/lib"
	"github.com/01-edu/public/test-go/tests/correct"
)

type (
	NodeI12  = student.NodeI
	NodeIS12 = correct.NodeI
)

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

func nodePushBackListInt12(l1 *NodeI12, l2 *NodeIS12, data int) {
	n1 := &NodeI12{Data: data}
	n2 := &NodeIS12{Data: data}

	if l1 == nil {
		l1 = n1
	} else {
		iterator1 := l1
		for iterator1.Next != nil {
			iterator1 = iterator1.Next
		}
		iterator1.Next = n1
	}

	if l2 == nil {
		l2 = n2
	} else {
		iterator2 := l2
		for iterator2.Next != nil {
			iterator2 = iterator2.Next
		}
		iterator2.Next = n2
	}
}

func comparFuncNodeInt12(l1 *NodeI12, l2 *NodeIS12) {
	for l1 != nil || l2 != nil {
		if (l1 == nil && l2 != nil) || (l1 != nil && l2 == nil) {
			lib.Fatalf("\nstudent list:%s\nlist:%s\n\nListSort() == %v instead of %v\n\n",
				printListStudent(l1), correct.PrintList(l2), l1, l2)
		}
		if l1.Data != l2.Data {
			lib.Fatalf("\nstudent list:%s\nlist:%s\n\nListSort() == %v instead of %v\n\n",
				printListStudent(l1), correct.PrintList(l2), l1.Data, l2.Data)
		}
		l1 = l1.Next
		l2 = l2.Next
	}
}

func main() {
	var link1 *NodeI12
	var link2 *NodeIS12

	type nodeTest struct {
		data []int
	}
	table := []nodeTest{}
	table = append(table, nodeTest{[]int{}})

	// just numbers/ints
	for i := 0; i < 2; i++ {
		table = append(table, nodeTest{lib.MultRandInt()})
	}
	table = append(table, nodeTest{[]int{5, 4, 3, 2, 1}})

	for _, arg := range table {
		for i := 0; i < len(arg.data); i++ {
			nodePushBackListInt12(link1, link2, arg.data[i])
		}
		aux1 := correct.ListSort(link2)
		aux2 := student.ListSort(link1)

		comparFuncNodeInt12(aux2, aux1)

		link1 = &NodeI12{}
		link2 = &NodeIS12{}
	}
}
