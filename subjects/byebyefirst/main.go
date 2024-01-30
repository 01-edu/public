package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ByeByeFirst([]string{}))
	fmt.Println(piscine.ByeByeFirst([]string{"one arg"}))
	fmt.Println(piscine.ByeByeFirst([]string{"first", "second"}))
	fmt.Println(piscine.ByeByeFirst([]string{"", "abcd", "efg"}))
}
