# ThirdTimeIsACharm

### Instructions

Write a function `ThirdTimeIsACharm()` that takes a string and returns another string with every third character.

- Prints the Output followed by newline `\n`.
- If the string is empty, return newline `\n`.
- If there is no third character, return `\n`

### Expected function

```go
func ThirdTimeIsACharm(str string) string {
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
    fmt.Print(ThirdTimeIsACharm("123456789"))
    fmt.Print(ThirdTimeIsACharm(""))
    fmt.Print(ThirdTimeIsACharm("a b c d e f")) //all spaces
    fmt.Print(ThirdTimeIsACharm("12"))

}
```

And its output :

```go
$ go run . | cat -e
369$
$
b e$
$
```
