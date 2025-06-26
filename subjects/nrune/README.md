## nrune

### Instructions

Write a function that returns the nth `rune` of a `string`. If not possible, it returns `0`.

### Expected function

```go
func NRune(s string, n int) rune {

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
	z01.PrintRune(piscine.NRune("Hello!", 3))
	z01.PrintRune(piscine.NRune("Salut!", 2))
	z01.PrintRune(piscine.NRune("Bye!", -1))
	z01.PrintRune(piscine.NRune("Bye!", 5))
	z01.PrintRune(piscine.NRune("Ola!", 4))
	z01.PrintRune('\n')
}
```

And its output :

```console
$ go run .
la!
$
```

### Notions

- [rune-literals](https://golang.org/ref/spec#Rune_literals)
