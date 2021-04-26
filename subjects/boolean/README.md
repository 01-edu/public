## boolean

### Instructions

Create a `.go` file.

- The code below has to be copied in that file.

- The necessary changes have to be applied so that the program works.

- The program must be submitted inside a folder with the name `boolean`.

### Code to be copied

```go
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
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

### Usage

```console
student@ubuntu:~/boolean$ go build
student@ubuntu:~/boolean$ ./boolean "not" "odd"
I have an even number of arguments
student@ubuntu:~/boolean$ ./boolean "not even"
I have an odd number of arguments
```
