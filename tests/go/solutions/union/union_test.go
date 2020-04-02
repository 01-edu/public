package main

import (
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

func TestUnion(t *testing.T) {
	arg1 := []string{"zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj"}
	arg2 := []string{"ddf6vewg64f", "gtwthgdwthdwfteewhrtag6h4ffdhsd"}
	arg3 := []string{""}
	arg4 := []string{"rien", "cette phrase ne cache rien"}
	arg5 := []string{" this is ", " wait shr"}
	arg6 := []string{" more ", "then", "two", "arguments"}

	str1 := z01.RandAlnum()
	str2 := strings.Join([]string{z01.RandAlnum(), str1, z01.RandAlnum()}, "")

	arg7 := []string{str1, str2}
	arg8 := []string{z01.RandAlnum(), z01.RandAlnum()}
	args := [][]string{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8}

	for _, v := range args {
		z01.ChallengeMainExam(t, v...)
	}
}
