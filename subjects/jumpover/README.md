# jumpover

### Instructions

Write a function `jumpover` that takes a string and returns another string with every third character.

- Prints the output followed by newline `\n`.
- If the string is empty, return newline `\n`.
- If there is no third character, return newline `\n`.

### Expected function

```go
func JumpOver(str string) string {
}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
    fmt.Print(JumpOver("1010101010"))
    fmt.Print(JumpOver(""))
    fmt.Print(JumpOver("t w e l v e"))
    fmt.Print(JumpOver("12"))
}
```

And its output :

```go
$ go run . | cat -e
101$
$
w v$
$
```
