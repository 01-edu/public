package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/public/go/tests/func/correct"

	"github.com/01-edu/public/go/tests/lib"
)

type (
	NodeI14  = student.NodeI
	NodeIS14 = correct.NodeI
)

func listToStringStu3(n *NodeI14) string {
	var res string
	it := n
	for it != nil {
		res += strconv.Itoa(it.Data) + "-> "
		it = it.Next
	}
	res += "<nil>"
	return res
}

func nodepushback1(l *NodeI14, data int) *NodeI14 {
	n := &NodeI14{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func nodepushback2(l *NodeIS14, data int) *NodeIS14 {
	n := &NodeIS14{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func comparFuncNodeInt14(l1 *NodeI14, l2 *NodeIS14, data []int) {
	for l1 != nil || l2 != nil {
		if (l1 == nil && l2 != nil) || (l1 != nil && l2 == nil) {
			lib.Fatalf("\ndata used to insert: %d\nstudent list:%s\nlist:%s\n\nSortListInsert() == %v instead of %v\n\n",
				data, listToStringStu3(l1), correct.PrintList(l2), l1, l2)
		}
		if l1.Data != l2.Data {
			lib.Fatalf("\ndata used to insert: %d\nstudent list:%s\nlist:%s\n\nSortListInsert() == %v instead of %v\n\n",
				data, listToStringStu3(l1), correct.PrintList(l2), l1.Data, l2.Data)
		}
		l1 = l1.Next
		l2 = l2.Next
	}
}

func move(l *NodeI14) *NodeI14 {
	p := l
	n := l.Next
	ret := n

	for n != nil && l.Data > n.Data {
		p = n
		n = n.Next
	}
	p.Next = l
	l.Next = n
	return ret
}

func sortStudentsList(l *NodeI14) *NodeI14 {
	head := l
	if head == nil {
		return nil
	}
	head.Next = sortStudentsList(head.Next)

	if head.Next != nil && head.Data > head.Next.Data {
		head = move(head)
	}
	return head
}

func main() {
	var link1 *NodeI14
	var link2 *NodeIS14

	type nodeTest struct {
		data     []int
		data_ref []int
	}
	table := []nodeTest{{
		data:     []int{},
		data_ref: []int{},
	}}

	for i := 0; i < 2; i++ {
		table = append(table, nodeTest{
			data:     lib.MultRandInt(),
			data_ref: lib.MultRandInt(),
		})
	}
	table = append(table,
		nodeTest{
			data:     []int{5, 4, 3, 2, 1},
			data_ref: lib.MultRandInt(),
		},
	)
	for _, arg := range table {
		for _, item := range arg.data {
			link2 = nodepushback2(link2, item)
			link1 = nodepushback1(link1, item)
		}

		link2 = correct.ListSort(link2)
		link1 = sortStudentsList(link1)

		for _, item := range arg.data_ref {
			link2 = correct.SortListInsert(link2, item)
			link1 = student.SortListInsert(link1, item)
		}

		comparFuncNodeInt14(link1, link2, arg.data_ref)
		link1 = &NodeI14{}
		link2 = &NodeIS14{}
	}
}
