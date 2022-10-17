## loafofbread

### Instructions

Write a function `LoafOfBread()` that takes a `string` and returns another one with words of 5 characters and skips the next character followed by newline `\n`.

- If there is a space in the middle of a `string` it should ignore it and get the next character until getting to a length of 5.
- If the `string` is less than 5 characters return "Invalid Output\n".

### Expected function

```go
func LoafOfBread(str string) string {

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
	fmt.Print(piscine.LoafOfBread("deliciousbread"))
	fmt.Print(piscine.LoafOfBread("This is a loaf of bread"))
	fmt.Print(piscine.LoafOfBread("loaf"))
}

```

And its output:

```go
$ go run . | cat -e
delic ousbr ad$
Thisi aloaf ofbre d$
Invalid Output$
```
