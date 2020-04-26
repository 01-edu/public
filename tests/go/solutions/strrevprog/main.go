package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	fmt.Println(correct.StrRev(os.Args[1]))
}
