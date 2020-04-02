package main

import (
	solutions ".."
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	fmt.Println(solutions.StrRev(os.Args[1]))
}
