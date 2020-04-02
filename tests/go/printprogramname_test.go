package student_test

import (
	"github.com/01-edu/z01"
	"strings"
	"testing"
)

func TestPrintProgramName(t *testing.T) {
	exercise := strings.ToLower(
		strings.TrimPrefix(t.Name(), "Test"))
	out, err1 := z01.MainOut("./student/printprogramname")
	if err1 != nil {
		t.Errorf(err1.Error())
	}

	correct, err2 := z01.MainOut("./solutions/printprogramname")

	if err2 != nil {
		t.Errorf(err2.Error())
	}

	arrOut := strings.Split(out, "/")
	ArrCor := strings.Split(correct, "/")
	if ArrCor[len(ArrCor)-1] != arrOut[len(arrOut)-1] {
		t.Errorf("./%s prints %q instead of %q\n",
			exercise, arrOut[len(arrOut)-1], ArrCor[len(ArrCor)-1])
	}
}
