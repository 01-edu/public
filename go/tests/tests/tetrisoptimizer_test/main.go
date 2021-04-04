package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/01-edu/public/go/tests/lib"
)

type (
	vect struct{ x, y int }

	// the tetromino is described with three vectors
	// example :
	//
	// [#][#][#]
	//    [#]
	//
	// tetromino{{1, 0}, {2, 0}, {1, 1}}
	tetromino [3]vect
)

// load a tetromino composed of the rune r in the map s, example :
//
// read(`......
// ......
// ......
// ....X.
// ..XXX.
// ......
// `,
//     'X',
// ) == tetromino{{-2, 1}, {-1, 1}, {0, 1}}
func read(s string, r rune) (t tetromino) {
	var origin vect
	i := 0
	first := true
	lines := strings.Split(s, "\n")
	for y, line := range lines {
		for x, rr := range line {
			if rr == r {
				if first {
					first = false
					origin = vect{x, y}
				} else {
					t[i] = vect{x - origin.x, y - origin.y}
					i++
					if i == 3 {
						return tetromino{}
					}
				}
			}
		}
	}
	return
}

func main() {
	var (
		timeout = 30 * time.Second
		samples = "./solutions/tetrisoptimizer/samples"
		student = "./student/tetrisoptimizer"
		exe     = filepath.Join(student, "a.out")
	)
	// load samples
	f, err := os.Open(samples)
	if err != nil {
		lib.Fatal("Cannot open directory", err)
	}
	defer f.Close()
	filenames, err := f.Readdirnames(0)
	if err != nil {
		lib.Fatal("Cannot read directory", err)
	}

	// separate samples into good (valid) and bad (invalid) files
	var goodFiles, badFiles []string
	for _, filename := range filenames {
		if strings.HasPrefix(filename, "good") {
			goodFiles = append(goodFiles, filepath.Join(samples, filename))
		} else if strings.HasPrefix(filename, "bad") {
			badFiles = append(badFiles, filepath.Join(samples, filename))
		}
	}
	sort.Strings(badFiles)
	sort.Strings(goodFiles)

	// compile student code
	cmd := exec.Command("go", "build", "-o", exe, "-trimpath", "-ldflags", "-s -w", student)
	cmd.Env = append(os.Environ(), "CGO_ENABLED=0", "GOARCH=amd64")
	if out, err := cmd.CombinedOutput(); err != nil {
		lib.Fatal("Cannot compile :", string(out))
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// test invalid cases
	for _, badFile := range badFiles {
		b, _ := exec.CommandContext(ctx, exe, badFile).CombinedOutput()
		if string(b) != "ERROR\n" {
			lib.Fatal(`Failed to handle bad format, should output : "ERROR\n"`)
		}
	}

	// test valid cases
	score := .0
	defer func() {
		fmt.Println("score =", score)
	}()
	for _, goodFile := range goodFiles {
		size, _ := strconv.Atoi(strings.Split(goodFile, "-")[2])
		b, err := exec.CommandContext(ctx, exe, goodFile).Output()
		if ctx.Err() == context.DeadlineExceeded {
			return
		}
		if err != nil {
			lib.Fatal("Failed to process a valid map : execution failed")
		}
		s := string(b)
		lines := strings.Split(s, "\n")
		if lines[len(lines)-1] != "" {
			lib.Fatal(`Failed to process a valid map : missing final '\n'`)
		}
		lines = lines[:len(lines)-1]
		for _, line := range lines {
			if len(line) != len(lines) {
				lib.Fatal("Failed to process a valid map : invalid square, it is expected as many lines as characters")
			}
		}
		if len(lines) < size {
			lib.Fatal("Failed to process a valid map : the square cannot be that small")
		}
		b, err = ioutil.ReadFile(goodFile)
		if err != nil {
			lib.Fatal("Failed to read a valid map")
		}
		pieces := strings.Split(string(b), "\n\n")
		surface := len(lines) * len(lines)
		if strings.Count(s, ".") != surface-len(pieces)*4 {
			lib.Fatal("Failed to process a valid map : the number of holes (character '.') is not correct")
		}
		letter := 'A'
		for _, piece := range pieces {
			if read(s, letter) != read(piece, '#') {
				lib.Fatal("Failed to process a valid map : a tetromino is missing")
			}
			letter += 1
		}
		score += float64(size*size) / float64(surface)
	}
}

// TODO:
// Fix path
// Add error cases :
//   wrong arguments
//   bad file types (links, folders)
