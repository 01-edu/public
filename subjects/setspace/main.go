package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.SetSpace("HelloWorld"))
	fmt.Println(piscine.SetSpace("HelloWorld12"))
	fmt.Println(piscine.SetSpace("Hello World"))
	fmt.Println(piscine.SetSpace(""))
	fmt.Println(piscine.SetSpace("LoremIpsumWord"))
}
