package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.SaveAndMiss("123456789", 3))
	fmt.Println(piscine.SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(piscine.SaveAndMiss("", 3))
	fmt.Println(piscine.SaveAndMiss("hello you all ! ", 0))
	fmt.Println(piscine.SaveAndMiss("what is your name?", 0))
	fmt.Println(piscine.SaveAndMiss("go Exercise Save and Miss", -5))
}
