## printwordstables

### Instructions

Write a function that prints the words of a `string` slice that will be returned a function `SplitWhiteSpaces`.

Each word is on a single line (each word ends with a `\n`).

### Expected function

```go
func PrintWordsTables(a []string) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import piscine ".."

func main() {
	a := piscine.SplitWhiteSpaces("Hello how are you?")
	piscine.PrintWordsTables(a)
}
```

And its output :

```console
student@ubuntu:~/printwordstables/test$ go build main.go
student@ubuntu:~/printwordstables/test$ ./main
Hello
how
are
you?
student@ubuntu:~/printwordstables/test$
```
