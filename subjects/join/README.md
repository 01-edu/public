## join

### Instructions

Write a function that returns the concatenation of all the `string`s of a slice of `string`s **separated** by the separator passed as the argument `sep`.

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
