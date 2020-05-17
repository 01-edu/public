package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		b, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(b)
	}
}
