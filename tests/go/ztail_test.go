package student_test

import (
	"github.com/01-edu/z01"
	"strings"
	"testing"
)

//Compares only the stdout of each program
//As long as the program executes like tail for the stdout
//and the error messages are send to stderr
//the program passes the test
func TestZTail(t *testing.T) {
	exercise := strings.ToLower(
		strings.TrimPrefix(t.Name(), "Test"))

	arg1 := []string{"-c", "23", "./student/" + exercise + "/main.go"}
	arg2 := []string{"./student/" + exercise + "/main.go", "-c", "23"}
	arg3 := []string{"-c", "jfdka", "23"}
	arg4 := []string{"-c", "23", "./student/" + exercise + "/README.md", "fjksdsf"}
	arg5 := []string{"-c", "23", "../../README.md", "fjksdsf", "-c", "43"}

	argArr := [][]string{arg1, arg2, arg3, arg4, arg5}

	for _, args := range argArr {
		correct, errC := z01.ExecOut("tail", args...)
		out, err := z01.MainOut("./student/"+exercise, args...)
		if out != correct {
			t.Errorf("./%s \"%s\" prints %q instead of %q\n",
				exercise, strings.Join(args, " "), out, correct)
		}
		if errC != nil && err == nil {
			t.Errorf("./%s \"%s\" prints %q instead of %q\n", exercise, strings.Join(args, " "), "", errC)
		}
		if err != nil && errC != nil && err.Error() != errC.Error() {
			t.Errorf("./%s %s prints %q instead of %q\n", exercise, strings.Join(args, " "), err, errC)
		}

	}
}
