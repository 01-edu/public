package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.CountChar("Hello World", 'l'))
	fmt.Println(piscine.CountChar("5  balloons", 5))
	fmt.Println(piscine.CountChar("   ", ' '))
	fmt.Println(piscine.CountChar("The 7 deadly sins", '7'))
}
