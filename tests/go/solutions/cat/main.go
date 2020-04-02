package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	size := len(os.Args)
	if size == 1 {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			log.Fatal(err)
		}
	} else {
		for i := 1; i < size; i++ {
			data, err := ioutil.ReadFile(os.Args[i])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Print(string(data))
		}
	}
}
