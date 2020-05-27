package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:])%2 == 0 {
		fmt.Println("I have an even number of arguments")
	} else {
		fmt.Println("I have an odd number of arguments")
	}
}
