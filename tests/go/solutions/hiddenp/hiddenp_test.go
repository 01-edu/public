package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestHiddenP(t *testing.T) {
	arg1 := []string{"fgex.;", "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6"}
	arg2 := []string{"abc", "2altrb53c.sse"}
	arg3 := []string{"abc", "btarc"}
	arg4 := []string{"DD", "DABC"}
	arg5 := []string{""}
	args := [][]string{arg1, arg2, arg3, arg4, arg5}

	for i := 0; i < 10; i++ {
		randomLowerLetter := z01.RandStr(1, z01.RuneRange('a', 'z'))
		randomUpperLetter := z01.RandStr(1, z01.RuneRange('A', 'Z'))
		extraArg := []string{randomLowerLetter, z01.RandLower()}
		extraArg2 := []string{randomUpperLetter, z01.RandUpper()}

		args = append(args, extraArg)
		args = append(args, extraArg2)

	}
	for i := 0; i < 10; i++ {
		randomLowerLetter := z01.RandStr(2, z01.RuneRange('a', 'z'))
		randomUpperLetter := z01.RandStr(2, z01.RuneRange('A', 'Z'))
		extraArg := []string{randomLowerLetter, z01.RandLower()}
		extraArg2 := []string{randomUpperLetter, z01.RandUpper()}

		args = append(args, extraArg)
		args = append(args, extraArg2)

	}
	for i := 0; i < 10; i++ {

		randomLowerLetter := z01.RandStr(1, z01.RuneRange('a', 'z'))
		randomUpperLetter := z01.RandStr(1, z01.RuneRange('A', 'Z'))
		extraArg := []string{randomLowerLetter, z01.RandLower()}
		extraArg2 := []string{randomUpperLetter, z01.RandUpper()}

		args = append(args, extraArg)
		args = append(args, extraArg2)

	}

	for _, v := range args {
		z01.ChallengeMainExam(t, v...)
	}
}
