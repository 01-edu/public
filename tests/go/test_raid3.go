package main

import (
	"os/exec"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	execFatal := func(name string, arg ...string) string {
		b, err := exec.Command(name, arg...).CombinedOutput()
		s := string(b)
		if err != nil {
			z01.Fatal(s)
		}
		return s
	}

	// executes the test and compares student with solutions
	executeTest := func(args string, x, y int) {
		var output, correct string

		if x == 0 && y == 0 {
			output = execFatal("sh", "-c", args+" | ./raid3_student")
			correct = execFatal("sh", "-c", args+" | ./raid3")
		} else {
			output = execFatal("sh", "-c", "./"+args+" "+strconv.Itoa(x)+" "+strconv.Itoa(y)+" | ./raid3_student")
			correct = execFatal("sh", "-c", "./"+args+" "+strconv.Itoa(x)+" "+strconv.Itoa(y)+" | ./raid3")
		}
		if output != correct {
			z01.Fatalf("./%s %d %d | ./raid3 prints %q instead of %q\n",
				args, x, y, output, correct)
		}
	}

	// builds all the files for testing, student, solution and relevant files
	execFatal("go", "build", "-o", "raid3", "solutions/raid3/main.go")
	execFatal("go", "build", "-o", "raid3_student", "student/raid3/main.go")
	execFatal("go", "build", "solutions/raid3/raid1aProg/raid1a.go")
	execFatal("go", "build", "solutions/raid3/raid1bProg/raid1b.go")
	execFatal("go", "build", "solutions/raid3/raid1cProg/raid1c.go")
	execFatal("go", "build", "solutions/raid3/raid1dProg/raid1d.go")
	execFatal("go", "build", "solutions/raid3/raid1eProg/raid1e.go")

	// testing all raids1
	table := []string{"raid1a", "raid1b", "raid1c", "raid1d", "raid1e"}
	for _, s := range table {
		x := z01.RandIntBetween(1, 50)
		y := z01.RandIntBetween(1, 50)
		executeTest(s, x, y)
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

// TODO: fix paths
