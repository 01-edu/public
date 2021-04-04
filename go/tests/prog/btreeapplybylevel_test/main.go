package main

import (
	"fmt"

	student "student"

	"github.com/01-edu/public/go/tests/func/correct"
)

func main() {
	root := &correct.TreeNode{Data: "04"}
	rootS := &student.TreeNode{Data: "04"}

	ins := []string{"01", "07", "05", "12", "02", "03", "10"}

	for _, v := range ins {
		root = correct.BTreeInsertData(root, v)
		rootS = student.BTreeInsertData(rootS, v)
	}

	correct.ChallengeTree("BTreeApplyByLevel", correct.BTreeApplyByLevel, student.BTreeApplyByLevel, root, rootS, fmt.Print)
}
