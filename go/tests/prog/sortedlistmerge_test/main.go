package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/public/go/tests/func/correct"

	"github.com/01-edu/public/go/tests/lib"
)

type (
	NodeI13  = student.NodeI
	NodeIS13 = correct.NodeI
)

func printListStudent1(n *NodeI13) string {
	var res string
	it := n
	for it != nil {
		res += strconv.Itoa(it.Data) + "-> "
		it = it.Next
	}

	res += "<nil>"
	return res
}

func nodePushBackListInt13(l *NodeI13, l1 *NodeIS13, data int) {
	n := &NodeI13{Data: data}
	n1 := &NodeIS13{Data: data}

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

func main() {
	var link1 *NodeI13
	var link2 *NodeI13
	var linkTest1 *NodeIS13
	var linkTest2 *NodeIS13
	type nodeTest struct {
		data1 []int
		data2 []int
	}

	table := []nodeTest{{
		data1: []int{},
	}}

	for i := 0; i < 3; i++ {
		val := nodeTest{
			data1: lib.MultRandInt(),
			data2: lib.MultRandInt(),
		}
		table = append(table, val)
	}
	table = append(table,
		nodeTest{
			data1: []int{3, 5, 7},
			data2: []int{1, -2, 4, 6},
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.data1); i++ {
			nodePushBackListInt13(link1, linkTest1, arg.data1[i])
		}
		for i := 0; i < len(arg.data2); i++ {
			nodePushBackListInt13(link2, linkTest2, arg.data2[i])
		}

		link1 = student.ListSort(link1)
		link2 = student.ListSort(link2)
		linkTest1 = correct.ListSort(linkTest1)
		linkTest2 = correct.ListSort(linkTest2)

		aux1 := student.SortedListMerge(link1, link2)
		aux2 := correct.SortedListMerge(linkTest1, linkTest2)

		if aux1 == nil && aux2 == nil {
		} else if aux1 != nil && aux2 == nil {
			lib.Fatalf("\nstudent merged lists:%s\nmerged lists:%s\n\nSortListMerge() == %v instead of %v\n\n",
				printListStudent1(aux1), correct.PrintList(aux2), aux1, aux2)
		} else if aux1.Data != aux2.Data {
			lib.Fatalf("\nstudent merged lists:%s\nmerged lists:%s\n\nSortListMerge() == %v instead of %v\n\n",
				printListStudent1(aux1), correct.PrintList(aux2), aux1, aux2)
		}

		link1 = &NodeI13{}
		link2 = &NodeI13{}
		linkTest1 = &NodeIS13{}
		linkTest2 = &NodeIS13{}
	}
}
