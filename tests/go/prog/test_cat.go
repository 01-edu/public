package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"../../lib"
)

func main() {
	pathFileName1 := "./student/cat/quest8.txt"
	pathFileName2 := "./student/cat/quest8T.txt"

	if _, err := os.Stat(pathFileName1); err != nil {
		lib.Fatalln(err)
	}
	if _, err := os.Stat(pathFileName2); err != nil {
		lib.Fatalln(err)
	}
	table := []string{pathFileName1, pathFileName1 + " " + pathFileName2, "asd"}

	for _, s := range table {
		lib.ChallengeMain("cat", strings.Fields(s)...)
	}
	if _, err := exec.Command("go", "build", "-o", "cat_student", "./student/cat/main.go").Output(); err != nil {
		lib.Fatal(string(err.(*exec.ExitError).Stderr))
	}
	if _, err := exec.Command("go", "build", "-o", "cat_solution", "./solutions/cat/main.go").Output(); err != nil {
		lib.Fatal(string(err.(*exec.ExitError).Stderr))
	}
	pwd, err := os.Getwd()
	if err != nil {
		lib.Fatalln(err)
	}

	for i := 0; i < 2; i++ {
		randStdin := lib.RandAlnum()
		cmd := exec.Command("sh", "-c", pwd+"/cat_solution")
		solutionResult := execStdin(cmd, randStdin)
		cmdS := exec.Command(pwd + "/cat_student")
		studentResult := execStdin(cmdS, randStdin)

		if solutionResult != studentResult {
			lib.Fatalf("./cat prints %s instead of %s\n", studentResult, solutionResult)
		}
	}
}

func execStdin(cmd *exec.Cmd, randomStdin string) string {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		lib.Fatalln(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		lib.Fatalln(err)
	}
	if err := cmd.Start(); err != nil {
		lib.Fatalln(err)
	}
	_, err = stdin.Write([]byte(randomStdin))
	if err != nil {
		lib.Fatalln(err)
	}
	stdin.Close()

	out, _ := ioutil.ReadAll(stdout)
	if err := cmd.Wait(); err != nil {
		lib.Fatalln(err)
	}
	return string(out)
}

// TODO: handle stdin in ChallengeMain
