# PrintAndMiss

### Instructions

Write a function called `PrintAndMiss()` that takes a string and an integer and prints a string containing the characters until reaching the integer, then skipping that same number of characters, and repeats until the end of the string.

- Prints the first half followed by newline `\n`.
- if the string is empty retun `Invalid Output`.


### Expected function

```go
func PrintAndMiss(arg string, loop int) string {
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
    fmt.PrintAndMiss("123456789", 3)
    fmt.PrintAndMiss("abcdefghijklmnopqrstuvwyz", 3)
    fmt.PrintAndMiss("", 3)
    fmt.PrintAndMiss("hello you all ! ", 2)
}
```

And its output :

```go
$ go run . | cat -e
123789$
abcghimnostuz$
Invalid Output$
heo u ll $
```