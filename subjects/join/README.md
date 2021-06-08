## join

### Instructions

Write a function that simulates the behaviour of the `Join` function in Go. This function returns the concatenation of all the strings of a slice of strings **separated** by a separator passed in argument.

### Expected function

```go
func Join(strs []string, sep string) string {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.Join(toConcat, ":"))
}
```

And its output :

```console
$ go run .
Hello!: How: are: you?
$
```

### Notions

- [strings/Join](https://golang.org/pkg/strings/#Join)
