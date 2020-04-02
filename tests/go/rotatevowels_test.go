package student_test

import (
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

func TestRotateVowels(t *testing.T) {

	Lower := z01.RuneRange('a', 'z')
	Upper := z01.RuneRange('A', 'Z')
	letters := Lower + Upper + " "
	var arr []string
	for i := 0; i < 10; i++ {
		str := z01.RandStr(z01.RandIntBetween(2, 20), letters)
		arr = append(arr, str)
	}
	arr = append(arr, "")

	for _, v := range arr {
		z01.ChallengeMain(t, strings.Fields(v)...)
	}
	z01.ChallengeMain(t, "Hello World")
	z01.ChallengeMain(t, "HEllO World", "problem solved")
	z01.ChallengeMain(t, "str", "shh", "psst")
	z01.ChallengeMain(t, "happy thoughts", "good luck")
	z01.ChallengeMain(t, "al's elEphAnt is overly underweight!")
	z01.ChallengeMain(t, "aEi", "Ou")

}
