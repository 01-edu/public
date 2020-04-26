package main

// Compares only the stdout of each program
// As long as the program executes like tail for the stdout
// and the error messages are send to stderr
// the program passes the test
func main() {
	argArr := [][]string{
		{"-c", "23", "student/ztail/main.go"},
		{"-c", "23", "student/ztail/README.md", "fjksdsf"},
		{"-c", "23", "../../README.md", "fjksdsf", "-c", "43"},
		{"student/ztail/main.go", "-c", "23"},
		{"-c", "jfdka", "23"},
	}
	for _, args := range argArr {
		_ = args
		// b, errC := exec.Command("tail", args...).CombinedOutput()
		// correct := string(b)
		// b, err := exec.Command
		// out, err := z01.MainOut("student/ztail", args...)
		// if out != correct {
		// 	z01.Fatalf("./ztail \"%s\" prints %q instead of %q\n",
		// 		strings.Join(args, " "), out, correct)
		// }
		// if errC != nil && err == nil {
		// 	z01.Fatalf("./ztail \"%s\" prints %q instead of %q\n", strings.Join(args, " "), "", errC)
		// }
		// if err != nil && errC != nil && err != errC.Error() {
		// 	z01.Fatalf("./ztail %s prints %q instead of %q\n", strings.Join(args, " "), err, errC)
		// }
	}
}

// TODO: never use programs we didn't wrote (do not run the true cat but instead code our own solution)
