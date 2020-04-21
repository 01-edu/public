package main

import (
	"os"
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	pathFileName := "./student/displayfile/quest8.txt"
	_, err := os.Stat(pathFileName)
	if err != nil {
		t.Fatal(err)
	}
	table := []string{"", pathFileName, "quest8.txt asdsada"}
	for _, s := range table {
		z01.ChallengeMain(strings.Fields(s)...)
	}
}
