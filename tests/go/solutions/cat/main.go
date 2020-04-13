package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {
	size := len(os.Args)
	if size == 1 {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			panic(err)
		}
	} else {
		for i := 1; i < size; i++ {
			data, err := ioutil.ReadFile(os.Args[i])
			if err != nil {
				panic(err)
			}
			fmt.Print(string(data))
		}
	}
}
