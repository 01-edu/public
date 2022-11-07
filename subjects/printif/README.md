## printif

### Instructions

Write a function that takes a `string` as an argument and returns the letter `G` if the argument length is greater than or equal to 3, otherwise returns `Invalid Input` followed by a newline `\n`.
- If it's an empty `string` return `Invalid Input` followed by a newline `\n`. 

### Expected function

```go
func PrintIf(str string) string {

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
	fmt.Print(piscine.PrintIf("abcdefz"))
	fmt.Print(piscine.PrintIf("T2d97d1"))
	fmt.Print(piscine.PrintIf(""))
	fmt.Print(piscine.PrintIf("14"))
}
```

And its output:

```console
$ go run . | cat -e
G$
G$
Invalid Output$
Invalid Output$
```
