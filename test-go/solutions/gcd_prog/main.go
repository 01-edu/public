package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	if len(os.Args) == 3 {
		a, _ := strconv.Atoi(os.Args[1])
		b, _ := strconv.Atoi(os.Args[2])
		fmt.Println(lib.GCD(a, b))
	}
}
