## printif

### Instructions

Write a function that takes a `string` as an argument and returns the letter `G` followed by a newline `\n` if the argument length is more or equal than 3, otherwise returns `Invalid Input` followed by a newline `\n`.

- If it's an empty string return `G` followed by a newline `\n`.

### Expected function

```go
func PrintIf(str string) string {

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
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}

```

And its output:

```console
$ go run . | cat -e
G$
G$
G$
Invalid Input$
```
