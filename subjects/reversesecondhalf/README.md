## reversesecondhalf

### Instructions

Write a function that takes a `string` as an argument and returns the second half reversed. If the length of the `string` is odd, round it up.

- Return the second half reversed followed by a newline `\n`.
- If the `string` is empty, return `Invalid Output`.
- If the `string` length equals one, return it, followed by a newline `\n`.

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
	"piscine"
)

func main() {
	fmt.Print(piscine.ReverseSecondHalf("This is the 1st half This is the 2nd half"))
	fmt.Print(piscine.ReverseSecondHalf(""))
	fmt.Print(piscine.ReverseSecondHalf("Hello World"))
}
```

And its output:

```console
$ go run . | cat -e
flah dn2 eht si sihT
Invalid Output
dlroW
$
```
