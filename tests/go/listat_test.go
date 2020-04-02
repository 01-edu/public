package student_test

import (
	"testing"

	solution "./solutions"
	student "./student"
	"github.com/01-edu/z01"
)

type Node5 = student.NodeL
type NodeS5 = solution.NodeL

func nodePushBackList5(l *Node5, l1 *NodeS5, data interface{}) (*Node5, *NodeS5) {
	n := &Node5{Data: data}
	n1 := &NodeS5{Data: data}
	it := l
	itS := l1

	if it == nil {
		it = n
	} else {
		iterator := it
		for iterator.Next != nil {
			iterator = iterator.Next
		}
		iterator.Next = n
	}
	if itS == nil {
		itS = n1
	} else {
		iterator1 := itS
		for iterator1.Next != nil {
			iterator1 = iterator1.Next
		}
		iterator1.Next = n1
	}
	return it, itS
}

// compare functions that consist on inserting using the structure node
func comparFuncNode5(solutionList *NodeS5, l *Node5, l1 *NodeS5, t *testing.T, arg int) {
	if l == nil && l1 == nil {
	} else if l != nil && l1 == nil {
		t.Errorf("\nListAt(%s, %d) == %v instead of %v\n\n",
			solution.ListToString(solutionList), arg, l, l1)
	} else if l.Data != l1.Data {
		t.Errorf("\nListAt(%s, %d) == %v instead of %v\n\n",
			solution.ListToString(solutionList), arg, l.Data, l1.Data)
	}
}

// exercise 7
// finds an element of a solution.ListS using a given position
func TestListAt(t *testing.T) {
	var link *Node5
	var link2 *NodeS5

	type nodeTest struct {
		data []interface{}
		pos  int
	}

	var table []nodeTest

	for i := 0; i < 4; i++ {
		val := nodeTest{
			data: solution.ConvertIntToInterface(z01.MultRandInt()),
			pos:  z01.RandIntBetween(1, 12),
		}
		table = append(table, val)
	}

	for i := 0; i < 4; i++ {
		val := nodeTest{
			data: solution.ConvertIntToStringface(z01.MultRandWords()),
			pos:  z01.RandIntBetween(1, 12),
		}
		table = append(table, val)
	}

	table = append(table,
		nodeTest{
			data: []interface{}{"I", 1, "something", 2},
			pos:  z01.RandIntBetween(1, 4),
		},
	)

	for _, arg := range table {
		for i := 0; i < len(arg.data); i++ {
			link, link2 = nodePushBackList5(link, link2, arg.data[i])
		}

		result := student.ListAt(link, arg.pos)
		result2 := solution.ListAt(link2, arg.pos)

		comparFuncNode5(link2, result, result2, t, arg.pos)
		link = nil
		link2 = nil
	}
}
