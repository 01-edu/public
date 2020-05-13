package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			panic(err)
		}
	} else {
		for _, arg := range os.Args[1:] {
			data, err := ioutil.ReadFile(arg)
			if err != nil {
				panic(err)
			}
			os.Stdout.Write(data)
		}
	}
}
