package main

import (
	"os"
	"os/exec"

	"../../lib"
)

var name = "student"

func test(newName string) {
	if err := os.Rename(name, newName); err != nil {
		lib.Fatalln(err)
	}
	b, err := exec.Command("./" + name).CombinedOutput()
	if err != nil {
		lib.Fatalln(b)
	}
	if string(b) != name+"\n" {
		lib.Fatalln("Failed to print the program name")
	}
	name = newName
}

func main() {
	test("student")
	test("choumi")
	test("🤦🏻‍♀️")
	test("€")
}
