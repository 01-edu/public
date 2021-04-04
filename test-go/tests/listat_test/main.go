package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
	"github.com/01-edu/public/test-go/tests/correct"
)

type (
	Node5  = student.NodeL
	NodeS5 = correct.NodeL
)

func nodePushBackList5(l1 *Node5, l2 *NodeS5, data interface{}) (*Node5, *NodeS5) {
	n1 := &Node5{Data: data}
	n2 := &NodeS5{Data: data}

	if l1 == nil {
		l1 = n1
	} else {
		iterator := l1
		for iterator.Next != nil {
			iterator = iterator.Next
		}
		iterator.Next = n1
	}
	if l2 == nil {
		l2 = n2
	} else {
		iterator1 := l2
		for iterator1.Next != nil {
			iterator1 = iterator1.Next
		}
		iterator1.Next = n2
	}
	return l1, l2
}

// compare functions that consist on inserting using the structure node
func comparFuncNode5(solutionList *NodeS5, l1 *Node5, l2 *NodeS5, arg int) {
	if l1 == nil && l2 == nil {
		return
	}
	if l1 != nil && l2 == nil {
		lib.Fatalf("\nListAt(%s, %d) == %v instead of %v\n\n",
			correct.ListToString(solutionList), arg, l1, l2)
	}
	if l1.Data != l2.Data {
		lib.Fatalf("\nListAt(%s, %d) == %v instead of %v\n\n",
			correct.ListToString(solutionList), arg, l1.Data, l2.Data)
	}
}

// finds an element of a correct.ListS using a given position
func main() {
	var link1 *Node5
	var link2 *NodeS5

	type nodeTest struct {
		data []interface{}
		pos  int
	}

	var table []nodeTest

	for i := 0; i < 4; i++ {
		table = append(table, nodeTest{
			data: correct.ConvertIntToInterface(lib.MultRandInt()),
			pos:  lib.RandIntBetween(1, 12),
		})
	}

	for i := 0; i < 4; i++ {
		table = append(table, nodeTest{
			data: correct.ConvertIntToStringface(lib.MultRandWords()),
			pos:  lib.RandIntBetween(1, 12),
		})
	}

	table = append(table, nodeTest{
		data: []interface{}{"I", 1, "something", 2},
		pos:  lib.RandIntBetween(1, 4),
	})

	for _, arg := range table {
		for i := 0; i < len(arg.data); i++ {
			link1, link2 = nodePushBackList5(link1, link2, arg.data[i])
		}

		result1 := student.ListAt(link1, arg.pos)
		result2 := correct.ListAt(link2, arg.pos)

		comparFuncNode5(link2, result1, result2, arg.pos)
		link1 = nil
		link2 = nil
	}
}
