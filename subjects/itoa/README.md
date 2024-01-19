## itoa

### Instructions

- Write a function that simulates the behavior of the `Itoa` function in Go. `Itoa` transforms a number represented as an`int` in a number represented as a `string`.

- For this exercise the handling of the signs + or - **does have** to be taken into account.

## Expected function

```go
func Itoa(n int) string {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
    fmt.Println(piscine.Itoa(12345))
    fmt.Println(piscine.Itoa(0))
    fmt.Println(piscine.Itoa(-1234))
    fmt.Println(piscine.Itoa(987654321))
}
```

And its output :

```console
$ go run .
12345
0
-1234
987654321
$
```

### Notions

- [strconv/Itoa](https://pkg.go.dev/strconv#Itoa)
