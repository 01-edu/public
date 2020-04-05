package student_test

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestSortParams(t *testing.T) {
	args := []string{"1", "a", "2", "A", "3", "b", "4", "C"}
	z01.ChallengeMain(t, args...)
}
