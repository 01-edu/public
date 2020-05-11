package main

import (
	"fmt"
	"os"
	"strconv"

	correct "../.."
)

func main() {
	if len(os.Args) == 3 {
		firstArg, err := strconv.Atoi(os.Args[1])
		secondArg, err2 := strconv.Atoi(os.Args[2])

		if err != nil || err2 != nil {
			fmt.Println(err)
		}
		correct.Raid1e(firstArg, secondArg)
	} else {
		fmt.Println("too much arguments")
	}
}
