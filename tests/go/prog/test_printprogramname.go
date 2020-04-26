package main

import (
	"os"
	"os/exec"

	"github.com/01-edu/z01"
)

var name = "student"

func test(newName string) {
	if err := os.Rename(name, newName); err != nil {
		z01.Fatalln(err)
	}
	b, err := exec.Command("./" + name).CombinedOutput()
	if err != nil {
		z01.Fatalln(b)
	}
	if string(b) != name+"\n" {
		z01.Fatalln("Failed to print the program name")
	}
	name = newName
}

func main() {
	test("student")
	test("choumi")
	test("🤦🏻‍♀️")
	test("€")
}
