package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestPrintComb(t *testing.T) {
	z01.Challenge(t, student.PrintComb, solutions.PrintComb)
}
