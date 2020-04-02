package student_test

import (
	"testing"

	solution "./solutions"
	student "./student"
	"github.com/01-edu/z01"
)

//the structs from the other packages
//struct just for the first exercise
type NodeF = student.Node
type NodeFS = solution.Node

//exercise 1
//simple struct, just insert a element in the struct
func TestCreateElem(t *testing.T) {
	n := &NodeFS{}
	n1 := &NodeF{}

	type nodeTest struct {
		data []int
	}

	table := []nodeTest{}

	for i := 0; i < 5; i++ {
		val := nodeTest{
			data: z01.MultRandIntBetween(-1000000, 10000000),
		}
		table = append(table, val)
	}

	table = append(table,
		nodeTest{
			data: []int{132423},
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.data); i++ {
			solution.CreateElem(n, arg.data[i])
			student.CreateElem(n1, arg.data[i])
		}

		if n.Data != n1.Data {
			t.Errorf("CreateElem == %d instead of %d\n", n, n1)
		}
	}
}
