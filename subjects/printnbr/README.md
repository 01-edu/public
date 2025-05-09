## printnbr

### Instructions

Write a function that prints an `int` passed in parameter.
All possible values of type `int` have to go through.
You cannot convert to `int64`.

### Expected function

```go
func PrintNbr(n int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	piscine.PrintNbr(-123)
	piscine.PrintNbr(0)
	piscine.PrintNbr(123)
	z01.PrintRune('\n')
}
```

And its output :

```console
$ go run .
-1230123
$
```

### Notions

- [01-edu/z01](https://github.com/01-edu/z01)
- [numeric types](https://golang.org/ref/spec#Numeric_types)
