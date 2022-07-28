# ReverseSecondHalf

### Instructions
Write a function `ReverseSecondHalf()` that takes a string as an argument and prints the second half reversed. If the length of the string is odd, round it up.

- Prints the second half reversed followed by newline `\n`.
- if the string is empty, return `Invalid Output`.
- If the string's length equals one, return it, followed by newline `\n`.

### Expected function

```go
func ReverseSecondHalf(str string) string {
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
    ReverseSecondHalf("This is the 1st half This is the 2nd half")
    ReverseSecondHalf("")
    ReverseSecondHalf("Hello World")
}
```

And its output :

```go
$ go run . | cat -e
flah dn2 eht si sihT$
Invalid Output$
dlroW$
```