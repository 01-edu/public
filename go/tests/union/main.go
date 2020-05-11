package main

import (
	"strings"

	"github.com/01-edu/public/go/lib"
)

func main() {
	arg1 := []string{"zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj"}
	arg2 := []string{"ddf6vewg64f", "gtwthgdwthdwfteewhrtag6h4ffdhsd"}
	arg3 := []string{""}
	arg4 := []string{"rien", "cette phrase ne cache rien"}
	arg5 := []string{" this is ", " wait shr"}
	arg6 := []string{" more ", "then", "two", "arguments"}

	str1 := lib.RandAlnum()
	str2 := strings.Join([]string{lib.RandAlnum(), str1, lib.RandAlnum()}, "")

	arg7 := []string{str1, str2}
	arg8 := []string{lib.RandAlnum(), lib.RandAlnum()}
	args := [][]string{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8}

	for _, v := range args {
		lib.ChallengeMain("union", v...)
	}
}
