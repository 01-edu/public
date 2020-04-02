package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 2 || len(os.Args) < 2 {
		fmt.Println()
		return
	}
	nbr, _ := strconv.Atoi(os.Args[1])
	fmt.Println(isPowerOf2(nbr))
}

//& computes the logical bitwise AND of its operands.
func isPowerOf2(n int) bool {
	return n != 0 && ((n & (n - 1)) == 0)
}
