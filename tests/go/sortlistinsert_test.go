package student_test

import (
	"strconv"
	"testing"

	solution "./solutions"
	solutions "./solutions"
	student "./student"
	"github.com/01-edu/z01"
)

type NodeI14 = student.NodeI
type NodeIS14 = solution.NodeI

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

func comparFuncNodeInt14(l *NodeI14, l1 *NodeIS14, t *testing.T, data []int) {
	for l != nil || l1 != nil {
		if (l == nil && l1 != nil) || (l != nil && l1 == nil) {
			t.Errorf("\ndata used to insert: %d\nstudent list:%s\nlist:%s\n\nSortListInsert() == %v instead of %v\n\n",
				data, listToStringStu3(l), solution.PrintList(l1), l, l1)
			return
		} else if l.Data != l1.Data {
			t.Errorf("\ndata used to insert: %d\nstudent list:%s\nlist:%s\n\nSortListInsert() == %v instead of %v\n\n",
				data, listToStringStu3(l), solution.PrintList(l1), l.Data, l1.Data)
			return
		}
		l = l.Next
		l1 = l1.Next
	}
}

// exercise 16
func TestSortListInsert(t *testing.T) {
	var link *NodeI14
	var link2 *NodeIS14

	type nodeTest struct {
		data     []int
		data_ref []int
	}
	table := []nodeTest{}

	table = append(table,
		nodeTest{
			data:     []int{},
			data_ref: []int{},
		})
	for i := 0; i < 2; i++ {
		val := nodeTest{
			data:     z01.MultRandInt(),
			data_ref: z01.MultRandInt(),
		}
		table = append(table, val)
	}
	table = append(table,
		nodeTest{
			data:     []int{5, 4, 3, 2, 1},
			data_ref: z01.MultRandInt(),
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.data); i++ {
			link2 = nodepushback2(link2, arg.data[i])
			link = nodepushback1(link, arg.data[i])
		}

		link2 = solutions.ListSort(link2)
		link = sortStudentsList(link)

		for i := 0; i < len(arg.data_ref); i++ {
			link2 = solution.SortListInsert(link2, arg.data_ref[i])
			link = student.SortListInsert(link, arg.data_ref[i])
		}

		comparFuncNodeInt14(link, link2, t, arg.data_ref)
		link = &NodeI14{}
		link2 = &NodeIS14{}
	}
}

func sortStudentsList(l *NodeI14) *NodeI14 {

	Head := l
	if Head == nil {
		return nil
	}
	Head.Next = sortStudentsList(Head.Next)

	if Head.Next != nil && Head.Data > Head.Next.Data {
		Head = move(Head)
	}
	return Head
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
