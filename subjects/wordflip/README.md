## wordflip

### Instructions

Write a function `WordFlip()` that takes a `string` as input and returns it in reverse order.

- The output should be followed by a newline `\n`.
- If the string is empty, return `Invalid Output`.
- Ignore multiple spaces between words and trim any leading or trailing spaces in the string.

### Expected function

```go
func WordFlip(str string) string {

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
	fmt.Print(piscine.WordFlip("First second last"))
	fmt.Print(piscine.WordFlip(""))
	fmt.Print(piscine.WordFlip("     "))
	fmt.Print(piscine.WordFlip(" hello  all  of  you! "))
}
```

And its output:

```console
$ go run . | cat -e
last second First$
Invalid Output$
$
you! of all hello$
```
