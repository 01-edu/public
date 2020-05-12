package correct

type Nodelist struct {
	Data int
	Next *Nodelist
}

func SortList(l *Nodelist, cmp func(a, b int) bool) *Nodelist {
	head := l
	if head == nil {
		return nil
	}
	head.Next = SortList(head.Next, cmp)

	if head.Next != nil && cmp(head.Data, head.Next.Data) {
		head = moveValue(head, cmp)
	}
	return head
}

func moveValue(l *Nodelist, cmp func(a, b int) bool) *Nodelist {
	p := l
	n := l.Next
	ret := n

	for n != nil && cmp(l.Data, n.Data) {
		p = n
		n = n.Next
	}
	p.Next = l
	l.Next = n
	return ret
}


func listToString4(n *correct.Nodelist) (res string) {
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

func insertNodeListSolution(l *correct.Nodelist, data int) *correct.Nodelist {
	n := &correct.Nodelist{Data: data}
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

func compare(l2 *correct.Nodelist, l1 *student.Nodelist, f func(a, b int) bool) {
	cmp := correct.GetName(f)

	for l1 != nil && l2 != nil {
		if l1.Data != l2.Data {
			lib.Fatalf("\nstudent list:%s\nlist:%s\nfunction cmp:%s\n\nSortListInsert() == %v instead of %v\n\n",
				listToStringStu4(l1), listToString4(l2), cmp, l1.Data, l2.Data)
			return
		}
		l1 = l1.Next
		l2 = l2.Next
	}
}

func main() {
	var linkSolutions *correct.Nodelist
	var linkStudent *student.Nodelist

	type nodeTest struct {
		data      []int
		functions []func(a, b int) bool
	}

	table := []nodeTest{}

	for i := 0; i < 4; i++ {
		table = append(table, nodeTest{
			data:      lib.MultRandIntBetween(1, 1234),
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
			solutionResult := correct.SortList(linkSolutions, fn)
			compare(solutionResult, studentResult, fn)
		}
		linkSolutions = &correct.Nodelist{}
		linkStudent = &student.Nodelist{}
	}
}
