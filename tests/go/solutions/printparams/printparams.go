package main

import (
	"fmt"
	"os"
)

func main() {
	for _, a := range os.Args[1:] {
		fmt.Println(a)
	}
}
