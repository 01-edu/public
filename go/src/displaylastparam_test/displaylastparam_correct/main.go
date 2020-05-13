package main

import (
	"fmt"
	"os"
)

func main() {
	size := len(os.Args)
	if size > 1 {
		fmt.Println(os.Args[size-1])
	}
}
