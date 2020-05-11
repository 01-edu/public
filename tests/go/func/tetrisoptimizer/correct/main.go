package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ./fillit filename")
		return
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	ok, tetriminos := validateFile(data)
	if !ok {
		fmt.Println("ERROR")
		return
	}
	board := solve(tetriminos)
	board.print()
}
