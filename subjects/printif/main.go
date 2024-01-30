package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.PrintIf("abcdefz"))
	fmt.Print(piscine.PrintIf("abc"))
	fmt.Print(piscine.PrintIf(""))
	fmt.Print(piscine.PrintIf("14"))
}
