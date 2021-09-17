## boolean

### Instructions

Create a new directory called `boolean`.

- The code below has to be copied in a file called `main.go` inside the `boolean` directory.

- The necessary changes have to be applied so that the program works.

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
$ go run . "not" "odd"
I have an even number of arguments
$ go run . "not even"
I have an odd number of arguments
```
