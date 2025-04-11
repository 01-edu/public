## lastrune

### Instructions

Write a function that returns the last `rune` of a `string`.

### Expected function

```go
func LastRune(s string) rune {

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
	z01.PrintRune(piscine.LastRune("Hello!"))
	z01.PrintRune(piscine.LastRune("Salut!"))
	z01.PrintRune(piscine.LastRune("Ola!"))
	z01.PrintRune('\n')
}
```

And its output :

```console
$ go run .
!!!
$
```

### Notions

- [rune-literals](https://golang.org/ref/spec#Rune_literals)
