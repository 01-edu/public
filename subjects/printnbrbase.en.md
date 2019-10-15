## putnbrbase

### Instructions

Write a function that prints an `int` in a `string` base passed in parameters.

If the base is not valid, the function prints `NV` (Not Valid):

Validity rules for a base :

-   A base must contain at least 2 characters.
-   Each character of a base must be unique.
-   A base should not contain `+` or `-` characters.

The function has to manage negative numbers. (as shown in the example)

### Expected function

```go
func PrintNbrBase(nbr int, base string) () {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"fmt"
	"github.com/01-edu/z01"
	piscine ".."
)

func main() {
	piscine.PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
125
-1111101
7D
-uoi
NV
student@ubuntu:~/piscine-go/test$
```
