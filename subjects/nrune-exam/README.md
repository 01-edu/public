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
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
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
