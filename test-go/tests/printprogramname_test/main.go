package main

import (
	"os"
	"os/exec"

	"github.com/01-edu/public/test-go/lib"
)

var oldName = "exe"

func test(name string) {
	if err := os.Rename(oldName, name); err != nil {
		lib.Fatalln(err)
	}
	b, err := exec.Command("./" + name).CombinedOutput()
	if err != nil {
		lib.Fatalln(err)
	}
	if string(b) != name+"\n" {
		lib.Fatalln("Failed to print the program name :", string(b))
	}
	oldName = name
}

func main() {
	test("choumi")
	test("🤦🏻‍♀️")
	test("€")
}
