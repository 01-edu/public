package student_test

import (
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

func TestBoolean(t *testing.T) {

	table := append(z01.MultRandWords(), "1 2 3 4 5")

	for _, s := range table {
		z01.ChallengeMain(t, strings.Fields(s)...)
	}
}
