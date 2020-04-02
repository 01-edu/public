package student_test

import (
	"os"
	"path"
	"strconv"
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

// builds all the files for testing, student, solution and relevant files
func mainOut(name string) (out string, err error) {
	names := []string{"main.go", "/raid1aProg/raid1a.go", "/raid1bProg/raid1b.go", "/raid1cProg/raid1c.go", "/raid1dProg/raid1d.go", "/raid1eProg/raid1e.go"}
	var pathRaid1 string

	for i := 0; i < len(names); i++ {
		pathRaid1 = path.Join(name, names[i])
		if strings.Contains(pathRaid1, "main.go") {
			_, err = z01.ExecOut("go", "build", "-o", "raid3", pathRaid1)
			if err != nil {
				return "", err
			}
		} else {
			_, err = z01.ExecOut("go", "build", pathRaid1)
		}
	}
	if err != nil {
		return "", err
	}
	return
}

func chall(t *testing.T) {
	_, err := mainOut("./solutions/raid3")
	if err != nil {
		t.Error(err)
	}
	_, err = z01.ExecOut("go", "build", "-o", "raid3_student", "./student/raid3/main.go")
}

// executes the test and compares student with solutions
func executeTest(t *testing.T, args string, x, y int) {
	var output, correct string
	var err error

	if x == 0 && y == 0 {
		output, err = z01.ExecOut("sh", "-c", args+" | ./raid3_student")
		if err != nil {
			t.Error(err)
		}
		correct, err = z01.ExecOut("sh", "-c", args+" | ./raid3")
		if err != nil {
			t.Error(err)
		}
	} else {
		output, err = z01.ExecOut("sh", "-c", "./"+args+" "+strconv.Itoa(x)+" "+strconv.Itoa(y)+" | ./raid3_student")
		if err != nil {
			t.Error(err)
		}
		correct, err = z01.ExecOut("sh", "-c", "./"+args+" "+strconv.Itoa(x)+" "+strconv.Itoa(y)+" | ./raid3")
		if err != nil {
			t.Error(err)
		}
	}
	if output != correct {
		t.Errorf("./%s %d %d | ./raid3 prints %q instead of %q\n",
			args, x, y, output, correct)
	}
}

func TestRaid3(t *testing.T) {
	table := []string{"raid1a", "raid1b", "raid1c", "raid1d", "raid1e"}
	chall(t)
	// testing all raids1
	for _, s := range table {
		x := z01.RandIntBetween(1, 50)
		y := z01.RandIntBetween(1, 50)
		executeTest(t, s, x, y)
	}

	// testing special case AA, AC, A, A
	//								A  C
	executeTest(t, "raid1e", 2, 1)
	executeTest(t, "raid1c", 2, 1)
	executeTest(t, "raid1d", 1, 2)
	executeTest(t, "raid1e", 1, 2)
	executeTest(t, "raid1e", 1, 1)
	executeTest(t, "echo", 0, 0)

	// removing all binary files after finishing the tests
	table = append(table, "raid3_student")
	table = append(table, "raid3")
	for _, v := range table {
		err := os.Remove(v)
		if err != nil {
			t.Error(err)
		}
	}
}
