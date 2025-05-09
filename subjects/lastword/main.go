package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(piscine.LastWord(" lorem,ipsum "))
	fmt.Print(piscine.LastWord(" "))
}
