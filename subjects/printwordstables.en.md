## printwordstables

### Instructions

Write a function that prints the words of a `string` array that will be returned a function `SplitWhiteSpaces`. (the testing will be done with ours)

Each word is on a single line.

Each word ends with a `\n`.

### Expected function

```go
func PrintWordsTables(table []string) {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import 	piscine ".."

func main() {
	str := "Hello how are you?"
	table := piscine.SplitWhiteSpaces(str)
	piscine.PrintWordsTables(table)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build main.go
student@ubuntu:~/piscine-go/test$ ./main
Hello
how
are
you?
student@ubuntu:~/piscine-go/test$
```
