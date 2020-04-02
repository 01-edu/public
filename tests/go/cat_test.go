package student_test

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	solutions "./solutions"

	"github.com/01-edu/z01"
)

//executes commands
func execC(name string, args ...string) (string, error) {
	out, err := exec.Command(name, args...).Output()

	output := string(out)
	if err == nil {
		return output, nil
	}
	if output == "" {
		return "", z01.Wrap(err, "Command failed")
	}
	return "", errors.New(output)
}

func TestCat(t *testing.T) {

	var table []string
	pathFileName := "./student/cat/quest8.txt"
	pathFileName2 := "./student/cat/quest8T.txt"

	solutions.CheckFile(t, pathFileName)
	solutions.CheckFile(t, pathFileName2)

	table = append(table, pathFileName, pathFileName+" "+pathFileName2, "asd")

	for _, s := range table {
		z01.ChallengeMain(t, strings.Fields(s)...)
	}
	_, err := execC("go", "build", "-o", "cat_student", "./student/cat/main.go")
	_, err = execC("go", "build", "-o", "cat_solution", "./solutions/cat/main.go")
	if err != nil {
		log.Fatal(err.Error())
	}
	pwd, err := os.Getwd()

	if err != nil {
		t.Errorf(err.Error())
	}

	for i := 0; i < 2; i++ {
		randStdin := z01.RandAlnum()
		cmd := exec.Command("sh", "-c", pwd+"/cat_solution")
		solutionResult := execStdin(cmd, randStdin)
		cmdS := exec.Command(pwd + "/cat_student")
		studentResult := execStdin(cmdS, randStdin)

		if solutionResult != studentResult {
			t.Errorf("./cat prints %s instead of %s\n", studentResult, solutionResult)
		}
	}
	execC("rm", "cat_student", "cat_solution")
}

func execStdin(cmd *exec.Cmd, randomStdin string) string {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	_, err = stdin.Write([]byte(randomStdin))
	if err != nil {
		log.Fatal(err)
	}
	stdin.Close()

	out, _ := ioutil.ReadAll(stdout)
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	return string(out)
}
