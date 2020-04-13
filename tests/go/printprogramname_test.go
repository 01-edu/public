package student_test

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestPrintProgramName(t *testing.T) {
	b, err := exec.Command("go", "run", "./student/printprogramname").Output()
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes.TrimSpace(b)) != "printprogramname" {
		t.Fatal("Failed to print the program name")
	}
}

// TODO: add more test cases (different program names), to do so compile then rename and test the binary several times
