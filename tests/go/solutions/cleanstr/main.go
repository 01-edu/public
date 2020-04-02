package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		re := regexp.MustCompile(`( +)`)
		fmt.Print(re.ReplaceAllString(strings.Trim(os.Args[1], " "), " "))
	}
	fmt.Println()
}
