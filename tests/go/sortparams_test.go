package student_test

import (
	"github.com/01-edu/z01"
	"testing"
)

func TestSortParams(t *testing.T) {
	args := []string{"1", "a", "2", "A", "3", "b", "4", "C"}
	z01.ChallengeMain(t, args...)
}
