package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Base(os.Args[0]))
}
