package solutions

import (
	"os"
	"testing"
)

//function to check if the file exists in a given path
func CheckFile(t *testing.T, path string) {
	_, err := os.Stat(path)

	if err != nil {
		t.Errorf(err.Error())
	}
}
