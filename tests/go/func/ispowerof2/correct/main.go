package main

import (
	"fmt"
	"os"
	"strconv"
)

// computes the logical bitwise AND of its operands.
func isPowerOf2(n int) bool {
	return n != 0 && ((n & (n - 1)) == 0)
}

func main() {
	if len(os.Args) == 2 {
		nbr, _ := strconv.Atoi(os.Args[1])
		fmt.Println(isPowerOf2(nbr))
	}
}
