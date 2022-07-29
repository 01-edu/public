# FifthAndSkip

### Instructions

Write a function `FifthAndSkip()` that takes a string and returns another one with words of 5 characters and skips the next character followed by newline `\n`. 

- If there is a space in the middle of a word it should ignore it and get the first character until getting to a length of 5.
- If the string less than 5 characters returns "Invalid Output"

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
)

func main() {
    fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
    fmt.Print(FifthAndSkip("This is a short sentence"))
    fmt.Print(FifthAndSkip("1234"))
}
```

And its output :

```go
$ go run . | cat -e
abcde ghijk mnopq stuwx yz$
Thisi ashor sente ce$
Invalid Output$
```
