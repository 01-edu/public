package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solution "./solutions"
	student "./student"
)

func TestRot14(t *testing.T) {
	type nodeTest struct {
		data []string
	}

	table := []nodeTest{}
	for i := 0; i < 5; i++ {
		val := nodeTest{
			data: z01.MultRandWords(),
		}
		table = append(table, val)
	}

	for _, arg := range table {
		for _, s := range arg.data {
			z01.Challenge(t, solution.Rot14, student.Rot14, s)
		}
	}
}
