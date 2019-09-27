## nrune

### Instructions

Write a function that returns the nth `rune` of a `string`.

### Expected function

```go
func NRune(s string, n int) rune {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"github.com/01-edu/z01"
	piscine ".."
)

func main() {
	z01.PrintRune(piscine.NRune("Hello!", 3))
	z01.PrintRune(piscine.NRune("Salut!", 2))
	z01.PrintRune(piscine.NRune("Ola!", 4))
	z01.PrintRune('\n')
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
la!
student@ubuntu:~/piscine-go/test$
```
