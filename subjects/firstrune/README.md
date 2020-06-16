## firstrune

### Instructions

Write a function that returns the first `rune` of a `string`.

### Expected function

```go
func FirstRune(s string) rune {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"github.com/01-edu/z01"

	piscine ".."
)

func main() {
	z01.PrintRune(piscine.FirstRune("Hello!"))
	z01.PrintRune(piscine.FirstRune("Salut!"))
	z01.PrintRune(piscine.FirstRune("Ola!"))
	z01.PrintRune('\n')
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
HSO
student@ubuntu:~/[[ROOT]]/test$
```
