package student_test

import (
	"strconv"
	"testing"

	solution "./solutions"
	student "./student"
	"github.com/01-edu/z01"
)

func listToString4(n *solution.Nodelist) string {
	var res string
	it := n
	for it != nil {
		res += strconv.Itoa(it.Data) + "-> "
		it = it.Next
	}
	res += "<nil>"
	return res
}

func listToStringStu4(n *student.Nodelist) string {
	var res string
	it := n
	for it != nil {
		res += strconv.Itoa(it.Data) + "-> "
		it = it.Next
	}
	res += "<nil>"
	return res
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

func compare(t *testing.T, l *solution.Nodelist, l1 *student.Nodelist, f func(a, b int) bool) {
	cmp := solution.GetName(f)

	for l1 != nil && l != nil {
		if l1.Data != l.Data {
			t.Errorf("\nstudent list:%s\nlist:%s\nfunction cmp:%s\n\nSortListInsert() == %v instead of %v\n\n",
				listToStringStu4(l1), listToString4(l), cmp, l1.Data, l.Data)
			return
		}
		l1 = l1.Next
		l = l.Next
	}
}

func TestSortList(t *testing.T) {
	var linkSolutions *solution.Nodelist
	var linkStudent *student.Nodelist

	type nodeTest struct {
		data      []int
		functions []func(a, b int) bool
	}

	table := []nodeTest{}

	for i := 0; i < 4; i++ {
		val := nodeTest{
			data:      z01.MultRandIntBetween(1, 1234),
			functions: []func(a, b int) bool{ascending, descending},
		}
		table = append(table, val)
	}

	table = append(table,
		nodeTest{
			data:      []int{2, 5, 3, 1, 9, 6},
			functions: []func(a, b int) bool{ascending, descending},
		})

	for _, arg := range table {
		for i := 0; i < len(arg.data); i++ {
			linkStudent = insertNodeListStudent(linkStudent, arg.data[i])
			linkSolutions = insertNodeListSolution(linkSolutions, arg.data[i])
		}

		for i := 0; i < len(arg.functions); i++ {
			studentResult := student.SortList(linkStudent, arg.functions[i])
			solutionResult := solution.SortList(linkSolutions, arg.functions[i])

			compare(t, solutionResult, studentResult, arg.functions[i])
		}
		linkSolutions = &solution.Nodelist{}
		linkStudent = &student.Nodelist{}
	}
}
