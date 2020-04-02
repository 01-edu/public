package student_test

import (
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

func TestPrintParams(t *testing.T) {
	table := append(z01.MultRandWords(), "choumi is the best cat")
	for _, s := range table {
		z01.ChallengeMain(t, strings.Fields(s)...)
	}
}
