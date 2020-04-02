package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
	"strconv"
)

func numberOfBytes(args []string) (int, []string) {
	n := len(args)
	nbytes := 0
	var files []string
	for i, v := range args {
		var err error
		_, err = strconv.Atoi(v)
		if v == "-c" {
			if i >= n-1 {
				fmt.Printf("tail: option requires an argument -- 'c'\nTry 'tail --help' for more information.")
				os.Exit(1)
			}
			arg := args[i+1]

			nbytes, err = strconv.Atoi(arg)

			if err != nil {
				fmt.Printf("tail: invalid number of bytes: `%s`\n", arg)
				os.Exit(1)
			}
			continue
		}

		if err != nil {
			files = append(files, v)
		}

	}
	return nbytes, files
}

func fileSize(fi *os.File) int64 {
	fil, err := fi.Stat()

	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return fil.Size()
}

func main() {

	n := len(os.Args)
	if n < 4 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	nbytes, files := numberOfBytes(os.Args[1:])

	printName := len(files) > 1

	//open files for reading only
	for j, f := range files {
		fi, err := os.Open(f)
		if err != nil {
			fmt.Printf("tail: cannot open '%s' for reading: No such file or directory\n", f)
			os.Exit(1)
		}
		if printName {
			fmt.Printf("==> %s <==\n", f)
		}
		read := make([]byte, int(nbytes))
		_, er := fi.ReadAt(read, fileSize(fi)-int64(nbytes))
		if er != nil {
			fmt.Println(er.Error())
		}

		for _, c := range read {
			z01.PrintRune(rune(c))
		}

		if j < len(files)-1 {
			z01.PrintRune('\n')
		}

		fi.Close()
	}
}
