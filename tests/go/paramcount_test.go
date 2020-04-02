package student_test

import (
	"strconv"
	"testing"

	"github.com/01-edu/z01"
)

func TestParamCount(t *testing.T) {
	arg1 := []string{"2", "5", "u", "19"}
	arg2 := []string{"2"}
	arg3 := []string{"1", "2", "3", "5", "7", "24"}
	arg4 := []string{"6", "12", "24"}

	args := [][]string{arg1, arg2, arg3, arg4}

	for i := 0; i < 10; i++ {
		var arg []string
		init := z01.RandIntBetween(0, 10)
		for i := init; i < init+z01.RandIntBetween(5, 10); i++ {
			arg = append(arg, strconv.Itoa(i))
		}
		args = append(args, arg)
	}

	for i := 0; i < 1; i++ {
		args = append(args, z01.MultRandWords())
	}

	for _, v := range args {
		z01.ChallengeMain(t, v...)
	}

	z01.ChallengeMain(t)
}
