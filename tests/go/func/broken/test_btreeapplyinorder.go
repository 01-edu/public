package main

import (
	"fmt"

	"./correct"
	"./student"
)

func main() {
	root := &correct.TreeNode{Data: "08"}
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
		root = correct.BTreeInsertData(root, arg)
		rootS = student.BTreeInsertData(rootS, arg)
	}

	correct.ChallengeTree("BTreeApplyInorder", correct.BTreeApplyInorder, student.BTreeApplyInorder, root, rootS, fmt.Println)
}
