package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestEightQueens(t *testing.T) {
	z01.Challenge(t, student.EightQueens, solutions.EightQueens)
}
