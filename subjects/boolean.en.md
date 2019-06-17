## Boolean

### Instructions

Create a `.go` file and copy the code below into that file
and add the code necessary so the program works.

- The program must be submitted inside a folder with the name `boolean`.

```go
func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}

func main() {
	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
```

### Expected output

```console
student@ubuntu:~/student/boolean$ go build
student@ubuntu:~/student/boolean$ ./boolean "not" "odd"
I have an even number of arguments
student@ubuntu:~/student/boolean$ ./boolean "not even"
I have an odd number of arguments
```
