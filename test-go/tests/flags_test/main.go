package main

import (
	"github.com/01-edu/public/test-go/lib"
)

type node struct {
	flags          []string
	flagsShorthand []string
	randArgFlag    []string
	randArg        []string
}

func main() {
	s := []string{"--insert=", "--order"}
	strShorthand := []string{"-i=", "-o"}
	var randflag []string
	var randflagarg []string
	for i := 0; i < 2; i++ {
		randflagarg = append(randflagarg, lib.RandWords())
		randflag = append(randflag, lib.RandWords())
	}

	node := &node{
		flags:          s,
		flagsShorthand: strShorthand,
		randArgFlag:    randflagarg,
		randArg:        randflag,
	}

	node.randArg = append(node.randArg, "")

	lib.ChallengeMain("flags", node.flagsShorthand[0]+"v2", "v1")
	lib.ChallengeMain("flags", node.flagsShorthand[1], "v1")
	lib.ChallengeMain("flags", "-h")
	lib.ChallengeMain("flags", "--help")
	lib.ChallengeMain("flags")

	for _, v2 := range node.randArgFlag {
		for _, v1 := range node.randArg {
			lib.ChallengeMain("flags", node.flags[0]+v2, node.flags[1], v1)
		}
	}
	for _, v2 := range node.randArgFlag {
		for _, v1 := range node.randArg {
			lib.ChallengeMain("flags", node.flagsShorthand[0]+v2, node.flagsShorthand[1], v1)
		}
	}
}
