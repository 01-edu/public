package student_test

import (
	"testing"

	"github.com/01-edu/z01"
)

type node struct {
	flags          []string
	flagsShorthand []string
	randArgFlag    []string
	randArg        []string
}

func TestFlags(t *testing.T) {
	str := []string{"--insert=", "--order"}
	strShorthand := []string{"-i=", "-o"}
	var randflag []string
	var randflagarg []string
	for i := 0; i < 2; i++ {
		randflagarg = append(randflagarg, z01.RandWords())
		randflag = append(randflag, z01.RandWords())
	}

	node := &node{
		flags:          str,
		flagsShorthand: strShorthand,
		randArgFlag:    randflagarg,
		randArg:        randflag,
	}

	node.randArg = append(node.randArg, "")

	z01.ChallengeMain(t, node.flagsShorthand[0]+"v2", "v1")
	z01.ChallengeMain(t, node.flagsShorthand[1], "v1")
	z01.ChallengeMain(t, "-h")
	z01.ChallengeMain(t, "--help")
	z01.ChallengeMain(t)

	for _, v2 := range node.randArgFlag {
		for _, v1 := range node.randArg {
			z01.ChallengeMain(t, node.flags[0]+v2, node.flags[1], v1)
		}
	}
	for _, v2 := range node.randArgFlag {
		for _, v1 := range node.randArg {
			z01.ChallengeMain(t, node.flagsShorthand[0]+v2, node.flagsShorthand[1], v1)
		}
	}
}
