package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	var upper bool
	if args[0] == "--upper" {
		upper = true
		args = args[1:]
	}
	for _, arg := range args {
		if nb, err := strconv.Atoi(arg); err != nil || nb < 1 || nb > 26 {
			fmt.Print(" ")
		} else {
			if upper {
				nb += 'A' - 1
			} else {
				nb += 'a' - 1
			}
			fmt.Printf("%c", rune(nb))
		}
	}
	fmt.Println()
}
