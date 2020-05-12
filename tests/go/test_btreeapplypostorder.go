package main

import (
	"fmt"

	solutions "./solutions"
	student "./student"
)

func main() {
	root := &solutions.TreeNode{Data: "08"}
	rootS := &student.TreeNode{Data: "08"}
	pos := []string{
		"x",
		"z",
		"y",
		"t",
		"r",
		"q",
		"01",
		"b",
		"c",
		"a",
		"d",
	}

	for _, arg := range pos {
		root = solutions.BTreeInsertData(root, arg)
		rootS = student.BTreeInsertData(rootS, arg)
	}

	solutions.ChallengeTree("BTreeApplyPostorder", solutions.BTreeApplyPostorder, student.BTreeApplyPostorder, root, rootS, fmt.Println)
}
