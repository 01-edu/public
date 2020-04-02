package student_test

import (
	"strconv"
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestBasicAtoi2(t *testing.T) {
	table := make([]string, 30)
	for i := range table {
		table[i] = strconv.Itoa(z01.RandPosZ())
	}
	table = append(table,
		strconv.Itoa(z01.MaxInt),
		"",
		"0",
		"Invalid123",
		"123Invalid",
		"Invalid",
		"1Invalid23",
		"123",
		"123.",
		"123.0",
	)
	for _, arg := range table {
		z01.Challenge(t, student.BasicAtoi2, solutions.BasicAtoi2, arg)
	}
}
