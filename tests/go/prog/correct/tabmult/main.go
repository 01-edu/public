package main

import (
	"fmt"
	"os"
	"strconv"
)

func Tabmul(nbr int) {
	i := 1
	for i < 10 {
		result := nbr * i
		fmt.Println(i, "x", nbr, "=", result)
		i++
	}
}

func main() {
	if len(os.Args) == 2 {
		number, _ := strconv.Atoi(os.Args[1])
		Tabmul(number)
	} else {
		fmt.Println()
	}
}
