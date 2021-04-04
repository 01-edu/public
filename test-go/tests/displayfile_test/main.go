package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	fileName := "quest8.txt"
	if err := ioutil.WriteFile(fileName, []byte("Almost there!!\n"), os.ModePerm); err != nil {
		panic(err)
	}
	table := []string{"", fileName, fileName + " asdsada"}
	for _, s := range table {
		lib.ChallengeMain("displayfile", strings.Fields(s)...)
	}
}
