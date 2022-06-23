## oddlength

### Instructions

Write a function that receives a slice of strings and returns a new slice with the strings in which the length is odd.

### Expected function

```go
func Oddlength(strings []string) []string {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
    "piscine"
    "fmt"
)

func main() {
    l := piscine.Oddlength(["Hi", "Hola", "Ola", "Ciao", "Salut", "Hallo"])
    fmt.Println(l)
}
```

And its output:

```console
$ go run .
[Ola Salut Hallo]
```
