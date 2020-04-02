package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	fileName := os.Args[1]

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(data))
}
