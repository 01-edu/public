package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		s := strings.Split(strings.TrimSpace(os.Args[1]), " ")[0]
		if s != "" {
			fmt.Println(s)
		}
	}
}
