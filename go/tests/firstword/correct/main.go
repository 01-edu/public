package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		fmt.Println(strings.Split(strings.TrimSpace(os.Args[1]), " ")[0])
	}
}
