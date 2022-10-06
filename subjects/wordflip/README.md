## wordflip

### Instructions

Write a function `WordFlip()` that receives a `string` and returns it with the first and last words flipped.

- Prints the Output followed by newline `\n`.
- If the string is empty, return `Invalid Output`.
- Ignore spaces if it's more then one space in between words and all spaces in the edge of the words or sentences.

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
