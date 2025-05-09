package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsCapitalized("Hello! How are you?"))
	fmt.Println(piscine.IsCapitalized("Hello How Are You"))
	fmt.Println(piscine.IsCapitalized("Whats 4this 100K?"))
	fmt.Println(piscine.IsCapitalized("Whatsthis4"))
	fmt.Println(piscine.IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(piscine.IsCapitalized(""))
}
