## fifthandskip

### Instructions

Write a function `FifthAndSkip()` that takes a `string` and returns another `string`. The function separates every five characters of the `string` with a space and removes the sixth one.

- If there is a space in the middle of a word it should ignore it and get the first character until getting to a length of 5.
- If the `string` is less than 5 characters returns "Invalid Output\n".

### Expected function

```go
func FifthAndSkip(str string) string {

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
	fmt.Print(piscine.FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(piscine.FifthAndSkip("This is a short sentence"))
	fmt.Print(piscine.FifthAndSkip("1234"))
}
```

And its output:

```console
$ go run . | cat -e
abcde ghijk mnopq stuwx z$
Thisi ashor sente ce$
Invalid Output$
```
