## printif

### Instructions

Write a function that takes a `string` as an argument and returns the letter `G` followed by a newline `\n` if the argument length is more than 3, otherwise returns `Invalid Input` followed by a newline `\n`.

- If it's an empty string return `G` followed by a newline `\n`.

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

	"piscine"
)

func main() {
	fmt.Print(piscine.PrintIf("abcdefz"))
	fmt.Print(piscine.PrintIf("abc"))
	fmt.Print(piscine.PrintIf(""))
	fmt.Print(piscine.PrintIf("14"))
}
```

And its output:

```console
$ go run . | cat -e
G$
Invalid Output$
G$
Invalid Output$
```
