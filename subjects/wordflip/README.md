# wordflip

### Instructions

Write a function `WordFlip()` that receives a string and returns it with the first and last words flipped.

- Prints the Output followed by newline `\n`.
- If the string is empty, return `Invalid Output`.
- Ignore spaces if it's more then onr space in between words and all spaces in the edge of the words or sentences.

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
)

func main() {
    fmt.Print(WordFlip("First second last"))
    fmt.Print(WordFlip(""))
    fmt.Print(WordFlip("     ")) //all spaces
    fmt.Print(WordFlip(" hello  all  of  you! "))

}
```

And its output :

```go
$ go run . | cat -e
last second First$
Invalid Output$
$
you! of all hello$
```
