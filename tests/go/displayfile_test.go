package student_test

import (
	"strings"
	"testing"

	solutions "./solutions"
	"github.com/01-edu/z01"
)

func TestDisplayfile(t *testing.T) {

	var table []string
	pathFileName := "./student/displayfile/quest8.txt"

	solutions.CheckFile(t, pathFileName)

	table = append(table, "", pathFileName, "quest8.txt asdsada")

	for _, s := range table {
		z01.ChallengeMain(t, strings.Fields(s)...)
	}
}
