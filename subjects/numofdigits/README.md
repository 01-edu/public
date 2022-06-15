## numofdigits

### Instructions

Write a function that returns number of digits in a nonnegative integer n.
- if you enter a negative input return 0.

### Expected function

```go
func Numofdigits(int n) int {

}
```

### Usage

```go
package main

import (
	"github.com/01-edu/z01"
	"piscine"
)

func main() {
	z01.PrintRune(piscine.Numofdigits(3))
	z01.PrintRune(piscine.Numofdigits(245))
	z01.PrintRune(piscine.Numofdigits(-1))
	z01.PrintRune(piscine.Numofdigits(885))
	z01.PrintRune(piscine.Numofdigits(8574))
	z01.PrintRune('\n')
}
```

And its output :

```console
$ go run .
1$
3$
0$
3$
4$
