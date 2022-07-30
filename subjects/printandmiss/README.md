# PrintAndMiss

### Instructions

Write a function called `PrintAndMiss()` that takes a string and an integer and prints a string containing the characters until reaching the integer, then skipping that same number of characters, and repeats until the end of the string.

- Prints the first half followed by newline `\n`.
- If the string is empty retun `Invalid Output`.
- If the integer is `0` ot it's negative return the string followed by a newline `\n`.


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
	PrintAndMiss("123456789", 3))
	PrintAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	PrintAndMiss("", 3))
	PrintAndMiss("hello you all ! ", 0))
	PrintAndMiss("what is your name?", 0))
	PrintAndMiss("go Exercise Print and Miss", -5))
}
```

And its output :

```go
$ go run . | cat -e
123789$
abcghimnostuz$
Invalid Output$
hello you all ! $
what is your name?$
Invalid Output$
```
