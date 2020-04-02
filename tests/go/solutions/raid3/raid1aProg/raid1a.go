package main

import (
	"fmt"
	"os"
	"strconv"

	solutions "../.."
)

func main() {
	if len(os.Args) == 3 {
		firstArg, err := strconv.Atoi(os.Args[1])
		secundArg, err2 := strconv.Atoi(os.Args[2])

		if err != nil || err2 != nil {
			fmt.Println(err.Error())
		}
		solutions.Raid1a(firstArg, secundArg)
	} else {
		fmt.Println("to much arguments")
	}
}
