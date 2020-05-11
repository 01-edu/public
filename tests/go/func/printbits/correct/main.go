package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		nbr, _ := strconv.Atoi(os.Args[1])
		printBits(byte(nbr))
	}
}

func printBits(octe byte) {
	fmt.Printf("%08b", octe)
}
