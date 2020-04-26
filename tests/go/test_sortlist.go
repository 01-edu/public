package main

import (
	"strconv"

	"github.com/01-edu/z01"

	solution "./solutions"
	student "./student"
)

func listToString4(n *solution.Nodelist) (res string) {
	it := n
	for it != nil {
		res += strconv.Itoa(it.Data) + "-> "
		it = it.Next
	}
	res += "<nil>"
	return
}

func listToStringStu4(n *student.Nodelist) (res string) {
	it := n
	for it != nil {
		res += strconv.Itoa(it.Data) + "-> "
		it = it.Next
	}
	res += "<nil>"
	return
}

func ascending(a, b int) bool {
	return a <= b
}

func descending(a, b int) bool {
	return a >= b
}

func insertNodeListSolution(l *solution.Nodelist, data int) *solution.Nodelist {
	n := &solution.Nodelist{Data: data}
	it := l
	if it == nil {
		it = n
	} else {
		iterator := it
		for iterator.Next != nil {
			iterator = iterator.Next
		}
		iterator.Next = n
	}
	return it
}

func insertNodeListStudent(l1 *student.Nodelist, data int) *student.Nodelist {
	n1 := &student.Nodelist{Data: data}
	it := l1
	if it == nil {
		it = n1
	} else {
		iterator := it
		for iterator.Next != nil {
			iterator = iterator.Next
		}
		iterator.Next = n1
	}
	return it
}

func compare(l2 *solution.Nodelist, l1 *student.Nodelist, f func(a, b int) bool) {
	cmp := solution.GetName(f)

	for l1 != nil && l2 != nil {
		if l1.Data != l2.Data {
			z01.Fatalf("\nstudent list:%s\nlist:%s\nfunction cmp:%s\n\nSortListInsert() == %v instead of %v\n\n",
				listToStringStu4(l1), listToString4(l2), cmp, l1.Data, l2.Data)
			return
		}
		l1 = l1.Next
		l2 = l2.Next
	}
}

func main() {
	var linkSolutions *solution.Nodelist
	var linkStudent *student.Nodelist

	type nodeTest struct {
		data      []int
		functions []func(a, b int) bool
	}

	table := []nodeTest{}

	for i := 0; i < 4; i++ {
		table = append(table, nodeTest{
			data:      z01.MultRandIntBetween(1, 1234),
			functions: []func(a, b int) bool{ascending, descending},
		})
	}

	table = append(table,
		nodeTest{
			data:      []int{2, 5, 3, 1, 9, 6},
			functions: []func(a, b int) bool{ascending, descending},
		})

	for _, arg := range table {
		for _, item := range arg.data {
			linkStudent = insertNodeListStudent(linkStudent, item)
			linkSolutions = insertNodeListSolution(linkSolutions, item)
		}

		for _, fn := range arg.functions {
			studentResult := student.SortList(linkStudent, fn)
			solutionResult := solution.SortList(linkSolutions, fn)
			compare(solutionResult, studentResult, fn)
		}
		linkSolutions = &solution.Nodelist{}
		linkStudent = &student.Nodelist{}
	}
}
