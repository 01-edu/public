package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.ReverseSecondHalf("This is the 1st half This is the 2nd half"))
	fmt.Print(piscine.ReverseSecondHalf(""))
	fmt.Print(piscine.ReverseSecondHalf("Hello World"))
}
