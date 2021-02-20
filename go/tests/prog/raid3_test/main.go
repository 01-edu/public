package main

import (
	"os/exec"
	"strconv"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	execFatal := func(name string, arg ...string) string {
		b, err := exec.Command(name, arg...).CombinedOutput()
		if _, ok := err.(*exec.ExitError); ok {
			lib.Fatalf("%s\n", b)
		} else {
			lib.Fatalf("%s\n", err)
		}
		return string(b)
	}

	// executes the test and compares student with solutions
	executeTest := func(name string, x, y int) {
		var output, correct string

		if x == 0 && y == 0 {
			output = execFatal("sh", "-c", name+" | ./raid3")
			correct = execFatal("sh", "-c", name+" | raid3_prog")
		} else {
			output = execFatal("sh", "-c", "./"+name+" "+strconv.Itoa(x)+" "+strconv.Itoa(y)+" | ./raid3")
			correct = execFatal("sh", "-c", "./"+name+" "+strconv.Itoa(x)+" "+strconv.Itoa(y)+" | raid3_prog")
		}
		if output != correct {
			lib.Fatalf("./%s %d %d | ./raid3 prints %q instead of %q\n",
				name, x, y, output, correct)
		}
	}

	// testing all raids1
	names := []string{"raid1a", "raid1b", "raid1c", "raid1d", "raid1e"}
	for _, name := range names {
		x := lib.RandIntBetween(1, 50)
		y := lib.RandIntBetween(1, 50)
		executeTest(name, x, y)
	}

	// testing special case AA, AC, A, A
	// 								A  C
	executeTest("raid1e", 2, 1)
	executeTest("raid1c", 2, 1)
	executeTest("raid1d", 1, 2)
	executeTest("raid1e", 1, 2)
	executeTest("raid1e", 1, 1)
	executeTest("echo", 0, 0)
}
