## hashcode

### Instructions

Write a function called `HashCode()` that takes a `string` as an argument and returns a new **hashed** `string`.

- Hash equation: (ASCII of current character + size of the string) % 127, so it can be in the limit of the ASCII size '127'.

- If the final number gives an unprintable character, add 33.


### Expected function

```go
func HashCode(dec string) string {
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
	fmt.Println(piscine.HashCode("A"))
	fmt.Println(piscine.HashCode("AB"))
	fmt.Println(piscine.HashCode("BAC"))
	fmt.Println(piscine.HashCode("Hello World"))
}
```

And its output:

```console
$ go run .
B
CD
EDF
Spwwz+bz}wo
```
