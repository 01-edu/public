## printfirsthalf

### Instructions

Write a function called `PrintFirstHalf()` that takes a `string` as an argument and returns printed the first half of a `string`.

- If the length of the `string` is odd round it down.
- Prints the first half followed by a newline `\n`.
- If the `string` is empty, return `Invalid Output`, followed by a newline `\n`.
- If the `string` length is equal to one, return it, followed by a newline `\n`.

### Expected function

```go
func PrintFirstHalf(str string) string {

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
	fmt.Print(piscine.PrintFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Print(piscine.PrintFirstHalf(""))
	fmt.Print(piscine.PrintFirstHalf("Hello World"))
}
```

And its output:

```console
$ go run . | cat -e
This is the 1st half$
Invalid Output$
Hello$
```
