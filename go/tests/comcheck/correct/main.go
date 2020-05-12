package main

import (
	"fmt"
	"os"
)

func main() {
	safeWords := []string{"01", "galaxy", "galaxy 01"}
	count := 0
	for _, arg := range os.Args[1:] {
		for _, s := range safeWords {
			if s == arg {
				count++
			}
		}
	}
	if count == 1 || count == 2 {
		fmt.Println("Alert!!!")
	}
}
