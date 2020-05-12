package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		if i, err := strconv.Atoi(os.Args[1]); err == nil && i > 0 {
			fmt.Printf("%x\n", i)
		} else {
			fmt.Println("ERROR")
		}
	}
}
