## foreach

### Instructions

Write a function `ForEach` that, for an `int` slice, applies a function on each element of that slice.

### Expected function

```go
func ForEach(f func(int), a []int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func PrintNbr(n int) {
	fmt.Print(n)
}

func main() {

	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
    fmt.Println()
}
```

And its output :

```console
$ go run .
123456
$
```

### Notions

- [Function literals](https://golang.org/ref/spec#Function_literals)
- [Function declaration](https://golang.org/ref/spec#Function_declarations)
- [Function types](https://golang.org/ref/spec#Function_types)
