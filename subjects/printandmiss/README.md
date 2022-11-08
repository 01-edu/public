## printandmiss

### Instructions

Write a function called `PrintAndMiss()` that takes a `string` and an `int` as an argument. The function should move through the `string` in sets determined by the `int`, printing the first set, omitting the second, printing the third, and so on, in a 'print' and 'miss' fashion until the end of the `string` is reached. Return a `string` containing the printed characters.

- If the `string` is empty or the `int` is negative return `Invalid Output` followed by newline `\n`.
- If the `int` is `0` return the `string` followed by a newline `\n`.

### Expected function

```go
func PrintAndMiss(arg string, num int) string {

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
	fmt.Print(piscine.PrintAndMiss("123456789", 3))
	fmt.Print(piscine.PrintAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Print(piscine.PrintAndMiss("", 3))
	fmt.Print(piscine.PrintAndMiss("hello you all ! ", 0))
	fmt.Print(piscine.PrintAndMiss("what is your name?", 0))
	fmt.Print(piscine.PrintAndMiss("go Exercise Print and Miss", -5))
}
```

And its output:

```console
$ go run . | cat -e
123789$
abcghimnostuz$
Invalid Output$
hello you all ! $
what is your name?$
Invalid Output$
```
