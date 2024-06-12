
# IS-Negative

### Instructions

Write a function named `StrisNegative()` that defines if a number (You should check if it's a number) is negative or positive.

- Your function print P if the number is positive
- Your function print F if the number is negative
- For the number 0 is zero print 0.
- Any input other than number should print !
- Your function should always print ('\n')at the end of the output.

### Expected function

```go
func StrisNegative(str string){
    //code
}
```

### Usage

```go
package main

import (
	"os"
    "piscine"
	"fmt"
)

func main() {
	argument := os.Args[1:]
	piscine.StrisNegative(argument[0])
}
```
And its output:

```console
$ go run . | cat -e
!$
$ go run . "7" | cat -e
P$
$ go run . "-9" | cat -e
N$
$ go run . a | cat -e
!$
$ go run . "-552" | cat -e
N$
$ go run . "58" | cat -e
P$
$ go run . "Hello World" | cat -e
!$
```
