package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	pathFileName1 := "./student/cat/quest8.txt"
	pathFileName2 := "./student/cat/quest8T.txt"

	if _, err := os.Stat(pathFileName1); err != nil {
		z01.Fatalln(err)
	}
	if _, err := os.Stat(pathFileName2); err != nil {
		z01.Fatalln(err)
	}
	table := []string{pathFileName1, pathFileName1 + " " + pathFileName2, "asd"}

	for _, s := range table {
		z01.ChallengeMain("cat", strings.Fields(s)...)
	}
	if _, err := exec.Command("go", "build", "-o", "cat_student", "./student/cat/main.go").Output(); err != nil {
		z01.Fatal(string(err.(*exec.ExitError).Stderr))
	}
	if _, err := exec.Command("go", "build", "-o", "cat_solution", "./solutions/cat/main.go").Output(); err != nil {
		z01.Fatal(string(err.(*exec.ExitError).Stderr))
	}
	pwd, err := os.Getwd()
	if err != nil {
		z01.Fatalln(err)
	}

	for i := 0; i < 2; i++ {
		randStdin := z01.RandAlnum()
		cmd := exec.Command("sh", "-c", pwd+"/cat_solution")
		solutionResult := execStdin(cmd, randStdin)
		cmdS := exec.Command(pwd + "/cat_student")
		studentResult := execStdin(cmdS, randStdin)

		if solutionResult != studentResult {
			z01.Fatalf("./cat prints %s instead of %s\n", studentResult, solutionResult)
		}
	}
}

func execStdin(cmd *exec.Cmd, randomStdin string) string {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		z01.Fatalln(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		z01.Fatalln(err)
	}
	if err := cmd.Start(); err != nil {
		z01.Fatalln(err)
	}
	_, err = stdin.Write([]byte(randomStdin))
	if err != nil {
		z01.Fatalln(err)
	}
	stdin.Close()

	out, _ := ioutil.ReadAll(stdout)
	if err := cmd.Wait(); err != nil {
		z01.Fatalln(err)
	}
	return string(out)
}

// TODO: handle stdin in ChallengeMain
