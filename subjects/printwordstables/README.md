## printwordstables

### Instructions

Write a function that receives a `string slice` and prints each element of the slice in a seperate line.

### Expected function

```go
func PrintWordsTables(a []string) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "piscine"

func main() {
	a := piscine.SplitWhiteSpaces("Hello how are you?")
	piscine.PrintWordsTables(a)
}
```

And its output :

```console
$ go run .
Hello
how
are
you?
$
```
