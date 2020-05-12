package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/public/go/lib"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("ERROR: " + err.Error())
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("ERROR: " + err.Error())
	}
	fmt.Println(lib.IntRange(a, b))
}
