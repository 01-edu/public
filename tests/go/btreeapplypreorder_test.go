package student_test

import (
	"fmt"
	"testing"

	solutions "./solutions"
	student "./student"
)

func TestBTreeApplyPreorder(t *testing.T) {
	root := &solutions.TreeNode{Data: "08"}
	rootS := &student.TreeNode{Data: "08"}
	var pos []string

	pos = append(pos,
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
	)

	for _, arg := range pos {
		root = solutions.BTreeInsertData(root, arg)
		rootS = student.BTreeInsertData(rootS, arg)
	}

	solutions.ChallengeTree(t, solutions.BTreeApplyPreorder, student.BTreeApplyPreorder, root, rootS, fmt.Println)
}
