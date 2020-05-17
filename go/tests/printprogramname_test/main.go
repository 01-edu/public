package main

import (
	"os"
	"os/exec"

	"lib"
)

var name = "printprogramname"

func test(newName string) {
	if err := os.Rename(name, newName); err != nil {
		lib.Fatalln(err)
	}
	b, err := exec.Command("./" + name).CombinedOutput()
	if err != nil {
		lib.Fatalln(string(b))
	}
	if string(b) != name+"\n" {
		lib.Fatalln("Failed to print the program name")
	}
	name = newName
}

func main() {
	test("choumi")
	test("🤦🏻‍♀️")
	test("€")
}
