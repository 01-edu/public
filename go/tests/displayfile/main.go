package main

import (
	"os"
	"strings"

	"github.com/01-edu/public/go/lib"
)

func main() {
	pathFileName := "./student/displayfile/quest8.txt"
	_, err := os.Stat(pathFileName)
	if err != nil {
		lib.Fatal(err)
	}
	table := []string{"", pathFileName, "quest8.txt asdsada"}
	for _, s := range table {
		lib.ChallengeMain("displayfile", strings.Fields(s)...)
	}
}
