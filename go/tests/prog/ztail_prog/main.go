package main

import (
	"flag"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	var bytes int64
	flag.Int64Var(&bytes, "c", 0, "output the last NUM bytes")
	flag.Parse()
	filenames := flag.Args()
	for i, filename := range filenames {
		file, err := os.Open(filename)
		check(err)
		defer file.Close()
		fileInfo, err := file.Stat()
		check(err)
		offset := fileInfo.Size() - bytes
		if offset < 0 {
			offset = 0
		}
		b := make([]byte, fileInfo.Size()-offset)
		_, err = file.ReadAt(b, offset)
		check(err)
		if len(filenames) > 1 {
			fmt.Println("==>", filename, "<==")
		}
		os.Stdout.Write(b)
		if i < len(filenames)-1 {
			fmt.Println()
		}
	}
}
