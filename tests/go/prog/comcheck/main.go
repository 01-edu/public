package main

import (
	"fmt"
	"os"
)

func main() {
	safeWords := []string{"01", "galaxy", "galaxy 01"}
	count := 0
	for i := 1; i < len(os.Args); i++ {
		for _, s := range safeWords {
			if os.Args[i] == s {
				count++
			}
		}
	}
	if count == 1 || count == 2 {
		fmt.Println("Alert!!!")
	}
}
