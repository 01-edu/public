package student_test

import (
	"strconv"
	"testing"

	solution "./solutions"
	student "./student"
	"github.com/01-edu/z01"
)

type NodeI13 = student.NodeI
type NodeIS13 = solution.NodeI

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

func TestSortedListMerge(t *testing.T) {
	var link *NodeI13
	var link2 *NodeI13
	var linkTest *NodeIS13
	var linkTest2 *NodeIS13
	type nodeTest struct {
		data  []int
		data2 []int
	}

	table := []nodeTest{}

	table = append(table,
		nodeTest{
			data: []int{},
		})
	for i := 0; i < 3; i++ {
		val := nodeTest{
			data:  z01.MultRandInt(),
			data2: z01.MultRandInt(),
		}
		table = append(table, val)

	}
	table = append(table,
		nodeTest{
			data:  []int{3, 5, 7},
			data2: []int{1, -2, 4, 6},
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.data); i++ {
			nodePushBackListInt13(link, linkTest, arg.data[i])
		}
		for i := 0; i < len(arg.data2); i++ {
			nodePushBackListInt13(link2, linkTest2, arg.data2[i])
		}

		link = student.ListSort(link)
		link2 = student.ListSort(link2)
		linkTest = solution.ListSort(linkTest)
		linkTest2 = solution.ListSort(linkTest2)

		aux := student.SortedListMerge(link, link2)
		aux2 := solution.SortedListMerge(linkTest, linkTest2)

		if aux == nil && aux2 == nil {
		} else if aux != nil && aux2 == nil {
			t.Errorf("\nstudent merged lists:%s\nmerged lists:%s\n\nSortListMerge() == %v instead of %v\n\n",
				printListStudent1(aux), solution.PrintList(aux2), aux, aux2)
		} else if aux.Data != aux2.Data {
			t.Errorf("\nstudent merged lists:%s\nmerged lists:%s\n\nSortListMerge() == %v instead of %v\n\n",
				printListStudent1(aux), solution.PrintList(aux2), aux, aux2)
		}

		link = &NodeI13{}
		link2 = &NodeI13{}
		linkTest = &NodeIS13{}
		linkTest2 = &NodeIS13{}

	}
}
