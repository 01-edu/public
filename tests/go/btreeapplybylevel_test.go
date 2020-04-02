package student_test

import (
	"fmt"
	"testing"

	solutions "./solutions"
	student "./student"
)

func TestBTreeApplyByLevel(t *testing.T) {
	root := &solutions.TreeNode{Data: "04"}
	rootS := &student.TreeNode{Data: "04"}

	ins := []string{"01", "07", "05", "12", "02", "03", "10"}

	for _, v := range ins {
		root = solutions.BTreeInsertData(root, v)
		rootS = student.BTreeInsertData(rootS, v)
	}

	solutions.ChallengeTree(t, solutions.BTreeApplyByLevel, student.BTreeApplyByLevel, root, rootS, fmt.Print)
}
