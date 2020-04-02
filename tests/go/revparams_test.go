package student_test

import (
	"github.com/01-edu/z01"
	"testing"
)

func TestRevParams(t *testing.T) {
	args := []string{"choumi", "is", "the", "best", "cat"}
	z01.ChallengeMain(t, args...)
}
