## concatparams

### Instructions

Write a function that takes the arguments received in parameters and returns them as a `string`.
The `string` is the result of all the arguments concatenated with a newline (`\n`) between.

### Expected function

```go
func ConcatParams(args []string) string {

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
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(piscine.ConcatParams(test))
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
