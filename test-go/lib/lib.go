package lib

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	MinInt = ^MaxInt
	MaxInt = 1<<63 - 1

	StrLen   = 13 // Default length of random strings
	SliceLen = 8  // Default length of slices
)

var (
	nsSince1970 = time.Now().UnixNano()
	bigRand     = rand.New(rand.NewSource(nsSince1970))

	// charsets
	Digit = RuneRange('0', '9')         // Decimal digit characters
	Lower = RuneRange('a', 'z')         // Lowercase latin alphabet characters
	Upper = RuneRange('A', 'Z')         // Uppercase latin alphabet characters
	ASCII = RuneRange(' ', '~')         // ASCII printable characters
	Space = strings.Repeat(" ", StrLen) // Spaces characters
	Basic = Lower + Upper               // Lower and Upper characters
	Alnum = Basic + Digit               // Basic and Digit characters
	Words = Alnum + Space               // Alnum and Space characters
)

func init() {
	rand.Seed(nsSince1970)
}

// RuneRange returns a string containing all the valid runes from a to b.
func RuneRange(a, b rune) string {
	var s []rune
	for {
		if utf8.ValidRune(a) {
			s = append(s, a)
		}
		if a == b {
			return string(s)
		}
		if a < b {
			a++
		} else {
			a--
		}
	}
}

// IntRange returns a slice containing all the int from a to b.
func IntRange(a, b int) (s []int) {
	for {
		s = append(s, a)
		if a == b {
			return
		}
		if a < b {
			a++
		} else {
			a--
		}
	}
}

// RandIntBetween returns a random int between a and b included.
func RandIntBetween(a, b int) int {
	if a > b {
		a, b = b, a
	}
	n := big.NewInt(int64(b))      // b
	n.Sub(n, big.NewInt(int64(a))) // b-a
	n.Add(n, big.NewInt(1))        // b-a+1
	n.Rand(bigRand, n)             // 0 <= n <= b-a
	n.Add(n, big.NewInt(int64(a))) // a <= n <= b
	return int(n.Int64())
}

// RandPosZ returns a random int between 0 and MaxInt included.
func RandPosZ() int { return RandIntBetween(0, MaxInt) }

// RandPos returns a random int between 1 and MaxInt included.
func RandPos() int { return RandIntBetween(1, MaxInt) }

// RandInt returns a random int between MinInt and MaxInt included.
func RandInt() int { return RandIntBetween(MinInt, MaxInt) }

// RandNeg returns a random int between MinInt and 1 included.
func RandNeg() int { return RandIntBetween(MinInt, 1) }

// RandNegZ returns a random int between MinInt and 0 included.
func RandNegZ() int { return RandIntBetween(MinInt, 0) }

// MakeIntFunc returns a slice of ints created by f.
func MakeIntFunc(f func() int) (s []int) {
	i := 0
	for i < SliceLen {
		s = append(s, f())
		i++
	}
	return
}

// MultRandPosZ returns a slice of random ints between 0 and MaxInt included.
func MultRandPosZ() []int { return MakeIntFunc(RandPosZ) }

// MultRandPos returns a slice of random ints between 1 and MaxInt included.
func MultRandPos() []int { return MakeIntFunc(RandPos) }

// MultRandInt returns a slice of random ints between MinInt and MaxInt included.
func MultRandInt() []int { return MakeIntFunc(RandInt) }

// MultRandNeg returns a slice of random ints between MinInt and 1 included.
func MultRandNeg() []int { return MakeIntFunc(RandNeg) }

// MultRandNegZ returns a slice of random ints between MinInt and 0 included.
func MultRandNegZ() []int { return MakeIntFunc(RandNegZ) }

// MultRandIntBetween returns a slice of random ints between a and b included.
func MultRandIntBetween(a, b int) []int {
	return MakeIntFunc(func() int { return RandIntBetween(a, b) })
}

// RandRune returns a random printable rune
// (although you may not have the corresponding glyph).
// One-in-ten chance to get a rune higher than 0x10000 (1<<16).
func RandRune() rune {
	ranges := unicode.PrintRanges
	table := ranges[rand.Intn(len(ranges))]
	if rand.Intn(10) == 0 {
		r := table.R32[rand.Intn(len(table.R32))]
		n := uint32(rand.Intn(int((r.Hi-r.Lo)/r.Stride) + 1))
		return rune(r.Lo + n*r.Stride)
	}
	r := table.R16[rand.Intn(len(table.R16))]
	n := uint16(rand.Intn(int((r.Hi-r.Lo)/r.Stride) + 1))
	return rune(r.Lo + n*r.Stride)
}

// RandStr returns a string with l random characters taken from chars.
// If chars is empty, the characters are random printable runes.
func RandStr(l int, chars string) string {
	if l <= 0 {
		return ""
	}
	dst := make([]rune, l)
	if chars == "" {
		for i := range dst {
			dst[i] = RandRune()
		}
	} else {
		src := []rune(chars)
		for i := range dst {
			r := rand.Intn(len(src))
			dst[i] = src[r]
		}
	}
	return string(dst)
}

// RandDigit returns a string containing random decimal digit characters.
func RandDigit() string { return RandStr(StrLen, Digit) }

// RandLower returns a string containing random lowercase latin alphabet characters.
func RandLower() string { return RandStr(StrLen, Lower) }

// RandUpper returns a string containing random uppercase latin alphabet characters.
func RandUpper() string { return RandStr(StrLen, Upper) }

// RandASCII returns a string containing random ASCII printable characters.
func RandASCII() string { return RandStr(StrLen, ASCII) }

// RandSpace returns a string containing random spaces characters.
func RandSpace() string { return RandStr(StrLen, Space) }

// RandBasic returns a string containing random lower and upper characters.
func RandBasic() string { return RandStr(StrLen, Basic) }

// RandAlnum returns a string containing random basic and digit characters.
func RandAlnum() string { return RandStr(StrLen, Alnum) }

// RandWords returns a string containing random alphanumeric and space characters.
func RandWords() string { return RandStr(StrLen, Words) }

// MakeStrFunc returns a slice of strings created by f.
func MakeStrFunc(f func() string) (s []string) {
	i := 0
	for i < StrLen {
		s = append(s, f())
		i++
	}
	return
}

// MultRandDigit returns a slice of strings containing random Decimal digit characters.
func MultRandDigit() []string { return MakeStrFunc(RandDigit) }

// MultRandLower returns a slice of strings containing random Lowercase latin alphabet.
func MultRandLower() []string { return MakeStrFunc(RandLower) }

// MultRandUpper returns a slice of strings containing random Uppercase latin alphabet.
func MultRandUpper() []string { return MakeStrFunc(RandUpper) }

// MultRandASCII returns a slice of strings containing random ASCII printable characters.
func MultRandASCII() []string { return MakeStrFunc(RandASCII) }

// MultRandSpace returns a slice of strings containing random Spaces characters.
func MultRandSpace() []string { return MakeStrFunc(RandSpace) }

// MultRandBasic returns a slice of strings containing random Lower and Upper characters.
func MultRandBasic() []string { return MakeStrFunc(RandBasic) }

// MultRandAlnum returns a slice of strings containing random Basic and Digit characters.
func MultRandAlnum() []string { return MakeStrFunc(RandAlnum) }

// MultRandWords returns a slice of strings containing random Alnum and Space characters.
func MultRandWords() []string { return MakeStrFunc(RandWords) }

func Format(a ...interface{}) string {
	ss := make([]string, len(a))
	for i, v := range a {
		switch v.(type) {
		case nil:
			ss[i] = "nil" // instead of "<nil>"
		case
			string,
			byte, // uint8
			rune: // int32

			// string     : a double-quoted string safely escaped with Go syntax
			// byte, rune : a single-quoted character literal safely escaped with Go syntax
			ss[i] = fmt.Sprintf("%q", v)
		default:
			// a Go-syntax representation of the value
			ss[i] = fmt.Sprintf("%#v", v)
		}
	}
	return strings.Join(ss, ", ")
}

var valueOf = reflect.ValueOf

func Call(fn interface{}, args []interface{}) []interface{} {
	// Convert args from []interface{} to []reflect.Value
	vals := make([]reflect.Value, len(args))
	for i, v := range args {
		if v != nil {
			vals[i] = valueOf(v)
		} else {
			vals[i] = reflect.Zero(reflect.TypeOf((*interface{})(nil)).Elem())
		}
	}

	vals = valueOf(fn).Call(vals)

	// Convert the return values from []reflect.Value to []interface{}
	result := make([]interface{}, len(vals))
	for i, v := range vals {
		result[i] = v.Interface()
	}
	return result
}

type Output struct {
	Results []interface{}
	Stdout  string
}

func Monitor(fn interface{}, args []interface{}) (out Output) {
	old := os.Stdout
	r, w, err := os.Pipe()
	panicIfNotNil(err)
	os.Stdout = w
	out.Results = Call(fn, args)
	outC := make(chan string)
	var buf strings.Builder
	go func() {
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	os.Stdout = old
	w.Close()
	out.Stdout = <-outC
	return out
}

func Challenge(name string, fn1, fn2 interface{}, args ...interface{}) {
	st1 := Monitor(fn1, args)
	st2 := Monitor(fn2, args)
	if !reflect.DeepEqual(st1.Results, st2.Results) {
		Fatalf("%s(%s) == %s instead of %s\n",
			name,
			Format(args...),
			Format(st1.Results...),
			Format(st2.Results...),
		)
	}
	if !reflect.DeepEqual(st1.Stdout, st2.Stdout) {
		Fatalf("%s(%s) prints:\n%s\ninstead of:\n%s\n",
			name,
			Format(args...),
			Format(st1.Stdout),
			Format(st2.Stdout),
		)
	}
}

func Fatal(a ...interface{}) {
	fmt.Fprint(os.Stderr, a...)
	os.Exit(1)
}

func Fatalln(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}

func Fatalf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

func panicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}

func ChallengeMainStdin(exercise, input string, args ...string) {
	run := func(pkg string) (string, int) {
		binaryPath := path.Join(os.TempDir(), "binaries", path.Base(pkg))
		panicIfNotNil(exec.Command("go", "build", "-o", binaryPath, pkg).Run())
		cmd := exec.Command(binaryPath, args...)
		if input != "" {
			cmd.Stdin = bytes.NewBufferString(input)
		}
		b, err := cmd.CombinedOutput()
		if err != nil {
			if ee, ok := err.(*exec.ExitError); ok {
				return string(b), ee.ExitCode()
			}
			panic(err)
		}
		return string(b), 0
	}
	console := func(out string) string {
		var quotedArgs []string
		for _, arg := range args {
			quotedArgs = append(quotedArgs, strconv.Quote(arg))
		}
		s := "\n$ "
		if input != "" {
			s += "echo -ne " + strconv.Quote(input) + " | "
		}
		return fmt.Sprintf(s+"./%s %s\n%s$ ", exercise, strings.Join(quotedArgs, " "), out)
	}
	code := func(code int) string {
		return fmt.Sprintf("echo $?\n%d\n$", code)
	}
	student, studentCode := run(path.Join("student", exercise))
	solution, solutionCode := run(path.Join("github.com/01-edu/public/test-go/solutions", exercise+"_prog"))
	if solutionCode == 0 {
		if studentCode != 0 {
			Fatalln("Your program fails (non-zero exit status) when it should not :\n" +
				console(student) +
				code(studentCode) + "\n\n" +
				"Expected :\n" +
				console(solution) +
				code(solutionCode))
		}
	} else {
		if studentCode == 0 {
			Fatalln("Your program does not fail when it should (with a non-zero exit status) :" + "\n" +
				console(student) +
				code(studentCode) + "\n\n" +
				"Expected :\n" +
				console(solution) +
				code(solutionCode))
		}
	}
	if student != solution {
		Fatalln("Your program output is not correct :\n" +
			console(student) + "\n\n" +
			"Expected :\n" +
			console(solution))
	}
}

// GCD returns greatest common divisor of a and b.
func GCD(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func ChallengeMain(exercise string, args ...string) {
	ChallengeMainStdin(exercise, "", args...)
}

// TODO: check unhandled errors on all solutions (it should contains "ERROR" on the first line to prove we correctly handle the error)
// TODO: remove the number of rand functions, refactor test cases (aka "table")
