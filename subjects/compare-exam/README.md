## compare

### Instructions

Write a function that behaves like the `Compare` function.

### Expected function

```go
func Compare(a, b string) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}
```

And its output :

```console
$ go run .
0
-1
1
$
```

### Notions

- [strings/Compare](https://golang.org/pkg/strings/#Compare)
