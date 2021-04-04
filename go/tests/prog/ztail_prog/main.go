package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	success = iota
	failure
)

var status = success

func notNil(err error) bool {
	if err != nil {
		status = failure
		fmt.Fprintln(os.Stderr, err)
		return true
	}
	return false
}

func main() {
	var bytes int64
	flag.Int64Var(&bytes, "c", 0, "output the last NUM bytes")
	flag.Parse()
	filenames := flag.Args()
	for i, filename := range filenames {
		file, err := os.Open(filename)
		if notNil(err) {
			continue
		}
		defer file.Close()
		fileInfo, err := file.Stat()
		if notNil(err) {
			continue
		}
		offset := fileInfo.Size() - bytes
		if offset < 0 {
			offset = 0
		}
		b := make([]byte, fileInfo.Size()-offset)
		_, err = file.ReadAt(b, offset)
		if notNil(err) {
			continue
		}
		if len(filenames) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Println("==>", filename, "<==")
		}
		os.Stdout.Write(b)
	}
	os.Exit(status)
}
