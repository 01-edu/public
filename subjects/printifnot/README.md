## printifnot

### Instructions

Write a function that takes a string and returns:

- `"G\n"` if the string's length is less than 3 (including empty string).

- `"Invalid Input\n"` otherwise.

### Expected function

```go
func PrintIfNot(str string) string {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}
```

And its output:

```console
$ go run . | cat -e
Invalid Input$
Invalid Input$
G$
G$
```
